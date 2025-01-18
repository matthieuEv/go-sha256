package sha256

func RightRotate(x uint32, n uint) uint32 {
	return (x >> n) | (x << (32 - n))
}

func RightShift(x uint32, n uint) uint32 {
	return x >> n
}

func Uint32ToHex(num uint32) string {
	var hex string
	for i := 0; i < 8; i++ {
		hex = binToHex[uint(num >> (4 * i) & 0xf)] + hex
	}
	return hex
}

func GetNbrBits(num uint8) int {
	var count int
	for num != 0 {
		num >>= 1
		count++
	}
	return count
}