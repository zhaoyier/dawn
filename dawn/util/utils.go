package util

import (
	"encoding/binary"
)

func BytesToInt32(buf []byte) int32 {
	return int32(binary.LittleEndian.Uint32(buf))
}

func BytesToInt64(buf []byte) int64  {
	return int64(binary.LittleEndian.Uint64(buf))
}

func Int32ToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}
