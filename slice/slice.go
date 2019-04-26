package slice

import "fmt"

func TestSlice() {
	var a = [3]string{"1", "2", "3"}
	var students = [...]string{"a", "b", "c", "d", "e"}

	fmt.Println(a)
	fmt.Println(students)

	for _, v := range students {
		fmt.Println(v)
	}

	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

	car := make([]int, 8)
	car = append(car, 4, 5, 6)
	fmt.Println(car)

	s := a[0:3]
	fmt.Println(s)

	fmt.Println(cap(s))
	fmt.Println(len(s))
}
