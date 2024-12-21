package main // Declares the package as 'main', which is necessary for executable programs in Go.

import "fmt" // Imports the "fmt" package, providing functions like Println for output.

func main() { // The 'main' function is the entry point of any Go program.
	fmt.Println("Hello, World") // Prints "Hello, World" to the console.
}

// Additional Information:

// 1. main.go:
//    - This is the file name where the program is saved. It contains the Go source code
//      and is conventionally named for standalone executable programs.
//    - In Go, the program must:
//      a) Belong to the 'main' package.
//      b) Contain a 'main()' function as the entry point.

// 2. go build main.go:
//    - This command compiles the program into a standalone executable file named
//      'main' (on Unix/Linux/macOS) or 'main.exe' (on Windows).
//    - Steps:
//      a) Go checks for syntax and errors.
//      b) Links dependencies like 'fmt'.
//      c) Produces a reusable binary file for execution.
//    - Example Usage:
//        $ go build main.go
//        $ ./main
//        Output: Hello, World

// 3. go run main.go:
//    - This command compiles and runs the program directly without creating a
//      permanent binary file.
//    - Useful for quickly testing or running code without saving the compiled binary.
//    - Internally, it:
//      a) Compiles the program into a temporary binary file.
//      b) Executes the temporary binary.
//      c) Deletes the temporary binary after execution.
//    - Example Usage:
//        $ go run main.go
//        Output: Hello, World

// 4. Difference between go build and go run:
//    - 'go build' creates a binary file that can be reused multiple times, ideal for production.
//    - 'go run' is for quick testing, as it compiles and runs the program on the fly without saving the binary.
