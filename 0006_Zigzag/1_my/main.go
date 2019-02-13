package main

func convert(s string, numRows int) string {
	ss := []rune(s)
	out := []rune{}

	if numRows == 1 {
		return s
	}
	for i := 0; i < numRows; i++ {
		for j := 0; true; j++ {
			// zig
			if j > 0 && i > 0 && i < numRows-1 {
				idx := i + (numRows-1-i)*2 + 2*(j-1)*(numRows-1)
				if idx >= len(ss) {
					break
				}
				out = append(out, ss[idx])
			}

			// zag
			idx := i + (numRows-1)*2*j
			if idx >= len(ss) {
				break
			}
			out = append(out, ss[idx])
		}
	}
	return string(out)
}
