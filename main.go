package main

import (
	"fmt"
	"net"
)

////////////////////////////////////////////////////////
//
//错误检查
//
////////////////////////////////////////////////////////
func checkError(err error, info string) (res bool) {

	if err != nil {
		fmt.Println(info + "  " + err.Error())
		return false
	}
	return true
}

////////////////////////////////////////////////////////
//
//服务器端接收数据线程
//参数：
//      数据连接 conn
//      通讯通道 messages
//
////////////////////////////////////////////////////////
func Handler(conn net.Conn, messages chan string, clientip chan string) {

	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for {
		lenght, err := conn.Read(buf)
		if checkError(err, "Connection") == false {
			conn.Close()
			break
		}
		if lenght > 0 {
			buf[lenght] = 0
		}
		//fmt.Println("Rec[",conn.RemoteAddr().String(),"] Say :" ,string(buf[0:lenght]))
		reciveStr := string(buf[0:lenght])
		messages <- reciveStr
		ip := conn.RemoteAddr().String()
		clientip <- ip

	}

}

////////////////////////////////////////////////////////
//
//服务器发送数据的线程
//
//参数
//      连接字典 conns
//      数据通道 messages
//
////////////////////////////////////////////////////////
func echoHandler(conns *map[string]net.Conn, messages chan string, clientip chan string) {

	for {
		msg := <-messages
		clientip := <-clientip
		fmt.Println("接收到的数据: ", msg)
		for key, value := range *conns {
			if clientip == key {
				switch msg {
				case "cpu":
					fmt.Println("send data to", "ip", key)
					sendmsg := cpuStat()
					_, err := value.Write([]byte(sendmsg))
					if err != nil {
						fmt.Println(err.Error())
						delete(*conns, key)
					}
				case "mem":
					fmt.Println("send data to", "ip", key)
					sendmsg := cpuStat()
					_, err := value.Write([]byte(sendmsg))
					if err != nil {
						fmt.Println(err.Error())
						delete(*conns, key)
					}
				case "io":
					fmt.Println("send data to", "ip", key)
					sendmsg := cpuStat()
					_, err := value.Write([]byte(sendmsg))
					if err != nil {
						fmt.Println(err.Error())
						delete(*conns, key)
					}
				}

			}
		}
	}

}

////////////////////////////////////////////////////////
//
//启动服务器
//参数
//  端口 port
//
////////////////////////////////////////////////////////
func StartServer(port string) {
	service := ":" + port //strconv.Itoa(port);
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, "ResolveTCPAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")
	//conns := make(map[string]net.Conn)
	//messages := make(chan string, 10)
	//clientip := make(chan string, 10)

	// //启动服务器广播线程
	// go echoHandler(&conns, messages, clientip)

	for {
		fmt.Println("Listening ...")
		conn, err := l.Accept()
		checkError(err, "Accept")
		//conns[conn.RemoteAddr().String()] = conn
		//启动一个新线程
		buf := make([]byte, 1024)
		lenght, err := conn.Read(buf)
		if checkError(err, "Connection") == false {
			conn.Close()
			break
		}
		if lenght > 0 {
			buf[lenght] = 0
		}
		//fmt.Println("Rec[",conn.RemoteAddr().String(),"] Say :" ,string(buf[0:lenght]))
		reciveStr := string(buf[0:lenght])
		// messages <- reciveStr
		// ip := conn.RemoteAddr().String()
		// clientip <- ip

		switch reciveStr {
		case "cpu":

			sendmsg := cpuStat()
			_, err := conn.Write([]byte(sendmsg))
			if err != nil {
				fmt.Println(err.Error())

			}
		case "mem":

			sendmsg := cpuStat()
			_, err := conn.Write([]byte(sendmsg))
			if err != nil {
				fmt.Println(err.Error())

			}
		case "io":

			sendmsg := cpuStat()
			_, err := conn.Write([]byte(sendmsg))
			if err != nil {
				fmt.Println(err.Error())

			}
		}

	}
}

func main() {

	StartServer("6061")

}
