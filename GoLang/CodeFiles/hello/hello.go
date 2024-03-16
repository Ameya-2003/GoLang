//main package

/*We organize Go programs in packages.

Each .go file first declares which package it is part of.

A package can be composed by multiple files, or just one file.

A program can contain multiple packages.

The main package is the entry point of the program and identifies an executable program.

*/

//fmt

/*We use the import keyword to import a package.

fmt is a built-in package provided by Go that provides input/output utility functions.

We have a large standard library ready to use that we can use for anything from network connectivity to math, crypto, image processing, filesystem access, and more.

*/

package main

import "fmt"

func amain() {
	fmt.Println("Hello, World!")
}
