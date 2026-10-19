package permutation

type BitOrder bool

const (
	LSBFirst BitOrder = false
	MSBFirst BitOrder = true
)

func Permute(value []byte, pBlock []int, start int, msbFirst BitOrder) []byte {
	resultSize := (len(pBlock) + 7) / 8
	result := make([]byte, resultSize)

	for i, src := range pBlock {
		bit := getBit(value, src-start, msbFirst)
		setBit(result, i, bit, msbFirst)
	}

	return result
}

func getBit(value []byte, index int, msbFirst BitOrder) bool {
	byteIndex := index / 8
	bitIndex := index % 8

	if msbFirst {
		bitIndex = 7 - bitIndex
	}

	return (value[byteIndex]>>bitIndex)&1 != 0
}

func setBit(value []byte, index int, bit bool, msbFirst BitOrder) {
	byteIndex := index / 8
	bitIndex := index % 8

	if msbFirst {
		bitIndex = 7 - bitIndex
	}

	if bit {
		value[byteIndex] |= 1 << bitIndex
	} else {
		value[byteIndex] &^= 1 << bitIndex
	}
}

// func main() {
// 	value := []byte{0xB2, 0x69}
// 	// p := []int{
// 	// 	8, 9, 10, 11, 12, 13, 14, 15,
// 	// 	0, 1, 2, 3, 4, 5, 6, 7,
// 	// }
// 	// p := []int{
// 	// 	9, 10, 11, 12, 13, 14, 15, 16,
// 	// 	1, 2, 3, 4, 5, 6, 7, 8,
// 	// }
// 	p := []int{
// 		16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
// 	}
// 	result := permute(value, p, 1, LSBFirst)

// 	fmt.Printf("%02X %02X\n", result[0], result[1])
// }
