package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// LoadINI 解析INI文件，将其内容加载到结构体或 map 中
func LoadINI(filename string, output interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 反射获取传入结构体或 map 的类型和值
	val := reflect.ValueOf(output).Elem()

	section := ""
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 跳过空行和注释
		if len(line) == 0 || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		// 处理节 (section)
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			section = strings.TrimSpace(line[1 : len(line)-1])
			continue
		}

		// 处理键值对
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// 在输出结构体或 map 中查找匹配的字段
		assignValue(val, section, key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// assignValue 根据反射设置字段值，适配结构体和 map
func assignValue(val reflect.Value, section, key, value string) {
	switch val.Kind() {
	case reflect.Struct:
		assignStructValue(val, section, key, value)
	case reflect.Map:
		assignMapValue(val, section, key, value)
	}
}

// assignStructValue 给结构体字段赋值
func assignStructValue(val reflect.Value, section, key, value string) {
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("ini")

		// 匹配 section 和 key
		if tag == fmt.Sprintf("%s.%s", section, key) {
			fieldVal := val.Field(i)
			if fieldVal.CanSet() {
				setFieldValue(fieldVal, value)
			}
		}
	}
}

// assignMapValue 给 map 赋值
func assignMapValue(val reflect.Value, section, key, value string) {
	if val.IsNil() {
		val.Set(reflect.MakeMap(val.Type()))
	}

	// 构造 map 的键值
	fullKey := section + "." + key
	val.SetMapIndex(reflect.ValueOf(fullKey), reflect.ValueOf(value))
}

// setFieldValue 根据类型设置字段的值
func setFieldValue(field reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, _ := strconv.ParseInt(value, 10, 64)
		field.SetInt(intValue)
	case reflect.Float32, reflect.Float64:
		floatValue, _ := strconv.ParseFloat(value, 64)
		field.SetFloat(floatValue)
	case reflect.Bool:
		boolValue, _ := strconv.ParseBool(value)
		field.SetBool(boolValue)
	}
}

func main() {
	// 示例一：使用结构体
	type Config struct {
		AppName  string `ini:"app.name"`
		AppPort  int    `ini:"app.port"`
		Database string `ini:"database.name"`
		User     string `ini:"database.user"`
		Password string `ini:"database.password"`
	}

	var cfg Config
	err := LoadINI("config.ini", &cfg)
	if err != nil {
		fmt.Println("Error loading INI file:", err)
		return
	}

	fmt.Println("Using struct:")
	fmt.Printf("AppName: %s\n", cfg.AppName)
	fmt.Printf("AppPort: %d\n", cfg.AppPort)
	fmt.Printf("Database: %s\n", cfg.Database)
	fmt.Printf("User: %s\n", cfg.User)
	fmt.Printf("Password: %s\n", cfg.Password)

	// 示例二：使用 map
	configMap := make(map[string]interface{})
	err = LoadINI("config.ini", &configMap)
	if err != nil {
		fmt.Println("Error loading INI file:", err)
		return
	}

	fmt.Println("\nUsing map:")
	for k, v := range configMap {
		fmt.Printf("%s: %v\n", k, v)
	}

	fmt.Println("===============================")

	reflectYaml()
}
