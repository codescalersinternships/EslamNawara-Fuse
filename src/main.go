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

	data := Fuse{
		Name: "Eslam",
		Age:  22,
		Sub: SubStruct{
			SomeValue:      20,
			SomeOtherValue: "some data",
		},
	}

	err = os.MkdirAll(mountPoint, 0777)
	CheckErr(err)

	err = Mount(data, mountPoint)
	CheckErr(err)
}

func CheckErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}
