package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

//func main() {
//	// defer statement is executed, and places
//	// fmt.Println("Bye") on a list to be executed prior to the function returning
//	defer fmt.Println("Bye")
//
//	// The next line is executed immediately
//	fmt.Println("Hi")
//
//	// fmt.Println*("Bye") is now invoked, as we are at the end of the function scope
//}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	if err := write("readme.txt", "This is a readme file from PK"); err != nil {
		log.Fatal("failed to write file:", err)
	}
	if err := fileCopy("readme.txt", "readme-copy.txt"); err != nil {
		log.Fatal("failed to copy file: %s")
	}

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	b()
}

func write(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.WriteString(file, text)
	if err != nil {
		//file.Close() --- We can avoid using two file.close statement by applying defer keyword
		return err
	}
	//file.Close()
	return nil
}

func fileCopy(source string, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	n, err := io.Copy(dst, src)
	if err != nil {
		return err
	}
	fmt.Printf("Copied %d bytes from %s to %s\n", n, source, destination)

	if err := src.Close(); err != nil {
		return err
	}

	return dst.Close()
}
