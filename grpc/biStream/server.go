package main

import (
  pb "bi.com/user/hello/bipb"
  "net"
  "io"
  "log"
  "google.golang.org/grpc"
)
type bi struct {
  pb.UnimplementedActionServer
}
func (s *bi) Hit(stream pb.Action_HitServer) error {
  for {
    in, err := stream.Recv()
    if err ==io.EOF {
      return nil
    }
    if err != nil {
      return err
    }
    if err := stream.Send(&pb.Y{Message: in.GetName()+"~"}); err != nil {
      return err
    }
  }
}
var opts []grpc.ServerOption
func main() {
  lis, err := net.Listen("tcp", "localhost:60000")
  if err != nil {
    log.Fatalf("Listen failed: %v",err)
  }
  server := grpc.NewServer(opts...)
  pb.RegisterActionServer(server, &bi{})
  server.Serve(lis)
}
