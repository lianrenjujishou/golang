package main

import pb "52fighting.com/user/chat/auto_generated"
import "net"
import "google.golang.org/grpc"
import "context"
import "log"

type server struct {
  pb.UnimplementedChatServer //这个类型是自动生成的, 并且已经实现了ChatServer中定义的方法,目的应该是我们在服务端中未实现ChatServer中定义方法的时候,自动调用该类型下的相应方法.
}
func (s *server) SayHello(ctx context.Context, in *pb.X) (*pb.Y, error) {
  log.Printf("Received: %s", in.GetName())
  return &pb.Y{Message: "Hello" + in.GetName()}, nil
}

func main() {
  lis, err := net.Listen("tcp", ":60000")  //监听
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }
  s := grpc.NewServer()  //新建实例
  pb.RegisterChatServer(s, &server{}) //将服务实例与实现接口的类型实例绑定
  if err := s.Serve(lis); err != nil { //循环接收数据
    log.Fatalf("Failed to serve: %v", err)
  }
}
