package codes

import "fmt"

func romanToInt(s string) int {
	fmt.Println(s)
	var value int
	var symbolToValue = map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	var specialCase = map[string]int{
		"I": 0,
		"X": 0,
		"C": 0,
	}

	for idx := 0; idx < len(s); idx++ {
		ss := string(s[idx])
		if _, ok := specialCase[ss]; ok {
			if idx < len(s)-1 {
				newss := string(s[idx]) + string(s[idx+1])
				if tempVal, ok := symbolToValue[newss]; ok {
					value += tempVal
					idx++
				} else {
					if tempVal, ok := symbolToValue[ss]; ok {
						value += tempVal
					}
				}
			} else {
				if tempVal, ok := symbolToValue[ss]; ok {
					value += tempVal
				}
			}
		} else {
			if tempVal, ok := symbolToValue[ss]; ok {
				value += tempVal
			}
		}
	}

	return value
}

func romanToInt1(s string) int {
	fmt.Println(s)
	var v, lv, cv int
	h := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		cv = h[s[i]]
		if cv < lv {
			v -= cv
		} else {
			v += cv
		}
		lv = cv
	}

	return v
}

func RomanToInt() {
	// fmt.Println(romanToInt("III"))
	// fmt.Println(romanToInt("LVIII"))
	// fmt.Println(romanToInt("MCMXCIV"))

	fmt.Println(romanToInt1("III"))
	fmt.Println(romanToInt1("LVIII"))
	fmt.Println(romanToInt1("MCMXCIV"))
	fmt.Println(romanToInt1("IL"))

}
