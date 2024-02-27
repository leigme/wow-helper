package file

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(zipFile, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	destDir = filepath.Join(destDir, strings.TrimSuffix(filepath.Base(zipFile), filepath.Ext(zipFile)))

	for _, f := range zipReader.File {
		err = unzip(f, destDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func unzip(file *zip.File, destDir string) error {
	path := filepath.Join(destDir, file.Name)

	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	if !file.FileInfo().IsDir() {
		fr, err := file.Open()
		defer fr.Close()
		if err != nil {
			return err
		}
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		defer fw.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(fw, fr)
		return err
	}
	return nil
}
