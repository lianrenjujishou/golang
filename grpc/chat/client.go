package main
import "google.golang.org/grpc"
import pb "52fighting.com/user/chat/auto_generated"
import "os"
import "context"
import "time"
import "log"
func main() {
  conn, err := grpc.Dial("localhost:60000", grpc.WithInsecure(), grpc.WithBlock())  //客户端向服务端播号,得到连接实例
  if err != nil {
    log.Fatalf("Failed to connect: %v", err)
  }
  defer conn.Close()  //关闭连接,释放资源
  c := pb.NewChatClient(conn) //新建客户端实例
  ctx, cancel := context.WithTimeout(context.Background(), time.Second) //上下文
  defer cancel()
  r, err := c.SayHello(ctx, &pb.X{Name: os.Args[1]}) //调用方法
  log.Printf("Greeting: %s", r.GetMessage())
}
//1.NewChatClient和NewChatServer在哪个文件里?
//2.被调用的函数肯定都是在_rpc.pb.go里.用到的类型,以及对类型中字段的处理函数都是在pb.go里.
