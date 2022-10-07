package db

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RDB = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func AddRedis(redisSet string, taskName string) bool {
	m := redis.Z{
		Member: taskName,
	}
	err := RDB.ZAdd(redisSet, m).Err()
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func RemoveAll() {
	err := RDB.FlushAll().Err()
	if err != nil {
		panic(err)
	}
}

// not used but keeping for reference
func CheckRedisByScore(redisSet string, taskScore int) bool {
	z := redis.ZRangeBy{Min: fmt.Sprintf("%f", float64(taskScore)-.01),
		Max: fmt.Sprintf("%f", float64(taskScore)+.01),
	}
	keys, _ := RDB.ZRangeByScore(redisSet, z).Result()
	if len(keys) == 0 {
		return false
	}
	return true
}

func DeleteRedis(redisSet string, taskName string) bool {
	err := RDB.ZRem(redisSet, taskName).Err()
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func PrintRedis(redisSet string) {
	fmt.Printf("--Printing %v list--\n", redisSet)
	var cursor uint64
	var k []string
	for {
		var keys []string
		var err error
		keys, cursor, err := RDB.ZScan(redisSet, cursor, "*", 10).Result()
		if err != nil {
			panic(err)
		}
		k = append(k, keys...)
		if cursor == 0 {
			break
		}
	}
	var i int
	for _, key := range k {
		if key != "0" {
			i++
			fmt.Printf("%v. %v\n", i, key)

		}
	}
}

func TransferSet(taskName string, set1 string, set2 string) bool {
	res := AddRedis(set2, taskName)
	if !res {
		return res
	}
	res = DeleteRedis(set1, taskName)
	return res

}
