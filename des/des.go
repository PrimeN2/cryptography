package des

import "MAI-crypto/permutation"

type DESKeyExpander struct{}

func (DESKeyExpander) ExpandKey(key []byte) [][]byte {
	if len(key) != 8 {
		panic("DES: Key must be 8 bytes")
	}
	permutedKey := permutation.Permute(
		key,
		pc1[:],
		1,
		permutation.MSBFirst,
	)

	c := read28(permutedKey, 0)
	d := read28(permutedKey, 28)

	roundKeys := make([][]byte, 16)

	for i, shift := range shifts {
		c = rotateLeft28(c, shift)
		d = rotateLeft28(d, shift)

		// Собираем C || D обратно в 56 бит.
		cd := make([]byte, 7)

		write28(cd, 0, c)
		write28(cd, 28, d)

		// 56 -> 48 бит
		roundKeys[i] = permutation.Permute(
			cd,
			pc2[:],
			1,
			permutation.MSBFirst,
		)
	}

	return roundKeys
}

var pc1 = [...]int{
	57, 49, 41, 33, 25, 17, 9,
	1, 58, 50, 42, 34, 26, 18,
	10, 2, 59, 51, 43, 35, 27,
	19, 11, 3, 60, 52, 44, 36,
	63, 55, 47, 39, 31, 23, 15,
	7, 62, 54, 46, 38, 30, 22,
	14, 6, 61, 53, 45, 37, 29,
	21, 13, 5, 28, 20, 12, 4,
}

var pc2 = [...]int{
	14, 17, 11, 24, 1, 5,
	3, 28, 15, 6, 21, 10,
	23, 19, 12, 4, 26, 8,
	16, 7, 27, 20, 13, 2,
	41, 52, 31, 37, 47, 55,
	30, 40, 51, 45, 33, 48,
	44, 49, 39, 56, 34, 53,
	46, 42, 50, 36, 29, 32,
}

var shifts = [...]int{
	1, 1, 2, 2, 2, 2, 2, 2,
	1, 2, 2, 2, 2, 2, 2, 1,
}

func read28(data []byte, offset int) uint32 {
	var result uint32

	for i := 0; i < 28; i++ {
		index := offset + i

		byteIndex := index / 8
		bitIndex := 7 - (index % 8)

		bit := (data[byteIndex] >> bitIndex) & 1

		result = (result << 1) | uint32(bit)
	}

	return result
}

func write28(data []byte, offset int, value uint32) {
	for i := 0; i < 28; i++ {
		index := offset + i

		byteIndex := index / 8
		bitIndex := 7 - (index % 8)

		bit := (value >> (27 - i)) & 1

		if bit != 0 {
			data[byteIndex] |= 1 << bitIndex
		}
	}
}
