package store

type URLStore interface {
	Save(code string, longUrl string)
	Get(code string) (string, bool)
}
