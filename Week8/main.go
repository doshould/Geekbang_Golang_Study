package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/hhxsv5/go-redis-memory-analysis"
)

var cli redis.UniversalClient

const (
	ip   string = "127.0.0.1"
	port uint16 = 6379
)

func init() {
	cli = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%v:%v", ip, port),
		Password:     "",
		DB:           0,
		PoolSize:     256,
		MinIdleConns: 100,
		MaxRetries:   5,
	})
}

func main() {
	Write(10000, "10_10k", InitValue(10))
	Write(100000, "10_100k", InitValue(10))
	Write(500000, "10_500k", InitValue(10))

	Write(10000, "1000_10k", InitValue(1000))
	Write(100000, "1000_100k", InitValue(1000))
	Write(500000, "1000_500k", InitValue(1000))

	Write(10000, "5000_10k", InitValue(5000))
	Write(100000, "5000_100k", InitValue(5000))
	Write(500000, "5000_500k", InitValue(5000))

	Analysis()
}

func Write(number int, key, value string) {
	for i := 0; i < number; i++ {
		k := fmt.Sprintf("%s:%v", key, i)
		cmd := cli.Set(k, value, -1)
		err := cmd.Err()
		if err != nil {
			fmt.Println(cmd.String())
		}
	}
}

func Analysis() {
	analysis, err := gorma.NewAnalysisConnection(ip, port, "")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer analysis.Close()
	analysis.Start([]string{":"})
	err = analysis.SaveReports("./report")
	if err != nil {
		fmt.Println("error:", err)
	}
}

func InitValue(size int) string {
	array := make([]byte, size)
	for i := 0; i < size; i++ {
		array[i] = 'x'
	}
	return string(array)
}
