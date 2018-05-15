package main

import (
	"context"
	"log"
	"os"
	"srv-db/g"
	"srv-db/handler"
	proto "srv-db/proto/db"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func runClient(service micro.Service) error {
	dbserviceClient := proto.NewDBClient("weiping.srv.dbservice", service.Client())
	rsp, err := dbserviceClient.GetUserById(context.TODO(), &proto.GetUserRequest{Id: 1})
	if err != nil {
		log.Println(err)
	}
	log.Println(rsp)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("weiping.srv.dbservice"),
		micro.Version("1.0.1"),
		micro.Flags(cli.BoolFlag{
			Name:  "run-client",
			Usage: "Luanch the client",
		}),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			if c.Bool("run-client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)
	g.InitDBConfig()

	proto.RegisterDBHandler(service.Server(), new(handler.DBServiceHandler))
	if err := service.Run(); err != nil {
		g.DB.Close()
		log.Fatal(err)
	}
}
