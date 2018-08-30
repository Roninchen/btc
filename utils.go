package main

import (
	"encoding/binary"
	"bytes"
	"os"
	"fmt"
)

func IntToByte(num int64) []byte {
	//func Write(w io.Writer, order ByteOrder, data interface{})
	var buffer bytes.Buffer
	err:=binary.Write(&buffer,binary.BigEndian,num)
	CheckErr("IntToByte",err)
	return buffer.Bytes()
}

func CheckErr(pos string,err error)  {
	if err !=nil {
		fmt.Println("error ,pos :",pos ,err)
		os.Exit(1)
	}
}

