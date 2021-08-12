package taxonomies

import (
	"fmt"
	"path/filepath"
	"strings"

	"ecksbee.com/telefacts-taxonomy-package/internal/actions"
	"ecksbee.com/telefacts/pkg/serializables"
)

func Remap(bytes []byte, remap map[string]string) error {
	if VolumePath == "" {
		return fmt.Errorf("empty VolumePath")
	}
	serializables.VolumePath = VolumePath
	unZipFiles, err := actions.Unzip(bytes)
	if err != nil {
		return err
	}
	for _, unZipFile := range unZipFiles {
		unzipped, err := actions.UnzipFile(unZipFile)
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
