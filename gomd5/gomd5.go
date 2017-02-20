package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main()  {
	hash := md5.New()
	hash.Write([]byte("123456"))

	rs := hex.EncodeToString(hash.Sum(nil))

	fmt.Println(rs)
	//e10adc3949ba59abbe56e057f20f883e
	//e10adc3949ba59abbe56e057f20f883e
}
