package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	// Configuration
	viper.AddConfigPath("./config")
	viper.SetConfigName("main")
	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

type SiteHeader struct {
	Title		string
	Keywords	string
	Description	string
}

func GetSiteHeader(moduleName string) SiteHeader {
	var header SiteHeader
	path := "site." + moduleName + ".head"
	err := viper.UnmarshalKey(path, &header)
	if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return header
}

func (h *SiteHeader) SetTitle(title string) {
	h.Title = title + h.Title
}