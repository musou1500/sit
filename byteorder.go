package sit

import (
	"encoding/binary"
)

func Htons(host uint16) uint16 {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, host)
	return binary.BigEndian.Uint16(bytes)
}

func Htonl(host uint32) uint32 {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(host))
	return binary.BigEndian.Uint32(bytes)
}
