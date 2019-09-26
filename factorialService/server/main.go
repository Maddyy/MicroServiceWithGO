package main
import (
	"net"
	factorialService "factorialService/proto"
	handlers "factorialService/handlers"
  "fmt"
  "google.golang.org/grpc"
)
func main(){
	lis ,err := net.Listen("tcp",":1111")
  if err!=nil{
	  fmt.Print(err)
  }
  defer lis.Close()
  factServ := handlers.FactorialServiceServer{}
  grpcServer := grpc.NewServer()
  factorialService.RegisterFactorialServiceServer(grpcServer,&factServ)

  if err :=grpcServer.Serve(lis);err !=nil{
	  fmt.Println(err)
  }
}