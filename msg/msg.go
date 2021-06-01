package msg

type Tag uint8

type IMsg interface {
	GetTag() Tag
	GetData() []byte
}

type Msg struct {
	tag  Tag
	data []byte
}

func NewMsg(tag Tag, data []byte) *Msg {
	return &Msg{tag: tag, data: data}
}

func (m Msg) GetTag() Tag {
	return m.tag
}

func (m Msg) GetData() []byte {
	return m.data
}
