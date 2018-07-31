// A summary: http://www.cnblogs.com/Martinium/articles/popcount.html
package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount23(x uint64) int {
	count := byte(0)
	for x != 0 {
		count += pc[byte(x)]
		x >>= 8
	}
	return int(count)
}

func PopCount24(x uint64) int {
	count := 0
	for x > 0 {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

func PopCount25(x uint64) int {
	count := 0
	for x > 0 {
		x = x & (x - 1)
		count++
	}

	return count
}
