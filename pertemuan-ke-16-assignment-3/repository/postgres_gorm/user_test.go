package postgres_gorm_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"

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

	expectedQueryString := regexp.QuoteMeta(`INSERT INTO "users" ("name","email","password","created_at","updated_at") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)

	t.Run("Positive Case", func(t *testing.T) {
		// Expected user data to insert
		user := &entity.User{
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// Set mock expectations for the transaction
		mock.ExpectQuery(expectedQueryString).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).
				AddRow(1)) // Mock the result of the INSERT operation

		// Call the CreateUser method
		createdUser, err := userRepo.CreateUser(context.Background(), user)

		// Assert the result
		require.NoError(t, err)
		require.NotNil(t, createdUser.ID)
		require.Equal(t, user.Name, createdUser.Name)
		require.Equal(t, user.Email, createdUser.Email)
	})

	t.Run("Negative Case", func(t *testing.T) {
		// Expected user data to insert
		user := &entity.User{
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// Set mock expectations for the transaction
		mock.ExpectQuery(expectedQueryString).
			WillReturnError(errors.New("db error"))
		mock.ExpectRollback()

		// Call the CreateUser method
		createdUser, err := userRepo.CreateUser(context.Background(), user)

		// Assert the result
		require.Error(t, err)
		require.Empty(t, createdUser)
	})
}

func TestUserRepository_GetUserByID(t *testing.T) {
	// Setup SQL mock
	mock, gormDB := setupSQLMock(t)
	userRepo := postgres_gorm.NewUserRepository(gormDB)

	expectedQueryString := regexp.QuoteMeta(`SELECT "id","name","email","password","created_at","updated_at" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT $2`)

	t.Run("Positive Case", func(t *testing.T) {
		mock.ExpectQuery(expectedQueryString).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
				AddRow(1, "John Doe", "john.doe@example.com"))

		user, err := userRepo.GetUserByID(context.Background(), 1)
		require.NoError(t, err)
		require.Equal(t, "John Doe", user.Name)
		require.Equal(t, "john.doe@example.com", user.Email)
	})

	t.Run("No data found Case", func(t *testing.T) {
		mock.ExpectQuery(expectedQueryString).
			WithArgs(1, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		user, err := userRepo.GetUserByID(context.Background(), 1)
		require.NoError(t, err)
		require.Empty(t, user)
	})

	t.Run("Negative Case", func(t *testing.T) {
		mock.ExpectQuery(expectedQueryString).
			WithArgs(1, 1).
			WillReturnError(errors.New("db down"))

		user, err := userRepo.GetUserByID(context.Background(), 1)
		require.Error(t, err)
		require.Empty(t, user)
	})
}

func TestUserRepository_UpdateUser(t *testing.T) {
	// Setup SQL mock
	mock, gormDB := setupSQLMock(t)
	userRepo := postgres_gorm.NewUserRepository(gormDB)

	expectedSelectQueryString := regexp.QuoteMeta(`SELECT "id","name","email","password","created_at","updated_at" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT $2`)

	expectedUpdateQueryString := regexp.QuoteMeta(`UPDATE "users" SET "name"=$1,"email"=$2,"password"=$3,"created_at"=$4,"updated_at"=$5 WHERE "id" = $6`)

	t.Run("Positive Case - select and update success", func(t *testing.T) {
		mock.ExpectQuery(expectedSelectQueryString).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
				AddRow(1, "John Doe", "john.doe@example.com"))

		mock.ExpectExec(expectedUpdateQueryString).
			WillReturnResult(sqlmock.NewResult(1, 1))

		user := entity.User{
			ID:    1,
			Name:  "Updated Name",
			Email: "updated.email@example.com",
		}

		updatedUser, err := userRepo.UpdateUser(context.Background(), user.ID, user)
		require.NoError(t, err)
		require.Equal(t, user.Name, updatedUser.Name)
		require.Equal(t, user.Email, updatedUser.Email)
	})

	t.Run("Negative Case - error on selecting rows", func(t *testing.T) {
		mock.ExpectQuery(expectedSelectQueryString).
			WithArgs(1, 1).
			WillReturnError(errors.New("database down"))

		user := entity.User{
			ID:    1,
			Name:  "Updated Name",
			Email: "updated.email@example.com",
		}

		updatedUser, err := userRepo.UpdateUser(context.Background(), user.ID, user)

		require.Error(t, err)
		require.Empty(t, updatedUser)
	})

	t.Run("Negative Case - error on updating rows", func(t *testing.T) {
		mock.ExpectQuery(expectedSelectQueryString).
			WithArgs(1, 1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
				AddRow(1, "John Doe", "john.doe@example.com"))

		mock.ExpectExec(expectedUpdateQueryString).
			WillReturnError(errors.New("db error"))
		mock.ExpectRollback()

		user := entity.User{
			ID:    1,
			Name:  "Updated Name",
			Email: "updated.email@example.com",
		}

		updatedUser, err := userRepo.UpdateUser(context.Background(), user.ID, user)

		require.Error(t, err)
		require.Empty(t, updatedUser)
	})
}

func TestUserRepository_DeleteUser(t *testing.T) {
	// Setup SQL mock
	mock, gormDB := setupSQLMock(t)
	userRepo := postgres_gorm.NewUserRepository(gormDB)

	expectedQueryString := regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1`)

	t.Run("Positive Case", func(t *testing.T) {
		mock.ExpectExec(expectedQueryString).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := userRepo.DeleteUser(context.Background(), 1)
		require.NoError(t, err)
	})

	t.Run("Negative Case", func(t *testing.T) {

		mock.ExpectExec(expectedQueryString).
			WithArgs(1).
			WillReturnError(errors.New("db error"))
		mock.ExpectRollback()

		err := userRepo.DeleteUser(context.Background(), 1)

		require.Error(t, err)
	})
}

func TestUserRepository_GetAllUsers(t *testing.T) {
	// Setup SQL mock
	mock, gormDB := setupSQLMock(t)
	userRepo := postgres_gorm.NewUserRepository(gormDB)

	expectedQueryString := regexp.QuoteMeta(`SELECT "id","name","email","password","created_at","updated_at" FROM "users"`)

	t.Run("Positive Case", func(t *testing.T) {
		mock.ExpectQuery(expectedQueryString).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
				AddRow(1, "John Doe", "john.doe@example.com").
				AddRow(2, "Jane Doe", "jane.doe@example.com"))

		users, err := userRepo.GetAllUsers(context.Background())
		require.NoError(t, err)
		require.Len(t, users, 2)
		require.Equal(t, "John Doe", users[0].Name)
		require.Equal(t, "john.doe@example.com", users[0].Email)
		require.Equal(t, "Jane Doe", users[1].Name)
		require.Equal(t, "jane.doe@example.com", users[1].Email)
	})

	t.Run("No data found Case", func(t *testing.T) {
		mock.ExpectQuery(expectedQueryString).
			WillReturnError(gorm.ErrRecordNotFound)

		users, err := userRepo.GetAllUsers(context.Background())
		require.NoError(t, err)
		require.Empty(t, users)
	})

	t.Run("Negative Case", func(t *testing.T) {
		mock.ExpectQuery(expectedQueryString).
			WillReturnError(errors.New("error db"))

		users, err := userRepo.GetAllUsers(context.Background())
		require.Error(t, err)
		require.Empty(t, users)
	})
}
