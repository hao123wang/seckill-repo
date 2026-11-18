package service

import (
	"context"

	pb "github.com/hao123wang/seckill-repo/user-srv/api/user/v1"
	"github.com/hao123wang/seckill-repo/user-srv/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserCase
}

func NewUserService(uc *biz.UserCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	_, err := s.uc.CreateUser(ctx, &biz.User{
		UserName: req.UserName,
		Pwd:      req.Pwd,
		Sex:      int(req.Sex),
		Age:      int(req.Age),
		Email:    req.Email,
		Contact:  req.Contact,
		Mobile:   req.Mobile,
		IdCard:   req.IdCard,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{Message: "trytest ok"}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	userInfo, err := s.uc.GetUerInfo(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserReply{
		Code:    0,
		Message: "success",
		Data: &pb.GetUserReplyData{
			UserName: userInfo.UserName,
			Pwd:      userInfo.Pwd,
			Sex:      int32(userInfo.Sex),
			Age:      int32(userInfo.Age),
			Email:    userInfo.Email,
			Contact:  userInfo.Contact,
			Mobile:   userInfo.Mobile,
			IdCard:   userInfo.IdCard,
		},
	}, nil
}
func (s *UserService) GetUserByName(ctx context.Context, req *pb.GetUserByNameRequest) (*pb.GetUserByNameReply, error) {
	userInfo, err := s.uc.GetUserInfoByName(ctx, req.UserName)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByNameReply{
		Code:    0,
		Message: "success",
		Data: &pb.GetUserReplyData{
			UserID:   userInfo.UserID,
			UserName: userInfo.UserName,
			Pwd:      userInfo.Pwd,
			Sex:      int32(userInfo.Sex),
			Age:      int32(userInfo.Age),
			Email:    userInfo.Email,
			Contact:  userInfo.Contact,
			Mobile:   userInfo.Mobile,
			IdCard:   userInfo.IdCard,
		},
	}, nil
}
