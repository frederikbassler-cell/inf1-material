package duration

import "fmt"

type Duration uint

func FromSeconds(t int) Duration {
	return Duration(t)
}

func FromMinuts(t int) Duration {
	return Duration(t * 60)
}

func FromHours(t int) Duration {

	return FromMinuts(t * 60)
}

func (t Duration) Seconds() int {
	return int(t)
}

func (t Duration) Minuts() int {
	return t.Seconds() / 60
}

func (t Duration) Hours() int {
	return t.Minuts() / 60
}

func (t *Duration) Scale(factor int) {
	//*t = Duration(*t * Duration(factor))
	*t = Duration(t.Seconds() * factor)
}

func Example() {
	//var a Length = 10
	a := FromHours(10)
	var b int = 2


	fmt.Println(a)
	a.Scale(b)
	fmt.Println(a)

	fmt.Println(a.Seconds())
	fmt.Println(a.Hours())

	// Output:
}
