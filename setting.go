package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Структура файла конфигурации
type setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPass     string
	PgBase     string
	Data       string
	Assets     string
	HTML       string
}

var cfg setting

// функция init открывает файл конфигурации и конвертирует его из формата JSON в формат string
func init() {
	file, err := os.Open("setting.cfg")
	if err != nil {
		fmt.Println("Ошибка загрузки файла конфигурации")
		panic(err.Error())
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Не удалось прочитать информацию о файле конфигурации")
		panic(err.Error())
	}
	readByte := make([]byte, stat.Size())
	_, err = file.Read(readByte)
	if err != nil {
		fmt.Println("Не удалось прочитать файл конфигурации")
		panic(err.Error())
	}
	err = json.Unmarshal(readByte, &cfg)
	if err != nil {
		fmt.Println("Не удалось считать данные о файле конфигурации")
		panic(err.Error())
	}
}
