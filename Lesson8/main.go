package main

import (
	"fmt"
	"Lesson8/readconfig"
	"reflect"
	"os"
)


func main()  {
	os.Setenv("DB_URL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	config := readconfig.ReadConfig()

	v := reflect.ValueOf(config)
	typeOfS := v.Type()
	for i := 0; i< v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}