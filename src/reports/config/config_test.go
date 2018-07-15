package config_test

import (
	"fmt"
	"testing"

	"github.com/RTradeLtd/mining-bootstrap/src/reports/config"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := config.LoadConfig("")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", cfg)
}
