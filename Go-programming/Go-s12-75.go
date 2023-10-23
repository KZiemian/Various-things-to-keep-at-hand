package main

import (
	"fmt"
	"reflect"
)

// func (v Value) Interface() interface{}

// func main() {
// 	// var x float64 = 3.4

// 	// v := reflect.ValueOf(x)

// 	// y, ok := v.Interface().(float64)

// 	// fmt.Printf("v: %v.\n", v)
// 	// fmt.Printf("y: %v.\n", y)
// 	// fmt.Printf("ok: %v.\n\n", ok)

// 	// fmt.Printf("v.Interface(): %v.\n", v.Interface())

// 	// fmt.Printf("value is %7.1e.\n", v.Interface())

// 	// var x float64 = 3.4
// 	// v := reflect.ValueOf(x)
// 	// v.SetFloat(7.1)

// 	// fmt.Printf("v: %v.\n", v)

// 	// v := reflect.ValueOf(x)
// 	// fmt.Println("settability of v:", v.CanSet())

// 	var x float64 = 3.4

// 	p := reflect.ValueOf(&x)

// 	fmt.Println("type of p:", p.Type())
// 	fmt.Println("settability of p:", p.CanSet())

// 	v := p.Elem()

// 	// fmt.Println("settability of v:", v.CanSet())

// 	v.SetFloat(7.1)

// 	fmt.Println(v.Interface())
// 	fmt.Println(x)
// }

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()

	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}

// f("hello", "world")
// go f("hello", "world")
// g()

timerChan := make(chan time.Time)

go func() {
	time.Sleep(deltaT)

	timeChan <- time.Now()
}()

completedAt := <-timeChan

select {
case v := <-ch1:
	fmt.Println("channel 1 sends", v)

case v := <-ch2:
	fmt.Println("channel 2 sends", v)

default:
	fmt.Println("neither channel was ready")
}

func Compose(f, g func(x float) float) func(x float) flota {
	return func(x float) float {
		return f(g(x))
	}
}

print(Compose(sin, cos)(0.5))
