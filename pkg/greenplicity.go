package application

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

	fil "github.com/filecoin-project/go-fil-markets"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
)

// FilecoinClient is a struct for interacting with the Filecoin network
type FilecoinClient struct {
	client *fil.Client
}

// NewFilecoinClient returns a new instance of the FilecoinClient
func NewFilecoinClient(apiAddr string) (*FilecoinClient, error) {
	client, err := fil.NewClient(apiAddr)
	if err != nil {
		return nil, err
	}
	fmt.Printf("API ADDR", apiAddr)
	return &FilecoinClient{client: client}, nil
}

// SignMessage signs a message with the provided private key
func (fc *FilecoinClient) SignMessage(privateKey crypto.PrivateKey, message []byte) ([]byte, error) {
	return crypto.Sign(privateKey, message)
}

// VerifySignature verifies a signature on a message
func (fc *FilecoinClient) VerifySignature(publicKey crypto.PublicKey, message []byte, signature []byte) error {
	ok, err := crypto.VerifySignature(publicKey, message, signature)

	if err != nil {
		return err
	}

	if !ok {
		return errors.New("signature verification failed")
	}

	return nil
}

// StoreData stores data on the Filecoin network
func (fc *FilecoinClient) StoreData(data []byte) (cid string, err error) {
	c, err := fc.client.Client.Upload(context.Background(), data)
	if err != nil {
		return "", err
	}
	return c.String(), nil
}

// RetrieveData retrieves data from the Filecoin network
func (fc *FilecoinClient) RetrieveData(cid string) ([]byte, error) {
	c, err := storagemarket.ParseCid(cid)
	if err != nil {
		return nil, err
	}
	data, err := fc.client.Client.Download(context.Background(), c)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GeneratePrivateKey generates a new private key
func (fc *FilecoinClient) GeneratePrivateKey() (crypto.PrivateKey, error) {
	privateKey := make([]byte, 32)
	_, err := rand.Read(privateKey)
	if err != nil {
		return nil, err
	}
	return crypto.PrivateKey(privateKey), nil
}

// CreateEnergyCertificate creates a new energy certificate
func (fc *FilecoinClient) CreateEnergyCertificate(issuer string, mwh *abi.TokenAmount, validUntil abi.ChainEpoch, privateKey crypto.PrivateKey) (*EnergyCertificate, error) {
	energyCertificate := &EnergyCertificate{
		Issuer:     issuer,
		MWh:        mwh,
		ValidUntil: validUntil,
	}

	data := energyCertificate.Bytes()
	signature, err := fc.SignMessage(privateKey, data)
	if err != nil {
		return nil, err
	}
	energyCertificate.Signature = signature

	return energyCertificate, nil
}

// EnergyCertificate is a struct representing an energy certificate
type EnergyCertificate struct {
	Issuer     string
	MWh        *abi.TokenAmount
	ValidUntil abi.ChainEpoch
	Signature  []byte
}

// Bytes returns the byte representation of an energy certificate
func (ec *EnergyCertificate) Bytes() []byte {
	return append([]byte(ec.Issuer), ec.MWh.Bytes()...)
}

// Verify verifies an energy certificate
func (ec *EnergyCertificate) Verify(data []byte, publicKey crypto.PublicKey) error {
	return fc.VerifySignature(publicKey, data, ec.Signature)
}

// Cid returns the CID of an energy certificate
func (ec *EnergyCertificate) Cid() string {
	return hex.EncodeToString(ec.Signature)
}

// String returns the string representation of an energy certificate
func (ec *EnergyCertificate) String() string {
	return fmt.Sprintf("Issuer: %s, MWh: %s, ValidUntil: %d, CID: %s", ec.Issuer, ec.MWh, ec.ValidUntil, ec.Cid())
}
