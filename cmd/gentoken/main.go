package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/maenet/go-appstoreconnect"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Generate signed token for App Store Connect API.\nusage: gentoken <key_id> <issuer_id> <private_key_file>\n")
}

func gentoken() error {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 3 {
		return fmt.Errorf("invalid arguments\nRun 'gentoken -help' for usage.")
	}

	args := flag.Args()
	keyID := args[0]
	issuerID := args[1]
	file := args[2]

	privateKey, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("failed to read private key file: %w", err)
	}

	token := appstoreconnect.NewToken(
		keyID,
		issuerID,
		time.Minute*5,
		[]string{},
	)

	signedToken, err := token.Sign(privateKey)
	if err != nil {
		return err
	}

	fmt.Println(signedToken)
	return nil
}

func main() {
	if err := gentoken(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
