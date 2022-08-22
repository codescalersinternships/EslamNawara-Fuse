# Fuse Filesystem
Filesystem in USErspace (FUSE) is a software interface for Unix and Unix-like computer operating systems that lets non-privileged users create their own file systems without editing kernel code. This is achieved by running file system code in user space while the FUSE module provides only a bridge to the actual kernel interfaces. 

## How to use 
- Clone the repo
```sh
$ git clone https://github.com/codescalersinternships/EslamNawara-Fuse
```
- Run the demo program
```sh
$ go run demo/main.go <mountPoint>
```

## Demo
This demo creates the file system from a struct hierarchy, you can change it as you wish, then hold 2 seconds then change the value of field of struct to be reflected on the filesystem.
```go
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
		data.Name = "not Eslam"
	}()

	err = fs.Mount(data, mountPoint)
	if err != nil {
		fmt.Println(err)
		return
	}
}
```
