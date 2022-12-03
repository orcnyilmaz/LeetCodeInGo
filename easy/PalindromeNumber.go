package easy

// https://leetcode.com/problems/palindrome-number/
func IsPalindrome(x int) bool {
	tempNum := x
	var reversedNum int
	carry := tempNum % 10

	for tempNum != 0 {
		reversedNum = reversedNum*10 + carry
		tempNum = tempNum / 10
		carry = tempNum % 10
	}

	return reversedNum == x
}
