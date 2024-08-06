package scenario

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/nskforward/ssl/internal/util"
)

func GenPair(commonName, output, privateKeyFilename, caKeyFilename, caCertFilename, domains string, days int) error {
	if commonName == "" {
		commonName = util.AskString("common name", true, false)
	}

	if output == "" {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		output = dir
	}

	fi, err := os.Stat(output)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("output must be a dir")
	}

	if privateKeyFilename == "" {
		privateKeyFilename = filepath.Join(output, "key.pem")
		err = GenKey(privateKeyFilename)
		if err != nil {
			return err
		}
	}

	var dnsNames = []string{}
	var ipAddresses = []net.IP{}

	for _, domain := range strings.Split(domains, ",") {
		d := strings.TrimSpace(domain)
		ip := net.ParseIP(d)
		if ip == nil {
			dnsNames = append(dnsNames, d)
		} else {
			ipAddresses = append(ipAddresses, ip)
		}
	}

	certFilename := filepath.Join(output, "cert.pem")
	err = GenCert(commonName, certFilename, privateKeyFilename, caKeyFilename, caCertFilename, days, dnsNames, ipAddresses)
	if err != nil {
		return err
	}

	f1, err := os.OpenFile(certFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f1.Close()

	f2, err := os.Open(caCertFilename)
	if err != nil {
		return err
	}
	defer f2.Close()

	io.Copy(f1, f2)

	return nil
}
