package msg

const (
	PKG_CRC32_SIZE = 4
)

const (
	PKG_CRC32_BEGIN = 0
	PKG_CRC32_END   = PKG_CRC32_BEGIN + PKG_CRC32_SIZE

	PKG_HEADER_SIZE
)

const (
	MSG_TYPE_SIZE = 1
	MSG_SEQ_SIZE  = 4
	MSG_LEN_SIZE  = 4

	MAX_MESSAGE_SIZE = 10240
)

const (
	MSG_HEADER_BEGIN = 0
	MSG_TYPE_BEGIN
	MSG_TYPE_END = MSG_TYPE_BEGIN + MSG_TYPE_SIZE
	MSG_SEQ_BEGIN
	MSG_SEQ_END = MSG_SEQ_BEGIN + MSG_SEQ_SIZE
	MSG_LEN_BEGIN
	MSG_LEN_END = MSG_LEN_BEGIN + MSG_LEN_SIZE
	MSG_HEADER_END

	MSG_HEADER_SIZE
)

const (
	TYPE_NORMAL = 0x01
	TYPE_FEC    = 0x02
	TYPE_REQ    = 0x03
	TYPE_RESP   = 0x04
	TYPE_ACK    = 0x80
	TYPE_PING   = 0x81
	TYPE_PONG   = 0x82
)

const (
	MSG_STATUS_INIT = 1 << iota
	MSG_STATUS_TRANSMITTED
	MSG_STATUS_ACKED
	MSG_STATUS_LOSS
)

// ack msg index
const (
	ACK_HEADER_BEGIN = 0
	ACK_TYPE_BEGIN
	ACK_TYPE_END = ACK_TYPE_BEGIN + MSG_TYPE_SIZE
	ACK_SEQ_BEGIN
	ACK_SEQ_END = ACK_SEQ_BEGIN + MSG_SEQ_SIZE
	ACK_NEXT_SEQ_BEGIN
	ACK_NEXT_SEQ_END = ACK_NEXT_SEQ_BEGIN + MSG_SEQ_SIZE
	ACK_HEADER_END

	ACK_HEADER_SIZE
)
