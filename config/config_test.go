package config_test

import (
	"testing"

	"github.com/ecator/gomeeting/config"
)

func TestConfig(t *testing.T) {
	if c, e := config.ParseConfig("../config.yml"); e == nil {
		t.Log(c)
	} else {
		t.Error(e)
	}

}
