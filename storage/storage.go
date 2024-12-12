package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"kemov/LinkKeeperBot/lib/e"
)

type Storage interface{
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

var ErrNoSavedPages = errors.New("Нет сохраненной страницы")

type Page struct {
	URL      string
	UserName string
}

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("Не могу вычислить хэш", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("Не могу вычислить хэш", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
