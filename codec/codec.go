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

//LvCodec length + value
type LvCodec struct {
}

func (l LvCodec) Encode(msg msg.IMsg) ([]byte, error) {
	data := msg.GetData()
	buf := bytes.NewBuffer(make([]byte, 0, len(data)+1))
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

func (l LvCodec) Decode(reader io.Reader) (msg.IMsg, error) {
	var (
		length uint8
	)
	if err := binary.Read(reader, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	data := make([]byte, length)
	if length > 0 {
		if _, err := reader.Read(data); err != nil {
			return nil, err
		}
	}
	return msg.NewMsg(data), nil
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
	if err := binary.Read(reader, binary.BigEndian, &tag); err != nil {
		return nil, err
	}
	if err := binary.Read(reader, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	data := make([]byte, length)
	if length > 0 {
		if _, err := reader.Read(data); err != nil {
			return nil, err
		}
	}
	return msg.NewMsg(data).SetTag(msg.Tag(tag)), nil
}
