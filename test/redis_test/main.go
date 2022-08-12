package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"math/rand"
	"strings"
	"time"
)

func GenerateSmsCode(width int) string {
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()
}


func main() {
	smsCode := GenerateSmsCode(6)
	fmt.Println(smsCode)
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("82.157.166.247:6379"),
	})
	rdb.Set("18222199046", smsCode, time.Duration(300)*time.Second)
}
