package sha256


func Sum256(msg []byte) string{
	var Temp1, Temp2 uint32
	var Majority, choice uint32
	var S0, S1 uint32
	msgLen := uint64(len(msg) * 8)
	
	lenMessageBlock := (msgLen + 64 + 1)
	if lenMessageBlock % 512 != 0 {
		lenMessageBlock += 512 - (lenMessageBlock % 512)
	}
	var messageBlock = make([]uint8, lenMessageBlock)
	
	if len(msg) != 0 {
		for i := 0; i < len(msg); i++ {
			messageBlock[i] = uint8(msg[i])
		}
	}
	// 2. Add the 1 bit to the end of the message
	messageBlock[len(msg)] = 0x80

	// 3. Add the message length in bits as a 64-bit big-endian integer
	for i:= 0; i <= 8; i++ {
		messageBlock[int(lenMessageBlock/8) - 1 - i] = uint8(msgLen >> (8 * i))
	}
	
	if DEBUG {
		PrintBlock(messageBlock)
		print("\n")
	}

	nbrChunks := int(lenMessageBlock / 512)

	h0 := ih0
	h1 := ih1
	h2 := ih2
	h3 := ih3
	h4 := ih4
	h5 := ih5
	h6 := ih6
	h7 := ih7

	for j:= 0; j < nbrChunks; j++ {
		if DEBUG {
			print("Chunk ",j+1,"\n")
		}

		var messageSchedule = make([]uint32, 64)

		// 4. Create the message schedule
		for i := 0+16*j; i < 16 + 16*j; i++ {
			messageSchedule[i-j*16] = uint32(messageBlock[0 + i*4]) << 24 | 
								uint32(messageBlock[1 + i*4]) << 16 | 
								uint32(messageBlock[2 + i*4]) << 8 | 
								uint32(messageBlock[3 + i*4])
		}

		
		// 5. Extend the message schedule
		for i := 16; i < 64; i++ {
			sigma0 := RightRotate(messageSchedule[i-15], 7) ^ 
					RightRotate(messageSchedule[i-15], 18) ^ 
					RightShift(messageSchedule[i-15], 3)

			sigma1 := RightRotate(messageSchedule[i-2],17) ^
					RightRotate(messageSchedule[i-2],19) ^
					RightShift(messageSchedule[i-2], 10)

			messageSchedule[i] = messageSchedule[i-16] + sigma0 + messageSchedule[i-7] + sigma1
		}
		if DEBUG {
			PrintSchedule(messageSchedule)
			print("\n")
		}

		// 6. Initialize the working variables
		a := h0
		b := h1
		c := h2
		d := h3
		e := h4
		f := h5
		g := h6
		h := h7
		// 7. update the working variables
		for i := 0; i < 64; i++ {

			S1 = RightRotate(e, 6) ^ RightRotate(e, 11) ^ RightRotate(e, 25)
			choice = (e & f) ^ (^e & g)
			S0 = RightRotate(a, 2) ^ RightRotate(a, 13) ^ RightRotate(a, 22)
			Temp1 = h + S1 + choice + messageSchedule[i] + kConstants[i]
			Majority = (a & b) ^ (a & c) ^ (b & c)
			Temp2 = S0 + Majority
			
			h = g
			g = f
			f = e
			e = d + Temp1
			d = c
			c = b
			b = a
			a = Temp1 + Temp2

		}
		h0 += a
		h1 += b
		h2 += c
		h3 += d
		h4 += e
		h5 += f
		h6 += g
		h7 += h
	}

	return Uint32ToHex(h0) + Uint32ToHex(h1) + 
		   Uint32ToHex(h2) + Uint32ToHex(h3) + 
		   Uint32ToHex(h4) + Uint32ToHex(h5) + 
		   Uint32ToHex(h6) + Uint32ToHex(h7)
}

// 0 -> 1 -> 0, 1
// 1 -> 11 -> 1, 3
// 10 -> 101 -> 2, 5
// 11 -> 111 -> 3, 7
// 100 -> 1001 -> 4, 9
// 101 -> 1011 -> 5, 11