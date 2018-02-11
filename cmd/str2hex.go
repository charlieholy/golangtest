package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := "Go is good language!我们都一样"
	//根据编码后长度来分配缓存空间
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, []byte(src))
	fmt.Printf("dst=%v\n", dst)
	fmt.Printf("dst(hex)=%s\n", string(dst))

	//根据解码长度来分配缓存空间
	dst2 := make([]byte, hex.DecodedLen(len(dst)))
	hex.Decode(dst2, dst)
	fmt.Printf("dst2=%v\n", dst2)
	fmt.Printf("dst2.(string)=%s\n", string(dst2))
}
