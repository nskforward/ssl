package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/nskforward/ssl/internal/scenario"
	"github.com/nskforward/ssl/internal/util"
)

func main() {

	var (
		commonName     = ""
		output         = ""
		keyFilename    = ""
		caKeyFilename  = ""
		caCertFilename = ""
		domains        = ""
		days           = 365
	)

	flag.StringVar(&commonName, "cn", "", "common name")
	flag.StringVar(&output, "o", "", "output dir/file path")
	flag.StringVar(&keyFilename, "key", "", "private key file path")
	flag.StringVar(&caKeyFilename, "cakey", "ca_key.pem", "CA private key file path")
	flag.StringVar(&caCertFilename, "cacert", "ca.pem", "CA cert file path")
	flag.StringVar(&domains, "domain", "", "domain list with ',' as separator")
	flag.IntVar(&days, "days", 365, "cert expired after these days")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		util.Fatal(fmt.Errorf("command not specified"))
	}

	cmd := args[0]
	// args = args[1:]

	switch strings.ToLower(cmd) {
	case "ca":
		err := scenario.GenPairCA(commonName, output, keyFilename, days)
		util.Fatal(err)
		return

	case "server":
		err := scenario.GenPair(commonName, output, keyFilename, caKeyFilename, caCertFilename, domains, days)
		util.Fatal(err)
		return
	}

	util.Fatal(fmt.Errorf("unknown command: %s", cmd))
}
