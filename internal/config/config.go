package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)


type HTTPServer struct{
	Addr string
}

//env-default:"production"

type Config struct{
	Env string `yaml:"env" env:"ENV" env-required:"true" `
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer string `yaml:"http_server"`
}


// in this function we are storing our file data at config path in form of config struct

func MustLoad() *Config{
	var configPath string
	
	configPath = os.Getenv("CONFIG_PATH")
	
	if configPath ==""{
		// if it is an empty string then we will check if it will passed in arguments duing running of the program or not
		flags:= flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath ==""{
			log.Fatal("Config path is not set")
		}
	}
    // checking if file is present on the given path

	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("Config path doesn't exist %s", configPath)
	}

	var cfg Config

	// reading file on configpath according to the config struct

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil{
		log.Fatalf("can not read config file: %s", err.Error())
	}

	// this function is returning the address of the config
	return &cfg
}