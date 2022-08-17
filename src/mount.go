package main

import (
	"encoding/json"
	"os"
	"reflect"
	"strconv"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/fatih/structs"
)

var (
	inodeCnt   uint64
	Attributes fuse.Attr
	dataMp     map[string]interface{}
)

const errNotPermitted = "Operation not permitted"

type FS struct {
	node fs.Node
}

type EntryGetter interface {
	GetDirentType() fuse.DirentType
}

func Mount(filePath, mountPoint string) error {
	var data []Fuse
	fileContent, err := os.ReadFile(filePath)
	CheckErr(err)

	err = json.Unmarshal(fileContent, &data)
	CheckErr(err)

	con, err := fuse.Mount(mountPoint, fuse.FSName("structFuse"), fuse.Subtype("tmpfs"))
	CheckErr(err)
	defer con.Close()

	inodeCnt = 2

	dataMp = structs.Map(data[0])

	err = fs.Serve(con, FS{})
	CheckErr(err)

	return nil

}

func (FS) Root() (fs.Node, error) {
	dir := NewDir()
	dir.Entries = createEntries(dataMp)
	return dir, nil
}

func createEntries(structMap interface{}) map[string]interface{} {
	entries := map[string]interface{}{}
	for key, val := range structMap.(map[string]interface{}) {
		if reflect.TypeOf(val).Kind() == reflect.Map {
			dir := NewDir()
			dir.Entries = createEntries(val)
			entries[key] = dir
		} else {
			var value string
			if reflect.TypeOf(val).Kind() == reflect.Int {
				value = strconv.Itoa(val.(int))
			} else {
				value = val.(string)
			}
			entries[key] = NewFile([]byte(value))
		}
	}
	return entries
}
