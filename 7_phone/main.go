package main

// helpful
// https://towardsdev.com/introduction-to-orm-using-gorm-in-golang-d1936a45ffdb

import (
	"fmt"
	"phone_db/db"
	"phone_db/phone"
)

var numbers = []string{
	"1234567890",
	"123 456 7891",
	"(123) 456 7892",
	"(123) 456-7893",
	"123-456-7894",
	"123-456-7890",
	"1234567892",
	"(123)456-7892",
	"(123)456-7898",
}

func main() {
	for _, n := range numbers {
		ph := phone.PhoneNumber{
			RawText: n,
		}
		ph.Normalize()

		db.UpdateUser(ph.Number)
		res := db.ReadUser(ph.Number)
		fmt.Println(res)
	}
}
