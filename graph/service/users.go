package service

import (
	"context"

	"github.com/MizukiShigi/graphql-study/graph/db"
	"github.com/MizukiShigi/graphql-study/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)


type userService struct {
	exec boil.ContextExecutor
}

func(u *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.ID.EQ(id),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.Name.EQ(name),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

func (u *userService) AddUserById(ctx context.Context, id string, name string) (*model.AddUserByIDPayload, error) {
	user := &db.User{
		ID: id,
		Name: name,
	}
	if err := user.Insert(ctx, u.exec, boil.Infer()); err != nil {
		return nil, err
	}

	return convertAddUserByIDPayload(user), nil
}

func (u *userService) GetUsersByIDs(ctx context.Context, IDs []string) ([]*model.User, error) {
	users, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.ID.IN(IDs),
	).All(ctx, u.exec)
	if err != nil {
		return nil, err
	}
	return convertUserSlice(users), nil
}

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID: user.ID,
		Name: user.Name,
	}
}

func convertUserSlice(users db.UserSlice) []*model.User {
	var userSlice []*model.User
	for _, user := range users {
		userSlice = append(userSlice, convertUser(user))
	}
	return userSlice
}

func convertAddUserByIDPayload(user *db.User) *model.AddUserByIDPayload {
	return &model.AddUserByIDPayload{
		User: convertUser(user),
	}
}