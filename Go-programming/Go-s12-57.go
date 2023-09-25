package main

import "fmt"

func main() {
	err := openFile()

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	defer closeFile()


	err = readData()

	if err != nil {
		fmt.Println(err.Error())

		return
	}


	err = writeData()

	if err != nil {
		fmt.Println(err.Error())

		return
	}


	fmt.Println("no errors")
}

func openFile() error {
	fmt.Println("opening file connection")

	return nil
}

func closeFile() {
	fmt.Println("closing file connection")
}

func readData() error {
	fmt.Println("reading file data")

	return nil
}

func writeData() error {
	fmt.Println("writing file data")

	return nil
}
