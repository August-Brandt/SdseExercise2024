package main

func main() {
	list := LinkedList{}
	list.Insert("Hello")
	list.Insert("There")
	list.Insert("General")
	list.Insert("Kenobi")
	list.Print()
	list.DeleteLast()
	list.Print()
}