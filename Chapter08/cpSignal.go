package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
)

var BUFFERSIZE int64
var FILESIZE int64
var BYTESWRITTEN int64

func Copy(src, dst string, BUFFERSIZE int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	FILESIZE = sourceFileStat.Size()

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
		BYTESWRITTEN = BYTESWRITTEN + int64(n)
	}
	return err
}

func progressInfo() {
	progress := float64(BYTESWRITTEN) / float64(FILESIZE) * 100
	fmt.Printf("Progress: %.2f%%\n", progress)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("usage: %s source destination BUFFERSIZE\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	source := os.Args[1]
	destination := os.Args[2]
	BUFFERSIZE, _ = strconv.ParseInt(os.Args[3], 10, 64)
	BYTESWRITTEN = 0

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGINFO:
				progressInfo()
			default:
				fmt.Println("Ignored:", sig)
			}
		}
	}()

	fmt.Printf("Copying %s to %s\n", source, destination)
	err := Copy(source, destination, BUFFERSIZE)
	if err != nil {
		fmt.Printf("File copying failed: %q\n", err)
	}
}
