package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/nanopack/hoarder/config"
)

const CLEAN_FREQ = 10

func removeOldKeys() error {
	if !config.GarbageCollect {
		return errors.New("Garbage collection not 'on' but cleanup called! Killing GC")
	}
	datas, err := driver.List()
	if err != nil {
		return err
	}

	now := time.Now()

	config.Log.Trace("Garbage Collector - Finding files...")
	for _, data := range datas {
		if data.ModTime.Unix() < (now.Unix() - int64(config.CleanAfter)) {
			config.Log.Debug("Cleaning key: %s", data.Name)
			if err := driver.Remove(data.Name); err != nil {
				return errors.New(fmt.Sprintf("Cleaning of '%s' failed - ", data.Name, err.Error()))
			}
		}
	}

	return nil
}

func startCollection() {
	tick := time.Tick(CLEAN_FREQ * time.Second)

	for _ = range tick {
		if err := removeOldKeys(); err != nil {
			config.Log.Error(err.Error())
			return
		}
	}
}
