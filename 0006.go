func convert(s string, numRows int) string {
	count := len(s)
	if count <= 1 {
		return s
	}

	n := numRows*2 - 2
	buf := bytes.Buffer{}

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < count; j += n {
			buf.WriteByte(s[i+j])
			if i != 0 && i != numRows-1 && j+n-i < count {
				buf.WriteByte(s[j+n-i])
			}
		}
	}
	return buf.String()
}
