package msg

type Tag uint8

type IMsg interface {
	GetData() []byte
	GetTag() Tag
}

type Msg struct {
	data []byte
	tag  Tag
}

func NewMsg(data []byte) *Msg {
	return &Msg{data: data}
}

func (m Msg) GetData() []byte {
	return m.data
}

func (m *Msg) SetTag(t Tag) *Msg {
	m.tag = t
	return m
}

func (m Msg) GetTag() Tag {
	return m.tag
}
