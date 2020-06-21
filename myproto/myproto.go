package myproto

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"sody.com/chat/mylog"
)

/**
编码方法：前四个字节表示内容长度，后面是内容
*/
func Encode(content string) ([]byte, error) {
	bs := []byte(content)
	bsLen := int32(len(bs))
	//byte buffer
	pkg := new(bytes.Buffer)

	//长度
	err := binary.Write(pkg, binary.LittleEndian, bsLen)
	if err != nil {
		mylog.GetLogger().Printf("binary write failed, err:%v\n", err)
		return nil, err
	}

	//真正要传输的内容
	err = binary.Write(pkg, binary.LittleEndian, bs)
	if err != nil {
		mylog.GetLogger().Printf("binary write failed, err:%v\n", err)
		return nil, err
	}

	return pkg.Bytes(), nil
}

func Decode(r *bufio.Reader) (string, error) {
	//先读前四个字节，长度
	lenBs, err := r.Peek(4)
	if err != nil {
		mylog.GetLogger().Printf("read length bytes failed, err:%v\n", err)
		return "", err
	}
	pkg := bytes.NewBuffer(lenBs)
	var contentLen int32
	err = binary.Read(pkg, binary.LittleEndian, &contentLen)
	if err != nil {
		mylog.GetLogger().Printf("read length failed, err:%v\n", err)
		return "", err
	}

	if int32(r.Buffered()) < contentLen+4 {
		mylog.GetLogger().Printf("package length invalid, need length:%d actual length:%d\n", contentLen+4, r.Buffered())
		tmpErr := errors.New("package length invalid")
		return "", tmpErr
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+contentLen))
	_, err = r.Read(pack)
	if err != nil {
		mylog.GetLogger().Printf("read content failed, err:%v\n", err)
		return "", err
	}
	return string(pack[4:]), nil

}
