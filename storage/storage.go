package storage

import (
	"crypto/sha1"
	"fmt"
	"github.com/rshelekhov/read-it-later-bot/lib/e"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
	// CreatedAt time.Time
}

func (p *Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("error hashing url", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("error hashing username", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

type storageErr string

func (e storageErr) Error() string {
	return string(e)
}

const (
	ErrNoSavedPages storageErr = "no saved pages"
)
