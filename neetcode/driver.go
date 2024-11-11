package neetcode

import (
	"fmt"
	"fun-with-golang/neetcode/array_hashing"
)

//

func Init() {
	test1 := []int{1, 2, 3, 3}
	test2 := []int{1, 2, 3, 4}

	fmt.Println(array_hashing.ContainsDuplicate(test1))
	fmt.Println(array_hashing.ContainsDuplicate(test2))
}
