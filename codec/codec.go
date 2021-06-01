package codec

import (
	"bytes"
	"encoding/binary"
	"github.com/fanjindong/payne/msg"
	"io"
)

type ICodec interface {
	Encoder
	Decoder
}

type Encoder interface {
	Encode(msg msg.IMsg) ([]byte, error)
}

type Decoder interface {
	Decode(reader io.Reader) (msg.IMsg, error)
}

//TlvCodec 标识域（Tag）+长度域（Length）+值域（Value）
type TlvCodec struct {
}

func (c TlvCodec) Encode(msg msg.IMsg) ([]byte, error) {
	data := msg.GetData()
	buf := bytes.NewBuffer(make([]byte, 0, len(data)+1+1))
	if err := binary.Write(buf, binary.BigEndian, uint8(msg.GetTag())); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, uint8(len(data))); err != nil {
		return nil, err
	}
	if len(data) != 0 {
		if _, err := buf.Write(data); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func (c TlvCodec) Decode(reader io.Reader) (msg.IMsg, error) {
	var (
		tag    uint8
		length uint8
	)
	binary.Read(reader, binary.BigEndian, &tag)
	binary.Read(reader, binary.BigEndian, &length)
	data := make([]byte, length)
	if length > 0 {
		if _, err := reader.Read(data); err != nil {
			return nil, err
		}
	}
	return msg.NewMsg(msg.Tag(tag), data), nil
}

//
////LvCodec length + value
//type LvCodec struct {
//}
//
//func (l LvCodec) Encode(data []byte) ([]byte, error) {
//	buf := bytes.NewBuffer(nil)
//	if err := binary.Write(buf, binary.BigEndian, uint8(len(data))); err != nil {
//		return nil, err
//	}
//	if _, err := buf.Write(data); err != nil {
//		return nil, err
//	}
//	return buf.Bytes(), nil
//}
//
//func (l LvCodec) Decode(reader io.Reader) ([]byte, error) {
//	var dataLen uint8
//	binary.Read(reader, binary.BigEndian, &dataLen)
//	data := make([]byte, dataLen)
//	if _, err := reader.Read(data); err != nil {
//		fmt.Println(dataLen, data, err)
//		return nil, err
//	}
//	return data, nil
//}
