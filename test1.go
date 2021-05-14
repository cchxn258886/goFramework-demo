package main

const (
	a = iota//0
	b = iota//1
)
const (
	name = "menglu"
	c    = iota //1
	d    = iota//2
)
func main() {
	//var ss = "ABCDEFGA";
	//print(strings.Count(ss, "A"))
	//print(len(ss))
	var cha1 = make(chan struct{})
	var s = "";
	go func() {
		s = "go func string!"
		println(s)
		close(cha1)
	}()
	<-cha1
	println("main",s)
}


