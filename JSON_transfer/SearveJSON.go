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
		//fmt.Println(request.URL.Query())

		A, existA :=request.URL.Query()["A"]
		if existA{
			values.FirstOperand,_=strconv.Atoi(A[0])
		}else{
			http.Error(writer,"Opearand A not found",http.StatusBadRequest)
			return
		}
		B, existB :=request.URL.Query()["B"]
		if existB{
			values.SecondOperand,_=strconv.Atoi(B[0])
		}else{
			http.Error(writer,"Opearand B not found",http.StatusBadRequest)
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
		//fmt.Println("POST method called")
		defer request.Body.Close()
		decoder := json.NewDecoder(request.Body)
		var values operands
		err := decoder.Decode(&values)
		if err!=nil{
			fmt.Println("From POST ",err)
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

func main()  {
	fmt.Println("Server is running.....")
	http.HandleFunc("/",handler)
	http.ListenAndServe(":9000",nil)
}