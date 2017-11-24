package test_test

import (

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type Operands struct {
	FirstOperand int
	SecondOperand int

}
type rspn struct {
	Sum int
	Sub int
	Mul int
	Div float64
}
var Param Operands
var expectedResult rspn
var url string
var _ = Describe("Server", func() {
	BeforeEach(func() {
		url="http://127.0.0.1:9000"
		Param.FirstOperand=73
		Param.SecondOperand=38

		expectedResult.Sum=111
		expectedResult.Sub=35
		expectedResult.Mul=2774
		expectedResult.Div=1.9210526315789473
	})

	Describe("Get Request test", func() {
		Context("With valid two parameter", func() {
			It("Should be successfull request with Status Code:200", func() {
				resp,err:=http.Get(url+"?FirstOperand="+strconv.Itoa(Param.FirstOperand)+"&SecondOperand="+strconv.Itoa(Param.SecondOperand))
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(200))
				defer resp.Body.Close()
				responseJSON,err2:=ioutil.ReadAll(resp.Body)
				Expect(err2).NotTo(HaveOccurred())
				var result rspn
				json.Unmarshal(responseJSON,&result)
				Expect(expectedResult).To(Equal(result))

			})
		})
		Context("With only FirstOperand", func() {
			It("Should return bad request with Status Code: 400", func() {
				resp,err:=http.Get(url+"?FirstOperand="+strconv.Itoa(Param.FirstOperand))
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(400))
			})
		})
		Context("With only SecondOperand", func() {
			It("Should return bad request with Status Code: 400", func() {
				resp,err:=http.Get(url+"?SecondOperand="+strconv.Itoa(Param.SecondOperand))
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(400))
			})
		})
		Context("With No Operands", func() {
			It("Should return bad request with Status Code: 400", func() {
				resp,err:=http.Get(url)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(400))
			})
		})
		Context("With wrong rquest format", func() {
			It("Should return bad request with Status Code: 400", func() {
				resp,err:=http.Get(url+"?A="+strconv.Itoa(Param.FirstOperand)+"&B="+strconv.Itoa(Param.SecondOperand))
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(400))
			})
		})
		Context("With wrong rquest with string as operand", func() {
			It("Should return bad request with Status Code: 400", func() {
				resp,err:=http.Get(url+"?FirstOperand=aa&SecondOperand=bb")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(400))
			})
		})
		Context("With wrong rquest with 0 as second operand", func() {
			It("Should return \"Internal Server Error\" with Status Code: 500", func() {
				resp,err:=http.Get(url+"?FirstOperand=0&SecondOperand=0")
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(500))
			})
		})

	})
	Describe("POST request test", func() {
		BeforeEach(func() {
			url="http://127.0.0.1:9000"
			Param.FirstOperand=73
			Param.SecondOperand=38

			expectedResult.Sum=111
			expectedResult.Sub=35
			expectedResult.Mul=2774
			expectedResult.Div=1.9210526315789473
		})
		Context("With valid two parameter", func() {
			It("Should be successfull request with Status Code:200", func() {
				client:=http.Client{}
				jsonInput,err:=json.MarshalIndent(Param,""," ")
				Expect(err).NotTo(HaveOccurred())
				request,err:=http.NewRequest("POST",url,bytes.NewBuffer(jsonInput))
				Expect(err).NotTo(HaveOccurred())
				resp,err:=client.Do(request)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(200))

				defer resp.Body.Close()
				responseJSON,err2:=ioutil.ReadAll(resp.Body)
				Expect(err2).NotTo(HaveOccurred())
				var result rspn
				json.Unmarshal(responseJSON,&result)
				Expect(expectedResult).To(Equal(result))

			})
		})
		Context("With second parameter is 0", func() {
			It("Should return \"Internal Server Error\" with Status Code:500", func() {
				client:=http.Client{}
				Param.SecondOperand=0;
				jsonInput,err:=json.MarshalIndent(Param,""," ")
				Expect(err).NotTo(HaveOccurred())
				request,err:=http.NewRequest("POST",url,bytes.NewBuffer(jsonInput))
				Expect(err).NotTo(HaveOccurred())
				resp,err:=client.Do(request)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(500))
			})
		})
	})
})
