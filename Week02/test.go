package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"src/github.com/axgle/mahonia"
)

var (
	path = "D:/matt_liu/学习/go语言/2020.11.28/data/"
	oFile = "kaifang.txt"
	goodFile = "kaifang_good.txt"
	badFile = "kaifang_bad.txt"
)

type Config struct {
	Server  `config:"server"`
	Mysql	`config:"mysql"`
}

type Server struct {
	Ip string
	Port int
}

type Mysql struct {
	UserName string
	PassWd string
	DataBase string
	Host string
	Port int
	TimeOut float64
}

func main() {


}

func cleanDatas()  {
	oldFile, err := os.Open(path + oFile)
	if err != nil {
		fmt.Println(err)
	}
	defer oldFile.Close()

	goodFile, err := os.OpenFile(path+goodFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer goodFile.Close()
	badFile, err := os.OpenFile(path+badFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer badFile.Close()

	myBytes := bufio.NewReader(oldFile)

	for {
		lineBytes, _, err := myBytes.ReadLine()
		if err == io.EOF {
			break
		}
		lineStr := string(lineBytes)
		enc := mahonia.NewDecoder("GBK")
		convertStr := enc.ConvertString(lineStr)
		fields := strings.Split(convertStr, ",")
		if len(fields) >= 2 && len(fields[1]) == 18 {
			goodFile.WriteString(convertStr+"\n")
			fmt.Println("good:",convertStr)
		}else {
			badFile.WriteString(convertStr+"\n")
			fmt.Println("bad", convertStr)
		}
	}

}
