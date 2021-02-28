package main

import (
	"fmt"
	"github.com/malwaredb/gosdhash"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <File> -- Show sdhash for a file\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage: %s <File1> <File2> -- Show sdhash similarity for two files\n", os.Args[0])
		os.Exit(1)
	}

	if len(os.Args) == 2 {
		filePath := os.Args[1]
		fName := filepath.Base(filePath)
		fileObj, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open %s: %s\n", filePath, err)
			os.Exit(2)
		}
		defer fileObj.Close()
		fileContents, err := io.ReadAll(fileObj)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read from %s: %s\n", filePath, err)
			os.Exit(2)
		}

		sdhash := sdhash.SDHash_From_Buffer(fName, fileContents)
		fmt.Printf("%s: %s\n", filePath, sdhash)
		return
	}

	if len(os.Args) == 3 {
		hash1 := sdhash.SDHash_From_FPath(os.Args[1])
		hash2 := sdhash.SDHash_From_FPath(os.Args[2])
		similarity := sdhash.SDHash_Compare_Hashes(hash1, hash2)
		fmt.Printf("Similarity: %d\n", similarity)
	}
}