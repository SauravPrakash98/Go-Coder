package bindfs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"syscall"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
	"github.com/hanwen/go-fuse/fuse/pathfs"
)

type bindFs struct {
	pathfs.FileSystem
	Root string
}

func (fs *bindFs) GetAttr(name string, context *fuse.Context) (a *fuse.Attr, code fuse.Status) {
	copypath := filepath.Join(fs.Root, name)
	var err error = nil
	st := syscall.Stat_t{}
	err = syscall.Stat(copypath, &st)
	if err != nil {
		return nil, fuse.ToStatus(err)
	}
	a = &fuse.Attr{}
	a.FromStat(&st)
	return a, fuse.OK
}

func (fs *bindFs) OpenDir(name string, context *fuse.Context) (stream []fuse.DirEntry, code fuse.Status) {
	f, err := os.Open(filepath.Join(fs.Root, name))
	if err != nil {
		return nil, fuse.ToStatus(err)
	}
	final := make([]fuse.DirEntry, 0, 100)
	for {
		inf, err := f.Readdir(100)
		for i := range inf {
			if inf[i] == nil {
				continue
			}
			n := inf[i].Name()
			d := fuse.DirEntry{
				Name: n,
			}
			s := fuse.ToStatT(inf[i])
			if s != nil {
				d.Mode = uint32(s.Mode)
				d.Ino = s.Ino
			} else {
				fmt.Printf("ReadDir entry %q for %q has no stat info", n, name)
				os.Exit(5)
			}
			final = append(final, d)
		}
		if len(final) < 100 || err == io.EOF {
			break
		}
		if err != nil {
			fmt.Print("directory could not be opened")
			os.Exit(4)
		}
	}
	f.Close()
	return final, fuse.OK
}

func (fs *bindFs) Open(name string, flags uint32, context *fuse.Context) (FILE nodefs.File, code fuse.Status) {
	text, err := os.OpenFile(filepath.Join(fs.Root, name), int(flags), 0)
	if err != nil {
		fmt.Print("Error in opening file")
		os.Exit(3)
	}
	return nodefs.NewLoopbackFile(text), fuse.OK
}
