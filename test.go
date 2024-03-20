package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d ans: %d\n", 53, minNonZeroProduct(53))

}
func minNonZeroProduct(p int) int {
	var total = int64(1<<p) - 1
	var mod int64 = 1e9 + 7
	ans := quickPow(int64(1<<p)-2, total/2, mod)
	ans = ans * ((int64(1<<p) - 1) % mod) % mod
	return int(ans)
}

func quickPow(a, b, mod int64) int64 {
	var ans int64 = 1
	a = a % mod
	for ; b != 0; b >>= 1 {
		if b&1 == 1 {
			ans = (ans * a) % mod
		}
		a = (a * a) % mod
	}
	fmt.Println("pow", ans)
	return ans % mod
}

//

// [001, 010, 011, 100, 101, 110, 111]

// 3 4  st.1
// 3 4  st.2
// 3 4  st.3

// 001 001 001
// 110 110 110 111

// [0001, 0010, 0011, 0100, 0101, 0110, 0111, 1000, 1001, 1010, 1011, 1100, 1101, 1110, 1111]

// st.1 7 8
// st.2 7 8
// st.3 7 8
// st.4 7 8

// 0001 0001 0001 0001 0001 0001 0001

// 1110 1110 1110 1110 1110 1110 1110 1110 1111
