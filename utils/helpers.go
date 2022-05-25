package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func TypeOf(x interface{}) {
	fmt.Printf("%T\n", x)
}

func ToByteSlice(x int64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, x)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
