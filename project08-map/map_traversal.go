package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	fmt.Println("======================================")

	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	fmt.Println("======================================")

	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)

	fmt.Println("======================================")
	fmt.Println(CountWordTime("how do you do"))

	fmt.Println("======================================")
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])

	fmt.Println("======================================")

	n := make(map[string]int)

	n["k1"] = 7
	n["k2"] = 13

	fmt.Println("map:", n)

	v1 := n["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(n))

	delete(n, "k2")
	fmt.Println("map:", n)

	_, prs := n["k2"]
	fmt.Println("prs:", prs)

	n1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n1)
}

func CountWordTime(text string) map[string]int {
	if len(text) == 0 {
		return nil
	}

	words := strings.Fields(text)
	wordTimes := make(map[string]int, len(words))

	// 遍历切片，统计单词出现的次数
	for _, word := range words {
		wordTimes[word]++
	}

	return wordTimes

}
