package mysql

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var (
	Argon2ParamVar = &Argon2Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
)

func CreateUser(ctx context.Context, username, password, email string, argon2Params *Argon2Params) (int, error) {
	var users []User
	err := DB.WithContext(ctx).Where("username = ?", username).Find(&users).Error
	if err != nil {
		return 0, err
	}

	if len(users) != 0 {
		return 0, errors.New("user already exists")
	}

	passWord, err := generateFromPassword(password, argon2Params)
	if err != nil {
		return 0, err
	}

	err = DB.WithContext(ctx).Create(&User{Username: username, Password: passWord, Email: email}).Error
	if err != nil {
		return 0, err
	}

	err = DB.WithContext(ctx).Where("username = ?", username).Find(&users).Error
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errors.New("user does not exist")
	}
	return users[0].ID, nil
}

func CheckUser(ctx context.Context, username, password string) (int, error) {
	var user User
	err := DB.WithContext(ctx).Where("username = ?", username).Find(&user).Error
	if err != nil {
		return 0, err
	}
	if user.ID == 0 {
		return 0, errors.New("user does not exist")
	}

	match, err := comparePasswordAndHash(password, user.Password)
	if err != nil {
		return 0, err
	}
	if !match {
		return 0, errors.New("password does not match")
	}

	return user.ID, nil
}

func ModifyUsername(ctx context.Context, oldUsername, NewUsername string) error {
	var user User
	err := DB.WithContext(ctx).Where("username = ?", oldUsername).First(user).Update("username", NewUsername).Error
	if err != nil {
		return err
	}
	return nil
}

func ModifyPassword(ctx context.Context, username, newPassword string) error {
	var user User
	password, err := generateFromPassword(newPassword, Argon2ParamVar)
	if err != nil {
		return err
	}

	err = DB.WithContext(ctx).Where("username = ?", username).First(user).Update("password", password).Error
	if err != nil {
		return err
	}
	return nil
}

func ModifyEmail(ctx context.Context, username, newEmail string) error {
	var user User
	err := DB.WithContext(ctx).Where("username = ?", username).First(user).Update("email", newEmail).Error
	if err != nil {
		return err
	}
	return nil
}

func generateFromPassword(password string, argon2Params *Argon2Params) (encodedHash string, err error) {
	salt, err := generateRandomBytes(argon2Params.SaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, argon2Params.Iterations, argon2Params.Memory, argon2Params.Parallelism, argon2Params.KeyLength)

	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, argon2Params.Memory, argon2Params.Iterations, argon2Params.Parallelism, base64Salt, base64Hash)

	return encodedHash, nil
}

func generateRandomBytes(saltLength uint32) ([]byte, error) {
	buf := make([]byte, saltLength)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func comparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	argon2Params, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	inputHash := argon2.IDKey([]byte(password), salt, argon2Params.Iterations, argon2Params.Memory, argon2Params.Parallelism, argon2Params.KeyLength)

	if subtle.ConstantTimeCompare(hash, inputHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (argon2Params *Argon2Params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, fmt.Errorf("invalid hash format")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, fmt.Errorf("incompatible hash version")
	}

	argon2Params = &Argon2Params{}
	if _, err := fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &argon2Params.Memory, &argon2Params.Iterations, &argon2Params.Parallelism); err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	argon2Params.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	argon2Params.KeyLength = uint32(len(hash))

	return argon2Params, salt, hash, nil
}
