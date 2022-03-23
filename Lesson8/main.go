package main

import (
	"fmt"
	"Lesson8/readconfig"
	"reflect"
)


func main()  {
	config := readconfig.ReadConfig()

	v := reflect.ValueOf(config)
	typeOfS := v.Type()
	for i := 0; i< v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}