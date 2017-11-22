package main

import (
	"testing"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type rspns struct {
	Sum int
	Sub int
	Mul int
	Div float64
}

var url ="http://127.0.0.1:9000/"

type operand struct {
	FirstOperand int
	SecondOperand int
}
func TestServeJSON(t *testing.T)  {
	t.Run("Test 1",runTest(t,url+"?FirstOperand=50&SecondOperand=5",operand{50,5},"GET",rspns{55,45,250,10},200))
	t.Run("Test 2",runTest(t,url+"?FirstOperand=50",operand{50,5},"GET",rspns{0,0,0,0},400))
	t.Run("Test 3",runTest(t,url,operand{0,0},"GET",rspns{0,0,0,0},400))
	t.Run("Test 4",runTest(t,url,operand{9,3},"POST",rspns{12,6,27,3},200))
	t.Run("Test 5",runTest(t,url,operand{0,0},"POST",rspns{0,0,0,0},500))

}

func runTest(t *testing.T, requestURL string,data operand,testType string, expectedResponse rspns,expectedStatusCode int)  func(t2 *testing.T){
	return func(t *testing.T){
	   if(testType=="GET"){
		 responseData,statusCode:=sendGETRequest(requestURL,t)
		 //fmt.Println("runTest:",responseData)
		 //fmt.Println("StatusCode: ",statusCode)
		 validateResponse(t,responseData,statusCode,expectedResponse,expectedStatusCode)

	   }else if testType=="POST"{
			response,statusCode:=sendPOSTRequest(requestURL,data,t)
			//fmt.Println(response,statusCode)
			validateResponse(t,response,statusCode,expectedResponse,expectedStatusCode)
		}else{
			t.Error("Unknown request type.")
	   }
}
}

func sendGETRequest(request string,t *testing.T) (rspns,int){
  response,err:=http.Get(request)
  if err!=nil{
  	t.Fatal("Request Error.")
  }else{
  	defer response.Body.Close()
  	result,err:=ioutil.ReadAll(response.Body)
  	if err!=nil{
  		t.Fatal("Error: reading from response body")
	}else{

		var ret rspns
		json.Unmarshal(result,&ret)
		return ret,response.StatusCode
	}
  }

  return rspns{0,0,0,0},500
}

func sendPOSTRequest(requestURL string,data operand,t *testing.T)  (rspns,int){

	jsonData,err:=json.MarshalIndent(data,""," ")
	if err!=nil{
		t.Fatal("JSON conversion Error! ",err)
	}else{
		client:=&http.Client{}
		request,err:=http.NewRequest("POST",requestURL,bytes.NewBuffer(jsonData))
		if err!=nil{
			t.Fatal(err)
		}else {
			resp,err2:=client.Do(request)
			if err2!=nil{
				t.Fatal(err2)
			}else{
				defer resp.Body.Close()
				result,err:=ioutil.ReadAll(resp.Body)
				if err!=nil{
					t.Fatal("Error: reading from response body")
				}else{

					var ret rspns
					json.Unmarshal(result,&ret)
					return ret,resp.StatusCode
				}
			}
		}
	}
	return rspns{0,0,0,0},500
}

func validateResponse(t *testing.T,responseData rspns,statusCode int, expectedResponse rspns, expectedStatusCode int)  {
	if expectedStatusCode!=statusCode{
		t.Fatal("Status code mismatch! Expected:",expectedStatusCode," Got: ",statusCode)
	}
	if expectedResponse!=responseData{
		t.Fatal("Response mismatch! Expected: ",expectedResponse,"Found: ",responseData)
	}
}