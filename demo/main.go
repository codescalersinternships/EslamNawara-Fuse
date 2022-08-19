package main

import (
	"fmt"
	"os"
	"time"

	"EslamNawara-Fuse/fs"
)

type SubStruct struct {
	SomeValue      int
	SomeOtherValue string
}

type Fuse struct {
	Name string
	Age  int
	Sub  SubStruct
}

func main() {
	var err error
	if len(os.Args) != 2 {
		fmt.Println("Mounting point not specified")
		return
	}
	mountPoint := os.Args[1]
	data := &Fuse{
		Name: "Eslam",
		Age:  22,
		Sub: SubStruct{
			SomeValue:      20,
			SomeOtherValue: "some data",
		},
	}

	err = os.MkdirAll(mountPoint, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		time.Sleep(2 * time.Second)
		data.Name = "not eslam"
	}()

	err = fs.Mount(data, mountPoint)
	if err != nil {
		fmt.Println(err)
		return
	}
}
