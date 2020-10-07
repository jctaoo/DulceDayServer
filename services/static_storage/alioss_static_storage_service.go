package static_storage

import (
	"DulceDayServer/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type AliOSSStaticStorageService struct {
	bucket *oss.Bucket
}


func NewAliOSSStaticStorageService(bucket *oss.Bucket) *AliOSSStaticStorageService {
	return &AliOSSStaticStorageService{bucket: bucket}
}

func (a AliOSSStaticStorageService) Save(reader io.Reader, key string) error {
	return a.bucket.PutObject(key, reader)
}

func (a AliOSSStaticStorageService) SaveImage(reader io.Reader, key string) error {
	option := oss.ContentType("image/jpg")
	return a.bucket.PutObject(key, reader, option)
}

func (a AliOSSStaticStorageService) GetFileUrl(key string) (url string, err error) {
	return a.bucket.SignURL(key, oss.HTTPGet, int64(config.SiteConfig.AliOssStaticStorageConfig.ResourceExpiresSec))
}

