package qsort

import "fmt"

func QSort(l []int) {

	// Sonderfall: Die Liste ist leer oder hat nur ein Element.
	if len(l) <= 1 {
		return
	}

	pivot := l[0]

	left := []int{}
	right := []int{}

	// Partitionieren der Liste:
	// KLeinere Elemente als das Pivot nach links, größere nach rechts
	for _, el := range l[1:] {
		if el < pivot {

			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}
	QSort(left)
	QSort(right)

	// Elemente in die ursprünliche Liste zurückkopieren
	for i, el := range left {
		l[i] = el
	}

	l[len(left)] = pivot

	for i, el := range right {
		l[i+len(left)+1] = el

	}

}

func ExampleQSort() {

	l1 := []int{42, 7, 19, 3, 88, 1, 56, 23, 91, 14, 67, 5, 99, 32, 11, 73, 28, 4, 60}

	QSort(l1)
	fmt.Println(l1)

	// Output:
	// [1 3 4 5 7 11 14 19 23 28 32 42 56 60 67 73 88 91 99]
}


