package backends_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/viper"

	"github.com/nanopack/hoarder/backends"
)

var (
	fsPath = "/tmp/hoarder_fs_test"
)

func TestMain(m *testing.M) {

	// start clean
	os.RemoveAll(fsPath)

	// configure
	viper.Set("backend", fsPath)

	// run
	rtn := m.Run()

	// end clean
	os.RemoveAll(fsPath)

	os.Exit(rtn)
}

// test init
func TestInit(t *testing.T) {
	backends.Initialize()
	if _, err := os.Stat(fsPath); err != nil {
		t.Error("Failed to INIT filesystem - ", err)
	}
}

// test write
func TestWrite(t *testing.T) {
	reader := bytes.NewBuffer([]byte("testdata"))

	if err := backends.Write("testfile", reader); err != nil {
		t.Error("Failed to WRITE file - ", err)
	}
}

// test list
func TestList(t *testing.T) {
	DataInfo, err := backends.List()
	if err != nil {
		t.Error("Failed to LIST file - ", err)
	}
	if DataInfo[0].Name != "testfile" {
		t.Errorf("Failed to LIST file - incorrect file found: %s", DataInfo[0].Name)
	}
}

// test read
func TestRead(t *testing.T) {
	reader, err := backends.Read("testfile")
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
	DataInfo, err := backends.Stat("testfile")
	if err != nil {
		t.Error("Failed to STAT file - ", err)
	}
	if DataInfo.Size != 8 {
		t.Errorf("Failed to STAT file - incorrect size: %d", DataInfo.Size)
	}
}

// test remove
func TestRemove(t *testing.T) {
	err := backends.Remove("testfile")
	if err != nil {
		t.Error("Failed to REMOVE file - ", err)
	}
}

// ensure remove idempotency
func TestRemove2(t *testing.T) {
	err := backends.Remove("testfile")
	if err != nil {
		t.Error("Failed to REMOVE file - ", err)
	}
}
