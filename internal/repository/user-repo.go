package repository

import (
	"context"

	"github.com/danush754/Task-Manager/internal/db"
)

type UserRepository interface {
	CreateUser(ctx context.Context, args db.CreateUserParams) (db.TblUser, error)
	GetUserByEmail(ctx context.Context, email string) (db.TblUser, error)
	GetUserById(ctx context.Context, id int32) (db.TblUser, error)
	GetUsersList(ctx context.Context) ([]db.TblUser, error)
}
