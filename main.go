package main

import (
	"fmt"
	"strings"
)


func main() {
	//numberMap := map[string]int{
	//	"零":0,
	//	"一":1,
	//	"二":2,
	//	"三":3,
	//	"四":4,
	//	"五":5,
	//	"六":6,
	//	"七":7,
	//	"八":8,
	//	"九":9,
	//	"十":10,
	//	"百":100,
	//	"千":1000,
	//	"万":10000,
	//}

	fmt.Println(chinesenNumber("一千三百二十一"))
	fmt.Println(chinesenNumber("一千零一十二"))

}

func chinesenNumber(s string) int {
	a := strings.
	return len(a)
}

//func StringToRuneArr(s string, arr []rune) {
//	src := []rune(s)
//	for i, v := range src {
//		if i >= len(arr) {
//			break
//		}
//		arr[i] = v
//	}
//}
//
//func main(){
//	str := "这是一个字符串ABCDEF"
//	var arr [10]rune
//	utility.StringToRuneArr(str, arr[:])
//	fmt.Println(string(arr[:]))
//}

//判断第一位
//1-10

//判断第二位为进制单位

//判断第三位是否位零
//1-9


//一
//十

//十一
//二十
//三十七

//一百
//一百零一
//一百一十
//一百一十二


//九十九万 九千 九百 九十 九
//九十九万
