package feistel

import "MAI-crypto/cipher"

type Feistel struct {
	keyExpander   cipher.KeyExpander
	roundFunction cipher.RoundFunction
	roundKeys     [][]byte
}

func NewFeistel(keyExpander cipher.KeyExpander, roundFunction cipher.RoundFunction) *Feistel {
	return &Feistel{
		keyExpander:   keyExpander,
		roundFunction: roundFunction,
	}
}

func (f *Feistel) SetKey(key []byte) {
	f.roundKeys = f.keyExpander.ExpandKey(key)
}

func (f *Feistel) Encrypt(block []byte) []byte {
	return f.crypt(block, false)
}

func (f *Feistel) Decrypt(block []byte) []byte {
	return f.crypt(block, true)
}

func (f *Feistel) crypt(block []byte, reverseKeys bool) []byte {
	if len(block)%2 != 0 {
		panic("Feistel: block size must be even")
	}

	if len(f.roundKeys) == 0 {
		panic("Feistel: key is not set")
	}

	half := len(block) / 2

	L := append([]byte(nil), block[:half]...)
	R := append([]byte(nil), block[half:]...)

	for i := 0; i < len(f.roundKeys); i++ {
		keyIndex := i

		if reverseKeys {
			keyIndex = len(f.roundKeys) - 1 - i
		}

		roundKey := f.roundKeys[keyIndex]

		fResult := f.roundFunction.Apply(R, roundKey)
		newR := xorBytes(L, fResult)

		L, R = R, newR
	}

	return append(
		append([]byte(nil), R...),
		L...,
	)
}

func xorBytes(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("Feistel: XOR operands must have equal length")
	}

	result := make([]byte, len(a))

	for i := range a {
		result[i] = a[i] ^ b[i]
	}

	return result
}
