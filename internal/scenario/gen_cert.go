package scenario

import (
	"bytes"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	math "math/rand"
	"net"
	"os"
	"time"
)

func GenCertCA(commonName, output, privateKeyPath string, days int) error {
	privateKey, err := LoadPrivateKey(privateKeyPath)
	if err != nil {
		return err
	}

	template := &x509.Certificate{
		SerialNumber: big.NewInt(math.Int63()),
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{"Internet Technologies LLC"},
			Country:      []string{"RU"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, days),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		SignatureAlgorithm:    x509.SHA256WithRSA,
	}

	data, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return fmt.Errorf("cannot create cert: %w", err)
	}

	cert, err := x509.ParseCertificate(data)
	if err != nil {
		return fmt.Errorf("cannot parse cert: %w", err)
	}

	var buf bytes.Buffer
	err = pem.Encode(&buf, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})
	if err != nil {
		return err
	}
	return os.WriteFile(output, buf.Bytes(), 0755)
}

func GenCert(commonName, output, privateKeyPath, caKeyPath, caCertPath string, days int, dnsNames []string, ipAddresses []net.IP) error {
	privateKey, err := LoadPrivateKey(privateKeyPath)
	if err != nil {
		return err
	}

	privateKeyCA, err := LoadPrivateKey(caKeyPath)
	if err != nil {
		return err
	}

	certCA, err := LoadCert(caCertPath)
	if err != nil {
		return err
	}

	template := &x509.Certificate{
		SerialNumber: big.NewInt(math.Int63()),
		Subject: pkix.Name{
			CommonName:   commonName,
			Organization: []string{"Internet Technologies LLC"},
			Country:      []string{"RU"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, days),
		IsCA:                  false,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		SignatureAlgorithm:    x509.SHA256WithRSA,
		DNSNames:              dnsNames,
		IPAddresses:           ipAddresses,
	}

	data, err := x509.CreateCertificate(rand.Reader, template, certCA, &privateKey.PublicKey, privateKeyCA)
	if err != nil {
		return fmt.Errorf("cannot create cert: %w", err)
	}

	cert, err := x509.ParseCertificate(data)
	if err != nil {
		return fmt.Errorf("cannot parse cert: %w", err)
	}

	var buf bytes.Buffer
	err = pem.Encode(&buf, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})
	if err != nil {
		return err
	}
	return os.WriteFile(output, buf.Bytes(), 0755)
}
