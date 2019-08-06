package config_test

import (
	"fmt"
	"testing"

	"github.com/ecator/gomeeting/config"
)

func TestConfig(t *testing.T) {
	if c, e := config.ParseConfig("../config.yml"); e == nil {
		fmt.Println(c)
	} else {
		fmt.Println(e)
	}

}
