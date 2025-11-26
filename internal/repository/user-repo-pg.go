package repository

import (
	"context"

	"github.com/danush754/Task-Manager/internal/db"
)

type userRepo struct {
	queries *db.Queries
}

func NewUserRepo(q *db.Queries) UserRepository {
	return &userRepo{queries: q}
}

func (ur *userRepo) CreateUser(ctx context.Context, args db.CreateUserParams) (db.TblUser, error) {
	return ur.queries.CreateUser(ctx, args)
}

func (ur *userRepo) GetUserByEmail(ctx context.Context, email string) (db.TblUser, error) {
	return ur.queries.GetUserByEmail(ctx, email)
}

func (ur *userRepo) GetUserById(ctx context.Context, id int32) (db.TblUser, error) {
	return ur.queries.GetUserById(ctx, id)
}

func (ur *userRepo) GetUsersList(ctx context.Context) ([]db.TblUser, error) {
	return ur.queries.GetUsersList(ctx)
}
