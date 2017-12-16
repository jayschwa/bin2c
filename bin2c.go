package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	bytesPerLine := flag.Uint("perline", 16, "Max bytes per line")
	isConst := flag.Bool("const", true, "Define data array as const")
	varName := flag.String("name", "data", "Name of data array")
	typeName := flag.String("type", "unsigned char", "Type of data array")
	flag.Parse()

	src := os.Stdin
	if filename := flag.Arg(0); filename != "" {
		var err error
		src, err = os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer src.Close()
	}

	if *isConst {
		fmt.Print("const ")
	}
	fmt.Printf("%s %s[] = {\n", *typeName, *varName)
	bytes := make([]byte, *bytesPerLine)
	for {
		n, err := src.Read(bytes)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Print("\t")
		for i := 0; i < n; i++ {
			fmt.Printf("0x%02X,", bytes[i])
		}
		fmt.Println()
	}
	fmt.Println("};")
}
