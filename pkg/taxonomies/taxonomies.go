package taxonomies

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"ecksbee.com/telefacts/pkg/serializables"
)

var (
	VolumePath string
)

type Meta struct {
	Name    string
	Zip     string
	Entries []string
	Remap   map[string]string
}

func NewTaxonomy(id string, tm Meta, bytes []byte) error {
	if VolumePath == "" {
		return fmt.Errorf("empty VolumePath")
	}
	err := Remap(bytes, tm.Remap)
	if err != nil {
		return err
	}
	err = Discover(tm.Entries)
	if err != nil {
		return err
	}
	serializables.VolumePath = VolumePath
	workingDir := filepath.Join(VolumePath, "taxonomies")
	pathStr := filepath.Join(workingDir, id)
	_, err = os.Stat(pathStr)
	if !os.IsNotExist(err) {
		return fmt.Errorf("%s cannot be overwritten", id)
	}
	err = os.Mkdir(pathStr, 0755)
	if err != nil {
		return err
	}
	meta := filepath.Join(pathStr, "_")
	file, _ := os.OpenFile(meta, os.O_CREATE, 0755)
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(tm)
}
