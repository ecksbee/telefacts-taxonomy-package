package actions

import (
	zipPkg "archive/zip"
	bytesPkg "bytes"
	"io"
)

func UnzipFile(unzipFile *zipPkg.File) ([]byte, error) {
	rc, err := unzipFile.Open()
	defer rc.Close()
	if err != nil {
		return nil, err
	}
	var buffer bytesPkg.Buffer
	_, err = io.Copy(&buffer, rc)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Unzip(bytes []byte) ([]*zipPkg.File, error) {
	bytesReader := bytesPkg.NewReader(bytes)
	zipReader, err := zipPkg.NewReader(bytesReader, bytesReader.Size())
	if err != nil {
		return nil, err
	}
	return zipReader.File, nil
}
