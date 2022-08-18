package main

type SubStruct struct {
    SomeValue      int    `json:"Value"`
    SomeOtherValue string `json:"Other"`
}

type Fuse struct {
    Name string    `fuse:"Name"`
    Age  int       `fuse:"Age"`
    Sub  SubStruct `fuse:"Sub"`
}

