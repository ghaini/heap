package main

import (
	"bufio"
	"fmt"
	"heap/heap"
	"log"
	"os"
	"strconv"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Heap Capacity: ")
	capacityString, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}

	capacityString= strings.ReplaceAll(capacityString, "\n", "")
	capacity, err := strconv.Atoi(capacityString)
	if err != nil {
		log.Println(err)
		return
	}

	maxHeap := heap.NewMaxHeap(capacity)
	fmt.Println(`
(1)-Insert Node 	[e.g: 1-50(node value)]
(2)-Delete Node		[e.g: 2-4(node key)]
(3)-Pop Root Node 		
(4)-Get Root Node 		
(5)-Print Heap 		
(6)-Exit`)
	for {
		fmt.Print(`Enter option number: `)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		input= strings.ReplaceAll(input, "\n", "")
		inputSlice := strings.Split(input, "-")
		switch inputSlice[0] {
		case "1":
			val, err := strconv.Atoi(inputSlice[1])
			if err != nil {
				log.Println(err)
				return
			}

			err = maxHeap.InsertNode(val)
			if err != nil {
				log.Println(err)
				return
			}
		case "2":
			val, err := strconv.Atoi(inputSlice[1])
			if err != nil {
				log.Println(err)
				return
			}

			err = maxHeap.DeleteNode(val)
			if err != nil {
				log.Println(err)
				return
			}
		case "3":
			root, err := maxHeap.PopRootNode()
			if err != nil {
				log.Println(err)
				return
			}

			fmt.Print(root)
		case "4":
			root, err := maxHeap.GetRootValue()
			if err != nil {
				log.Println(err)
				return
			}

			fmt.Print(root)
		case "5":
			fmt.Println(maxHeap.ToString())
		case "6":
			fmt.Println("Bye")
			return
		}
	}

}
