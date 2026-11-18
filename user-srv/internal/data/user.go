package data

import (
	"context"

	"github.com/hao123wang/seckill-repo/user-srv/internal/biz"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{data: data}
}

func (r *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	err := r.data.db.Model(&biz.User{}).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) FindByID(ctx context.Context, userID int64) (*biz.User, error) {
	var user biz.User
	err := r.data.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByName(ctx context.Context, userName string) (*biz.User, error) {
	var user biz.User
	err := r.data.db.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
