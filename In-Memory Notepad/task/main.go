package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the maximum number of notes: ")
	scanner.Scan()
	temp := scanner.Text()
	cap, _ := strconv.Atoi(temp)
	notes := make([]string, 0, cap)
	for {
		fmt.Print("\nEnter a command and data: ")
		scanner.Scan()
		line := scanner.Text()
		data := strings.SplitN(line, " ", 2)
		command := data[0]
		if line == "exit" {
			fmt.Println("[Info] Bye!\n")
			return
		}
		if command == "create" {
			if len(notes) < cap {
				if len(line) < 7 {
					fmt.Println("[Error] Missing note argument")
				} else {
					notes = append(notes, line[7:])
					fmt.Println("[OK] The note was successfully created")
				}
			} else {
				fmt.Println("[Error] Notepad is full")
			}
		} else if command == "list" {
			if len(notes) == 0 {
				fmt.Println("[Info] Notepad is empty")
			} else {
				for i, note := range notes {
					fmt.Printf("[Info] %d: %s\n", i+1, note)
				}
			}
		} else if command == "clear" {
			notes = make([]string, 0, cap)
			fmt.Println("[OK] All notes were successfully deleted")
		} else if command == "update" {
			if len(data) <= 1 {
				fmt.Println("[Error] Missing position argument")
				continue
			}

			parts := strings.SplitN(data[1], " ", 2)

			if len(parts) < 1 {
				fmt.Println("[Error] Missing position argument")
				continue
			}

			if len(parts) < 2 {
				fmt.Println("[Error] Missing note argument")
				continue
			}

			i, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", parts[0])
				continue
			}

			if i > cap || i < 1 {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", i, cap)
				continue
			}

			if i > len(notes) {
				fmt.Println("[Error] There is nothing to update")
				continue
			}

			notes[i-1] = parts[1]
			fmt.Printf("[OK] The note at position %d was successfully updated\n", i)
		} else if command == "delete" {
			if len(data) < 2 {
				fmt.Println("[Error] Missing position argument")
				continue
			}

			i, err := strconv.Atoi(data[1])
			if i >= len(notes) {
				fmt.Println("[Error] There is nothing to delete")
			} else if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", data[1])
			} else if i > cap || i < 1 {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", i, cap)
			} else {
				notes = append(notes[:i-1], notes[i:]...)
				fmt.Printf("[OK] The note at position %d was successfully deleted\n", i)
			}
		} else {
			fmt.Println("[Error] Unknown command")
		}
	}
}
