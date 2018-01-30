package main

import ("fmt"
	"bytes"
	"encoding/binary"
	"encoding/hex"
)
func main() {
	pi := 3.1415926
	buf := bytes.Buffer{}
	//使用的是小编码，低地址对应低字节
	binary.Write(&buf, binary.LittleEndian, &pi)
	//常量浮点数默认是float64
	fmt.Printf("buf=%#v\n", buf.Bytes())

	var rpi float64
	binary.Read(&buf, binary.LittleEndian, &rpi)
	fmt.Printf("rpi=%#v\n", rpi)


}
