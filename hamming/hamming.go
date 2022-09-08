package hamming

import "errors"

func Distance(a, b string) (int, error) {
	distance := 0

	if len(a) != len(b) {
		return 0, errors.New("mismatch")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				distance = distance + 1
			}
		}
		return distance, nil
	}

}
