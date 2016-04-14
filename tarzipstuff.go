package icecubeapp

import (
	"archive/tar"
	"io"
	"log"
	"os"
)

func handleError(_e error) {
	if _e != nil {
		log.Fatal(_e)
	}
}

// TarGzWrite tars and gzips files
func TarGzWrite(_path string, tw *tar.Writer, fi os.FileInfo) {
	fr, err := os.Open(_path)
	handleError(err)
	defer fr.Close()

	h := new(tar.Header)
	h.Name = fi.Name()
	h.Size = fi.Size()
	h.Mode = int64(fi.Mode())
	h.ModTime = fi.ModTime()

	err = tw.WriteHeader(h)
	handleError(err)

	_, err = io.Copy(tw, fr)
	handleError(err)
}
