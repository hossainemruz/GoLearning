package main

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
)

type result struct{
	Sum int
	Sub int
	Mul int
	Div float64
}
type operands struct {
	FirstOperand int
	SecondOperand int
}
type tests struct {
	F int
	S int
}

func calculate(values operands)(result)  {
	var response result
	response.Sum=values.FirstOperand+values.SecondOperand
	response.Sub=values.FirstOperand-values.SecondOperand
	response.Mul=values.FirstOperand*values.SecondOperand
	response.Div=float64(values.FirstOperand)/float64(values.SecondOperand)
	return  response

}
func handler(writer http.ResponseWriter, request *http.Request)  {
	if request.Method == "GET"{
		var values operands
		var errx error
		fmt.Println("GET request found")

		A, existA :=request.URL.Query()["FirstOperand"]
		if existA{
			values.FirstOperand,errx=strconv.Atoi(A[0])
			if errx!=nil{
				http.Error(writer,"Operands should be integer",http.StatusBadRequest)
				return
			}
		}else{
			http.Error(writer,"Opearand FirstOperand not found",http.StatusBadRequest)
			return
		}
		B, existB :=request.URL.Query()["SecondOperand"]
		if existB{
			values.SecondOperand,errx=strconv.Atoi(B[0])
			if errx!=nil{
				http.Error(writer,"Operands should be integer",http.StatusBadRequest)
				return
			}
		}else{
			http.Error(writer,"Opearand SecondOperand not found",http.StatusBadRequest)
			return
		}
		response:=calculate(values)
		responseJSON, err:= json.MarshalIndent(response,""," ")
		if err!=nil{
			http.Error(writer,"Conversion Error",http.StatusInternalServerError)
		}else{
			fmt.Fprintln(writer, string(responseJSON))
		}

	}
	if request.Method == "POST"{
		fmt.Println("POST request found")
		defer request.Body.Close()
		decoder := json.NewDecoder(request.Body)
		var values operands
		err := decoder.Decode(&values)
		if err!=nil{
			http.Error(writer,"Wrong JSON",http.StatusBadRequest)
		}else {
			response:=calculate(values)
			responseJSON, err:= json.MarshalIndent(response,""," ")
			if err!=nil{
				http.Error(writer,"Conversion Error",http.StatusInternalServerError)
			}else{
				fmt.Fprintln(writer, string(responseJSON))
			}
		}
	}
}
func StartServer(url string) {
	fmt.Println("Server is starting......")
}
func main()  {
	fmt.Println("Server is running.....")
	http.HandleFunc("/",handler)
	http.ListenAndServe(":9000",nil)
}