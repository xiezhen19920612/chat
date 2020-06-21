package main

import (
	"bufio"
	"bytes"
	"sody.com/chat/mylog"
	"sody.com/chat/myproto"
)

func main() {
	bys, err := myproto.Encode("sodyxie hahaha")
	if err != nil {
		mylog.GetLogger().Printf("encode failed, err:%v\n", err)
		return
	}
	mylog.GetLogger().Println("encode success")

	bysReader := bytes.NewBuffer(bys)
	b := bufio.NewReader(bysReader)
	content, err := myproto.Decode(b)
	if err != nil {
		mylog.GetLogger().Printf("decode failed, err:%v\n", err)
		return
	}

	mylog.GetLogger().Println(content)

}
