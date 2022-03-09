package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"

	"github.com/oneir0313/leet/FreqStack"
	"github.com/oneir0313/leet/LRUCache"
	problems "github.com/oneir0313/leet/problems"
)

var (
	name string
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
		param := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		in := []reflect.Value{reflect.ValueOf(param)}
		res := meth.Call(in)
		check(6, res[0].Interface())
	case "SmallestRange":
		param := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
		in := []reflect.Value{reflect.ValueOf(param)}
		res := meth.Call(in)
		check([]int{20, 24}, res[0].Interface())
	case "AllPathsSourceTarget":
		param := [][]int{{1, 2}, {3}, {3}, {}}
		in := []reflect.Value{reflect.ValueOf(param)}
		res := meth.Call(in)
		check([][]int{{0, 1, 3}, {0, 2, 3}}, res[0].Interface())
	case "FreqStack":
		obj := FreqStack.Constructor()
		obj.Push(5)
		obj.Push(7)
		obj.Push(5)
		obj.Push(7)
		obj.Push(4)
		obj.Push(5)
		check(5, obj.Pop())
		check(7, obj.Pop())
		check(5, obj.Pop())
		check(4, obj.Pop())
	case "LRUCache":
		obj := LRUCache.Constructor(2)
		obj.Put(1, 1)
		obj.Put(2, 2)
		check(1, obj.Get(1))
		obj.Put(3, 3)
		check(-1, obj.Get(2))
		obj.Put(4, 4)
		check(-1, obj.Get(1))
		check(3, obj.Get(3))
		check(4, obj.Get(4))
	case "GenerateParenthesis":
		param := 3
		in := []reflect.Value{reflect.ValueOf(param)}
		res := meth.Call(in)
		check([]string{"((()))", "(()())", "(())()", "()(())", "()()()"}, res[0].Interface())
	case "UniquePaths":
		in := []reflect.Value{reflect.ValueOf(23), reflect.ValueOf(12)}
		res := meth.Call(in)
		check(193536720, res[0].Interface())
	case "UniquePathsWithObstacles":
		in := []reflect.Value{reflect.ValueOf([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})}
		res := meth.Call(in)
		check(2, res[0].Interface())
		//[[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
	case "UniquePathsIII":
		in := []reflect.Value{reflect.ValueOf([][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}})}
		res := meth.Call(in)
		check(2, res[0].Interface())
	default:
		// current debug problem
		meth.Call(nil)
	}
}

func check(a interface{}, b interface{}) {
	fmt.Printf("your output: %v, expected answer: %v \n", b, a)

	if reflect.TypeOf(a).Kind() == reflect.Slice {
		if reflect.DeepEqual(a, b) {
			fmt.Printf("answer is correct!\n")
			return
		}
		panic("answer is wrong!")
	}
	if a == b {
		fmt.Printf("answer is correct!\n")
		return
	}
	panic("answer is wrong!")
}
