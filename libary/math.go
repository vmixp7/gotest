package libary

import (
	"fmt"
	"math"
)

// 來計算圓的面積
func MathArea(radius float64) float64 {

	// 使用 math.Pi 來計算圓的面積：A = π * r^2
	area := math.Pi * math.Pow(radius, 2)

	fmt.Printf("圓的半徑是: %.2f\n", radius)
	fmt.Printf("圓的面積是: %.2f\n", area)
	return area
}

// 位移運算符的使用
func MathMove() {
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("第一行 - c 的值为 %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("第二行 - c 的值为 %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	c = a << 3 /* 240 = 1111 0000 左移2位就是a乘以2的3次方 */
	fmt.Printf("第四行 - c 的值为 %d\n", c)

	c = a >> 3 /* 15 = 0000 1111 右移2位就是a除以2的3次方 */
	fmt.Printf("第五行 - c 的值为 %d\n", c)
}
