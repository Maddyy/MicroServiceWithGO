package main

import (
	"google.golang.org/grpc"
	"fmt"
	factorialService "factorialService/proto"
	"context"
)
func main(){
 
	var conn *grpc.ClientConn
	conn,err :=grpc.Dial(":1111",grpc.WithInsecure())
	if err!=nil{
		fmt.Println(err)
	}

   defer conn.Close()

   factServ := factorialService.NewFactorialServiceClient(conn)
   result,err2 := factServ.Factorial(context.Background(),&factorialService.FactRequest{Num:4})

   if err2 != nil {
	fmt.Println(err2)
   } else {
	fmt.Println("Result :",result.Result)
   }
}