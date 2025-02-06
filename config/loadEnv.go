package config

import "github.com/spf13/viper"

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	DevDBHost         string `mapstructure:"DEV_POSTGRES_HOST"`
	DevDBUserName     string `mapstructure:"DEV_POSTGRES_USER"`
	DevDBUserPassword string `mapstructure:"DEV_POSTGRES_PASSWORD"`
	DevDBName         string `mapstructure:"DEV_POSTGRES_DB"`
	DevDBPort         string `mapstructure:"DEV_POSTGRES_PORT"`
	AppEnv            string `mapstructure:"APP_ENV"`

	CloudinaryCloudName string `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	CloudinaryApiKey string `mapstructure:"CLOUDINARY_API_KEY"`
	CloudinarySecretKey string `mapstructure:"CLOUDINARY_SECRET_KEY"`

}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
