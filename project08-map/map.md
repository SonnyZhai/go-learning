## 映射关系容器-map

map是一种`无序`的基于key-value的数据结构，Go语言中的map是`引用类型`，必须初始化才能使用。

---

> <h3 style="text-align: center;"> 1. map定义 </h3>

`1. map[KeyType]ValueType`

map类型的变量默认初始值为nil,需要使用make()函数来分配内存。

`2. make(map[KeyType]ValueType, [cap])`

cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

---

> <h3 style="text-align: center;"> 2. map基本操作 </h3>

```go
func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap) // map[张三:90 小明:100]
	fmt.Println(scoreMap["小明"]) // 100
	fmt.Printf("type of a:%T\n", scoreMap) // type of a:map[string]int
}
```

```go
func main() {
	userInfo := map[string]string{
		"username": "沙河小王子",
		"password": "123456",
	}
	fmt.Println(userInfo) // map[password:123456 username:沙河小王子]
}
```

---

> <h3 style="text-align: center;"> 3. 判断某个键是否存在 </h3>

`value, ok := map[key]`

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
```

---

> <h3 style="text-align: center;"> 4. map的遍历 </h3>

**遍历map时的元素顺序与添加键值对的顺序无关。**

`for range`

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k, v := range scoreMap {
		fmt.Println(k, v) // 张三 90 小明 100 娜扎 60
	}
}
```

`遍历key`

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k := range scoreMap {
		fmt.Println(k) // 张三 小明 娜扎
	}
}
```

---

> <h3 style="text-align: center;"> 5. 使用delete()函数删除键值对 </h3>

`delete(map, key)`

```go
func main(){
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	delete(scoreMap, "小明")//将小明:100从map中删除
	for k,v := range scoreMap{
		fmt.Println(k, v)
	}
}
```

---

> <h3 style="text-align: center;"> 6. 按照指定顺序遍历map </h3>

```go
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
}
```

---

> <h3 style="text-align: center;"> 7. 元素为map类型的切片 </h3>

```go
func main() {
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
}
```

---

> <h3 style="text-align: center;"> 8. 值为切片类型的map </h3>

```go
func main() {
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
}
```