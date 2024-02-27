package file

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func Zip(src, dest string) error {
	bd, err := os.Create(filepath.Join(dest, filepath.Base(src)+".zip"))
	if err != nil {
		return err
	}
	defer bd.Close()
	writer := zip.NewWriter(bd)
	defer writer.Close()
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(src, path)

		if err != nil {
			return err
		}
		zipFile, err := writer.Create(relPath)
		if err != nil {
			return err
		}
		fsFile, err := os.Open(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(zipFile, fsFile)
		return nil
	})
}
