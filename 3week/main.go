package main

import (
	prof "3week/pprof"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
)

func main() {
	log.Println(prof.FindPrime(109), "res")

	contents, err := ioutil.ReadFile("note.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print(prof.Solve(contents))

	//marshall json
	//omitempty
	type User struct {
		Id     int    `json:"id",string`
		Name   string `json:",omitempty"`
		Income int    `json:"-"`
	}
	//case 1 send object -> client
	// user := &User{1, "NameStruct", 22}
	// // resz, _ := json.Marshal(s)

	counterfeit := `[
		{"id":12, "name":"alisa"},
		{"id":11, "address": "pankov12", "age" : 18}
	]`

	//2 case json receive with random key - use empty interface

	var interf interface{}
	data := []byte(counterfeit)
	json.Unmarshal(data, &interf)

	//  m := make(map[string]interface{})

	// log.Println(interf, "json random key, use interface")

	//type assertion - any struct or field struct -> check -> then get value
	// утверждение типа - является то что пришло, тем типом который я проверяю - swithc case or IF

	//dynamic values - work
	//reflect ------------ expensive
	// способность программы исследовать собственную структуру, в особенности через типы
	// обход полей у структур, в рантайме -
	//динамическая распоковка в стуруктуру
	//cast - []byte type int, value, type string value, - reflect switch case -> readBuf - cast

	// printReflect(user)

	//or use codogeneration
	//выполняется не в рантайме, генерирует друггую программу
	// дешевле чем reflect и буыстрее, созадаем дерево,и по нашим правилам созадаем структуры и фукнции,

	// Увеличение скорости работы кода, то есть для замены рефлексии;
	// Уменьшение рутинной работы программиста (и ошибок, связанных с ней);
	// Реализацию оберток по заданным правилам.

}

func printReflect(i interface{}) error {
	value := reflect.ValueOf(i).Elem() //get all value inside struct
	log.Println(value.NumMethod(), "coutn obj in interface")
	for i := 0; i < value.NumField(); i++ {
		valueField := value.Field(i)
		typeField := value.Type().Field(i)
		log.Println(valueField, typeField.Name, typeField.Tag, typeField.Type.Kind(), "val field in structure")
	}
	return nil
}

// empty interface

//slice, map, chan - link type
// task marhsal encoding - use - empty interface,
//codogeneration?
