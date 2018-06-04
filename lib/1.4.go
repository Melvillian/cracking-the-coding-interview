package lib

// Palindrome returns true if the characters in word are some permutation of a palindrome.
// Rather than checking all permutations, which is slow, it checks if the following statement is true:
// "word is a palindrome iff if at most 1 character appears an odd number of times within it"
func Palindrome(word string) (bool) {
	m := make(map[int32]bool);

	for _, c := range word {
		isOdd := m[c];

		if (isOdd) {
			m[c] = false
		} else if (!isOdd) {
			m[c] = true;
		} else { // m[c] hasn't been set yet
			m[c] = true;
		}
	}

	// if we see more than 1 character that occurs an odd number of times, it's not a palindrome
	hasAlreadySeenAnOdd := false;
	for _, isOdd := range m {
		if (isOdd && !hasAlreadySeenAnOdd) {
			hasAlreadySeenAnOdd = true;
		} else if (isOdd && hasAlreadySeenAnOdd) {
			return false;
		}
	}

	return true;
}