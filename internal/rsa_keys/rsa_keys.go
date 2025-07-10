package rsa_keys

import (
	"crypto/rsa"
	"encoding/pem"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/youmark/pkcs8"
	"os"
)

// LoadRSAPrivateKeyFromFile carga la clave privada RSA desde un archivo PEM.
func LoadRSAPrivateKeyFromFile(filePath string) (*rsa.PrivateKey, error) {
	signBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("leyendo el archivo privado de firma: %s", err)
	}

	block, _ := pem.Decode(signBytes)
	if err != nil {
		fmt.Printf("Error al decodificar el archivo privado de firma: %s", err)
		return nil, err
	}

	if block == nil {
		fmt.Printf("No se pudo decodificar el archivo privado de firma: %s", err)
		return nil, fmt.Errorf("no se pudo decodificar el archivo privado")
	}

	if block.Type == "ENCRYPTED PRIVATE KEY" {
		return pkcs8.ParsePKCS8PrivateKeyRSA(block.Bytes, []byte("123456789"))
	}

	return jwt.ParseRSAPrivateKeyFromPEM(signBytes)
}

// LoadRSAPublicKeyFromFile carga la clave pública RSA desde un archivo PEM.
func LoadRSAPublicKeyFromFile(filePath string) (*rsa.PublicKey, error) {
	keyBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return jwt.ParseRSAPublicKeyFromPEM(keyBytes)
}
