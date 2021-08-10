package install

import "ecksbee.com/telefacts-taxonomy-package/pkg/taxonomies"

func Run(taxonomyPackage string) error {
	id := ""
	name := ""
	bytes := []byte{}
	entries := []string{}

	return taxonomies.NewTaxonomy(id, taxonomies.TaxonomyMeta{
		Name:    name,
		Zip:     taxonomyPackage,
		Entries: entries,
	}, bytes)
}
