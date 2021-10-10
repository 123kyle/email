package email

import (
	"archive/zip"
	"io"
	"os"
)

func CompressToZip(dstPath string, srcPaths ...string) error {
	zf, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer zf.Close()
	zw := zip.NewWriter(zf)
	for i := range srcPaths {
		srcInfo, err := os.Stat(srcPaths[i])
		if err != nil {
			return err
		}
		if srcInfo.IsDir() {
			// todo 压缩文件夹的内容
			continue
		}
		srcZw, err := zw.Create(srcInfo.Name())
		if err != nil {
			return err
		}
		srcFd, err := os.Open(srcPaths[i])
		if err != nil {
			return err
		}
		if _, err := io.Copy(srcZw, srcFd); err != nil {
			_ = srcFd.Close()
			return err
		}
		_ = srcFd.Close()
	}
	return zw.Close()
}
