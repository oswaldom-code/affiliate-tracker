package affiliate_tracker

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"

	"github.com/oswaldom-code/affiliate-tracker/src/adapters/repository"
	"github.com/oswaldom-code/affiliate-tracker/src/services/affiliate_tracker_services/types"
	"github.com/oswaldom-code/affiliate-tracker/src/services/ports"
)

type Tracker interface {
	GenerateIdentifier(referralLink types.ReferralLink) (string, error)
	IdentifierDecoding(encryptedData string, referred types.Referred) (string, error)
	encrypter(data string) (string, error)
	decrypt(encryptedDataStr string) (string, error)
}

type tracker struct {
	r ports.Repository
}

// service entry point
func NewTrackerService() Tracker {
	return &tracker{r: repository.NewRepository()}
}

// IdentifierDecoding generates a unique identifier for the given reference (agent, url)
func (t *tracker) IdentifierDecoding(encryptedData string, referred types.Referred) (string, error) {
	// replace spaces
	encryptedData = strings.Replace(encryptedData, " ", "+", -1)
	result, err := t.decrypt(encryptedData)
	if err != nil {
		return "", err
	}
	// assign to object referred the already decrypted url and agent values
	// This object already contains the information of the HTTP request headers
	referred.Agent = strings.Split(result, ":::")[0]
	referred.Url = strings.Split(result, ":::")[1]
	// TODO Persist the information on DB
	// returns the url of the job offer
	return referred.Url, nil
}

// GenerateIdentifier generates a unique identifier for the given reference (agent, url)
func (t *tracker) GenerateIdentifier(referralLink types.ReferralLink) (string, error) {
	return t.encrypter(referralLink.ToString())
}

// encrypter implements AES Crypt to encode text
func (t *tracker) decrypt(encryptedDataStr string) (string, error) {
	// TODO: get key from config
	key := []byte("0123456789abcdef")
	encryptedDataByte, _ := base64.StdEncoding.DecodeString(encryptedDataStr)
	originalData, err := AesDecrypt(encryptedDataByte, key)
	if err != nil {
		return "", err
	}
	return string(originalData), nil
}

// encrypter implements AES Crypt to encode text
func (t *tracker) encrypter(data string) (string, error) {
	// TODO: get key from config
	key := []byte("0123456789abcdef")
	result, err := AesEncrypt([]byte(data), key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

/*
	Helper functions to encrypt and decrypt data
*/

func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS7UnPadding(originalData []byte) []byte {
	length := len(originalData)
	unpadding := int(originalData[length-1])
	return originalData[:(length - unpadding)]
}

func AesEncrypt(originalData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	originalData = PKCS7Padding(originalData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encryptedData := make([]byte, len(originalData))
	blockMode.CryptBlocks(encryptedData, originalData)
	return encryptedData, nil
}

func AesDecrypt(encryptedData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	originalData := make([]byte, len(encryptedData))
	blockMode.CryptBlocks(originalData, encryptedData)
	originalData = PKCS7UnPadding(originalData)
	return originalData, nil
}

/*
	END Helper functions
*/
