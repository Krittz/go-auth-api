package repository

import (
	"database/sql"
	"go-auth-api/internal/user/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id int64) (*model.User, error)
}
type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}
func (r *userRepo) Create(user *model.User) error {
	query := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password)
	return err
}
func (r *userRepo) FindByEmail(email string) (*model.User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE email = ?`
	row := r.db.QueryRow(query, email)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepo) FindByID(id int64) (*model.User, error) {
	query := `SELECT id, name, email, password, created_at FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
