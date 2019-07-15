package main

import (
	"fmt"
)


func main() {
	numberMap := map[string]int{
		"零":0,
		"一":1,
		"二":2,
		"三":3,
		"四":4,
		"五":5,
		"六":6,
		"七":7,
		"八":8,
		"九":9,
		"十":10,
		"百":100,
		"千":1000,
		"万":10000,
	}
	fmt.Println(chinesenNumber("一千三百二十一"))
	fmt.Println(chinesenNumber("一千零一十二"))
	fmt.Println("十三")
}

func chinesenNumber(s string) int {
	a := 0
	for k,v := range s {
		fmt.Println(k,":",string(v))

	}
	return a
}

