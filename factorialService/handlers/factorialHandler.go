package handlers

import (
	factorialService "factorialService/proto"
	"context"
)
type FactorialServiceServer struct{
}

func (f *FactorialServiceServer) Factorial(ctx context.Context,in *factorialService.FactRequest)(*factorialService.FactResponse,error){
   num := in.Num
   var  Factorialx int64 
   var  index int64 
   Factorialx =1

   for index = 1; index <= num; index++ {
	Factorialx=Factorialx*index
}
   return &factorialService.FactResponse {Result:Factorialx},nil
}