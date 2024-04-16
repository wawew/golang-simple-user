package repository

import (
	"database/sql"

	"github.com/wawew/golang-simple-user/entity"
)

type IUserRepository interface {
	GetUserByID(id string) (*entity.User, error)
	GetAllUsers() ([]*entity.User, error)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) IUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetUserByID(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRow("SELECT Userid, Name FROM Users WHERE Userid = $1", id).Scan(&user.Userid, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetAllUsers() ([]*entity.User, error) {
	rows, err := r.db.Query("SELECT Userid, Name FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Userid, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
