package email

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func CompressToGzip(srcPath, dstPath string) error {
	sf, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer sf.Close()
	sfInfo, err := sf.Stat()
	if err != nil {
		return err
	}
	df, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer df.Close()
	writer := gzip.NewWriter(df)
	writer.Name = sfInfo.Name()
	writer.ModTime = sfInfo.ModTime()
	_, err = io.Copy(writer, sf)
	if err != nil {
		return err
	}
	if err = writer.Close(); err != nil {
		return err
	}
	return nil
}

func DecompressFromGzip(srcPath string) (string, error) {
	sf, err := os.Open(srcPath)
	if err != nil {
		return "", err
	}
	defer sf.Close()
	gr, err := gzip.NewReader(sf)
	if err != nil {
		return "", err
	}
	dir := filepath.Dir(srcPath)
	dstPath := filepath.Join(dir, gr.Name)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	_, err = io.Copy(dst, gr)
	if err != nil {
		return "", err
	}
	return dstPath, nil
}
