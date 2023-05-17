package cache

import "fmt"

func TaskViewKey(id uint) string {
	return fmt.Sprintf("view:task:%d", id)
}
