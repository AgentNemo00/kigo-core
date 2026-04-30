package util

import (
	"encoding/json"
	"encoding/binary"
)

func MapToStruct(in any, out any) error {
    data, err := json.Marshal(in)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, out)
}

func Uint16ToBytesBE(value uint16) []byte {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, value)
	return buf
}

func Uint32ToBytesBE(value uint32) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, value)
	return buf
}

