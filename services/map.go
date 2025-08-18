package services

import (
	"fmt"
)

type obj struct {
	param int
}

func MapTest1() {
	m := make(map[int]*obj)
	m[1] = &obj{param: 1}
	m[2] = &obj{param: 2}
	m[3] = &obj{param: 3}

	for _, v := range m {
		if v.param == 2 {
			v = &obj{param: 999} // v = &obj{param: 999} 只是把 局部變數 v 改指向新的物件，並不會改變 m[2] 存的指標。這一步不會影響 map 裡的內容
		}
		v.param += 70 // v 已經被重新賦值（例如那次 param == 2 時），那修改的是新的 obj，不會影響 map
	}
	fmt.Println(m[1].param)
	fmt.Println(m[2].param)
	fmt.Println(m[3].param)
}
