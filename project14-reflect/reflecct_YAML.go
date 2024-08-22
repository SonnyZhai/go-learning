package main

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"gopkg.in/yaml.v3"
)

// LoadYAML 解析 YAML 文件，将内容加载到传入的结构体或 map 中
func LoadYAML(filename string, output interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(output)
	if err != nil {
		return err
	}

	return nil
}

// printStruct 遍历并打印结构体中的所有字段和值
func printStruct(v reflect.Value, indent string) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		name := fieldType.Name

		switch field.Kind() {
		case reflect.Struct:
			fmt.Printf("%s%s:\n", indent, name)
			printStruct(field, indent+"  ")
		case reflect.Slice:
			fmt.Printf("%s%s:\n", indent, name)
			for j := 0; j < field.Len(); j++ {
				fmt.Printf("%s  - %v\n", indent, field.Index(j).Interface())
			}
		default:
			fmt.Printf("%s%s: %v\n", indent, name, field.Interface())
		}
	}
}

// parseConfig 是通用的 YAML 解析函数，能够解析到结构体或 map 中
func parseConfig(config interface{}) {
	v := reflect.ValueOf(config).Elem()
	fmt.Println("Parsed configuration:")
	if v.Kind() == reflect.Struct {
		printStruct(v, "")
	} else if v.Kind() == reflect.Map {
		printMap(v, "")
	} else {
		fmt.Printf("Unsupported type: %v\n", v.Kind())
	}
}

// printMap 遍历并打印 map 中的所有键值对
func printMap(v reflect.Value, indent string) {
	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		switch value.Kind() {
		case reflect.Map:
			fmt.Printf("%s%s:\n", indent, key)
			printMap(value, indent+"  ")
		case reflect.Slice:
			fmt.Printf("%s%s:\n", indent, key)
			for i := 0; i < value.Len(); i++ {
				fmt.Printf("%s  - %v\n", indent, value.Index(i).Interface())
			}
		default:
			fmt.Printf("%s%s: %v\n", indent, key, value.Interface())
		}
	}
}

// CustomTime 自定义类型，用于演示自定义类型解析
type CustomTime struct {
	time.Time
}

// UnmarshalYAML 实现自定义类型的 UnmarshalYAML 接口
func (ct *CustomTime) UnmarshalYAML(value *yaml.Node) error {
	var timeString string
	if err := value.Decode(&timeString); err != nil {
		return err
	}

	// 自定义时间格式
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		return err
	}

	ct.Time = parsedTime
	return nil
}

// 示例配置结构体
type Config struct {
	AppName    string         `yaml:"app_name"`
	AppPort    int            `yaml:"app_port"`
	Database   DatabaseConfig `yaml:"database"`
	Services   []string       `yaml:"services"`
	StartTimes []CustomTime   `yaml:"start_times"`
}

// DatabaseConfig 嵌套的结构体示例
type DatabaseConfig struct {
	DBName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func reflectYaml() {
	// 示例一：解析到结构体
	var cfg Config
	err := LoadYAML("config.yaml", &cfg)
	if err != nil {
		fmt.Println("Error loading YAML file:", err)
		return
	}

	parseConfig(&cfg)

	// 示例二：解析到 map
	configMap := make(map[string]interface{})
	err = LoadYAML("config.yaml", &configMap)
	if err != nil {
		fmt.Println("Error loading YAML file:", err)
		return
	}

	parseConfig(&configMap)
}
