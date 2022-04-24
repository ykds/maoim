package protocal

import (
	"encoding/binary"
)

const (
	_opSize     = 2
	_ackSize    = 2
	_seqSize    = 4
	_lengthSize = 4
	_headerSize = _opSize + _ackSize + _seqSize + _lengthSize

	_opOffset     = _opSize
	_ackOffset    = _opOffset + _ackSize
	_seqOffset    = _ackOffset + _seqSize
	_lengthOffset = _seqOffset + _lengthSize
)

func (p *Proto) Pack() []byte {
	bodyLen := len(p.Body)
	buf := make([]byte, _headerSize+bodyLen)
	binary.BigEndian.PutUint16(buf[:_opOffset], uint16(p.Op))
	binary.BigEndian.PutUint16(buf[_opOffset:_ackOffset], uint16(p.Ack))
	binary.BigEndian.PutUint32(buf[_ackOffset:_seqOffset], uint32(p.Seq))
	binary.BigEndian.PutUint32(buf[_seqOffset:_lengthOffset], uint32(bodyLen))
	if bodyLen != 0 {
		copy(buf[_lengthOffset:], p.Body)
	}
	return buf
}

func (p *Proto) PackHeartBeat() []byte {
	buf := make([]byte, _headerSize)
	binary.BigEndian.PutUint16(buf[:_opOffset], uint16(OpHeartBeat))
	return buf
}

func (p *Proto) Unpack(data []byte) {
	if len(data) == 0 {
		return
	}

	p.Op = int32(binary.BigEndian.Uint32(data[:_opOffset]))
	p.Ack = int32(binary.BigEndian.Uint32(data[_opOffset:_ackOffset]))
	p.Seq = int32(binary.BigEndian.Uint32(data[_ackOffset:_seqOffset]))
	length := int(binary.BigEndian.Uint32(data[_seqOffset:_lengthOffset]))

	body := data[_lengthOffset:]

	if length != len(body) {
		p.Body = nil
	} else {
		p.Body = body
	}
}
