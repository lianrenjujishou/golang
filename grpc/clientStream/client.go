package main

import (
  "google.golang.org/grpc"
  "log"
  "context"
  pb "sc.com/user/hello/scpb"
)

func main() {
  conn, err := grpc.Dial("localhost:60000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Failed to connect: %v", err)
  }
  defer conn.Close()

  c := pb.NewScClient(conn)
  stream, err := c.Greeting(context.Background())
  if err != nil {
    log.Fatalf("failed: %v", err)
  }
  for _,point := range []*pb.X{{Name: "杨毅"},{Name: "赵敏"}}{
    if err := stream.Send(point); err != nil {
      log.Fatalf("%v.Send(%v) = %v", stream, point, err)
    }
  }
  reply, err := stream.CloseAndRecv()
  log.Printf("summary: %v,%v", reply, err)
}
