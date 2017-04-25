package strings

import "testing"

func Test_InterceptNumber(t *testing.T) {
	if numbers := InterceptNumber("123a456bca"); numbers[0] != 123 && numbers[1] != 456 {
		t.Fatal()
	}

	if numbers := InterceptNumber("123.9aaaa"); numbers[0] != 123.9 {
		t.Fatal()
	}
}
