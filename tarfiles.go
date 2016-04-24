package icecubeapp

import (
	"archive/tar"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func tarit(w http.ResponseWriter, r *http.Request) {
	tarball := tar.NewWriter(w)
	defer tarball.Close()

	info, err := os.Stat("files")
	if err != nil {
		return
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base("files")
	}

	filepath.Walk("files", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, "files"))
		}

		if err := tarball.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(tarball, file)
		return err
	})
}
