package backup

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"path/filepath"
)

func DirHash(path string) (string, error) {
	hash := md5.New()
	if err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		info, err := d.Info()
		if err != nil {
			return err
		}
		fmt.Fprintf(hash, "%v", path)
		fmt.Fprintf(hash, "%v", d.IsDir())
		fmt.Fprintf(hash, "%v", d.Type())
		fmt.Fprintf(hash, "%v", d.Name())
		fmt.Fprintf(hash, "%v", info.ModTime())
		fmt.Fprintf(hash, "%v", info.Size())
		return nil
	}); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
