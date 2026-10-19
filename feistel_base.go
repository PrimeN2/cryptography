package main

import (
	"MAI-crypto/feistel"
	"fmt"
)

type TestKeyExpander struct{}

func (TestKeyExpander) ExpandKey(key []byte) [][]byte {
	return [][]byte{
		{0x01},
		{0x02},
		{0x03},
	}
}

type TestRoundFunction struct{}

func (TestRoundFunction) Apply(block []byte, roundKey []byte) []byte {
	return []byte{block[0] ^ roundKey[0]}
}

func main() {
	f := feistel.NewFeistel(
		TestKeyExpander{},
		TestRoundFunction{},
	)

	f.SetKey([]byte{0xAA})

	plaintext := []byte{0x10, 0x20}

	ciphertext := f.Encrypt(plaintext)
	fmt.Printf("Encrypted: %02X %02X\n", ciphertext[0], ciphertext[1])

	decrypted := f.Decrypt(ciphertext)
	fmt.Printf("Decrypted: %02X %02X\n", decrypted[0], decrypted[1])
}
