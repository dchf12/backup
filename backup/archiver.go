package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type Archiver interface {
	DestFmt() func(int64) string
	Archive(src, dest string) error
}

type zipper struct{}

var ZIP Archiver = (*zipper)(nil)

func (z *zipper) DestFmt() func(int64) string {
	return func(i int64) string {
		return "%d.zip"
	}
}

// Archive creates a zip file of the specified directory and its contents
// and places it at the specified destination
// panic if it's not exist
func (z *zipper) Archive(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0o777); err != nil {
		return err
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer func() {
		if err := out.Close(); err != nil {
			panic(err)
		}
	}()
	w := zip.NewWriter(out)
	defer func() {
		if err := w.Close(); err != nil {
			panic(err)
		}
	}()
	return filepath.WalkDir(src, func(path string, info os.DirEntry, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		f, err := w.Create(path)
		if err != nil {
			return err
		}
		_, _ = io.Copy(f, in)
		return nil
	})

}
