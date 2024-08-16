package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func Main() {
	left1 := dispatchCoin()
	left := dispatchCoinWithMap()
	fmt.Println("金币分配结果：", distribution)
	fmt.Println("剩下：", left)
	fmt.Println("剩下：", left1)
}

func dispatchCoin() int {
	totalCoins := coins

	for _, user := range users {
		for _, userLetter := range user {
			switch userLetter {
			case 'e', 'E':
				distribution[user] += 1
				totalCoins -= 1
			case 'i', 'I':
				distribution[user] += 2
				totalCoins -= 2
			case 'o', 'O':
				distribution[user] += 3
				totalCoins -= 3
			case 'u', 'U':
				distribution[user] += 4
				totalCoins -= 4
			}
		}
	}

	return totalCoins
}

func dispatchCoinWithMap() int {

	// 定义金币分配规则,
	// 一个 rune 值实际上是一个 int32 类型的整数，用于表示一个 Unicode 码点。
	// 因此，一个 rune 可以表示任何一个字符，包括 ASCII 字符、汉字、特殊符号等。
	rules := map[rune]int{
		'e': 1, 'E': 1,
		'i': 2, 'I': 2,
		'o': 3, 'O': 3,
		'u': 4, 'U': 4,
	}

	totalCoins := coins

	// 遍历每个用户，计算他们应该分到的金币
	for _, user := range users {
		for _, letter := range user {
			if value, ok := rules[letter]; ok {
				distribution[user] += value
				totalCoins -= value
			}
		}
	}

	return totalCoins
}
