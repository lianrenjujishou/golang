package main
import pb "work.out/user/hello/sspb"
import "google.golang.org/grpc"
import "os"
import "io"
import "log"
import "context"


var opts []grpc.DialOption
func main() {
   conn, err := grpc.Dial("localhost:60000", grpc.WithInsecure())
   if err != nil {
     log.Fatalf("Failed to Dial: %v", err)
   }
   defer conn.Close()

   c := pb.NewTransportClient(conn) 
   feature, err := c.ListFeatures(context.Background(), &pb.X{Name: os.Args[1]})
   if err != nil {
     log.Fatalf("Receive Failed: %v", err)
   }
   for {
       result, err := feature.Recv() 
     if err == io.EOF {
       break
     }
     if err != nil {
       log.Fatalf("Receive failed : %v", err)
     }
     log.Println(result)
   }
}
//0.重要的一点是,虽说客户端和服务端的函数名称一样, 但两边的signature可能是不同的.
//1.对于server-side stream来说, 客户端只是发一次, 因此直接在客户端调用函数时将请求的信息作为函数参数一次就发过来了,这个通过_gprc中客户端调用函数的格式定义也能看出来; 而服务端这边因为是流式的, 所以要不停的调用stream.Send()响应给客户端.而客户端这边使用stream.Recv()来接收每次发过来的数据. 总结: a.将请求的参数直接放到函数中;b.客户端这边的Recv(); 服务端这边的Send();
//2.对于client-side stream来说, 因为客户端是要不断的发送数据,所以是先和服务端拿到一个通道stream,然后再利用stream.Send()发送数据,最后啥时候不想发了,或是发完了,再使用stream.CloseAndRecv()告诉通道那边的服务端; 而因为客户端会通过通道Send()所以服务端肯定要有相应的stream.Recv()来接收, 并且要实刻检测客户端不发数据时利用stream.SendAndClose(xxxx)将xxxx返回给客户端(如统计信息). 总结:对于cliend-side stream来说, 不仅有1中的Send()和Recv()还有CloseAndRecv()和CloseAndRecv(xxxx).
//server-side stream 和 client-side stream都是通过检测Recv()的返回值为io.EOF判断接收结束的
//stream猜测用的就是golang中的通道知识
