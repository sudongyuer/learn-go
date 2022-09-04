package main

func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	println("The value:", m["Answer"])
	m["Answer"] = 48
	println("The value:", m["Answer"])
	delete(m, "Answer")
	println("The value:", m["Answer"])
	v, ok := m["Answer"]
	println("The value:", v, "Present?", ok)
}
