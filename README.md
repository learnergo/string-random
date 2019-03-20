```
package main

import (
	"fmt"
	random "github.com/learnergo/string-random"
)

func main() {
	fmt.Println("大小写字母数字组合：", random.Rand(12))
	fmt.Println("大写字母组合：", random.UpperLetterRand(12))
	fmt.Println("小写字母组合：", random.LowerLetterRand(12))
	fmt.Println("数字组合：", random.DigitRand(12))
	scope := "12345abcdefghijklmnopqrstuvwxyz" // EOS账户，去除'.'
	rand := random.New(scope)
	fmt.Println("自定义组合：", rand.Rand(12))
}
```