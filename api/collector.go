package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/nanopack/hoarder/config"
)

// garbage collection will have this many seconds before consistancy
const CLEAN_FREQ = 10

// removeOldKeys removes keys older than specified age
func removeOldKeys() error {
	datas, err := driver.List()
	if err != nil {
		return err
	}

	now := time.Now()

	config.Log.Trace("Garbage Collector - Finding files...")
	for _, data := range datas {
		// CleanAfter.Value defaults to Now() to ensure no files are deleted in case
		// cobra decides to change how 'Command.Flag().Changed' works. It does this
		// because no files, written by hoarder, will have a modified time before the
		// Unix epoch began
		if data.ModTime.Unix() < (now.Unix() - int64(config.CleanAfter.Value)) {
			config.Log.Debug("Cleaning key: %s", data.Name)
			if err := driver.Remove(data.Name); err != nil {
				return errors.New(fmt.Sprintf("Cleaning of '%s' failed - ", data.Name, err.Error()))
			}
		}
	}

	return nil
}

// startCollection calls removeOldKeys at set intervals
func startCollection() {
	tick := time.Tick(CLEAN_FREQ * time.Second)

	for _ = range tick {
		if err := removeOldKeys(); err != nil {
			config.Log.Error(err.Error())
		}
	}
}
