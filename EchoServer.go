package GoLearning

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func  server()  {
	fmt.Println("This is server")
	//listen to a port
	ln, err :=net.Listen("tcp",":9999") // net.Listen() return two values connection and err
	if err!=nil{
		fmt.Println(err)
			return
	}

	for{
		//accept connection
		c,err :=ln.Accept()
		if err!=nil{
			fmt.Println(err)
			continue
		}
		//handle connection
		go handleServerConnection(c)
	}

}

func handleServerConnection(c net.Conn)  {
	var msg string
	for i:=0;;i++ {

		err := gob.NewDecoder(c).Decode(&msg) //decode received message
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("recieved ", msg)
		}
	}
	c.Close()
}

func client()  {
	fmt.Println("This is client")
	//connect to server
	c, err := net.Dial("tcp","127.0.0.1:9999")
	if err!=nil{
		fmt.Println(err)
		return
	}
	for i:=1;;i++{
		msg := "hello Emruz "
		//send message
		err = gob.NewEncoder(c).Encode(msg)
		if err!=nil{
			fmt.Println(err)
		}else {
			fmt.Println("send ", msg)
		}
		time.Sleep(time.Second*2)
	}
	c.Close()

}

func main()  {
	go server();
	go client();

	var input string
	fmt.Scanln(&input)
}