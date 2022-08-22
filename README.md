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

	"fuse/src"
)

//change these structs as you wish to represent your file system
type structure struct {
	String       string
	Int          int
	Bool         bool
	SubStructure subStructure
}

type subStructure struct {
	Float float32
}

func Routine(input *structure) {
	time.Sleep(time.Second * 5)
	input.String = "new string"
}

func main() {
	var err error
	if len(os.Args) != 2 {
		fmt.Println("too few arguments")
		fmt.Println(len(os.Args))
		os.Exit(1)
	}
	mountPoint := os.Args[1]
	// fill the struct as you wish with data to be represented in the files
	input := &structure{
		String: "str",
		Int:    18,
		Bool:   true,
		SubStructure: subStructure{
			Float: 1.3,
		},
	}

	err = os.MkdirAll(mountPoint, 0777)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	go Routine(input)
	err = fs.Mount(mountPoint, input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
```
