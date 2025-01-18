package sha256

func PrintBlock(messageBlock []uint8) {
	// print the message block
	print("MessageBlock:\n")
	for i := 0; i < len(messageBlock); i++ {
		if i % 8 == 0 {
			print(" ")
		}
		if i % 32 == 0 {
			print("\n")
		}
		byteIndex := i / 8
		bitIndex := 7 - (i % 8)
		if (messageBlock[byteIndex]>>bitIndex)&1 == 1 {
			print("\033[31m1\033[0m") // red
		} else {
			print("0")
		}
	}
	print("\n")
}

func PrintSchedule(block []uint32) {
	print("Message Schedule:\n")
	for i := 0; i < 32 *64; i++ {
		// if i % 8 == 0 {
		// 	print(" ")
		// }
		if i % 32 == 0 {
			print("\n")
			print("w[", i/32, "] = ")
			if i/32 < 10 {
				print(" ")
			}
		}
		byteIndex := i / 32
		bitIndex := 31 - (i % 32)
		if (block[byteIndex]>>bitIndex)&1 == 1 {
			print("\033[31m1\033[0m") // red
		} else {
			print("0")
		}
	}
	print("\n")
}

func printUint8(num uint8) {
	for i := 0; i < 8; i++ {
		bitIndex := 7 - i
		if (num>>bitIndex)&1 == 1 {
			print("\033[31m1\033[0m") // red
		} else {
			print("0")
		}
	}
	print("\n")
}

func printUint32(num uint32) {
	for i := 0; i < 32; i++ {
		bitIndex := 31 - i
		if (num>>bitIndex)&1 == 1 {
			print("\033[31m1\033[0m") // red
		} else {
			print("0")
		}
	}
	print("\n")
}

func printUint64(num uint64) {
	for i := 0; i < 64; i++ {
		bitIndex := 63 - i
		if (num>>bitIndex)&1 == 1 {
			print("\033[31m1\033[0m") // red
		} else {
			print("0")
		}
	}
	print("\n")
}