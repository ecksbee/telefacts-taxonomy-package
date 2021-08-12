package install

import (
	"io/ioutil"

	"ecksbee.com/telefacts-taxonomy-package/internal/actions"
	"ecksbee.com/telefacts-taxonomy-package/pkg/taxonomies"
)

func Run(taxonomyPackage string, volumePath string) error {
	id := ""
	bytes, err := ioutil.ReadFile(taxonomyPackage)
	if err != nil {
		return err
	}
	unZipFiles, err := actions.Unzip(bytes)
	if err != nil {
		return err
	}
	name, err := name(unZipFiles)
	if err != nil {
		return err
	}
	entries, err := entries(unZipFiles)
	if err != nil {
		return err
	}
	remap, err := remap(unZipFiles)
	if err != nil {
		return err
	}
	cleanedBytes, err := clean(unZipFiles)
	if err != nil {
		return err
	}

	taxonomies.VolumePath = volumePath
	return taxonomies.NewTaxonomy(id, taxonomies.Meta{
		Name:    name,
		Zip:     taxonomyPackage,
		Entries: entries,
		Remap:   remap,
	}, cleanedBytes)
}
