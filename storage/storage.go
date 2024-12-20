package storage

import (
	"kemov/LinkKeeperBot/lib/e"
	
	"crypto/sha1"
	"context"
	"errors"
	"fmt"
	"io"
)

type Storage interface {
	Save(ctx context.Context, p *Page) error
	PickRandom(ctx context.Context, userName string) (*Page, error)
	Remove(ctx context.Context, p *Page) error
	IsExists(ctx context.Context, p *Page) (bool, error)
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
