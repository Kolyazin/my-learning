package main

import (
	"fmt"
	"Lesson9/readconfig"
	"reflect"
)


func main()  {
	config, err := readconfig.ReadConfig()

	if err != nil {
		fmt.Println("Ошибка при чтении конфигурации.")
		return
	}

	v := reflect.ValueOf(config)
	typeOfS := v.Type()
	for i := 0; i< v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}