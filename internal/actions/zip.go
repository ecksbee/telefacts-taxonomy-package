package actions

import (
	"archive/zip"
	"bytes"
)

func Zip(unZipFiles []*zip.File) ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	for _, file := range unZipFiles {
		f, err := writer.Create(file.Name)
		if err != nil {
			return nil, err
		}
		data, err := UnzipFile(file)
		if err != nil {
			return nil, err
		}
		_, err = f.Write([]byte(data))
		if err != nil {
			return nil, err
		}
	}
	err := writer.Close()
	return buf.Bytes(), err
}
