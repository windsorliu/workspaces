package note

import "fmt"

//常數宣告
func Const() {
	const (
		c0 = 8
		c1 = iota //當前行數，預設從0開始
		c2        //預設為上一行的值，也就是iota
		c3
		c4 = 12
		c5 = iota
	)

	fmt.Printf("c0=%v,c1=%v,c2=%v,c3=%v,c4=%v,c5=%v", c0, c1, c2, c3, c4, c5)
	//c0=8,c1=1,c2=2,c3=3,c4=12,c5=5
}

//Array
func Array() {
	a := [...]int{9, 8, 7, 6, 5}

	for i, v := range a {
		fmt.Printf("a[%v]=%v\n", i, v)
	}

	var b [3][4]int = [3][4]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
	}

	for i, v := range b {
		for i2, v2 := range v {
			fmt.Printf("b[%v][%v]=%v\n", i, i2, v2)
		}
	}
}

func Slice() {

}
