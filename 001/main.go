package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Start")
	openFile("sample.txt")
	fmt.Println("=============")
	useIoutilReadFile("sample.txt")
	fmt.Println("End")
}

func openFile(filename string) {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	buf := make([]byte, 1024)

	for {
		n, err := fp.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buf))
	}
}

func useIoutilReadFile(fileName string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
