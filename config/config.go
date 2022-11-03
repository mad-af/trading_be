package config

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var Env struct {
	HttpPort string `json:"HTTP_PORT"`
	ApiUrl   string `json:"API_URL"`
	// POSTGRESQL
	PostgreHost     string `json:"POSTGRES_HOST"`
	PostgreUser     string `json:"POSTGRES_USER"`
	PostgrePassword string `json:"POSTGRES_PASSWORD"`
	PostgreDbName   string `json:"POSTGRES_DBNAME"`
	PostgrePort     string `json:"POSTGRES_PORT"`
	PostgreSslMode  string `json:"POSTGRES_SSLMODE"`
	// KEY
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	rootApp := strings.TrimSuffix(path, "/config")
	fmt.Println(rootApp)
	os.Setenv("APP_PATH", rootApp)

	var myEnv map[string]string
	myEnv, err = godotenv.Read()
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(myEnv)
	json.Unmarshal(b, &Env)

	signBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/key/private.key", os.Getenv("APP_PATH")))
	if err != nil {
		panic("Error when load private key. " + err.Error())
	}
	Env.PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic("Error when load private key. " + err.Error())
	}

	verifyBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/key/public.pem", os.Getenv("APP_PATH")))
	if err != nil {
		panic("Error when load public key. " + err.Error())
	}
	Env.PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic("Error when load public key. " + err.Error())
	}
}
