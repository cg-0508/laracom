package handler

import (
	"fmt"
	pb "github.com/cg-0508/laracom/user-service/proto/user"
	"github.com/cg-0508/laracom/user-service/repo"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
)

type UserService struct {
	Repo repo.Repository
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	fmt.Print(111)
	// 对密码进行哈希加密
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	fmt.Print(222)

	if err != nil {
		log.Fatalf("hashedPass error: %v", err)
		return err
	}
	fmt.Print(333)

	req.Password = string(hashedPass)
	fmt.Print(444)

	if err := srv.Repo.Create(req); err != nil {
		log.Fatalf("创建用户失败: %v", err)
		return err
	}
	fmt.Print(555)

	res.User = req
	fmt.Print(666)

	return nil
}