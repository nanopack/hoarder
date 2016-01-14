package backends_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/nanopack/hoarder/backends"
)

var (
	fsPath = "/tmp/hoarder_fs_test"
	driver = &backends.Filesystem{Path: fsPath}
)

func TestMain(m *testing.M) {
	// start clean
	os.RemoveAll(fsPath)

	// run
	rtn := m.Run()

	// end clean
	os.RemoveAll(fsPath)

	os.Exit(rtn)
}

// test init
func TestInit(t *testing.T) {
	driver.Init()
	if _, err := os.Stat(fsPath); err != nil {
		t.Error("Failed to INIT filesystem - ", err)
	}

	if driver.Path != fsPath {
		t.Error("Failed to INIT filesystem - bad Path setting")
	}
}

// test write
func TestWrite(t *testing.T) {
	reader := bytes.NewBuffer([]byte("testdata"))

	if err := driver.Write("testfile", reader); err != nil {
		t.Error("Failed to WRITE file - ", err)
	}
}

// test list
func TestList(t *testing.T) {
	fileInfo, err := driver.List()
	if err != nil {
		t.Error("Failed to LIST file - ", err)
	}
	if fileInfo[0].Name != "testfile" {
		t.Error("Failed to LIST file - incorrect file found: %s", fileInfo[0].Name)
	}
}

// test read
func TestRead(t *testing.T) {
	reader, err := driver.Read("testfile")
	if err != nil {
		t.Error("Failed to READ file - ", err)
	}

	data, _ := ioutil.ReadAll(reader)
	if string(data) != "testdata" {
		t.Errorf("Failed to READ file - incorrect contents: %s", data)
	}
}

// test stat
func TestStat(t *testing.T) {
	fileInfo, err := driver.Stat("testfile")
	if err != nil {
		t.Error("Failed to STAT file - ", err)
	}
	if fileInfo.Size != 8 {
		t.Errorf("Failed to STAT file - incorrect size: %d", fileInfo.Size)
	}
}

// test remove
func TestRemove(t *testing.T) {
	err := driver.Remove("testfile")
	if err != nil {
		t.Error("Failed to REMOVE file - ", err)
	}
}

// ensure remove idempotency
func TestRemove2(t *testing.T) {
	err := driver.Remove("testfile")
	if err != nil {
		t.Error("Failed to REMOVE file - ", err)
	}
}
