package main

import (
	"testing"
	"io/ioutil"
	"path/filepath"
	"os"
	"path"
)

func TestCp(t *testing.T) {
	//tpath, _ := os.Getwd()
	//testdir := filepath.Join(tpath, "test-get")

	base := ROOT + "/cp/"
	pwd, err := ioutil.TempDir("", "test")
	Nil(t, err)
	localBase := filepath.Join(pwd, "cp")
	func() {
		SetUp()
		err := os.MkdirAll(localBase, 0755)
		err = os.MkdirAll(localBase+"/test", 0755)
		Nil(t, err)
	}()
	defer TearDown()

	err = os.Chdir(localBase)
	Nil(t, err)
	Upx("mkdir", base)
	Upx("cd", base)
	// upx put localBase/FILE upBase/FILE
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	t.Log("local:", getwd)
	t.Log("localbase:", localBase)

	// upx put /path/to/dir /path/to/dir/

	putDir(t, localBase, base+"/putdir/", base+"/putdir/")
	CreateFile("FILE")
	oldPath := filepath.Join(base, "FILE")
	putFile(t, filepath.Join(localBase, "FILE"), "", path.Join(base, "FILE"))
	newPath := base + "putdir/"
	t.Log("dir", localBase+"test", base)
	t.Log(oldPath, newPath)
	mvFile(t, oldPath, newPath)
}
