package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide an integer")
		os.Exit(1)
	}
	aNumber, _ := strconv.ParseInt(os.Args[1], 10, 64)

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, aNumber)
	if err != nil {
		fmt.Println("Little Endian:", err)
	}

	fmt.Printf("%d is %x in Little Endian\n", aNumber, buf)

	err = binary.Write(buf, binary.BigEndian, aNumber)
	if err != nil {
		fmt.Println("Big Endian:", err)
	}
	fmt.Printf("And %x in Big Endian\n", buf)
}
