package static_storage

import "io"

type Service interface {
	Save(reader io.Reader, key string) error
	SaveImage(reader io.Reader, key string) error
	GetFileUrl(key string) (url string, err error)
}
