package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 1000000; i++ {
		sum := int64(i) + time.Now().UnixNano()
		str := strconv.FormatInt(sum, 10)
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	}
}
