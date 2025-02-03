package main

import (
	"bufio"
	"fmt"
	"os"
)

// Files in Go
// -----------
// In Go, the `os` and `io` packages provide functionality to work with files and directories.
// 
// Operations supported:
// - Reading a file
// - Writing to a file
// - Creating a new file
// - Copying a file
// - Deleting a file
// - Reading directories

func main() {
	// Method - 1 to read a file
	// Open an existing file for reading
	f, err := os.Open("example.txt")
	if err != nil {
		// Log the error
		panic(err)
	}

	// Get file information
	fileInfo, err := f.Stat()
	if err != nil {
		// Log the error
		panic(err)
	}

	// Print file details
	fmt.Println("file name:-", fileInfo.Name())
	fmt.Println("file or folder:-", fileInfo.IsDir())
	fmt.Println("file size:-", fileInfo.Size())
	fmt.Println("file permission:-", fileInfo.Mode())
	fmt.Println("File modified at:-", fileInfo.ModTime())

	defer f.Close()

	// Read file content into a buffer
	buf := make([]byte, fileInfo.Size())
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("data", d, buf)
	for i := 0; i < len(buf); i++ {
		fmt.Println("data", string(buf[i]))
	}

	// Method - 2 to read a file
	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// Read directories
	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	izhar, err := dir.ReadDir(2)
	for _, file := range izhar {
		fmt.Println(file.Name())
	}

	// Method - 1 to create a file
	/* file, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	// Write to a file using slices
	byte := []byte("Hello golang")
	file.Write(byte)
	*/

	// Method - 2 to create and copy a file
	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		e := writer.WriteByte(b)
		if e != nil {
			panic(err)
		}
	}

	writer.Flush()
	fmt.Println("written to new file successfully")

	// Delete a file
	err1 := os.Remove("example2.txt")
	if err1 != nil {
		panic(err1)
	}
}
