package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
	// uuid "github.com/satori/go.uuid"
	// "golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string		`json:"id,omitempty"`
	Email    string    	`json:"email" gorm:"unique;not null;default:null"`
	Password string    	`json:"password" gorm:"not null;default:null"`

}

// func (u *User) HashPassword() error {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(bytes)
// 	return nil
// }

// func (u *User) CheckPasswordHash(password string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, id string) error
}


var ErrRepo = errors.New("unable to handle User-Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) UserRepository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	sql := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)`

	if user.Email == "" || user.Password == "" {
		return ErrRepo
	}

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (User, error) {
	var user User
	err := repo.db.QueryRow("SELECT id, email, password FROM users WHERE id=$1", id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return User{}, ErrRepo
	}

	return user, nil
}

func (repo *repo) UpdateUser(ctx context.Context, user User) error {
	sql := `
		UPDATE users
		SET email = $2, password = $3
		WHERE id = $1`

	if user.ID == "" || user.Email == "" || user.Password == ""{
		return ErrRepo
	}

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) DeleteUser(ctx context.Context, id string) error {
	sql := `DELETE FROM users WHERE id = $1`

	_, err := repo.db.ExecContext(ctx, sql, id)
	if err != nil {
		return err
	}
	return nil
}

