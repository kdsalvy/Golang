package loops

// Loop ans its usage
func Loop() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			continue
		}
		println("continuing...")
	}

	for n := 0; n < 5; n++ {
		println(n)
	}

	// infinite loop
	i = 0
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}

	// Looping over collections
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}

	ports := map[string]int{"http": 80, "https": 443}
	for k, v := range ports {
		println(k, v)
	}
}
