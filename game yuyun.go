// membuat game tictactoe yuyun
package main
import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)
var (
	board=[3][3]string{
		{"","","_",},
		{"","","_",},
		{"","","_",},
	}
	currentPlayer = "X"
	gameOver =false
)
func main() {
	clearScreen()
	fmt.Println("Welcome to Citcattoe!\n")
	printBoard()

	for !gameOver {
		handleTurn(currentPlayer)
		printBoard()
		if checkWin(currentPlayer) {
			fmt.Printf("Player %s wins!\n", currentPlayer)
			gameOver = true
			break
		}
		switchPlayer()
	}
	fmt.Println("\nGame Over. Thank for playing!")
}
func printBoard() {
	fmt.Println()
	for_, row:= range board {
		fmt.Println(strings.Join(row[:], " "))
    }
	fmt.Println()
    }

	func handle turn(player String) {
	     var row,col int
		 for{
			fmt.Printf("Player %S'S turn.\n",player)
			fmt.Print("Enter row (0-2):")
			fmt.Scanln(&col)
			fmt.Print("Enter column(0-2):")
			fmt.Scanln(&col)

			if isvalidMove(row,col) {
				board[row][col]= player
				break
				}else{
					fmt.Println("Invalid move.Please try again.")
				}
			}
		 }
		 func isvalidMove(row, col int) bool {
		if row < 0 || row > 2|| col < 0 || col >2 {
			return false
		}
		return true
		 }
		 func checkWinner(player string) bool {
			for i : 0; i < 3;i++ {
				if board[i][0] == player && board[1][i] == player && board[2][i] == player {
					return true
				}
				if board[0][i]== player && board[1][i]== player board[2][i]== player {
					return true
				}
				if board[0][0]== player && board[1][1]== player board[2][2]== player {
					return true
			}
			if board[0][2]== player && board[1][1]== player board[2][0]== player {
				return true
		 }
		 return false
		}
	}
		func checkDraw() bool {
			for _, row := range board {
			for _, cell := range row {
				if cell =="_" {
					return false
				}	
			}
		}
		return true
	}
	func switchPlayer() {
		if currentPlayer =="X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
	func clearScreen() {
		cmd :=exec.Command("clear")
		//assuming Unix-like system
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}