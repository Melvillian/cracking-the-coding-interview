package lib

func Palindrome(word string) (bool) {
	m := make(map[int32]bool);
	for _, c := range word {
		k := m[c];

		if (k == true) {
			m[c] = false
		} else {
			m[c] = true;
		}
	}

	isFirstOdd := false;
	for _, bit := range m {
		if (bit == true && isFirstOdd == false) {
			isFirstOdd = true;
		} else if (bit == true && isFirstOdd == true) {
			return false;
		}
	}

	return true;
}