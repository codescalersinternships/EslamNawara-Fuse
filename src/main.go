package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	if len(os.Args) != 2 {
		fmt.Println("Mounting point not specified")
		return
	}
	mountPoint := os.Args[1]
	var data []Fuse
	data = append(data, Fuse{
		Name: "Eslam",
		Age:  22,
		Sub: SubStruct{
			SomeValue:      20,
			SomeOtherValue: "some data",
		},
	})

	err = os.MkdirAll(mountPoint, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = Mount(data, mountPoint)
	if err != nil {
		fmt.Println(err)
		return
	}
}
