package Hellofs

import (
	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"

	"testing"
)

/* type Myfs struct {
	pathfs.FileSystem
} */

var fs Myfs

func TestGetAttr(t *testing.T) {
	/*
		nfs := pathfs.NewPathNodeFs(&Myfs{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
		_, _, err := nodefs.MountRoot("mnt", nfs.Root(), nil)
		if err != nil {
			os.Exit(2)
		}
		server.Serve()
	*/
	Attr, status := fs.GetAttr("", nil)
	if Attr.Mode != fuse.S_IFDIR && status != fuse.OK {
		t.Errorf("Root Directory cannot be accessed")
	}
}

func TestOpenDir(t *testing.T) {

	dir, status := fs.OpenDir("dir1", nil)
	if dir != nil && status != fuse.OK {
		t.Errorf("Invalid Operation")
	}
}

func TestOpen(t *testing.T) {

	text, status := fs.Open("file.txt", 0, nil)
	if text != nodefs.NewDataFile([]byte("HELLO WORLD")) && status != fuse.OK {
		t.Errorf("Error in opening file")
	}

}
