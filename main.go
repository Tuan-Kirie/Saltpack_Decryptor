package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/keys-pub/keys"
	"github.com/keys-pub/keys/saltpack"
)

func main() {
	// аргументы
	hexKey := flag.String("key", "", "hex encoded 32-byte key")
	inFile := flag.String("in", "", "input file")
	outFile := flag.String("out", "", "output file (optional)")

	flag.Parse()

	if *hexKey == "" || *inFile == "" {
		log.Fatal("usage: -key <hex> -in <file> [-out <file>]")
	}

	// hex → []byte
	seedBytes, err := hex.DecodeString(*hexKey)
	if err != nil {
		log.Fatal(err)
	}

	if len(seedBytes) != 32 {
		log.Fatalf("seed must be 32 bytes, got %d", len(seedBytes))
	}

	var seed [32]byte
	copy(seed[:], seedBytes)

	recipientKey := keys.NewX25519KeyFromSeed(&seed)

	// читаем файл
	encrypted, err := os.ReadFile(*inFile)
	if err != nil {
		log.Fatal(err)
	}

	out, _, err := saltpack.SigncryptOpen(encrypted, false, saltpack.NewKeyring(recipientKey))
	if err != nil {
		log.Fatal(err)
	}

	// вывод
	if *outFile != "" {
		err = os.WriteFile(*outFile, []byte(out), 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print(string(out))
	}
}
