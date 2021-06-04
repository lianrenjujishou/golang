package main

import (
  pb "bi.com/user/hello/bipb"
  "google.golang.org/grpc"
  "io"
  "log"
  "context"
)

func main() {
  conn, err := grpc.Dial("localhost:60000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Dial Failed: %v",err)
  }
  defer conn.Close()
  c := pb.NewActionClient(conn)

  stream, err := c.Hit(context.Background())
  waitc := make(chan struct{})
  go func() {
    for {
      in, err := stream.Recv()
      if err == io.EOF {
        close(waitc)
	return
      }
      log.Printf("Received: %v", in.Message)
    }
  }()
  for _, note := range []*pb.X{{Name: "A"}, {Name: "B"}, {Name: "C"}}{
    if err := stream.Send(note); err != nil {
      log.Fatalf("Failed to send a note: %v", err)
    }
  }
  stream.CloseSend()
  <-waitc
}


//因为双向都可以读写,所以就要考虑使用协程了,将Recv()相关语句放里面, 这样不会阻断代码的执行
