package taxonomies

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"path/filepath"
	"sync"
	"time"

	"ecksbee.com/telefacts-taxonomy-package/internal/actions"
	"ecksbee.com/telefacts/pkg/attr"
	"ecksbee.com/telefacts/pkg/serializables"
	"github.com/joshuanario/r8lmt"
)

var (
	wLock     sync.Mutex
	out       chan interface{} = make(chan interface{})
	in        chan interface{} = make(chan interface{})
	dur       time.Duration    = 200 * time.Millisecond
	throttled bool             = false
)

func startSECThrottle() {
	if !throttled {
		r8lmt.Throttler(out, in, dur, false)
		throttled = true
	}
}

func Discover(entries []string) error {
	if VolumePath == "" {
		return fmt.Errorf("empty VolumePath")
	}
	startSECThrottle()
	serializables.VolumePath = VolumePath
	for _, entry := range entries {
		url, err := serializables.UrlToFilename(entry)
		if err != nil {
			return err
		}
		filepath := filepath.Join(VolumePath, "concepts", url)
		schemaFile, err := serializables.ReadSchemaFile(filepath)
		if err != nil {
			return err
		}
		if schemaFile == nil {
			continue
		}
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			importSchema(schemaFile)
		}()
		go func() {
			defer wg.Done()
			includeSchema(schemaFile)
		}()
		wg.Wait()
	}
	return nil
}

func includeSchema(file *serializables.SchemaFile) {
	if file == nil {
		return
	}
	includes := file.Include
	var wg sync.WaitGroup
	wg.Add(len(includes))
	for _, iitem := range includes {
		go func(item struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr "xml:\",any,attr\""
		}) {
			defer wg.Done()
			if item.XMLName.Space != attr.XSD {
				return
			}
			schemaLocationAttr := attr.FindAttr(item.XMLAttrs, "schemaLocation")
			if schemaLocationAttr == nil || schemaLocationAttr.Value == "" {
				return
			}
			if attr.IsValidUrl(schemaLocationAttr.Value) {
				go discoverRemoteURL(schemaLocationAttr.Value)
				return
			}
		}(iitem)
	}
	wg.Wait()
}

func importSchema(file *serializables.SchemaFile) {
	if file == nil {
		return
	}
	imports := file.Import
	var wg sync.WaitGroup
	wg.Add(len(imports))
	for _, iitem := range imports {
		go func(item struct {
			XMLName  xml.Name
			XMLAttrs []xml.Attr "xml:\",any,attr\""
		}) {
			defer wg.Done()
			if item.XMLName.Space != attr.XSD {
				return
			}
			namespaceAttr := attr.FindAttr(item.XMLAttrs, "namespace")
			if namespaceAttr == nil || namespaceAttr.Value == "" {
				return
			}
			schemaLocationAttr := attr.FindAttr(item.XMLAttrs, "schemaLocation")
			if schemaLocationAttr == nil || schemaLocationAttr.Value == "" {
				return
			}
			if attr.IsValidUrl(schemaLocationAttr.Value) {
				go discoverRemoteURL(schemaLocationAttr.Value)
				return
			}
		}(iitem)
	}
	wg.Wait()
}

func throttle(urlString string) {
	urlStruct, err := url.Parse(urlString)
	if urlStruct.Hostname() != "sec.gov" {
		return
	}
	if err != nil {
		return
	}
	in <- struct{}{}
	<-out
}

func discoverRemoteURL(url string) {
	body, err := actions.Scrape(url, throttle)
	if err != nil {
		return
	}
	dest, err := serializables.UrlToFilename(url)
	if err != nil {
		return
	}
	wLock.Lock()
	defer wLock.Unlock()
	actions.WriteFile(dest, body)
	discoveredSchema, err := serializables.ReadSchemaFile(dest)
	if err != nil {
		return
	}
	var wwg sync.WaitGroup
	wwg.Add(2)
	go func() {
		defer wwg.Done()
		importSchema(discoveredSchema)
	}()
	go func() {
		defer wwg.Done()
		includeSchema(discoveredSchema)
	}()
	wwg.Wait()
}
