package main
import "google.golang.org/grpc"
import "net"
import pb "work.out/user/hello/sspb"
import "log"

type transport struct{
  pb.UnimplementedTransportServer
  savedFeatures []*pb.Y
}

func (t *transport) ListFeatures(x *pb.X, stream pb.Transport_ListFeaturesServer) error {
  log.Println(x)
  for _, feature := range t.savedFeatures{
    if err := stream.Send(feature); err != nil {
      return err
    }
  }
  return nil
} 
 var opts []grpc.ServerOption
func main() {
  lis, err := net.Listen("tcp","localhost:60000")
  if err !=nil {
    log.Fatalf("listen failed: %v", err)
  }
  grpcServer := grpc.NewServer(opts...)
  t := transport{savedFeatures: []*pb.Y{{Message: "A"},{Message: "B"}}}
  pb.RegisterTransportServer(grpcServer, &t)
  grpcServer.Serve(lis)
}
