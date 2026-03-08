package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FilesystemStorage struct {
	basePath string
}

func NewFilesystemStorage(basePath string) *FilesystemStorage {
	return &FilesystemStorage{basePath: basePath}
}

func (fs *FilesystemStorage) Put(_ context.Context, path string, reader io.Reader) error {
	fullPath := filepath.Join(fs.basePath, path)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0o755); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, reader); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func (fs *FilesystemStorage) Get(_ context.Context, path string) (io.ReadCloser, error) {
	fullPath := filepath.Join(fs.basePath, path)

	file, err := os.Open(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("object not found: %s", path)
		}
		return nil, fmt.Errorf("open file: %w", err)
	}

	return file, nil
}

func (fs *FilesystemStorage) Delete(_ context.Context, path string) error {
	fullPath := filepath.Join(fs.basePath, path)

	if err := os.Remove(fullPath); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("delete file: %w", err)
	}

	return nil
}

func (fs *FilesystemStorage) Exists(_ context.Context, path string) (bool, error) {
	fullPath := filepath.Join(fs.basePath, path)

	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("stat file: %w", err)
	}

	return true, nil
}
