// Package util ...
package util

//
import (
	"fmt"

	"github.com/spf13/viper"
)

// GetURI ...
func GetURI() string {
	return fmt.Sprintf("%v:%v", viper.GetString("host"), viper.GetString("port"))
}
