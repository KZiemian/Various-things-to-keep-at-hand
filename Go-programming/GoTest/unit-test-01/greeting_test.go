package main

import (
	"fmt"
	"testing"
)

// func Test_SayHello_ValidArgument(t *testing.T) {
// 	name := "Mert"
// 	expected := fmt.Sprintf("Hello %s.", name)
// 	result := sayHello(name)

// 	if result != expected {
// 		t.Errorf("\"sayHello('%s')\" FAILED, expected -> %v, got -> %v",
// 			name, expected, result)
// 	} else {
// 		t.Logf("\"sayHello('%s')\" SUCCEDED, expected -> %v, got -> %v",
// 			name, expected, result)
// 	}

// }

func Test_SayHello_ValidArgument(t *testing.T) {
	inputs := []struct {
		name   string
		result string
	}{
		{name: "Yemeksepeti", result: "Hello Yemeksepeti."},
		{name: "Banabi", result: "Hello Banabi."},
		{name: "Yemek", result: "Hello Yemek."},
	}

	for _, item := range inputs {
		result := sayHello(item.name)

		if result != item.result {
			t.Errorf("\"sayHello('%s')\" failed, expected -> %v, got -> %v", item.name, item.result, result)
		} else {
			t.Logf("\"sayHello('%s')\" succeded, expected -> %v, got -> %v",
				item.name, item.result, result)
		}
	}
}

func Test_EmptyArguments(t *testing.T) {
	expected := "Hello Anonymous."
	result := sayHello("")

	if result != expected {
		t.Errorf("\"sayHello(\"\")\" failed, expected -> %v, got -> %v",
			expected, result)
	} else {
		t.Logf("\"sayHello(\"\")\" succeded, expected -> %v, got -> %v",
			expected, result)
	}

	expected = "Bye Bye Anonymous!"
	result = sayGoodBye("")

	if result != expected {
		t.Errorf("\"sayGoodBye(\"\")\" failed, expected -> %v, got -> %v",
			expected, result)
	} else {
		t.Logf("\"sayGoodBye(\"\")\" succeded, expected -> %v, got -> %v",
			expected, result)
	}

}

func Test_SayGoodBye(t *testing.T) {
	name := "Mert"
	expected := fmt.Sprintf("Bye Bye %s!", name)
	result := sayGoodBye(name)

	if result != expected {
		t.Errorf("\"sayGoodBye('%s')\" FAILED, expected -> %v, got -> %v",
			name, expected, result)
	} else {
		t.Logf("\"sayGoodBye('%s')\" SUCCEDED, expected -> %v, got -> %v",
			name, expected, result)
	}
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sayHello("Yemeksepeti")
	}
}

func ExamplesayHello() {
	fmt.Println(sayHello("Mert"))
	// Output: Hello Mert.
}
