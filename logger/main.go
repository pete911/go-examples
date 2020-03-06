package main

func main() {

	name := "production"
	Logf("accessing %s environment", name)
	Error("cannot access environment")
}
