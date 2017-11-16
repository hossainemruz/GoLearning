package GoLearning

import (
	"net/http"
	"io"	
)

func showPage(writer http.ResponseWriter, request *http.Request)  {
	io.WriteString(writer,`<DOCTYPE html>
		<head>
			<title>Gretting</title>
		</head>
		<body>
			<h1>Welcome to Go. Lets Go...</h1>
		</body>
	</html>`,)
}
func main()  {
	http.HandleFunc("/",showPage)
	http.ListenAndServe(":9000",nil)
}