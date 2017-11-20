package main

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"net/http"
	"bytes"
)
func main()  {
  ymlInput,err:=ioutil.ReadFile("input.yml")
  if err!=nil{
  	fmt.Println(err)
  }else{
  	jsonInput,err:=yaml.YAMLToJSON(ymlInput)
  	if err!=nil{
  		fmt.Println(err)
	}else{
		url:="http://127.0.0.1:9000"
		request, err:=http.NewRequest("POST",url,bytes.NewBuffer(jsonInput))
		if err!=nil{
			fmt.Println(err)
		}else{
			client:=http.Client{}
			response,requestError:=client.Do(request)
			if requestError!=nil{
				fmt.Println(requestError)
			}else {
				defer response.Body.Close()
				jsonOutput,_:=ioutil.ReadAll(response.Body)
				//fmt.Println(string(jsonOutput))

				ymlOutput,conversionError:=yaml.JSONToYAML(jsonOutput)
				if conversionError!=nil{
					fmt.Println(conversionError)
				}else {
					ioutil.WriteFile("output.yml",ymlOutput,0644)
					fmt.Println("Output sucessfully written on output.yml")
				}
			}

		}
	}
  }
}