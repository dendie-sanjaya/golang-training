package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"praisindo/entity"
	"time"

	"golang.org/x/exp/slog"
)

const (
	basedir = ""
)

var (
	httpClient *http.Client
	logger     *slog.Logger
)

func main() {
	ctx := context.Background()
	startTime := time.Now()
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	httpClient = &http.Client{Timeout: 10 * time.Second}

	slog.InfoContext(ctx, "Start retrieving users data")
	users, err := GetUsers(ctx)
	if err != nil {
		slog.WarnContext(ctx, "error when hit GetUsers", slog.Any("error", err))
		return
	}
	slog.InfoContext(ctx, fmt.Sprintf("Finished retrieving users data %d collected", len(users)))

	slog.InfoContext(ctx, "Start generating csv users")
	reportFile, _ := os.Create(basedir + "/users.csv")
	reportFileWriter := bufio.NewWriter(reportFile)
	// cetak header file
	_, _ = fmt.Fprintf(reportFileWriter, "ID,Name,Username,Email,Phone,Website\n")
	_ = reportFileWriter.Flush()

	for _, user := range users {
		_, _ = fmt.Fprintf(reportFileWriter, "%d,%s,%s,%s,%s,%s\n",
			user.ID, user.Name, user.Username, user.Email, user.Phone, user.Website)
		_ = reportFileWriter.Flush()
	}

	slog.InfoContext(ctx, fmt.Sprintf("Finish generating csv users. Elapsed Time: %d ms", time.Since(startTime).Milliseconds()))
}

func GetUsers(ctx context.Context) ([]entity.User, error) {
	const endpoint = "https://jsonplaceholder.typicode.com/users"
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		slog.WarnContext(ctx, "error when hit http.NewRequestWithContext", slog.Any("error", err))
		return nil, err
	}

	res, err := httpClient.Do(httpReq)
	if err != nil {
		slog.WarnContext(ctx, "error when hit httpClient.Do", slog.Any("error", err))
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	if res.StatusCode != http.StatusOK {
		errStatusCode := errors.New("not receiving status OK when hit API")
		slog.WarnContext(ctx, errStatusCode.Error(), slog.Any("error", errStatusCode), slog.Any("res.StatusCode", res.StatusCode))
		return nil, errStatusCode
	}

	var users []entity.User
	if err = json.NewDecoder(res.Body).Decode(&users); err != nil {
		slog.WarnContext(ctx, "error when hit Decode(&users)", slog.Any("error", err))
		return nil, err
	}

	return users, nil
}
