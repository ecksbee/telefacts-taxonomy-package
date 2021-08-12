package test_test

import (
	"path/filepath"
	"testing"

	"ecksbee.com/telefacts-taxonomy-package/pkg/install"
)

func Test_Run(t *testing.T) {
	tp := filepath.Join(".", "us-gaap-2020-01-31.zip")
	volume := filepath.Join(".", "data")
	err := install.Run(tp, volume)
	if err != nil {
		t.Fatal(err)
	}
}
