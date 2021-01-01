package main

func main() {

	var result = Pop()
	println(result)
}

func Pop() (st *Stack) {
	v := 0
	for ix:= len(st)-1; ix>=0; ix-- {
	if v=st[ix]; v!=0 {
	st[ix] = 0
	return v
	}
	}
	return 0
}