package cli

import (
	"context"
	"log"

	proto "srv-db/proto/db"

	micro "github.com/micro/go-micro"
)

func RunClient(service micro.Service) error {
	dbserviceClient := proto.NewDBClient("weiping.srv.dbservice", service.Client())
	rsp, err := dbserviceClient.GetUserById(context.TODO(), &proto.GetUserRequest{Id: 1})
	if err != nil {
		log.Println(err)
	}
	log.Println(rsp)
	return nil
}
