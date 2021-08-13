package test_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"ecksbee.com/telefacts-taxonomy-package/pkg/install"
)

func Test_Run_USGAAP2020(t *testing.T) {
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

func Test_Run_ESEF2017(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	tp := filepath.Join(wd, "esef_taxonomy_2017.zip")
	volume := filepath.Join(wd, "data")
	id, err := install.Run(tp, volume)
	if err != nil {
		t.Fatal(err)
	}
	if id == "" {
		t.Fatal(fmt.Errorf("empty id generated"))
	}
}
