package main

import (
	"os"
	"fmt"
)

func main() {

	args := os.Args

	if len(args) < 2 || len(args) > 4 {
		
		printUsage(args[0])

	} else {

		db := SetupDatabase()

		switch args[1] {

		case "add":

			if len(args) != 4 {
				printUsage(args[0])
			} else {
 				fmt.Println(AddLink(db, args[2], args[3]))
			}

		case "delete":

			if len(args) != 3 {
				printUsage(args[0])
			} else {
				fmt.Println(DeleteLink(db, args[2]))
			}

		case "show":
			fmt.Println(ShowLinks(db))

		case "serve":
			StartServer(db)

		default: 
			printUsage(args[0])
		}
	}
}

func printUsage(prog string) {

	fmt.Println("Usage:")
	fmt.Print("\t")
	fmt.Print(prog)
	fmt.Println(" add <short URL> <target URL>")
	fmt.Print("\t")
	fmt.Print(prog)
	fmt.Println(" delete <short URL>")
	fmt.Print("\t")
	fmt.Print(prog)
	fmt.Println(" show")
	fmt.Print("\t")
	fmt.Print(prog)
	fmt.Println(" serve")
}