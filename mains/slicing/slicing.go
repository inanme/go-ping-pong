package main

func main() {
	//useAppendToWrite()
	what00()
}

// a[low : high : max]
func what00() {
	m0 := [5]int{1, 2, 3, 4, 5}

	{
		m := m0[1:3:5] //it controls the resulting slice's capacity by setting it to max - low
		println("--------------")
		println(len(m))
		println(cap(m))
	}
	{
		m := m0[:0:0]
		println("--------------")
		println(len(m))
		println(cap(m))
	}

	{
		println("--------------0:")
		m := m0[0:]
		println(len(m))
		println(cap(m))
	}

	{
		println("--------------:5")
		m := m0[:5] //elem. < 5
		println(len(m))
		println(cap(m))
	}

	{
		println("--------------2:5")
		m := m0[2:5] // elem. 2,3,4
		println(len(m))
		println(cap(m))
	}

	{
		println("--------------")
		m := m0[2:2:2]
		println(len(m))
		println(cap(m))
	}

	{
		println("--------------")
		println(len(m0))
		println(cap(m0))
	}
}

func useAppendToWrite() {
	m := make([]int, 0, 3)
	//m[0] = 10 //is not good for writing
	m = append(m, 10)
	println(len(m))
	println(cap(m))

	for _, i2 := range m {
		println(i2)
	}
}
