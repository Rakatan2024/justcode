package lecture02

func intToRoman(num int) string {
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hund := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thou := []string{"", "M", "MM", "MMM"}
	return thou[num/1000] + hund[num%1000/100] + tens[num%100/10] + ones[num%10]
}
