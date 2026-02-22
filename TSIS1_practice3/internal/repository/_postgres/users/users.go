package users

import (
	"database/sql"
	"errors"

	"golang/internal/repository/_postgres"
	"golang/pkg/modules"
)

var ErrUserNotFound = errors.New("user not found")

type Repository struct {
	db *_postgres.Dialect
}

func NewUserRepository(db *_postgres.Dialect) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetUsers() ([]modules.User, error) {
	var users []modules.User
	err := r.db.DB.Select(&users, "SELECT * FROM users ORDER BY id")
	return users, err
}

func (r *Repository) GetUserByID(id int) (*modules.User, error) {
	var u modules.User
	err := r.db.DB.Get(&u, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (r *Repository) CreateUser(u modules.User) (int, error) {
	var id int
	q := `INSERT INTO users (name, email, age) VALUES ($1,$2,$3) RETURNING id`
	err := r.db.DB.QueryRow(q, u.Name, u.Email, u.Age).Scan(&id)
	return id, err
}

func (r *Repository) UpdateUser(id int, u modules.User) error {
	res, err := r.db.DB.Exec(
		"UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4",
		u.Name, u.Email, u.Age, id,
	)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}

func (r *Repository) DeleteUser(id int) error {
	res, err := r.db.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}
