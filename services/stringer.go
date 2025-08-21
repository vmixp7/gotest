package services

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// 为 IPAddr 添加一个 "String() string" 方法
func (ip IPAddr) String() string {
	parts := make([]string, len(ip))
	for i, b := range ip {
		parts[i] = fmt.Sprintf("%d", b)
	}
	return strings.Join(parts, ".")
}
func StringerTest1() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
