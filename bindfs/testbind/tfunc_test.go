package bindfs

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

var fs bindFs

func TestGetAttr(t *testing.T) {
	str, err_0 := filepath.Abs("testfolder")
	if err_0 != nil {
		fmt.Print("Error in getting the absolute path")
		os.Exit(2)
	}
	//	fmt.Println("This is absolute path ", str)
	attr, status := fs.GetAttr(str, nil)
	var err error
	st := syscall.Stat_t{}
	err = syscall.Stat(str, &st)
	if err != nil {
		fmt.Print("Error is listing files and directories")
	}
	if status != fuse.OK || st.Ino != attr.Ino {
		t.Errorf("Unable to list out the files and folders")
	}
}

//func TestOpenDir(t *Testing.T) {}

func TestOpen(t *testing.T) {
	str, err_0 := filepath.Abs("testfolder/dir1/text0.txt")
	if err_0 != nil {
		fmt.Print("error in getting the absolute path")
	}
	file, status := fs.Open(str, 0, nil)
	text, err := os.Open(str)
	if err != nil {
		fmt.Print("Error in Opening file")
		os.Exit(1)
	}
	if file != nodefs.NewLoopbackFile(text) || status != fuse.OK {
		t.Errorf("Unable to open the specified file")
	}

}
