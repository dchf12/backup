package backup

import (
	"os"
	"testing"
)

func TestDirHash(t *testing.T) {
	t.Parallel()
	tmpDir := t.TempDir()
	if err := os.WriteFile(tmpDir+"/file1", []byte("file1"), 0644); err != nil {
		t.Fatal(err)
	}
	hash1, err := DirHash(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(tmpDir+"/file1", []byte("file1 changed"), 0644); err != nil {
		t.Fatal(err)
	}
	hash2, err := DirHash(tmpDir)
	if err != nil {
		t.Fatal(err)
	}
	if hash1 == hash2 {
		t.Fatal("hashes should be different")
	}
}
