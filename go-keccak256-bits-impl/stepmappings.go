package keccak

func theta(a [25 * 64]bool) [25 * 64]bool {
	var c0, c1, c2, c3, c4, d [64]bool
	var r [25 * 64]bool

	copy(c0[:], xor(xor(xor(xor(a[0:1*64], a[5*64:6*64]), a[10*64:11*64]), a[15*64:16*64]), a[20*64:21*64]))
	copy(c1[:], xor(xor(xor(xor(a[1*64:2*64], a[6*64:7*64]), a[11*64:12*64]), a[16*64:17*64]), a[21*64:22*64]))
	copy(c2[:], xor(xor(xor(xor(a[2*64:3*64], a[7*64:8*64]), a[12*64:13*64]), a[17*64:18*64]), a[22*64:23*64]))
	copy(c3[:], xor(xor(xor(xor(a[3*64:4*64], a[8*64:9*64]), a[13*64:14*64]), a[18*64:19*64]), a[23*64:24*64]))
	copy(c4[:], xor(xor(xor(xor(a[4*64:5*64], a[9*64:10*64]), a[14*64:15*64]), a[19*64:20*64]), a[24*64:25*64]))

	copy(d[:], xor(c4[:], or(leftShift(c1[:], 1), rightShift(c1[:], (64-1)))))
	copy(r[0:1*64], xor(a[0:1*64], d[:]))
	copy(r[5*64:6*64], xor(a[5*64:6*64], d[:]))
	copy(r[10*64:11*64], xor(a[10*64:11*64], d[:]))
	copy(r[15*64:16*64], xor(a[15*64:16*64], d[:]))
	copy(r[20*64:21*64], xor(a[20*64:21*64], d[:]))

	copy(d[:], xor(c0[:], or(leftShift(c2[:], 1), rightShift(c2[:], (64-1)))))
	copy(r[1*64:2*64], xor(a[1*64:2*64], d[:]))
	copy(r[6*64:7*64], xor(a[6*64:7*64], d[:]))
	copy(r[11*64:12*64], xor(a[11*64:12*64], d[:]))
	copy(r[16*64:17*64], xor(a[16*64:17*64], d[:]))
	copy(r[21*64:22*64], xor(a[21*64:22*64], d[:]))

	copy(d[:], xor(c1[:], or(leftShift(c3[:], 1), rightShift(c3[:], (64-1)))))
	copy(r[2*64:3*64], xor(a[2*64:3*64], d[:]))
	copy(r[7*64:8*64], xor(a[7*64:8*64], d[:]))
	copy(r[12*64:13*64], xor(a[12*64:13*64], d[:]))
	copy(r[17*64:18*64], xor(a[17*64:18*64], d[:]))
	copy(r[22*64:23*64], xor(a[22*64:23*64], d[:]))

	copy(d[:], xor(c2[:], or(leftShift(c4[:], 1), rightShift(c4[:], (64-1)))))
	copy(r[3*64:4*64], xor(a[3*64:4*64], d[:]))
	copy(r[8*64:9*64], xor(a[8*64:9*64], d[:]))
	copy(r[13*64:14*64], xor(a[13*64:14*64], d[:]))
	copy(r[18*64:19*64], xor(a[18*64:19*64], d[:]))
	copy(r[23*64:24*64], xor(a[23*64:24*64], d[:]))

	copy(d[:], xor(c3[:], or(leftShift(c0[:], 1), rightShift(c0[:], (64-1)))))
	copy(r[4*64:5*64], xor(a[4*64:5*64], d[:]))
	copy(r[9*64:10*64], xor(a[9*64:10*64], d[:]))
	copy(r[14*64:15*64], xor(a[14*64:15*64], d[:]))
	copy(r[19*64:20*64], xor(a[19*64:20*64], d[:]))
	copy(r[24*64:25*64], xor(a[24*64:25*64], d[:]))
	return r
}

func rhopi(a [25 * 64]bool) [25 * 64]bool {
	var t, tAux [64]bool
	var r [25 * 64]bool

	copy(r[0:1*64], a[0:1*64])

	copy(t[:], a[1*64:2*64])

	copy(tAux[:], a[10*64:11*64])
	copy(r[10*64:11*64], or(leftShift(t[:], 1), rightShift(t[:], 64-1)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[7*64:8*64])
	copy(r[7*64:8*64], or(leftShift(t[:], 3), rightShift(t[:], 64-3)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[11*64:12*64])
	copy(r[11*64:12*64], or(leftShift(t[:], 6), rightShift(t[:], 64-6)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[17*64:18*64])
	copy(r[17*64:18*64], or(leftShift(t[:], 10), rightShift(t[:], 64-10)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[18*64:19*64])
	copy(r[18*64:19*64], or(leftShift(t[:], 15), rightShift(t[:], 64-15)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[3*64:4*64])
	copy(r[3*64:4*64], or(leftShift(t[:], 21), rightShift(t[:], 64-21)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[5*64:6*64])
	copy(r[5*64:6*64], or(leftShift(t[:], 28), rightShift(t[:], 64-28)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[16*64:17*64])
	copy(r[16*64:17*64], or(leftShift(t[:], 36), rightShift(t[:], 64-36)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[8*64:9*64])
	copy(r[8*64:9*64], or(leftShift(t[:], 45), rightShift(t[:], 64-45)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[21*64:22*64])
	copy(r[21*64:22*64], or(leftShift(t[:], 55), rightShift(t[:], 64-55)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[24*64:25*64])
	copy(r[24*64:25*64], or(leftShift(t[:], 2), rightShift(t[:], 64-2)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[4*64:5*64])
	copy(r[4*64:5*64], or(leftShift(t[:], 14), rightShift(t[:], 64-14)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[15*64:16*64])
	copy(r[15*64:16*64], or(leftShift(t[:], 27), rightShift(t[:], 64-27)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[23*64:24*64])
	copy(r[23*64:24*64], or(leftShift(t[:], 41), rightShift(t[:], 64-41)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[19*64:20*64])
	copy(r[19*64:20*64], or(leftShift(t[:], 56), rightShift(t[:], 64-56)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[13*64:14*64])
	copy(r[13*64:14*64], or(leftShift(t[:], 8), rightShift(t[:], 64-8)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[12*64:13*64])
	copy(r[12*64:13*64], or(leftShift(t[:], 25), rightShift(t[:], 64-25)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[2*64:3*64])
	copy(r[2*64:3*64], or(leftShift(t[:], 43), rightShift(t[:], 64-43)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[20*64:21*64])
	copy(r[20*64:21*64], or(leftShift(t[:], 62), rightShift(t[:], 64-62)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[14*64:15*64])
	copy(r[14*64:15*64], or(leftShift(t[:], 18), rightShift(t[:], 64-18)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[22*64:23*64])
	copy(r[22*64:23*64], or(leftShift(t[:], 39), rightShift(t[:], 64-39)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[9*64:10*64])
	copy(r[9*64:10*64], or(leftShift(t[:], 61), rightShift(t[:], 64-61)))
	copy(t[:], tAux[:])

	copy(tAux[:], a[6*64:7*64])
	copy(r[6*64:7*64], or(leftShift(t[:], 20), rightShift(t[:], 64-20)))
	copy(t[:], tAux[:])

	copy(r[1*64:2*64], or(leftShift(t[:], 44), rightShift(t[:], 64-44)))
	return r
}

func chi(a [25 * 64]bool) [25 * 64]bool {
	var c0, c1, c2, c3, c4 [64]bool
	var r [25 * 64]bool

	copy(c0[:], a[0:1*64])
	copy(c1[:], a[1*64:2*64])
	copy(c2[:], a[2*64:3*64])
	copy(c3[:], a[3*64:4*64])
	copy(c4[:], a[4*64:5*64])

	copy(r[0:1*64], xor(a[0:1*64], and(xorSingle(c1[:]), c2[:])))
	copy(r[1*64:2*64], xor(a[1*64:2*64], and(xorSingle(c2[:]), c3[:])))
	copy(r[2*64:3*64], xor(a[2*64:3*64], and(xorSingle(c3[:]), c4[:])))
	copy(r[3*64:4*64], xor(a[3*64:4*64], and(xorSingle(c4[:]), c0[:])))
	copy(r[4*64:5*64], xor(a[4*64:5*64], and(xorSingle(c0[:]), c1[:])))

	copy(c0[:], a[5*64:6*64])
	copy(c1[:], a[6*64:7*64])
	copy(c2[:], a[7*64:8*64])
	copy(c3[:], a[8*64:9*64])
	copy(c4[:], a[9*64:10*64])

	copy(r[5*64:6*64], xor(a[5*64:6*64], and(xorSingle(c1[:]), c2[:])))
	copy(r[6*64:7*64], xor(a[6*64:7*64], and(xorSingle(c2[:]), c3[:])))
	copy(r[7*64:8*64], xor(a[7*64:8*64], and(xorSingle(c3[:]), c4[:])))
	copy(r[8*64:9*64], xor(a[8*64:9*64], and(xorSingle(c4[:]), c0[:])))
	copy(r[9*64:10*64], xor(a[9*64:10*64], and(xorSingle(c0[:]), c1[:])))

	copy(c0[:], a[10*64:11*64])
	copy(c1[:], a[11*64:12*64])
	copy(c2[:], a[12*64:13*64])
	copy(c3[:], a[13*64:14*64])
	copy(c4[:], a[14*64:15*64])

	copy(r[10*64:11*64], xor(a[10*64:11*64], and(xorSingle(c1[:]), c2[:])))
	copy(r[11*64:12*64], xor(a[11*64:12*64], and(xorSingle(c2[:]), c3[:])))
	copy(r[12*64:13*64], xor(a[12*64:13*64], and(xorSingle(c3[:]), c4[:])))
	copy(r[13*64:14*64], xor(a[13*64:14*64], and(xorSingle(c4[:]), c0[:])))
	copy(r[14*64:15*64], xor(a[14*64:15*64], and(xorSingle(c0[:]), c1[:])))

	copy(c0[:], a[15*64:16*64])
	copy(c1[:], a[16*64:17*64])
	copy(c2[:], a[17*64:18*64])
	copy(c3[:], a[18*64:19*64])
	copy(c4[:], a[19*64:20*64])

	copy(r[15*64:16*64], xor(a[15*64:16*64], and(xorSingle(c1[:]), c2[:])))
	copy(r[16*64:17*64], xor(a[16*64:17*64], and(xorSingle(c2[:]), c3[:])))
	copy(r[17*64:18*64], xor(a[17*64:18*64], and(xorSingle(c3[:]), c4[:])))
	copy(r[18*64:19*64], xor(a[18*64:19*64], and(xorSingle(c4[:]), c0[:])))
	copy(r[19*64:20*64], xor(a[19*64:20*64], and(xorSingle(c0[:]), c1[:])))

	copy(c0[:], a[20*64:21*64])
	copy(c1[:], a[21*64:22*64])
	copy(c2[:], a[22*64:23*64])
	copy(c3[:], a[23*64:24*64])
	copy(c4[:], a[24*64:25*64])

	copy(r[20*64:21*64], xor(a[20*64:21*64], and(xorSingle(c1[:]), c2[:])))
	copy(r[21*64:22*64], xor(a[21*64:22*64], and(xorSingle(c2[:]), c3[:])))
	copy(r[22*64:23*64], xor(a[22*64:23*64], and(xorSingle(c3[:]), c4[:])))
	copy(r[23*64:24*64], xor(a[23*64:24*64], and(xorSingle(c4[:]), c0[:])))
	copy(r[24*64:25*64], xor(a[24*64:25*64], and(xorSingle(c0[:]), c1[:])))
	return r
}
