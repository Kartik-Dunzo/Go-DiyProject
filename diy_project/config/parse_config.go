package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var Config_parse Config

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
	Product struct {
		TopProductCount        int
		MaxProductsCart        int
		MaxProductQuantitySold int
	}
	Server struct {
		IP   string
		Port string
	}
}

func LoadConfiguration(file string) {

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&Config_parse)
	fmt.Println(Config_parse)

}
