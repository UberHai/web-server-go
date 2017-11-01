package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Shawn Stevens"
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	tpl := fmt.Sprint(`
	<!DOCTYPE HTML>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
		<h1>` + name + `</h1>
		<p>Welcome!</p>
	</body>
	</html>
	`)

	fmt.Println(tpl)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file:", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
