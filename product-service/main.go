package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/cg-0508/laracom/product-service/db"
	"github.com/cg-0508/laracom/product-service/handler"
	"github.com/cg-0508/laracom/product-service/model"
	pb "github.com/cg-0508/laracom/product-service/proto/product"
	"github.com/cg-0508/laracom/product-service/repo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
)



// 启动 HTTP 服务监听客户端数据采集
func prometheusBoot() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":9092", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}




func main()  {
	// 创建数据库连接，程序退出时断开连接
	database, err := db.CreateConnection()
	defer database.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// 数据库迁移
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.ProductImage{})
	database.AutoMigrate(&model.Brand{})
	database.AutoMigrate(&model.Category{})
	database.AutoMigrate(&model.Attribute{})
	database.AutoMigrate(&model.AttributeValue{})
	database.AutoMigrate(&model.ProductAttribute{})

	// 初始化 Repo 实例用于后续数据库操作
	productRepo := &repo.ProductRepository{database}
	imageRepo := &repo.ImageRepository{database}
	brandRepo := &repo.BrandRepository{database}
	categoryRepo := &repo.CategoryRepository{database}
	attributeRepo := &repo.AttributeRepository{database}


	// 以下是 Micro 创建微服务流程
	srv := micro.NewService(
		micro.Name("laracom.service.product"),
		micro.Version("latest"),  // 新增接口版本参数
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)
	srv.Init()

	// 注册处理器
	pb.RegisterProductServiceHandler(srv.Server(), &handler.ProductService{productRepo})
	pb.RegisterImageServiceHandler(srv.Server(), &handler.ImageService{imageRepo})
	pb.RegisterBrandServiceHandler(srv.Server(), &handler.BrandService{brandRepo})
	pb.RegisterCategoryServiceHandler(srv.Server(), &handler.CategoryService{categoryRepo})
	pb.RegisterAttributeServiceHandler(srv.Server(), &handler.AttributeService{attributeRepo})

	// 采集监控数据
	prometheusBoot()

	// 启动商品服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}