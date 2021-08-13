package test_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"ecksbee.com/telefacts-taxonomy-package/pkg/install"
)

func Test_Run(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	tp := filepath.Join(wd, "us-gaap-2020-01-31.zip")
	volume := filepath.Join(wd, "data")
	id, err := install.Run(tp, volume)
	if err != nil {
		t.Fatal(err)
	}
	if id == "" {
		t.Fatal(fmt.Errorf("empty id generated"))
	}
}
