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
	IdentifierDecoding(encryptedData string, referred types.ReferredDTO) (string, error)
	encrypter(data string) (string, error)
	decrypt(encryptedDataStr string) (string, error)
	GetAll() ([]types.ReferredDTO, error)
	GetById(id int64) (types.ReferredDTO, error)
	GetByAgentId(agentId string) ([]types.ReferredDTO, error)
}

type tracker struct {
	r ports.Repository
}

// service entry point
func NewTrackerService() Tracker {
	return &tracker{r: repository.NewRepository()}
}

// IdentifierDecoding generates a unique identifier for the given reference (agent, url)
func (t *tracker) IdentifierDecoding(encryptedData string, referredDto types.ReferredDTO) (string, error) {
	// replace spaces
	encryptedData = strings.Replace(encryptedData, " ", "+", -1)
	result, err := t.decrypt(encryptedData)
	if err != nil {
		return "", err
	}
	// assign to object referred the already decrypted url and agent values
	// This object already contains the information of the HTTP request headers
	referredDto.AgentId = strings.Split(result, ":::")[0]
	referredDto.JobUrl = strings.Split(result, ":::")[1]
	_ = t.r.Save(referredDto.ToModel())
	// returns the url of the job offer
	return referredDto.JobUrl, nil
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

func (t *tracker) GetAll() ([]types.ReferredDTO, error) {
	referedsList, err := t.r.GetAll()
	referedDTO := types.ReferredDTO{}
	return referedDTO.ReferredModelsToReferredDTOs(referedsList), err
}

func (t *tracker) GetById(id int64) (types.ReferredDTO, error) {
	refered, err := t.r.GetById(id)
	referedDTO := types.ReferredDTO{}
	referedDTO.ReferredModelToReferredDTO(refered)
	// parse model data to application types
	return referedDTO, err
}

func (t *tracker) GetByAgentId(agentId string) ([]types.ReferredDTO, error) {
	referedsList, err := t.r.GetByAgentId(strings.ToUpper(agentId))
	referedDTO := types.ReferredDTO{}
	return referedDTO.ReferredModelsToReferredDTOs(referedsList), err
}
