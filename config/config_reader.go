package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type tomlConfig struct {
	AppConfigs map[string]appConfig
	DataBaseConfigs map[string]dataBaseConfig
	CacheConfigs map[string]cacheConfig
}

type appConfig struct {
	AppName string
	AppAddress string
	AuthTokenExpiresTime int64
	AuthTokenSecret string
}

type dataBaseConfig struct {
	Host string
	Port string
	Collection string
	User       string
	Password   string
	QueryParameters string
}

type cacheConfig struct {
	Host string
	Port string
	BlackListName string
}

func ReadConfig(tomlPath string, isProduction bool) {
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
		// TODO
	}
	if isProduction {
		SiteConfig.AppName = config.AppConfigs["production"].AppName
		SiteConfig.AppAddress = config.AppConfigs["production"].AppAddress
		SiteConfig.AuthTokenSecret = config.AppConfigs["production"].AuthTokenSecret
		SiteConfig.AuthTokenSecret = config.AppConfigs["production"].AuthTokenSecret
		SiteConfig.DataBaseConfig.Host = config.DataBaseConfigs["production"].Host
		SiteConfig.DataBaseConfig.Port = config.DataBaseConfigs["production"].Port
		SiteConfig.DataBaseConfig.Collection = config.DataBaseConfigs["production"].Collection
		SiteConfig.DataBaseConfig.User = config.DataBaseConfigs["production"].User
		SiteConfig.DataBaseConfig.Password = config.DataBaseConfigs["production"].Password
		SiteConfig.DataBaseConfig.QueryParameters = config.DataBaseConfigs["production"].QueryParameters
		SiteConfig.CacheConfig.Host = config.CacheConfigs["production"].Host
		SiteConfig.CacheConfig.Port = config.CacheConfigs["production"].Port
		SiteConfig.CacheConfig.BlackListName = config.CacheConfigs["production"].BlackListName
	} else {
		SiteConfig.AppName = config.AppConfigs["development"].AppName
		SiteConfig.AppAddress = config.AppConfigs["development"].AppAddress
		SiteConfig.AuthTokenSecret = config.AppConfigs["development"].AuthTokenSecret
		SiteConfig.AuthTokenSecret = config.AppConfigs["development"].AuthTokenSecret
		SiteConfig.DataBaseConfig.Host = config.DataBaseConfigs["development"].Host
		SiteConfig.DataBaseConfig.Port = config.DataBaseConfigs["development"].Port
		SiteConfig.DataBaseConfig.Collection = config.DataBaseConfigs["development"].Collection
		SiteConfig.DataBaseConfig.User = config.DataBaseConfigs["development"].User
		SiteConfig.DataBaseConfig.Password = config.DataBaseConfigs["development"].Password
		SiteConfig.DataBaseConfig.QueryParameters = config.DataBaseConfigs["development"].QueryParameters
		SiteConfig.CacheConfig.Host = config.CacheConfigs["development"].Host
		SiteConfig.CacheConfig.Port = config.CacheConfigs["development"].Port
		SiteConfig.CacheConfig.BlackListName = config.CacheConfigs["development"].BlackListName
	}
}