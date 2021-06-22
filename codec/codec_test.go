package codec

import (
	"bytes"
	"github.com/fanjindong/payne/msg"
	"reflect"
	"testing"
)

func TestTlvCodec(t *testing.T) {
	type args struct {
		msg msg.IMsg
	}
	tests := []struct {
		name    string
		args    args
		want    msg.IMsg
		wantErr bool
	}{
		{name: "1", args: args{msg: msg.NewMsg([]byte("1")).SetTag(msg.Tag(0))}, want: msg.NewMsg([]byte("1")).SetTag(msg.Tag(0))},
		{name: "2", args: args{msg: msg.NewMsg([]byte("你好")).SetTag(msg.Tag(1))}, want: msg.NewMsg([]byte("你好")).SetTag(msg.Tag(1))},
	}
	codec := TlvCodec{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodeGot, err := codec.Encode(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//t.Log(encodeGot)
			decodeGot, err := codec.Decode(bytes.NewReader(encodeGot))
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(decodeGot, tt.want) {
				t.Errorf("LvCodec() got = %v, want %v", decodeGot, tt.want)
			}
		})
	}
}

//
//func TestLvCodec(t *testing.T) {
//	type args struct {
//		data []byte
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    []byte
//		wantErr bool
//	}{
//		{name: "1", args: args{data: []byte("1")}, want: []byte("1")},
//		{name: "2", args: args{data: []byte("你好")}, want: []byte("你好")},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			l := LvCodec{}
//			got, err := l.Encode(tt.args.data)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			got, err = l.Decode(bytes.NewReader(got))
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("LvCodec() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
