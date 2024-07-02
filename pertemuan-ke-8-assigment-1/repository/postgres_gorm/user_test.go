package postgres_gorm_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"praisindo/entity"
	"praisindo/repository/postgres_gorm"
)

func setupSQLMock(t *testing.T) (sqlmock.Sqlmock, *gorm.DB) {
	// Setup SQL mock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// Setup GORM with the mock DB
	gormDB, gormDBErr := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if gormDBErr != nil {
		t.Fatalf("failed to open GORM connection: %v", gormDBErr)
	}
	return mock, gormDB
}

func TestUserRepository_CreateUser(t *testing.T) {
	// Setup SQL mock
	mock, gormDB := setupSQLMock(t)

	// Initialize userRepository with mocked GORM connection
	userRepo := postgres_gorm.NewUserRepository(gormDB)
	expectedQueryString := regexp.QuoteMeta(`INSERT INTO "users" ("name","email","created_at","updated_at") VALUES ($1,$2,$3,$4) RETURNING "name"`)

	t.Run("Positive Case", func(t *testing.T) {
		// Expected user data to insert
		user := &entity.User{
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// Set mock expectations for the transaction
		mock.ExpectQuery(expectedQueryString).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(1)) // Mock the result of the INSERT operation

		// Call the CreateUser method
		createdUser, err := userRepo.CreateUser(context.Background(), user)

		// Assert the result
		require.NoError(t, err)
		require.NotNil(t, createdUser.ID)
		require.Equal(t, user.Name, createdUser.Name)
		require.Equal(t, user.Email, createdUser.Email)
	})

}
