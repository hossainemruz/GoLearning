package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"encoding/base64"
	"strings"
)

type result struct {
	Sum int
	Sub int
	Mul int
	Div float64
}
type operands struct {
	FirstOperand  int
	SecondOperand int
}
type tests struct {
	F int
	S int
}

func calculate(values operands) result {
	var response result
	response.Sum = values.FirstOperand + values.SecondOperand
	response.Sub = values.FirstOperand - values.SecondOperand
	response.Mul = values.FirstOperand * values.SecondOperand
	response.Div = float64(values.FirstOperand) / float64(values.SecondOperand)
	return response

}
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Recieved new request....")
	// fmt.Println(request)
	if isAuthorised(writer, request) == false {
		writer.Header().Add("WWW-Authenticate", `Basic realm="Authorization Required"`)
		http.Error(writer, "401 Unauthorized", http.StatusUnauthorized)
	} else {
		if request.Method == "GET" {
			var values operands
			//fmt.Println(request.URL.Query())

			A, existA := request.URL.Query()["FirstOperand"]
			if existA {

				valA, errA := strconv.Atoi(A[0])
				if errA != nil {
					http.Error(writer, "FirstOperand not found", http.StatusBadRequest)
					return
				}
				values.FirstOperand = valA
			} else {
				http.Error(writer, "FirstOperand not found", http.StatusBadRequest)
				return
			}
			B, existB := request.URL.Query()["SecondOperand"]
			if existB {
				valB, errB := strconv.Atoi(B[0])
				if errB != nil {
					http.Error(writer, "SecondOperand not found", http.StatusBadRequest)
					return
				}
				values.SecondOperand = valB
			} else {
				http.Error(writer, "SecondOperand not found", http.StatusBadRequest)
				return
			}
			response := calculate(values)
			fmt.Println(response)
			responseJSON, err := json.MarshalIndent(response, "", " ")
			if err != nil {
				http.Error(writer, "Conversion Error", http.StatusInternalServerError)
			} else {
				fmt.Fprintln(writer, string(responseJSON))
			}

		}
		if request.Method == "POST" {
			//fmt.Println("POST method called")
			defer request.Body.Close()
			decoder := json.NewDecoder(request.Body)
			var values operands
			err := decoder.Decode(&values)
			if err != nil {
				fmt.Println(err)
			} else {
				response := calculate(values)
				fmt.Println(response)
				responseJSON, err := json.MarshalIndent(response, "", " ")
				if err != nil {
					http.Error(writer, "Conversion Error", http.StatusInternalServerError)
				} else {
					fmt.Fprintln(writer, string(responseJSON))
				}
			}
		}
	}

}

func isAuthorised(writer http.ResponseWriter, request *http.Request) bool {
	authorizationHeader := strings.SplitN(request.Header.Get("Authorization"), " ", 2)
	 fmt.Println(authorizationHeader)
	if len(authorizationHeader) != 2 {
		return false
	}
	baseCredential, err := base64.StdEncoding.DecodeString(authorizationHeader[1])
	if err != nil {
		return false
	} else {
		credential := strings.SplitN(string(baseCredential), ":", 2)
		 fmt.Println(credential)
		if credential[0] == "emruz" && credential[1] == "1234" {
			return true
		} else {
			return false
		}
	}
	return false
}

func main() {
	fmt.Println("Server is running.....")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
