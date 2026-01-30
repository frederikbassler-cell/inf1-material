package time

import (
	"fmt"
	"time"
)

func Example() {

	now := time.Now().Unix()
	minute := (now / 60) % 60
	hour := (now / 60 / 60) % 24

	fmt.Println(now)
	fmt.Println(minute)
	fmt.Println(hour)
}
