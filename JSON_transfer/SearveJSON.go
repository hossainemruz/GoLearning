package main

import (
	"fmt"
	"net/http"
)

type result struct{
	Sum int
	Sub int
	Mul int
	div float32
}
type operands struct {
	firstOperand int
	secondOperand int
}
func handler(writer http.ResponseWriter, request *http.Request)  {
	if request.Method == "GET"{
		var values operands
		A, exist :=request.URL.Query()["A"]
		if exist{
			values.firstOperand=A
		}else{
			http.Error(writer,"Opearand not found",)
		}
	}
}

func main()  {
	fmt.Println("Server is running.....")
	http.HandleFunc("/",handler)
	http.ListenAndServe(":9000",nil)
}