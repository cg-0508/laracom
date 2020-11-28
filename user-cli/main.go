package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	pb "github.com/cg-0508/laracom/user-service/proto/user"
	"golang.org/x/net/context"
	"log"
	"os"
)

/**
go run main.go --registry=mdns --name="aaa" --email="asda@1.com" --password="asdasdasdsa"
go run main.go --registry=mdns --name="bbb" --email="bbb@1.com" --password="bbbb" --userId=3194f72b-8ecc-4396-9fb8-961f78f7d6fe
 */
func main()  {

	// 初始化客户端服务，定义命令行参数标识
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "name",
				Usage: "Your Name",
			},
			&cli.StringFlag{
				Name:  "email",
				Usage: "Your Email",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Your Password",
			},
			&cli.StringFlag{
				Name:  "userId",
				Usage: "userId",
			},
		),
	)

	// 远程服务客户端调用句柄
	client := pb.NewUserService("laracom.service.user", service.Client())

	// 运行客户端命令调用远程服务逻辑设置
	service.Init(
		micro.Action(func(c *cli.Context) error {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			userId := c.String("userId")

			log.Println("参数:", name, email, password, userId)

			// 调用用户服务
			if userId != "" {
				r, err := client.Update(context.TODO(), &pb.User{
					Id: userId,
					Name: name,
					Email: email,
					Password: password,
				})
				if err != nil {
					log.Fatalf("更新用户失败: %v", err)
					return err
				}
				log.Printf("更新用户成功: %s", r.User.Id)

			}else{
				r, err := client.Create(context.TODO(), &pb.User{
					Name: name,
					Email: email,
					Password: password,
				})
				if err != nil {
					log.Fatalf("创建用户失败: %v", err)
					return err
				}
				log.Printf("创建用户成功: %s", r.User.Id)
				// 调用用户认证服务
				var token *pb.Token
				token, err = client.Auth(context.TODO(), &pb.User{
					Email: email,
					Password: password,
				})
				if err != nil {
					log.Fatalf("用户登录失败: %v", err)
				}
				log.Printf("用户登录成功：%s", token.Token)

				// 调用用户验证服务
				token, err = client.ValidateToken(context.TODO(), token)
				if err != nil {
					log.Fatalf("用户认证失败: %v", err)
				}
				log.Printf("用户认证成功：%s", token.Valid)

				getAll, err := client.GetAll(context.Background(), &pb.Request{})
				if err != nil {
					log.Fatalf("获取所有用户失败: %v", err)
					return err
				}
				for _, v := range getAll.Users {
					log.Println(v)
				}

				// 重置密码
				resetResp, err := client.CreatePasswordReset(context.Background(), &pb.PasswordReset{
					Email: email,
					Token: "password_reset_token",
				})
				log.Printf("密码重置token记录：%v", resetResp)
				// 验证重置token
				validateResetToken, err := client.ValidatePasswordResetToken(context.Background(), &pb.Token{
					Token: "password_reset_token",
				})
				log.Printf("密码重置token验证结果：%v", validateResetToken.Valid)

			}


			os.Exit(0)
			return nil
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatalf("用户客户端启动失败: %v", err)
	}
}