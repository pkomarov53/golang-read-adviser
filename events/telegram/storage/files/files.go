package files

import (
	"encoding/gob"
	"os"
	"path/filepath"
	"read-adviser/events/telegram/storage"
	"read-adviser/lib/e"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIfErr("Unable to save", err) }()

	fPath := filepath.Join(s.basePath, page.Username)

	if err := os.MkdirAll(fPath, defaultPerm); err != nil {
		return err
	}

	fName, err := fileName(page)
	if err != nil {
		return err
	}

	fPath = filepath.Join(fPath, fName)

	file, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	if err := gob.NewEncoder(file).Encode(page); err != nil {
		return err
	}

	return nil
}

func (s Storage) PickRandom(username string) (page *storage.Page, err error) {

}

func fileName(p *storage.Page) (string, error) {
	return p.Hash()
}
