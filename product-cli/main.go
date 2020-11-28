package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	pb "github.com/cg-0508/laracom/product-service/proto/product"
	"golang.org/x/net/context"
	"log"
	"os"
)

/**
go run main.go --registry=mdns
 */
func main()  {

	// 初始化客户端服务，定义命令行参数标识
	service := micro.NewService()

	// 远程服务客户端调用句柄
	client := pb.NewProductService("laracom.service.product", service.Client())

	// 运行客户端命令调用远程服务逻辑设置
	service.Init(
		micro.Action(func(c *cli.Context) error {
			r, err := client.Create(context.TODO(), &pb.Product{
				Id: 11,
				BrandId: 11,
				Sku: "2222",
				Name: "asdas",
				Slug: "asdasdas",
				Description: "asdasdsad",
				Price: 11,
				SalePrice: 11,
				Status: 11,
			})
			if err != nil {
				log.Fatalf("创建商品失败: %v", err)
				return err
			}
			log.Printf("创建商品成功: %s", r.Product.Id)

			os.Exit(0)
			return nil
		}),
	)

	if err := service.Run(); err != nil {
		log.Fatalf("用户客户端启动失败: %v", err)
	}
}