package main

import (
	"fmt"
	"reflect"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/fatih/structs"
)

const errNotPermitted = "Operation not permitted"

type FS struct {
	dataMp map[string]interface{}
}

type EntryGetter interface {
	GetDirentType() fuse.DirentType
}

func Mount(data Fuse, mountPoint string) error {
	con, err := fuse.Mount(mountPoint)
	CheckErr(err)
	defer con.Close()

	err = fs.Serve(con, NewFS(data))
	CheckErr(err)

	return nil

}
func NewFS(data Fuse) *FS {
	return &FS{
		dataMp: structs.Map(data),
	}
}

func (fs *FS) Root() (fs.Node, error) {
	dir := NewDir()
	dir.Entries = createEntries(fs.dataMp)
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
			entries[key] = NewFile([]byte(fmt.Sprint(reflect.ValueOf(val))))
		}
	}
	return entries
}
