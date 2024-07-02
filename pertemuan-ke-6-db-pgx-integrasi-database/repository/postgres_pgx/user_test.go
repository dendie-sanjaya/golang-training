package postgres_pgx_test

import (
	"context"
	"errors"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/require"
	"praisindo/entity"
	"praisindo/repository/postgres_pgx"
	"regexp"
	"testing"
	"time"
)

func TestUserRepository_CreateUser(t *testing.T) {
	// Setup mock DB
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v\n", err)
	}
	defer mock.Close()

	repo := postgres_pgx.NewUserRepository(mock)
	t.Run("Positive", func(t *testing.T) {
		// Positive case: Test successful user creation
		user := &entity.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "secret",
		}

		/*
			mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING name")).
				WithArgs(user.Name, user.Email, user.Password).
				WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(1))
		*/

		createdUser, err := repo.CreateUser(context.Background(), user)
		require.NoError(t, err)
		//require.NotNil(t, createdUser.ID)
		require.Equal(t, "John Doe", createdUser.Name) // Ensure user details are correctly set
		require.Equal(t, 200, createdUser.ID)          // Ensure user details are correctly set

	})

	t.Run("Negative", func(t *testing.T) {
		// Negative case: Test database query error
		user := &entity.User{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "secret",
		}
		mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id")).
			WithArgs(user.Name, user.Email, user.Password).
			WillReturnError(errors.New("database error"))

		_, err = repo.CreateUser(context.Background(), user)
		require.Error(t, err)
		require.EqualError(t, err, "database error")
	})
}

func TestUserRepository_GetUserByID(t *testing.T) {
	// Setup mock DB
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v\n", err)
	}
	defer mock.Close()

	repo := postgres_pgx.NewUserRepository(mock)

	t.Run("Positive", func(t *testing.T) {
		// Positive case: Test successful retrieval of user by ID
		expectedUser := entity.User{
			ID:        1,
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  "secret",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1")).
			WithArgs(expectedUser.ID).
			WillReturnRows(pgxmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
				AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email, expectedUser.Password, expectedUser.CreatedAt, expectedUser.UpdatedAt))

		user, err := repo.GetUserByID(context.Background(), expectedUser.ID)
		require.NoError(t, err)
		require.Equal(t, expectedUser.ID, user.ID)
		require.Equal(t, expectedUser.Name, user.Name)
		// Ensure other fields are correctly set
	})

	t.Run("Negative", func(t *testing.T) {
		// Negative case: Test error when user is not found
		nonExistentUserID := 999
		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1")).
			WithArgs(nonExistentUserID).
			WillReturnError(errors.New("user not found"))

		_, err = repo.GetUserByID(context.Background(), nonExistentUserID)
		require.Error(t, err)
		require.EqualError(t, err, "user not found")
	})
}

func TestUserRepository_UpdateUser(t *testing.T) {
	// Setup mock DB
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v\n", err)
	}
	defer mock.Close()

	repo := postgres_pgx.NewUserRepository(mock)

	t.Run("PositiveCase", func(t *testing.T) {
		userID := 1
		user := entity.User{
			ID:        userID,
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  "newpassword",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3")).
			WithArgs(user.Name, user.Email, userID).
			WillReturnResult(pgxmock.NewResult("UPDATE", 1))

		updatedUser, err := repo.UpdateUser(context.Background(), userID, user)
		require.NoError(t, err)
		require.Equal(t, user.ID, updatedUser.ID)
		require.Equal(t, user.Name, updatedUser.Name)
		// Ensure other fields are correctly set
	})

	t.Run("NegativeCase", func(t *testing.T) {
		userID := 1
		user := entity.User{
			ID:        userID,
			Name:      "John Doe",
			Email:     "john@example.com",
			Password:  "newpassword",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		mock.ExpectExec(regexp.QuoteMeta("UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3")).
			WithArgs(user.Name, user.Email, userID).
			WillReturnError(errors.New("update failed"))

		_, err := repo.UpdateUser(context.Background(), userID, user)
		require.Error(t, err)
		require.EqualError(t, err, "update failed")
	})
}

func TestUserRepository_DeleteUser(t *testing.T) {
	// Setup mock DB
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v\n", err)
	}
	defer mock.Close()

	repo := postgres_pgx.NewUserRepository(mock)

	t.Run("PositiveCase", func(t *testing.T) {
		userID := 1
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id = $1")).
			WithArgs(userID).
			WillReturnResult(pgxmock.NewResult("DELETE", 1))

		err := repo.DeleteUser(context.Background(), userID)
		require.NoError(t, err)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		userID := 1
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM users WHERE id = $1")).
			WithArgs(userID).
			WillReturnError(errors.New("deletion failed"))

		err := repo.DeleteUser(context.Background(), userID)
		require.Error(t, err)
		require.EqualError(t, err, "deletion failed")
	})
}

func TestUserRepository_GetAllUsers(t *testing.T) {
	// Setup mock DB
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v\n", err)
	}
	defer mock.Close()

	repo := postgres_pgx.NewUserRepository(mock)

	t.Run("PositiveCase", func(t *testing.T) {
		expectedUsers := []entity.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "password1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{ID: 2, Name: "Jane Doe", Email: "jane@example.com", Password: "password2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}
		rows := pgxmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
			AddRow(expectedUsers[0].ID, expectedUsers[0].Name, expectedUsers[0].Email, expectedUsers[0].Password, expectedUsers[0].CreatedAt, expectedUsers[0].UpdatedAt).
			AddRow(expectedUsers[1].ID, expectedUsers[1].Name, expectedUsers[1].Email, expectedUsers[1].Password, expectedUsers[1].CreatedAt, expectedUsers[1].UpdatedAt)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password, created_at, updated_at FROM users")).
			WillReturnRows(rows)

		users, err := repo.GetAllUsers(context.Background())
		require.NoError(t, err)
		require.Equal(t, expectedUsers, users)
	})

	t.Run("QueryError", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password, created_at, updated_at FROM users")).
			WillReturnError(errors.New("query failed"))

		_, err := repo.GetAllUsers(context.Background())
		require.Error(t, err)
		require.EqualError(t, err, "query failed")
	})

	t.Run("ScanError", func(t *testing.T) {
		rows := pgxmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
			AddRow(1, "John Doe", "john@example.com", "password1", time.Now(), time.Now()).
			RowError(0, errors.New("scan error"))

		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password, created_at, updated_at FROM users")).
			WillReturnRows(rows)

		res, err := repo.GetAllUsers(context.Background())
		require.NoError(t, err)
		require.Equal(t, 0, len(res))
	})
}
