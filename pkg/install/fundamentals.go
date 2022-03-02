package install

import (
	"ecksbee.com/telefacts-taxonomy-package/internal/actions"
	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/serializables"
)

func DownloadUTR(throttle func(string)) error {
	utr, err := actions.Scrape(attr.UTR, throttle)
	if err != nil {
		return err
	}
	dest, err := serializables.UrlToFilename(attr.UTR)
	if err != nil {
		return err
	}
	return actions.WriteFile(dest, utr)
}

func DownloadLRR(throttle func(string)) error {
	lrr, err := actions.Scrape(attr.LRR, throttle)
	if err != nil {
		return err
	}
	dest, err := serializables.UrlToFilename(attr.LRR)
	if err != nil {
		return err
	}
	return actions.WriteFile(dest, lrr)
}
