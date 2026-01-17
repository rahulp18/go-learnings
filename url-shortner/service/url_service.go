package service

import (
	"math/rand"
	"time"

	"github.com/rahulp18/url-shortner/store"
)

type URLService struct {
	store store.URLStore
}

func NewURLService(store store.URLStore) *URLService {
	rand.NewSource(time.Now().UnixMicro())
	return &URLService{store: store}
}
func (s *URLService) Shorten(longURL string) string {
	code := generateCode()
	s.store.Save(code, longURL)
	return code
}
func (s *URLService) Resolve(code string) (string, bool) {
	return s.store.Get(code)
}
func generateCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
