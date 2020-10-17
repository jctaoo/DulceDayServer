package auth

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

type EncryptionAdaptor interface {
	toHash(resource string) string
	verify(resource string, target string) bool
	generateVerificationCode() string
}

type EncryptionAdaptorImpl struct {
}

func NewEncryptionAdaptorImpl() *EncryptionAdaptorImpl {
	return &EncryptionAdaptorImpl{}
}

func (e EncryptionAdaptorImpl) toHash(resource string) string {
	bytes := []byte(resource)
	hash := md5.Sum(bytes)
	return fmt.Sprintf("%x", hash)
}

func (e EncryptionAdaptorImpl) verify(resource string, target string) bool {
	bytes := []byte(resource)
	hash := md5.Sum(bytes)
	hashString := fmt.Sprintf("%x", hash)
	return hashString == target
}

func (e EncryptionAdaptorImpl) generateVerificationCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}
