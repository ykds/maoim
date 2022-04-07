package websocket

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const (
	_finBit = 1 << 7
	rsv1Bit = 1 << 6
	rsv2Bit = 1 << 5
	rsv3Bit = 1 << 4
	opBit   = 0x0f // 15

	maskBit = 1 << 7
	lenBit  = 0x7f
)

const (
	ContinuationFrame = 0x0
	TextFrame         = 0x1
	BinaryFrame       = 0x2
	CloseFrame        = 0x8
	PingFrame         = 0x9
	PongFrame         = 0xA
)

type Conn struct {
	rwc     io.ReadWriteCloser
	rb      *bufio.Reader
	wb      *bufio.Writer
	maskKey []byte
}

func newConn(rwc io.ReadWriteCloser, rb *bufio.Reader, wb *bufio.Writer) *Conn {
	return &Conn{
		rwc: rwc,
		rb:  rb,
		wb:  wb,
	}
}

func (c *Conn) GetConn() io.ReadWriteCloser {
	return c.rwc
}

func (c *Conn) ReadWebSocket() (op int, payload []byte, err error) {
	var (
		fin            bool
		partialPayload []byte
		finOp      int
	)

	for {
		fin, op, partialPayload, err = c.readFrame()
		if err != nil {
			return
		}

		switch op {
		case BinaryFrame, TextFrame, ContinuationFrame:
			payload = append(payload, partialPayload...)
			if op != ContinuationFrame {
				finOp = op
			}
			if fin {
				op = finOp
				return
			}
		case PingFrame:
			err = c.WriteWebsocket(PongFrame, partialPayload)
			return
		case CloseFrame:
			err = c.WriteWebsocket(CloseFrame, partialPayload)
			c.Close()
			return
		default:
			err = fmt.Errorf("unknown control message, fin=%t, op=%d", fin, op)
		}
		return
	}
}

func (c *Conn) WriteWebsocket(msgType int, msg []byte) error {
	err := c.writeHeader(msgType, len(msg))
	if err != nil {
		return err
	}
	err = c.writeBody(msg)
	if err != nil {
		return err
	}
	return c.Flush()
}

func (c *Conn) writeHeader(msgType int, msgLen int) error {
	buf := make([]byte, 14)

	buf[0] = 0
	buf[0] |= _finBit | byte(msgType)

	buf[1] = 0
	switch {
	case msgLen <= 125:
		buf[1] |= byte(msgLen)
		buf = buf[:2]
	case msgLen <= 65535:
		buf[1] |= 126
		binary.BigEndian.PutUint16(buf[2:4], uint16(msgLen))
		buf = buf[:4]
	default:
		buf[1] |= 127
		binary.BigEndian.PutUint64(buf[2:10], uint64(msgLen))
		buf = buf[:10]
	}
	_, err := c.wb.Write(buf)
	return err
}

func (c *Conn) writeBody(body []byte) error {
	_, err := c.wb.Write(body)
	return err
}

func (c *Conn) readFrame() (fin bool, op int, data []byte, err error) {
	var (
		dataLen int64
	)
	firstByte, err := c.rb.ReadByte()
	if err != nil {
		return
	}

	fin = (firstByte & _finBit) != 0
	if rsv := firstByte & (rsv1Bit | rsv2Bit | rsv3Bit); rsv != 0 {
		return false, 0, nil, errors.New("不支持rsv")
	}
	op = int(firstByte & opBit)

	secondByte, err := c.rb.ReadByte()
	mask := (secondByte & maskBit) != 0

	switch secondByte & lenBit {
	case 126:
		p := make([]byte, 2)
		_, err = c.rb.Read(p)
		if err != nil {
			return
		}
		dataLen = int64(binary.BigEndian.Uint16(p))
	case 127:
		p := make([]byte, 8)
		_, err = c.rb.Read(p)
		if err != nil {
			return
		}
		dataLen = int64(binary.BigEndian.Uint64(p))
	default:
		dataLen = int64(secondByte & lenBit)
	}

	if mask {
		maskKey := make([]byte, 4)
		_, err = c.rb.Read(maskKey)
		if err != nil {
			return
		}
		if c.maskKey == nil {
			c.maskKey = make([]byte, 4)
		}
		copy(c.maskKey, maskKey)
	}

	if dataLen > 0 {
		data = make([]byte, int(dataLen))
		_, err = c.rb.Read(data)
		if err != nil {
			return
		}
		if mask {
			maskBytes(c.maskKey, data)
		}
	}
	return
}

func (c *Conn) Flush() error {
	return c.wb.Flush()
}

func (c *Conn) Close() error {
	return c.rwc.Close()
}
