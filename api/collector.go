package api

import (
	"fmt"
	"time"

	"github.com/jcelliott/lumber"
	"github.com/spf13/viper"
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

	lumber.Debug("Garbage Collector - Finding files...")
	for _, data := range datas {

		// CleanAfter.Value defaults to Now() to ensure no files are deleted in case
		// cobra decides to change how 'Command.Flag().Changed' works. It does this
		// because no files, written by hoarder, will have a modified time before the
		// Unix epoch began
		if data.ModTime.Unix() < (now.Unix() - int64(viper.GetInt("clean-after"))) {
			fmt.Printf("Cleaning key: %s\n", data.Name)
			if err := driver.Remove(data.Name); err != nil {
				return fmt.Errorf("Cleaning of '%s' failed - %v", data.Name, err.Error())
			}
		}
	}

	return nil
}

// StartCollection calls removeOldKeys at set intervals
func StartCollection() {
	tick := time.Tick(CLEAN_FREQ * time.Second)

	for _ = range tick {
		if err := removeOldKeys(); err != nil {
			fmt.Println(err.Error())
		}
	}
}
