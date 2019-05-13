package main

import (
	"log"
	"os"
	"srv-db/g"
	"srv-db/handler"
	proto "srv-db/proto/db"

	localCli "srv-db/cli"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("cotapi"),
		micro.Version("1.0.1"),
		micro.Flags(cli.BoolFlag{
			Name:  "run-client",
			Usage: "Luanch the client",
		}),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			if c.Bool("run-client") {
				localCli.RunClient(service)
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
