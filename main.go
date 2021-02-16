package main

import (
	"github.com/PonyWilliam/go-works/domain/repository"
	services2 "github.com/PonyWilliam/go-works/domain/server"
	"github.com/PonyWilliam/go-works/handler"
	works "github.com/PonyWilliam/go-works/proto"
	"strconv"
	"time"

	common "github.com/PonyWilliam/go-common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
)
var QPS = 1000
func main() {
	consulConfig,err := common.GetConsualConfig("127.0.0.1",8500,"/micro/config")
	//配置中心
	if err != nil{
		log.Fatal(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(
		func(options *registry.Options){
			options.Addrs = []string{"127.0.0.1"}
			options.Timeout = time.Second * 10
		},
	)

	srv := micro.NewService(
		micro.Name("go.micro.service.works"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	db,err := gorm.Open("mysql",
		mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host + ":"+ strconv.FormatInt(mysqlInfo.Port,10) +")/"+mysqlInfo.DataBase+"?charset=utf8&parseTime=True&loc=Local",
	)
	if err != nil{
		log.Error(err)

	}
	defer db.Close()
	db.SingularTable(true)
	srv.Init()
	rp := repository.NewWorkerRepository(db)
	err =rp.InitTable()
	if err!=nil{
		err := rp.InitTable()
		if err!=nil{
			log.Error(err)
		}
	}

	WorkServices := services2.NewWorkerServices(repository.NewWorkerRepository(db))
	err = works.RegisterWorksHandler(srv.Server(),&handler.Works{WorkService:WorkServices})

	if err:=srv.Run();err!=nil{
		log.Fatal(err)
	}
}