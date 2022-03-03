package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	ConnectionStringDb = ""
	Port               = 0
)

// Load Inicializar as variáveis de ambiente
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv(apiPort))
	if err != nil {
		Port = 9000
	}

	ConnectionStringDb = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		os.Getenv(dbUser),
		os.Getenv(dbPassword),
		os.Getenv(dbDatabase))
}
