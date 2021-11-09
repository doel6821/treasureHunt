package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type coord struct {
	x int
	y int
}

type possibility []coord

var board = [][]string{
	{"#","#","#","#","#","#","#","#"},
	{"#",".",".",".",".",".",".","#"},
	{"#",".","#","#","#",".",".","#"},
	{"#",".",".",".","#",".","#","#"},
	{"#","X","#",".",".",".",".","#"},
	{"#","#","#","#","#","#","#","#"},
}




func main () {
	poss := findPossibility(board)
	treasure := treasureLocation(poss)
	start := coord{x:4, y:1}
	end := coord{}
	printBoard(board)
	fmt.Println("Possibility Treasure : ")
	fmt.Println(findPossibility(board))
	fmt.Println(`
		X is your start position
		# is obstacle
		. is possibility treasure location
		use "A" for moving up/north
		use "B" for moving right/east
		use "C" for moving down/south
		use "D" for moving left/west
	`)

	for treasure != end {
		var x string
		fmt.Printf("input your move : ")
		fmt.Scanf("%s", &x)
		end = move(start, x)
		// fmt.Println(start)
		// fmt.Println(end)
		if checkObstacle(board, end) {
			board[start.x][start.y] = "."
			board[end.x][end.y] = "X"
			start = end
		} else {
			fmt.Println("Ups... You can't move there... ")
		}
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		printBoard(board)
	}
	board[start.x][start.y] = "$"
	fmt.Println("You win, You Found the TREASURE")
	printBoard(board)

}

func printBoard (board [][]string) {
	for _, b := range board {
		fmt.Println(b)
	}
}

func findPossibility (board [][]string) possibility {
	var res possibility
	for i, b := range board {
		for j, c := range b {
			if c == "." {
				res = append(res, coord{i,j})
			}
		}
	}

	return res
}

func randInt(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max-min)
}

func treasureLocation (possibility []coord) coord {
	loc := randInt(0,16)
	var res coord

	for i, p := range possibility{
		if i == loc {
			res = p
		}
	}
	return res

}

func move (start coord, x string) coord {
	var end coord
	switch x {
	case "A" :
		end = coord{
			x : start.x - 1,
			y : start.y,
		}
	case "B" :
		end = coord{
			x : start.x,
			y : start.y + 1,
		}
	case "C" :
		end = coord{
			x : start.x + 1,
			y : start.y,
		}
	case "D" :
		end = coord{
			x : start.x,
			y : start.y -1,
		}
	}
	return end
}

func checkObstacle (board [][]string, pos coord ) bool {
	if board[pos.x][pos.y] == "#" {
		return false
	} else {
		return true
	}
}

