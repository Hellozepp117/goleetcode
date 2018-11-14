package functions

func intToRoman(num int) string {
	var ans string
	for i := 1; i <= num/1000; i++ {
		ans += "M"
	}
	num = num % 1000
	if num/900 == 1 {
		ans += "CM"
		num -= 900
	}
	if num/500 == 1 {
		ans += "D"
		num -= 500
	}
	if num/400 == 1 {
		ans += "CD"
		num -= 400
	}

	for i := 1; i <= num/100; i++ {
		ans += "C"
	}
	num = num % 100
	if num/90 == 1 {
		ans += "XC"
		num -= 90
	}
	if num/50 == 1 {
		ans += "L"
		num -= 50
	}
	if num/40 == 1 {
		ans += "XL"
		num -= 40
	}
	for i := 1; i <= num/10; i++ {
		ans += "X"
	}
	num = num % 10

	if num/9 == 1 {
		ans += "IX"
		num -= 9
	}
	if num/5 == 1 {
		ans += "V"
		num -= 5
	}
	if num/4 == 1 {
		ans += "IV"
		num -= 4
	}
	for i := 1; i <= num; i++ {
		ans += "I"
	}
	return ans
}
