package main

import (
	common "github.com/PonyWilliam/go-common"
	"github.com/PonyWilliam/go-works/domain/repository"
	services2 "github.com/PonyWilliam/go-works/domain/server"
	"github.com/PonyWilliam/go-works/handler"
	works "github.com/PonyWilliam/go-works/proto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	"github.com/opentracing/opentracing-go"
	"strconv"

	"time"
)
func main() {
	consulConfig,err := common.GetConsualConfig("1.116.62.214",8500,"/micro/config")
	//配置中心
	if err != nil{
		log.Fatal(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(
		func(options *registry.Options){
			options.Addrs = []string{"1.116.62.214"}
			options.Timeout = time.Second * 10
		},
	)

		t,io,err := common.NewTracer("go.micro.service.works",":6831")
		if err != nil{
			log.Fatal(err)
		}
		defer io.Close()
		opentracing.SetGlobalTracer(t)
	srv := micro.NewService(
		micro.Name("go.micro.service.works"),
		micro.Version("latest"),
		micro.Address(":8083"),
		micro.Registry(consulRegistry),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(common.QPS)),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
		micro.WrapClient(hystrix.NewClientWrapper()),
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
	go common.PrometheusBoot("5006")//　协程
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