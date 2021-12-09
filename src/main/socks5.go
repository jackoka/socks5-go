package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"socks5-go/src/tcp"
	"strconv"
	"strings"
)

func HandleConn(clientConn net.Conn) {
	defer clientConn.Close()
	addr := clientConn.RemoteAddr().String()
	fmt.Println(addr, " conncet sucessful")

	buf := make([]byte, 2048)

	var targetConn net.Conn

	var clientBuf []byte

	for {
		//读取用户数据
		len, err := clientConn.Read(buf)
		if err != nil {
			fmt.Println("断开与客户机的链接................")
			fmt.Println("err = ", err)
			return
		}

		clientBuf = buf[:len]
		fmt.Printf("client----> length: %d,data: %d, [%s]\n", len, clientBuf, addr)

		if len == 3 {
			// 握手，核对版本信息
			/*VER := clientBuf[0]
			NMETHODS := clientBuf[1]
			METHODS  := clientBuf[2]
			fmt.Printf("VER: [%d] \nNMETHODS: %d \nMETHODS:%d\n", VER,NMETHODS,METHODS)*/

			// 参数1：5:版本
			// 参数2:
			//00:无需身份验证
			//01:GSSAPI
			//02:用户名/密码
			//03:至X'7F'IANA已分配
			//80:到X'FE'保留给私有方法
			//FF:无可接受的方法

			message := []byte{5, 0}
			clientConn.Write(message)
			fmt.Printf("<------%d\n\n", message)

		} else if len > 4 {
			//建立连接 或者 发送数据
			//协议版本
			VER := clientBuf[0]

			//CONNECT X'01'
			//BIND X'02'
			//UDP ASSOCIATE X'03'
			CMD := clientBuf[1]

			//固定 0
			RSV := clientBuf[2]

			//ATYP   address type of following address
			//IP V4 address: X'01'
			//DOMAINNAME: X'03'
			//IP V6 address: X'04'
			ATYP := clientBuf[3]

			if VER == 5 && CMD == 1 && RSV == 0 {
				//socks5

				if ATYP == 1 {
					//ip v4
				} else if ATYP == 3 {
					//域名
					//4: VER(1位) +  CMD(1位) + RSV(1位) + ATYP(1位)
					//2: DST_PORT(2位)
					domainLen := len - 2
					DST_ADDR := string(clientBuf[5:domainLen])
					DST_PORT := binary.BigEndian.Uint16(clientBuf[len-2 : len])

					targetConn = tcp.ConnectToTarget(DST_ADDR, DST_PORT)

					message := clientBuf[:len]
					message[1] = 0
					clientConn.Write(message)
					fmt.Printf("<------%d\n\n", message)

					go tcp.ReceiveFromTarget(targetConn, clientConn)
				} else if ATYP == 4 {
					// ip v6
				}
			} else {
				tcp.ForwardTcp(targetConn, clientBuf)
			}

		}

	}
}

func main() {
	port := 50000

	if len(os.Args) > 1 {
		param1 := os.Args[1]
		fmt.Printf("参数: %s\n", param1)

		paramArray := strings.Split(param1, "=")
		port, _ = strconv.Atoi(paramArray[1])
	}

	targetHost := "0.0.0.0" + ":" + strconv.Itoa(port)

	socket, err := net.Listen("tcp", targetHost)
	if err != nil {
		fmt.Println("开启监听失败,错误原因: ", err)
		return
	}
	defer socket.Close()
	fmt.Printf("开启监听: %s\n", targetHost)

	for {
		conn, err := socket.Accept()
		if err != nil {
			fmt.Println("建立链接失败,错误原因: ", err)
			return
		}
		go HandleConn(conn)
	}
}