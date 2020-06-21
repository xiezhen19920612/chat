package myclient

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sody.com/chat/config"
	"sody.com/chat/mylog"
	"strings"
)

func StartClient() {
	endpoint := config.GetEndPoint()
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		mylog.GetLogger().Printf("dial failed, err:%v\n", err)
		return
	}
	defer conn.Close()
	mylog.GetLogger().Printf("dial success at endpoint:%s\n", endpoint)

	//监听命令行输入
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("input:")
		content, err := r.ReadString('\n')
		if err != nil {
			mylog.GetLogger().Printf("read failed, err:%v\n", err)
			break
		}
		//quit signal
		if strings.Trim(content, "\r\n") == "Q" {
			mylog.GetLogger().Println("bye")
			break
		}

		//send to server
		conn.Write([]byte(content))
		mylog.GetLogger().Printf("send to server success, content:%s", content)
	}
}
