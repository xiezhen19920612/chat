package myserver

import (
	"bufio"
	"io"
	"net"
	"sody.com/chat/config"
	"sody.com/chat/mylog"
)

func StartServer() {
	endpoint := config.GetEndPoint()
	ss, err := net.Listen("tcp", endpoint)
	if err != nil {
		mylog.GetLogger().Printf("listen failed, err:%v\n", err)
		return
	}
	defer ss.Close()
	mylog.GetLogger().Printf("listen success at endpoint %s\n", endpoint)

	//开启go程处理
	for {
		conn, err := ss.Accept()
		if err != nil {
			mylog.GetLogger().Printf("accept failed, err:%v\n", err)
			break
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	//阻塞读取
	for {
		r := bufio.NewReader(conn)
		msg, err := r.ReadString('\n')
		if err == io.EOF {
			mylog.GetLogger().Printf("conn closed\n")
			break
		}
		if err != nil {
			mylog.GetLogger().Printf("read failed, err:%v\n", err)
			break
		}

		mylog.GetLogger().Printf("from client:%s", msg)
	}
}
