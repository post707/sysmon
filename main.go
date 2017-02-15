package main

import (
	"fmt"
	"net"
)

////////////////////////////////////////////////////////
//
//启动服务器
//参数
//端口 port
//
////////////////////////////////////////////////////////
func StartServer(port string) {
	service := ":" + port //strconv.Itoa(port);
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, "ResolveTCPAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")

	//读取配置文件
	fileConfig := new(Config)
	fileConfig.InitConfig("sysmon.conf")
	switch fileConfig.Mymap["type->code"] {
	case "0":
		server(l)
	case "1":
		server(l)
	default:
		server(l)
	}

}

func main() {

	StartServer("6061")

}

////////////////////////////////////////////////////////
//错误检查
////////////////////////////////////////////////////////
func server(l net.Listener) {
	for {
		fmt.Println("Listening ...")
		conn, err := l.Accept()
		checkError(err, "Accept")
		buf := make([]byte, 1024)
		lenght, err := conn.Read(buf)
		if checkError(err, "Connection") == false {
			conn.Close()
			break
		}
		if lenght > 0 {
			buf[lenght] = 0
		}
		reciveStr := string(buf[0:lenght])
		fmt.Println("sendmsg to " + conn.RemoteAddr().String())
		switch reciveStr {
		case "cpu":
			sendmsg := cpuStat()
			_, err := conn.Write([]byte(sendmsg))
			if err != nil {
				fmt.Println(err.Error())
			}
		case "mem":
			sendmsg := memStat()
			_, err := conn.Write([]byte(sendmsg))
			if err != nil {
				fmt.Println(err.Error())
			}
		case "io":
			sendmsg := ioStat()
			_, err := conn.Write([]byte(sendmsg))
			if err != nil {
				fmt.Println(err.Error())
			}
		}

	}
}

////////////////////////////////////////////////////////
//错误检查
////////////////////////////////////////////////////////
func checkError(err error, info string) (res bool) {

	if err != nil {
		fmt.Println(info + "  " + err.Error())
		return false
	}
	return true
}

func permission() {

}
