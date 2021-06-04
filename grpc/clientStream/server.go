package main
import (
  pb "sc.com/user/hello/scpb"
  "net"
  "google.golang.org/grpc"
  "log"
  "io"
)

type sc struct {
  pb.UnimplementedScServer
}
func (s *sc) Greeting(stream pb.Sc_GreetingServer) error {
  for {
    result, err := stream.Recv()
    if err == io.EOF {
      return stream.SendAndClose(&pb.Y{Message: "666"})  //服务端有个SendAndClose(*),而客户端也有个SendAndClose()但不带参数,问题是客户端调的这个方法我在protoc生成的那两个文件中找不到,只能找到服务端的,通过dlv调试找到了第55行就是,看一下代码,了解继承相关的内容就能看明白了 ;  另外,server-side中有Send,而client-side有SendAndClose
    }
    if err != nil {
      log.Fatalf("Failed: %v", err)
    }
    log.Printf("Received: %v", result)
  }
}
var opts []grpc.ServerOption
func main() {
  lis, err := net.Listen("tcp", "localhost:60000")
  if err != nil {
    log.Fatalf("listen failed: %v", err)
  }
  server := grpc.NewServer(opts...)
  pb.RegisterScServer(server, &sc{})
  server.Serve(lis)
}

