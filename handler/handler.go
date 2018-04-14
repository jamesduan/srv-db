package handler

import (
	"context"

	// example "db/proto/example"
	"log"
	proto "srv-db/proto/db"
)

type DBServiceHandler struct{}

func (d *DBServiceHandler) GetUserById(ctx context.Context, req *proto.GetUserRequest, rsp *proto.GetUserResponse) error {
	log.Println("getUserById")
	return nil
}

func (d *DBServiceHandler) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest, rsp *proto.DeleteUserResponse) error {
	log.Println("delete user")
	return nil
}

func (d *DBServiceHandler) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest, rsp *proto.UpdateUserResponse) error {
	log.Println("update user")
	return nil
}

//  rpc GetUserById(GetUserRequest) returns (GetUserResponse) {}
// 	rpc UpdateUser(UpdateUserRequest) returns (UpdateUserResponse) {}
// 	rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse) {}
// 	rpc Stream(StreamingRequest) returns (stream StreamingResponse) {}
// 	rpc PingPong(stream Ping) returns (stream Pong) {}

// Call is a single request handler called via client.Call or the generated client code
// func () Call(ctx context.Context, req *proto., rsp *example.Response) error {
// 	log.Log("Received Example.Call request")
// 	rsp.Msg = "Hello " + req.Name
// 	return nil
// }

// // Stream is a server side stream handler called via client.Stream or the generated client code
// func (e *Example) Stream(ctx context.Context, req *example.StreamingRequest, stream example.Example_StreamStream) error {
// 	log.Logf("Received Example.Stream request with count: %d", req.Count)

// 	for i := 0; i < int(req.Count); i++ {
// 		log.Logf("Responding: %d", i)
// 		if err := stream.Send(&example.StreamingResponse{
// 			Count: int64(i),
// 		}); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // PingPong is a bidirectional stream handler called via client.Stream or the generated client code
// func (e *Example) PingPong(ctx context.Context, stream example.Example_PingPongStream) error {
// 	for {
// 		req, err := stream.Recv()
// 		if err != nil {
// 			return err
// 		}
// 		log.Logf("Got ping %v", req.Stroke)
// 		if err := stream.Send(&example.Pong{Stroke: req.Stroke}); err != nil {
// 			return err
// 		}
// 	}
// }
