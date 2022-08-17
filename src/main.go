package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	if len(os.Args) != 3 {
		fmt.Println("too few arguments")
		fmt.Println(len(os.Args))
		return
	}
	filePath := os.Args[1]
	mountPoint := os.Args[2]

	err = os.MkdirAll(mountPoint, 0777)
	CheckErr(err)

	err = Mount(filePath, mountPoint)
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
