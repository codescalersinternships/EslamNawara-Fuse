package main

import (
	"fmt"
	"reflect"
	"strconv"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"github.com/fatih/structs"
)

const errNotPermitted = "Operation not permitted"

type FS struct {
	dataMp map[string]any
}

type EntryGetter interface {
	GetDirentType() fuse.DirentType
}

func Mount(data []Fuse, mountPoint string) error {
	con, err := fuse.Mount(mountPoint)
	if err != nil {
		return err
	}

	defer con.Close()

	err = fs.Serve(con, NewFS(data))
	if err != nil {
		return err
	}

	return nil

}
func NewFS(data []Fuse) *FS {
	return &FS{
		dataMp: createDataMap(data),
	}
}

func (fs *FS) Root() (fs.Node, error) {
	dir := NewDir()
	dir.Entries = createEntries(fs.dataMp)
	return dir, nil
}

func createEntries(structMap any) map[string]any {
	entries := map[string]any{}
	for key, val := range structMap.(map[string]any) {
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

func createDataMap(data []Fuse) map[string]any {
	dataMap := make(map[string]any)
	for i, elem := range data {
		dataMap[strconv.Itoa(i)] = structs.Map(elem)
	}
	return dataMap
}
