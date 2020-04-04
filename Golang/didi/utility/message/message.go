package message

const (
	MSGTYPE   = 4 //消息头部4byte，表示类型
	MSGLENGTH = 4 //消息中间8byte，表示消息内容长度
	MSGMAXLEN = 2048

	MAXPKGSIZE = MSGTYPE + MSGLENGTH + MSGMAXLEN
)

type Message interface {
	Serialize() ([]byte, error)
	MessageNumber() int32
}
