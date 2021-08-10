package taxonomies

import (
	zipPkg "archive/zip"
	bytesPkg "bytes"
	"fmt"
	"path/filepath"
	"strings"

	"ecksbee.com/telefacts-taxonomy-package/internal/actions"
	"ecksbee.com/telefacts/pkg/serializables"
)

//todo read bytes as zip and remap to volume/concepts
func Remap(bytes []byte, remap map[string]string) error {
	if VolumePath == "" {
		return fmt.Errorf("empty VolumePath")
	}
	serializables.VolumePath = VolumePath
	bytesReader := bytesPkg.NewReader(bytes)
	zipReader, err := zipPkg.NewReader(bytesReader, bytesReader.Size())
	if err != nil {
		return err
	}
	unZipFiles := zipReader.File
	for _, unZipFile := range unZipFiles {
		unzipped, err := actions.Unzip(unZipFile)
		if err != nil {
			return err
		}
		dir := filepath.Dir(unZipFile.Name)
		var url string
		if remap, found := remap[dir]; found {
			url = strings.Replace(unZipFile.Name, dir, remap, 1)
		}
		dest, err := serializables.UrlToFilename(url)
		if err != nil {
			return err
		}
		err = actions.WriteFile(dest, unzipped)
		if err != nil {
			return err
		}
	}
	return nil
}
