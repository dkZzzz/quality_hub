package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	err := generateKeyPair()
	if err != nil {
		panic(err)
	}

	privateBytes, err := os.ReadFile("../../private.pem")
	// fmt.Println(privateBytes)
	if err != nil {
		panic(err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		panic(err)
	}

	publicBytes, err := os.ReadFile("../../public.pem")
	// fmt.Println(publicBytes)
	if err != nil {
		panic(err)
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		panic(err)
	}
}

// 生成JWT
func GenerateJWT(username string, userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims = jwt.MapClaims{
		"username": username,
		"userid":   userID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("token is invalid")
	}
	return nil
}

func generateKeyPair() error {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// 将私钥保存为PEM格式
	privateKeyFile, err := os.Create("../../private.pem")
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return err
	}

	// 生成公钥
	publicKey := &privateKey.PublicKey

	// 将公钥保存为PEM格式
	publicKeyFile, err := os.Create("../../public.pem")
	if err != nil {
		return err
	}
	defer publicKeyFile.Close()

	publicKeyPEM, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicKeyBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyPEM,
	}
	if err := pem.Encode(publicKeyFile, publicKeyBlock); err != nil {
		return err
	}

	return nil
}
