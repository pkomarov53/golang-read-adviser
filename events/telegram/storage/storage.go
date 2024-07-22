package storage

import (
	"crypto/sha1"
	"fmt"
	"io"
	"read-adviser/lib/e"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(username string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	Username string

	// CreatedAt time.Time
}

func (p Page) Hash() (string, error) {
	hash := sha1.New()

	if _, err := io.WriteString(hash, p.URL); err != nil {
		return "", e.Wrapper("Unable to calculate hash", err)
	}

	if _, err := io.WriteString(hash, p.Username); err != nil {
		return "", e.Wrapper("Unable to calculate hash", err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
