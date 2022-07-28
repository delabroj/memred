package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mem := newMapStack()

outer:
	for {
		reader := bufio.NewReader(os.Stdin)
		rawCommand, _ := reader.ReadString('\n')
		command := strings.TrimSuffix(rawCommand, "\n")
		commandParts := strings.Split(command, " ")

		switch commandParts[0] {
		case "":

		case "END":
			if len(commandParts) != 1 {
				fmt.Println("ERR: END does not expect arguments")
				continue
			}
			break outer

		case "SET":
			if len(commandParts) != 3 {
				fmt.Println("ERR: SET expects 2 arguments")
				continue
			}
			mem.set(commandParts[1], commandParts[2])

		case "GET":
			if len(commandParts) != 2 {
				fmt.Println("ERR: GET expects 1 argument")
				continue
			}
			fmt.Println(mem.get(commandParts[1]))

		case "UNSET":
			if len(commandParts) != 2 {
				fmt.Println("ERR: UNSET expects 1 argument")
				continue
			}
			mem.unSet(commandParts[1])

		case "NUMEQUALTO":
			if len(commandParts) != 2 {
				fmt.Println("ERR: NUMEQUALTO expects 1 argument")
				continue
			}
			fmt.Println(mem.numEqualTo(commandParts[1]))

		case "BEGIN":
			if len(commandParts) != 1 {
				fmt.Println("ERR: BEGIN does not expect arguments")
				continue
			}
			mem.begin()

		case "COMMIT":
			if len(commandParts) != 1 {
				fmt.Println("ERR: COMMIT does not expect arguments")
				continue
			}
			ok := mem.commit()
			if !ok {
				fmt.Println("NO TRANSACTION")
			}

		case "ROLLBACK":
			if len(commandParts) != 1 {
				fmt.Println("ERR: ROLLBACK does not expect arguments")
				continue
			}
			ok := mem.rollBack()
			if !ok {
				fmt.Println("NO TRANSACTION")
			}

		// case "STACK":
		// 	fmt.Print(mem)

		default:
			fmt.Println("ERR: Unknown command")
		}
	}
}
