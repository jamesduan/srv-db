package main

import (
	"log"
	"os"

	"github.com/micro/cli"
	proto "github.com/micro/examples/service/proto"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, res *proto.HelloResponse) error {
	res.Greeting = "Hello " + req.Name
	return nil
}

func runClient(service micro.Service) error {
	greeter := proto.NewGreeterClient("JamesGreeter", service.Client())
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "JamesDuan"})
	if err != nil {
		return err
	}
	log.Println(rsp.Greeting)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("dbservice"),
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

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
