package app

import (
	"fmt"
	"github.com/labstack/gommon/color"
	conf "github.com/spf13/viper"
)

func setupConfig(base_path string) (err error) {
	conf.SetConfigName("config")
	conf.AddConfigPath(base_path)
	if err := conf.ReadInConfig(); err != nil {
		return fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow("config file not found."))
	}

	return err
}
