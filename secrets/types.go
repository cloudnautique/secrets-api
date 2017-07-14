package secrets

import (
	"github.com/rancher/go-rancher/client"
	"github.com/rancher/secrets-api/pkg/aesutils"
)

type SecretCollection struct {
	client.Collection
	Data []Secret `json:"data,omitempty"`
}

type BulkSecretInput struct {
	client.Resource
	Data []*UnencryptedSecret `json:"data,omitempty"`
}

type BulkEncryptedSecret struct {
	client.Resource
	Data      []*EncryptedSecret `json:"data,omitempty"`
	RewrapKey string             `json:"rewrapKey,omitempty"`
}

type BulkRewrappedSecret struct {
	client.Resource
	Data []*RewrappedSecret `json:"data,omitempty"`
}

type UnencryptedSecret struct {
	client.Resource
	Backend     string `json:"backend"`
	ClearText   string `json:"clearText,omitempty"`
	ExternalURI string `json:"externalURI,omitempty"`
	KeyName     string `json:"keyName"`
}

type EncryptedSecret struct {
	client.Resource
	Backend             string `json:"backend"`
	CipherText          string `json:"cipherText,omitempty"`
	EncryptionAlgorithm string `json:"encryptionAglorigthm"`
	ExternalURI         string `json:"externalURI,omitempty"`
	HashAlgorithm       string `json:"hashAlgorithm"`
	KeyName             string `json:"keyName"`
	RewrapKey           string `json:"rewrapKey,omitempty"`
	Signature           string `json:"signature"`
	tmpKey              aesutils.AESKey
}

type RewrappedSecret struct {
	client.Resource
	RewrapText string `json:"rewrapText,omitempty"`
	SecretName string `json:"name,omitempty"`
}

type Secret struct {
	client.Resource
}

type EncryptedData struct {
	EncryptedKey        RSAEncryptedData `json:"encryptedKey,omitempty"`
	EncryptedText       string           `json:"encryptedText,omitempty"`
	EncryptionAlgorithm string           `json:"encryptionAlgorithm,omitempty"`
	HashAlgorithm       string           `json:"hashAlgorithm,omitempty"`
	Signature           string           `json:"signature,omitempty"`
}

type RSAEncryptedData struct {
	EncryptedText       string `json:"encryptedText,omitempty"`
	EncryptionAlgorithm string `json:"encryptionAlgorithm,omitempty"`
	HashAlgorithm       string `json:"hashAlgorithm,omitempty"`
}
