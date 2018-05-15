package handler

import (
	"context"
	"fmt"

	// example "db/proto/example"
	"log"
	"srv-db/g"
	proto "srv-db/proto/db"
)

type DBServiceHandler struct{}

func (d *DBServiceHandler) GetUserById(ctx context.Context, req *proto.GetUserRequest, rsp *proto.GetUserResponse) error {
	log.Println("getUserById")
	sql := fmt.Sprintf("select * from user where id=%d", req.Id)
	// log.Println(sql)
	rows, err := g.DB.Query(sql)
	// log.Println(rows)
	for rows.Next() {
		var (
			id          int64
			name        string
			email       string
			sex         string
			portraitURL string
			createTime  string
			updateTime  string
		)
		err := rows.Scan(&id, &name, &email, &sex, &portraitURL, &createTime, &updateTime)
		if err != nil {
			log.Println(err)
		}
		rsp.Id = id
		rsp.Name = name
		rsp.Email = email
		rsp.Sex = sex
		rsp.PortraitUrl = portraitURL
		rsp.CreateTime = createTime
		rsp.UpdateTime = updateTime
		// log.Println(id, name, email, sex, portraitURL, createTime, updateTime)
	}
	if err != nil {
		log.Println("exec", sql, "fail", err)
		return err
	}
	defer rows.Close()
	return nil
}

func (d *DBServiceHandler) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest, rsp *proto.DeleteUserResponse) error {
	log.Println("delete user")
	sql := fmt.Sprintf("delete from user where id=%d", req.Id)
	_, err := g.DB.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func (d *DBServiceHandler) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest, rsp *proto.UpdateUserResponse) error {
	log.Println("update user")
	log.Println(req)
	sql := fmt.Sprintf(
		"update user set name='%s', email='%s', sex='%s', portrait_url='%s', create_time='%s', update_time='%s' where id=%d",
		req.Name,
		req.Email,
		req.Sex,
		req.PortraitUrl,
		req.CreateTime,
		req.UpdateTime,
		req.Id,
	)
	log.Println(sql)
	result, err := g.DB.Exec(sql)
	log.Println(result.RowsAffected())
	if err != nil {
		log.Println(err)
		return err
	}
	// log.Println(result.LastInsertId())
	return nil
}

func (d *DBServiceHandler) CreateUser(ctx context.Context, req *proto.CreateUserRequest, rsp *proto.CreateUserResponse) error {
	log.Println("Create User")
	sql := fmt.Sprintf(
		"insert into user(name, email, sex, portrait_url, create_time, update_time) values ('%s', '%s', '%s', '%s', '%s', '%s')",
		req.Name,
		req.Email,
		req.Sex,
		req.PortraitUrl,
		req.CreateTime,
		req.UpdateTime,
	)
	dbrsp, err := g.DB.Exec(sql)
	if err != nil {
		log.Panicln(err)
		return err
	}
	log.Println(dbrsp.RowsAffected())
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
