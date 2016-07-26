// Package collector handles the cleanup of old stored data.
package collector

import (
	"fmt"
	"time"

	"github.com/jcelliott/lumber"
	"github.com/spf13/viper"

	"github.com/nanopack/hoarder/backends"
)

// garbage collection will have this many seconds before consistancy
const CLEAN_FREQ = 10

// removeOldKeys removes keys older than specified age
func removeOldKeys() error {
	datas, err := backends.List()
	if err != nil {
		return err
	}

	now := time.Now()

	lumber.Debug("Garbage Collector - Finding files...")
	for _, data := range datas {

		// clean-after defaults to Now() to ensure no files are deleted in case
		// cobra decides to change how 'Command.Flag().Changed' works. It does this
		// because no files, written by hoarder, will have a modified time before the
		// Unix epoch began
		if data.ModTime.Unix() < (now.Unix() - int64(viper.GetInt("clean-after"))) {
			lumber.Debug("Cleaning key: ", data.Name)
			if err := backends.Remove(data.Name); err != nil {
				return fmt.Errorf("Cleaning of '%s' failed - %v", data.Name, err.Error())
			}
		}
	}

	return nil
}

// Start calls removeOldKeys at set intervals
func Start() {
	tick := time.Tick(CLEAN_FREQ * time.Second)

	for _ = range tick {
		if err := removeOldKeys(); err != nil {
			fmt.Println(err.Error())
		}
	}
}
