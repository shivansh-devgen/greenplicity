package certificates

import (
	"time",
	"errors"
)

// EnergyCertificate represents an energy certificate
type EnergyCertificate struct {
	ID          string
	ProducerID  string
	MWh         float64
	CreatedAt   time.Time
	ValidUntil  time.Time
	Certificate []byte // certificate file or hash
}

// NewEnergyCertificate creates a new energy certificate
func NewEnergyCertificate(producerID string, mwh float64, validUntil time.Time, certificate []byte) *EnergyCertificate {
	return &EnergyCertificate{
		ID:         generateCertificateID(),
		ProducerID: producerID,
		MWh:        mwh,
		CreatedAt:  time.Now(),
		ValidUntil: validUntil,
		Certificate:certificate,
	}
}

func (e *EnergyCertificate) VerifyCertificate() error {
    // check if the certificate has expired
    if time.Now().After(e.ValidUntil) {
        return errors.New("certificate has expired")
    }

    // code to verify the authenticity of the certificate, it could be by checking a digital signature or using a trusted certificate authority
    // this can vary depending on how the certificate was generated
    // In this example I am going to check if the certificate file is in the right format
    if !strings.HasSuffix(string(e.Certificate), ".pem") {
        return errors.New("Invalid certificate format")
    }
    return nil
}
// VerifyCertificate verifies the authenticity of an energy certificate
func generateCertificateID() string {
    // code to generate unique ID, in this example we will use uuid package to generate unique id
    uuid := uuid.New()
    return uuid.String()
}
