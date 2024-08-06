package scenario

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nskforward/ssl/internal/util"
)

func GenPairCA(commonName, output, privateKeyFilename string, days int) error {
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
		privateKeyFilename = filepath.Join(output, "ca_key.pem")
		err = GenKey(privateKeyFilename)
		if err != nil {
			return err
		}
	}

	certFilename := filepath.Join(output, "ca.pem")
	return GenCertCA(commonName, certFilename, privateKeyFilename, days)
}
