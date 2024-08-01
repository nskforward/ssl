package scenario

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/nskforward/ssl/internal/util"
)

func GenPair(commonName, output, privateKeyFilename, caKeyFilename, caCertFilename string, domains string) error {
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
		privateKeyFilename = filepath.Join(output, fmt.Sprintf("%s.key", commonName))
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

	certFilename := filepath.Join(output, fmt.Sprintf("%s.cert", commonName))
	return GenCert(commonName, certFilename, privateKeyFilename, caKeyFilename, caCertFilename, dnsNames, ipAddresses)
}
