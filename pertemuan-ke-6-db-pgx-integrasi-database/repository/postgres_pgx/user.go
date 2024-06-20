package postgres_pgx

import (
	"context"
	"log"
	"praisindo/entity"
	"praisindo/service"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// PgxPoolIface defines a little interface for pgxpool functionality.
// Since in the real implementation we can use pgxpool.Pool,
// this interface exists mostly for testing purpose.
type PgxPoolIface interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Ping(ctx context.Context) error
}

type userRepository struct {
	db PgxPoolIface
}

func NewUserRepository(db PgxPoolIface) service.IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id"
	var id int
	err := r.db.QueryRow(ctx, query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.User{}, err
	}
	user.ID = id
	return *user, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1"
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	query := "UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3"
	_, err := r.db.Exec(ctx, query, user.Name, user.Email, id)
	if err != nil {
		log.Printf("Error updating user: %v\n", err)
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		log.Printf("Error deleting user: %v\n", err)
		return err
	}
	return nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	query := "SELECT id, name, email, password, created_at, updated_at FROM users"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Printf("Error getting all users: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("Error scanning user: %v\n", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
