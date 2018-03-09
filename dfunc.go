package Hellofs

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type Myfs struct {
	pathfs.FileSystem
}

func (fs *Myfs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
	switch name {
	case "":
		return &fuse.Attr{
			Mode: fuse.S_IFDIR,
		}, fuse.OK
	case "dir1":
		return &fuse.Attr{
			Mode: fuse.S_IFDIR,
		}, fuse.OK
	case "file.txt":
		return &fuse.Attr{
			Mode: fuse.S_IFREG, Size: uint64(len(name)),
		}, fuse.OK

	}
	return nil, fuse.ENOENT
}

func (fs *Myfs) OpenDir(name string, context *fuse.Context) (dir []fuse.DirEntry, code fuse.Status) {
	switch name {
	case "":
		return []fuse.DirEntry{{Name: "dir1", Mode: fuse.S_IFDIR}, {Name: "file.txt", Mode: fuse.S_IFREG}}, fuse.OK
	case "dir1":
		//return []fuse.DirEntry{{Name: "file.txt", Mode: fuse.S_IFREG}}, fuse.OK
		return nil, fuse.OK
	}
	return nil, fuse.ENOENT
}

func (fs *Myfs) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
	if name != "file.txt" {
		return nil, fuse.ENOENT
	}
	return nodefs.NewDataFile([]byte("HELLO WORLD")), fuse.OK
}
