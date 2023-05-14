package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Id      string
	Name    string
	Address string
	Job     string
	Reason  string
}

var students []Student = []Student{
	{
		Id:      "S3",
		Name:    "Raihan",
		Address: "Jalan Pinus Raya",
		Reason:  "Karena saya mau jadi backend engineer",
	},
	{
		Id:      "S2",
		Name:    "Yovinka",
		Address: "Jalan Sangiang",
		Reason:  "Karena saya mau buat aplikasi pajak",
	},
	{
		Id:      "S1",
		Name:    "Faliq",
		Address: "Jalan Indralaya",
		Reason:  "Karena golang dibutuhkan",
	},
	{
		Id:      "S1",
		Name:    "Honesty",
		Address: "Jalan Bekasi Timur",
		Reason:  "Karena saya menjadi developer",
	},
	{
		Id:      "S1",
		Name:    "Yohan",
		Address: "Jalan Sangiang Peluit",
		Reason:  "Karena saya ingin menjadi software engineer",
	},
}

func main() {
	var inputs = os.Args

	result, err := FindStudent(inputs[1])

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Printf("%+v \n", result)

}

func FindStudent(name string) (Student, error) {
	for _, value := range students {
		if value.Name == name {
			return value, nil
		}
	}

	return Student{}, errors.New("Student not found")
}
