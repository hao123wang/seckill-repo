package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v3/log"
)

// User is a User model.
type User struct {
	UserID     string     `gorm:"column:id;type:varchar(64);primaryKey;not null"`
	UserName   string     `gorm:"column:user_name;type:varchar(64);not null"`
	Pwd        string     `gorm:"column:pwd;type:varchar(255);not null"`
	Sex        int        `gorm:"column:sex;type:tinyint;not null;default:0"`
	Age        int        `gorm:"column:age;type:int;not null;default:0"`
	Email      string     `gorm:"column:email;type:varchar(128);default:null"`
	Contact    string     `gorm:"column:contact;type:varchar(128);default:null"`
	Mobile     string     `gorm:"column:mobile;type:varchar(20);default:null"`
	IdCard     string     `gorm:"column:id_card;type:varchar(32);default:null"`
	CreateTime time.Time  `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP"`
	ModifyTime *time.Time `gorm:"column:modify_time;type:datetime;default:null"`
}

func (User) TableName() string {
	return "users"
}

type UserRepo interface {
	Save(ctx context.Context, user *User) (*User, error)
	FindByID(ctx context.Context, userID int64) (*User, error)
	FindByName(ctx context.Context, username string) (*User, error)
}

type UserCase struct {
	userRepo UserRepo
}

func NewUserCase(userRepo UserRepo) *UserCase {
	return &UserCase{userRepo: userRepo}
}

func (uc *UserCase) CreateUser(ctx context.Context, user *User) (*User, error) {
	log.InfoContext(ctx, "biz create user", "username", user.UserName)
	return uc.userRepo.Save(ctx, user)
}

func (uc *UserCase) GetUerInfo(ctx context.Context, userID int64) (*User, error) {
	log.InfoContext(ctx, "biz get user by id", "userID", userID)
	return uc.userRepo.FindByID(ctx, userID)
}

func (uc *UserCase) GetUserInfoByName(ctx context.Context, username string) (*User, error) {
	log.InfoContext(ctx, "biz get user by name", "username", username)
	return uc.userRepo.FindByName(ctx, username)
}
