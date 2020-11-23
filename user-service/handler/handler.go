package handler

import (
	"errors"
	"fmt"
	pb "github.com/cg-0508/laracom/user-service/proto/user"
	"github.com/cg-0508/laracom/user-service/repo"
	"github.com/cg-0508/laracom/user-service/service"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"log"
)

type UserService struct {
	Repo repo.Repository
	Token service.Authable
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


func (srv *UserService) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with:", req.Email, req.Password)
	// 获取用户信息
	user, err := srv.Repo.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}

	// 校验用户输入密码是否于数据库存储密码匹配
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	// 生成 jwt token
	token, err := srv.Token.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (srv *UserService) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {

	// 校验用户亲求中的token信息是否有效
	claims, err := srv.Token.Decode(req.Token)

	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("无效的用户")
	}

	res.Valid = true

	return nil
}
