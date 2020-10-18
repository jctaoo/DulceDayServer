package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type tomlConfig struct {
	AppConfigs                 map[string]appConfig
	DataBaseConfigs            map[string]dataBaseConfig
	CacheConfigs               map[string]cacheConfig
	AliOssStaticStorageConfigs map[string]aliOssStaticStorageConfig
}

type appConfig struct {
	AppName                      string
	AppAddress                   string
	AuthTokenExpiresTime         int64
	VerificationTokenExpiresTime int64
	AuthTokenSecret              string
	DefaultDeviceName            string
	AvatarSizeMB                 int
}

type dataBaseConfig struct {
	Host            string
	Port            string
	Collection      string
	User            string
	Password        string
	QueryParameters string
}

type cacheConfig struct {
	Host string
	Port string
	DB   int
	BlackListName,
	RevokeTokenListName,
	IPBlackListName,
	InActiveTokenListName string
	VerificationCodeListName string
}

type aliOssStaticStorageConfig struct {
	Endpoint           string
	AccessKeyId        string
	AccessKeySecret    string
	BucketName         string
	ResourceExpiresSec int
}

func ReadConfigOrExit(tomlPath string, isProduction bool) {
	var configStr string
	{
		file, err := os.Open(tomlPath)
		content, err := ioutil.ReadAll(file)
		configStr = string(content)
		err = file.Close()
		if err != nil {
			fmt.Println("Some Error Occurred When Read Config File: ", err)
		}
	}

	var config tomlConfig
	if _, err := toml.Decode(configStr, &config); err != nil {
		logrus.Error("解析配置时发生错误: ", err)
		os.Exit(-1)
		return
	}
	if isProduction {
		SiteConfig.AppName = config.AppConfigs["production"].AppName
		SiteConfig.AppAddress = config.AppConfigs["production"].AppAddress
		SiteConfig.AuthTokenSecret = config.AppConfigs["production"].AuthTokenSecret
		SiteConfig.AuthTokenExpiresTime = config.AppConfigs["production"].AuthTokenExpiresTime
		SiteConfig.VerificationTokenExpiresTime = config.AppConfigs["production"].VerificationTokenExpiresTime
		SiteConfig.DefaultDeviceName = config.AppConfigs["production"].DefaultDeviceName
		SiteConfig.AvatarSizeMB = config.AppConfigs["production"].AvatarSizeMB
		SiteConfig.DataBaseConfig.Host = config.DataBaseConfigs["production"].Host
		SiteConfig.DataBaseConfig.Port = config.DataBaseConfigs["production"].Port
		SiteConfig.DataBaseConfig.Collection = config.DataBaseConfigs["production"].Collection
		SiteConfig.DataBaseConfig.User = config.DataBaseConfigs["production"].User
		SiteConfig.DataBaseConfig.Password = config.DataBaseConfigs["production"].Password
		SiteConfig.DataBaseConfig.QueryParameters = config.DataBaseConfigs["production"].QueryParameters
		SiteConfig.CacheConfig.Host = config.CacheConfigs["production"].Host
		SiteConfig.CacheConfig.Port = config.CacheConfigs["production"].Port
		SiteConfig.CacheConfig.DB = config.CacheConfigs["production"].DB
		SiteConfig.CacheConfig.BlackListName = config.CacheConfigs["production"].BlackListName
		SiteConfig.CacheConfig.RevokeTokenListName = config.CacheConfigs["production"].RevokeTokenListName
		SiteConfig.CacheConfig.IPBlackListName = config.CacheConfigs["production"].IPBlackListName
		SiteConfig.CacheConfig.InActiveTokenListName = config.CacheConfigs["production"].InActiveTokenListName
		SiteConfig.CacheConfig.VerificationCodeListName = config.CacheConfigs["production"].VerificationCodeListName
		SiteConfig.AliOssStaticStorageConfig.AccessKeyId = config.AliOssStaticStorageConfigs["production"].AccessKeyId
		SiteConfig.AliOssStaticStorageConfig.AccessKeySecret = config.AliOssStaticStorageConfigs["production"].AccessKeySecret
		SiteConfig.AliOssStaticStorageConfig.Endpoint = config.AliOssStaticStorageConfigs["production"].Endpoint
		SiteConfig.AliOssStaticStorageConfig.BucketName = config.AliOssStaticStorageConfigs["production"].BucketName
		SiteConfig.AliOssStaticStorageConfig.ResourceExpiresSec = config.AliOssStaticStorageConfigs["production"].ResourceExpiresSec
	} else {
		SiteConfig.AppName = config.AppConfigs["development"].AppName
		SiteConfig.AppAddress = config.AppConfigs["development"].AppAddress
		SiteConfig.AuthTokenSecret = config.AppConfigs["development"].AuthTokenSecret
		SiteConfig.AuthTokenExpiresTime = config.AppConfigs["development"].AuthTokenExpiresTime
		SiteConfig.VerificationTokenExpiresTime = config.AppConfigs["development"].VerificationTokenExpiresTime
		SiteConfig.DefaultDeviceName = config.AppConfigs["development"].DefaultDeviceName
		SiteConfig.AvatarSizeMB = config.AppConfigs["development"].AvatarSizeMB
		SiteConfig.DataBaseConfig.Host = config.DataBaseConfigs["development"].Host
		SiteConfig.DataBaseConfig.Port = config.DataBaseConfigs["development"].Port
		SiteConfig.DataBaseConfig.Collection = config.DataBaseConfigs["development"].Collection
		SiteConfig.DataBaseConfig.User = config.DataBaseConfigs["development"].User
		SiteConfig.DataBaseConfig.Password = config.DataBaseConfigs["development"].Password
		SiteConfig.DataBaseConfig.QueryParameters = config.DataBaseConfigs["development"].QueryParameters
		SiteConfig.CacheConfig.Host = config.CacheConfigs["development"].Host
		SiteConfig.CacheConfig.Port = config.CacheConfigs["development"].Port
		SiteConfig.CacheConfig.DB = config.CacheConfigs["development"].DB
		SiteConfig.CacheConfig.BlackListName = config.CacheConfigs["development"].BlackListName
		SiteConfig.CacheConfig.RevokeTokenListName = config.CacheConfigs["development"].RevokeTokenListName
		SiteConfig.CacheConfig.IPBlackListName = config.CacheConfigs["development"].IPBlackListName
		SiteConfig.CacheConfig.InActiveTokenListName = config.CacheConfigs["development"].InActiveTokenListName
		SiteConfig.CacheConfig.VerificationCodeListName = config.CacheConfigs["development"].VerificationCodeListName
		SiteConfig.AliOssStaticStorageConfig.AccessKeyId = config.AliOssStaticStorageConfigs["development"].AccessKeyId
		SiteConfig.AliOssStaticStorageConfig.AccessKeySecret = config.AliOssStaticStorageConfigs["development"].AccessKeySecret
		SiteConfig.AliOssStaticStorageConfig.Endpoint = config.AliOssStaticStorageConfigs["development"].Endpoint
		SiteConfig.AliOssStaticStorageConfig.BucketName = config.AliOssStaticStorageConfigs["development"].BucketName
		SiteConfig.AliOssStaticStorageConfig.ResourceExpiresSec = config.AliOssStaticStorageConfigs["development"].ResourceExpiresSec
	}
}
