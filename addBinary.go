package leetcode

func addBinary(a string, b string) string {
	sb := []byte{}
	var carry byte
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry == 1 {
		if i >= 0 {
			carry += a[i] - '0'
			i--
		}
		if j >= 0 {
			carry += b[j] - '0'
			j--
		}

		sb = append([]byte{(carry % 2) + 48}, sb...)
		carry /= 2
	}

	return string(sb)
}
