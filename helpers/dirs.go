package helpers

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

func CreateDirectories(prjPath string, dirs []string, perm os.FileMode) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(filepath.Join(prjPath, dir), perm); err != nil {
			return err
		}
	}
	return nil
}

func CopyEmbeddedFiles(efs embed.FS, srcDir, destDir string) error {
	return fs.WalkDir(efs, srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}

		data, err := efs.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(destPath, data, os.ModePerm)
	})
}
