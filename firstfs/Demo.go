package main

import (
	"flag"
	"fmt"
	"os"

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
			Mode: fuse.S_IFDIR | 0755,
		}, fuse.OK
	case "dir1":
		return &fuse.Attr{
			Mode: fuse.S_IFDIR,
		}, fuse.OK
	case "file.txt":
		return &fuse.Attr{
			Mode: fuse.S_IFREG | 0644, Size: uint64(len("HELLO WORLD")),
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
	/*	if flags != 0 && fuse.O_ANYWRITE != 0 {
			return nil, fuse.EPERM
		}
	*/
	return nodefs.NewDataFile([]byte("HELLO WORLD")), fuse.OK
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Printf("Write the name of mount point\n")
		os.Exit(1)
	}
	//	var newfs Myfs
	//	newfs.fs = pathfs.NewDefaultFileSystem()

	nfs := pathfs.NewPathNodeFs(&Myfs{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
	server, _, err := nodefs.MountRoot(flag.Arg(0), nfs.Root(), nil)
	if err != nil {
		os.Exit(2)
	}
	server.Serve()

}
