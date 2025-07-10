package env

import (
	"crypto/rsa"
	"dbu-api/internal/rsa_keys"
	"encoding/json"
	"log"
	"os"
	"sync"
)

var (
	once   sync.Once
	config = &Configuration{}
)

type KeyStore struct {
	PrivateKey *rsa.PrivateKey `json:"private_key"`
	PublicKey  *rsa.PublicKey  `json:"public_key"`
}

type Configuration struct {
	App        App        `json:"app"`
	Router     Router     `json:"router"`
	Log        Log        `json:"log"`
	KeyConf    Key        `json:"key"`
	DB         DB         `json:"db"`
	Key        KeyStore   `json:"key_store"`
	IA         IAConfig   `json:"ia"`
	Department Department `json:"department"`
}

type App struct {
	ServiceName string `json:"service_name"`
	Port        int    `json:"port"`
}

type Router struct {
	AllowedDomains string `json:"allowed_domains"`
	LoggerHttp     bool   `json:"logger_http"`
}

type Log struct {
	Path           string `json:"path"`
	ReviewInterval int    `json:"review_interval"`
}
type Key struct {
	Private string `json:"private"`
	Public  string `json:"public"`
}

type IAConfig struct {
	UrlApi string `json:"url_api"`
	Model  string `json:"model"`
	Token  string `json:"token"`
}

type DB struct {
	Engine   string `json:"engine"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Instance string `json:"instance"`
	IsSecure bool   `json:"is_secure"`
	SSLMode  string `json:"ssl_mode"`
}

type Department struct {
	RequirementId int `json:"requirement_id"`
}

func NewConfiguration() *Configuration {
	fromFile()
	return config
}

// LoadConfiguration lee el archivo configuration.json
// y lo carga en un objeto de la estructura Configuration
func fromFile() {
	once.Do(func() {
		b, err := os.ReadFile("config.json")
		if err != nil {
			log.Fatalf("no se pudo leer el archivo de configuración: %s", err.Error())
		}

		err = json.Unmarshal(b, config)
		if err != nil {
			log.Fatalf("no se pudo parsear el archivo de configuración: %s", err.Error())
		}

		if config.DB.Engine == "" {
			log.Fatal("no se ha cargado la información de configuración")
		}

		privateKey, err := rsa_keys.LoadRSAPrivateKeyFromFile(config.KeyConf.Private)
		if err != nil {
			log.Fatalf("no se ha cargado la la llave privada %s", err)
		}

		publicKey, err := rsa_keys.LoadRSAPublicKeyFromFile(config.KeyConf.Public)
		if err != nil {
			log.Fatalf("no se ha cargado la la llave publica %s", err)
		}
		config.Key.PrivateKey = privateKey
		config.Key.PublicKey = publicKey
	})
}
