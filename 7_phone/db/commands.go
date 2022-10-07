package db

import (
	"fmt"
)

var table = "phone_numbers"

func CreateUser(n string) {
	number := Number{PhoneNumber: n}
	number.FillDefaults()

	db, err := GetDatabaseConnection()
	if err != nil {
		panic(1) //placehodler where real logs would go
	}
	result := db.Table(table).Create(&number)
	if result.Error != nil && result.RowsAffected != 1 {
		panic(1)
	}
}

func ReadUser(m string) string {
	var ns []Number
	db, err := GetDatabaseConnection()
	if err != nil {
		panic(1) //placehodler where real logs would go
	}
	_ = db.Table(table).Where(fmt.Sprintf("phone_number  = '%v'", m)).Find(&ns)
	if len(ns) == 0 {
		fmt.Println("No Matching Results")
		return ""
	}
	return ns[0].PhoneNumber
}

func UpdateUser(m string) {
	db, err := GetDatabaseConnection()
	if err != nil {
		panic(1) //placehodler where real logs would go
	}
	var n Number
	chk := ReadUser(m)
	if chk == "" {
		fmt.Println("Record does not exist, creating")
		CreateUser(m)
		return
	}
	result := db.Table(table).First(&n, fmt.Sprintf("phone_number = '%v'", m))
	if result.Error != nil {
		panic(1)
	} else {
		tx := db.Table(table).Save(&n)
		if tx.RowsAffected == 1 {
			fmt.Printf("updated number %v\n", m)
		}
	}
}
