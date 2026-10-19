package cipher

type KeyExpander interface {
	ExpandKey(key []byte) [][]byte
}

type RoundFunction interface {
	Apply(block []byte, roundKey []byte) []byte
}

type BlockCipher interface {
	SetKey(key []byte)

	EncryptBlock(block []byte)
	DecryptBlock(block []byte)
}
