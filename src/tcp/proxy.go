package tcp

import (
	"fmt"
	"net"
	"strconv"
)

func ClientHandleError(err error, when string) {
	if err != nil {
		fmt.Printf("连接接目标服务器异常: %s,when: %s\n", err, when)
	}

}

func ConnectToTarget(addr string, port uint16) net.Conn {
	targetHost := addr + ":" + strconv.Itoa(int(port))

	conn, err := net.Dial("tcp", targetHost)
	ClientHandleError(err, "client conn error")
	return conn
}

func ReceiveFromTarget(targetConn net.Conn, clientConn net.Conn) {
	if targetConn == nil {
		fmt.Println("conn is nil...............")
		return
	}
	buffer := make([]byte, 1024)

	for {
		n, err := targetConn.Read(buffer)
		if err != nil {
			targetConn.Close()
			clientConn.Close()

			fmt.Printf("连接接目标服务器异常: %s\n", err)
			fmt.Println("断开与目标服务器链接......\n")
			break
		}
		if n == 0 {
			break
		}

		backData := buffer[0:n]
		clientConn.Write(backData)
	}
}

func ForwardTcp(targetConn net.Conn, data []byte) {
	if targetConn == nil {
		fmt.Println("conn is nil...............")
		return
	}
	if data == nil {
		return
	}
	if len(data) == 0 {
		return
	}

	fmt.Printf("发送数据 -----> 目标服务器: %d\n\n", data)
	targetConn.Write(data)
}