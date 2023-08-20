package collections

import (
	"fmt"

	"github.com/kd/learn/go/collections/model"
)

// Demo is exported function. Exported function should start with capital letter and comment is must with first letter as function name
func Demo() {
	// arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("Array", arr)

	// slices
	s1 := arr[:]
	fmt.Println("Slice", s1)
	s2 := arr[1:]
	fmt.Println("Slice2", s2)
	s3 := arr[:2]
	fmt.Println("Slice3", s3)
	s4 := arr[1:2]
	fmt.Println("Slice4", s4)

	// maps
	m := map[string]int{"foo": 47}
	fmt.Println("Map", m)
	fmt.Println("Get key foo", m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println("Map after deleting", m)

	// structs
	var u model.User
	u.ID = 1
	u.FirstName = "Saurabh"
	u.LastName = "Kedia"
	fmt.Println("Struct", u)

	u2 := model.User{ID: 2,
		FirstName: "Arthur",
		LastName:  "Dent",
	}
	fmt.Println(u2)
}
