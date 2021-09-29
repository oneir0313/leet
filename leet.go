package main

import (
	"flag"
	"fmt"
	"leet/FreqStack"
	"leet/LRUCache"
	problems "leet/problems"
	"os"
	"reflect"
)

var (
	name    string
)

// -h list
func init() {
	// usage: -p [problem name]
	flag.StringVar(&name, "p", "Default", "the problem `name`")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: leet [Problem Name] \n")
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	m := problems.Problem{}

	fmt.Println("problem name: ", name)
	meth := reflect.ValueOf(m).MethodByName(name)
	if (meth.Kind() != reflect.Func) && name != "FreqStack" && name != "LRUCache" {
		fmt.Println("can not find problem name")
		return
	}

	switch name {
	case "Trap":
		param := []int{0,1,0,2,1,0,1,3,2,1,2,1}
		in := []reflect.Value{reflect.ValueOf(param)}
		res := meth.Call(in)
		fmt.Printf("Answer: %v \n", res[0])
		fmt.Printf("Expected: 6")
	case "SmallestRange":
		param := [][]int{{4,10,15,24,26},{0,9,12,20},{5,18,22,30}}
		in := []reflect.Value{reflect.ValueOf(param)}
		res := meth.Call(in)
		fmt.Printf("Answer: %v \n", res[0])
		fmt.Printf("Expected: [20 24]")
	case "AllPathsSourceTarget":
		param := [][]int{{1,2},{3},{3},{}}
		in := []reflect.Value{reflect.ValueOf(param)}
		res := meth.Call(in)
		fmt.Printf("Answer: %v \n", res[0])
		fmt.Printf("Expected: [[0 1 3] [0 2 3]]")
	case "FreqStack":
		obj := FreqStack.Constructor();
		obj.Push(5);
		obj.Push(7);
		obj.Push(5);
		obj.Push(7);
		obj.Push(4);
		obj.Push(5);
		fmt.Printf("pop: %v, Expected: 5 \n", obj.Pop())
		fmt.Printf("pop: %v, Expected: 7 \n", obj.Pop())
		fmt.Printf("pop: %v, Expected: 5 \n", obj.Pop())
		fmt.Printf("pop: %v, Expected: 4 \n", obj.Pop())
	case "LRUCache":
		obj := LRUCache.Constructor(2);
		obj.Put(1,1);
		obj.Put(2,2);
		fmt.Printf("get 1: %v \n", obj.Get(1))
		obj.Put(3,3);
		fmt.Printf("get 2: %v \n", obj.Get(2))
		obj.Put(4,4);
		fmt.Printf("get 1: %v \n", obj.Get(1))
		fmt.Printf("get 3: %v \n", obj.Get(3))
		fmt.Printf("get 4: %v \n", obj.Get(4))
	default:
		// current debug problem
		meth.Call(nil)
	}
}