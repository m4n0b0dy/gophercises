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
	var v int
	for i, _ := range make([]int, 1000) {
		if !CheckRedisByScore("todos", i) {
			v = i
			break
		}
	}
	m := redis.Z{
		Member: taskName,
		Score:  float64(v),
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
	for i := 0; i < len(k); i += 2 {
		var section []string
		if i > len(k)-2 {
			section = k[i:]
		} else {
			section = k[i : i+2]
		}
		fmt.Printf("%v. %v\n", section[1], section[0])
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
