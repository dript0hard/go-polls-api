package password

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	SALT_SIZE         int    = 32
	KEY_LEN           int    = 128
	PBKDF2_ITERATIONS int    = 15000
	ALGORITHM_HASH    string = "pbkdf2_sha512"

	PASSWORD_MIN_LENGTH uint8 = 8
)

type PasswordHasher interface {
	HashPassword(password string) *HashResult
	ValidatePassword(password string, hash *HashResult) bool
}

type HashResult struct {
	Algo       string
	Iterations int
	Salt       string
	CipherText string
}

// Hash format that will be saved in the database.
func (h HashResult) String() string {
	return fmt.Sprintf("%s$%s$%s$%s", h.Algo, strconv.Itoa(h.Iterations), h.Salt, h.CipherText)
}

// Deserialize it from the database string.
func NewHashResult(hashString string) *HashResult {
	split := strings.Split(hashString, "$")
	iterations, _ := strconv.Atoi(split[1])
	return &HashResult{
		Algo:       split[0],
		Iterations: iterations,
		Salt:       split[2],
		CipherText: split[3],
	}
}

type PBKDF2Password struct {
	Algo       string
	HashFunc   func() hash.Hash
	SaltSize   int
	KeyLen     int
	Iterations int
}

func NewPBKDF2Password(algoRep string, hashFunc func() hash.Hash, saltSize int, keyLen int, iter int) *PBKDF2Password {
	return &PBKDF2Password{
		Algo:       algoRep,
		HashFunc:   hashFunc,
		SaltSize:   saltSize,
		KeyLen:     keyLen,
		Iterations: iter,
	}
}

//What is going to be used for this project.split
func NewPBKDF2PasswordSha512() *PBKDF2Password {
	return NewPBKDF2Password(ALGORITHM_HASH, sha512.New, SALT_SIZE, KEY_LEN, PBKDF2_ITERATIONS)
}

// This is how the salts are gonna be generated might change later.
func genSalt() []byte {
	saltBytes := make([]byte, SALT_SIZE)
	rand.Read(saltBytes)
	return saltBytes
}

func (p *PBKDF2Password) HashPassword(password string) *HashResult {
	// This gets a new salt.
	salt := genSalt()
	// We create the pbkdf2 key.
	key := pbkdf2.Key([]byte(password), salt, p.Iterations, p.KeyLen, p.HashFunc)
	//  Its base64
	cipher := base64.StdEncoding.EncodeToString(key)

	// hash result from that.
	saltString := base64.StdEncoding.EncodeToString(salt)
	return &HashResult{Algo: p.Algo,
		Iterations: p.Iterations,
		Salt:       saltString,
		CipherText: cipher}
}

func (p PBKDF2Password) ValidatePassword(password string, hash *HashResult) bool {
	// Use the existing salt from the saved password.
	saltBytes, _ := base64.StdEncoding.DecodeString(hash.Salt)

	// Hash the given password with the salt and the hash function of choice.
	key := pbkdf2.Key([]byte(password), saltBytes, p.Iterations, p.KeyLen, p.HashFunc)
	newCipher := base64.StdEncoding.EncodeToString(key)

	return newCipher == hash.CipherText

}
