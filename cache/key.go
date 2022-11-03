package cache

import (
	"fmt"
	"strconv"
)

const (
	RankKey = "rank"
)

func TaskViewKey(id uint) string {
	fmt.Printf("view:task:%s", strconv.Itoa(int(id)))
	return fmt.Sprintf("view:task:%s", strconv.Itoa(int(id)))
}
