package affiliate_tracker

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/oswaldom-code/affiliate-tracker/src/adapters/repository"
	"github.com/oswaldom-code/affiliate-tracker/src/services/ports"
)

type Tracker interface {
	ProcessInputUrl(referralLink ReferralLink) string
	encrypter(data string) string
	decrypt()
}

type tracker struct {
	r ports.Repository
}

func NewTrackerService() Tracker {
	return &tracker{r: repository.NewRepository()}
}

func (t *tracker) ProcessInputUrl(referralLink ReferralLink) string {
	return t.encrypter(referralLink.ToString())
}

func (t *tracker) encrypter(data string) string {
	// TODO: get key from config
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	// Creamos un bloque de cifrado utilizando la clave
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// Creamos un objeto OFB con el bloque de cifrado y el vector de inicialización
	ciphertext := make([]byte, aes.BlockSize+len([]byte(data)))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	// Creamos un objeto OFB con el bloque de cifrado y el vector de inicialización
	stream := cipher.NewOFB(block, iv)
	// Ciframos el texto utilizando el stream
	buf := make([]byte, len("Hola mundo"))
	stream.XORKeyStream(buf, []byte("Hola mundo"))

	// Codificamos el texto cifrado en formato base64
	encoded := base64.StdEncoding.EncodeToString(buf)
	// Imprimimos el texto cifrado en formato base64
	fmt.Println("Texto cifrado:", encoded)
	return encoded
}

func (t *tracker) decrypt() {

}
