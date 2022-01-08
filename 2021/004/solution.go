package main

import (
	"fmt"
	"strconv"
	// "strconv"
	"strings"

	"github.com/opti21/AOC/utils"
)

func main() {
	input := utils.InputFile("./input.txt")

	resultA := make(chan int64)
	resultB := make(chan int64)

	go puzzleA(input, resultA)
	go puzzleB(input, resultB)

	fmt.Println(<-resultA)
	fmt.Println(<-resultB)
}

type board [5][5]string

func puzzleA(input []string, result chan int64) {
    boardInputs := input[2:]
    boards := []board{}

    // Make boards
    for i := 0; i < len(boardInputs) - 1; i += 6 {
        newBoard := board{}
            //fmt.Println(boardInputs[i])
            // Iterate through each row in the group of 5 rows
             for rowI := 0; rowI < 5; rowI++ {
                 currLine := boardInputs[i + rowI]
                    // fmt.Println(i + rowI)
                 // add each num in row to new board
                    // fmt.Println(currLine)
                 for colI := 0; colI < 14; colI += 3 {
                     // fmt.Printf("row: %v, col: %v, num: %s \n", rowI, colI / 3, currLine[colI: colI + 2])
                     newBoard[rowI][colI / 3] = currLine[colI: colI + 2]
               }
             }
        boards = append(boards, newBoard)
    }

    // Choose number and mark it on the boards
    chosenNums := strings.Split(input[0], ",")
    rowBingo := false
    colBingo := false
    bingo := [5]string{" *"," *"," *"," *"," *",}
    var winningBoard board
    var winningNum int

    for _, chosenNumStr := range chosenNums {

        chosenNum, err := strconv.Atoi(chosenNumStr)
        if err != nil {
            panic(err)
        }

        fmt.Println("---NEW---")
        fmt.Println(chosenNum)
        // Go through each board
        for boardI, board := range boards {
            if rowBingo || colBingo {
                break
            }
            // Check each num in board to see it matches chosen num
            for rowI := 0; rowI < 5; rowI++ {
                for colI := 0; colI < 5; colI++ {
                    // remove extra spaces
                    if board[rowI][colI] != " *" {
                        strippedNum := strings.TrimSpace(board[rowI][colI])
                        currNum, currErr := strconv.Atoi(strippedNum)
                        if currErr != nil {
                            panic(currErr)
                        }

                        if currNum == chosenNum {
                            boards[boardI][rowI][colI] = " *"
                            // fmt.Printf(" * ")
                        } else {
                            boards[boardI][rowI][colI] = board[rowI][colI]
                            // fmt.Printf("%v ", currNum)
                        }
                    } else {
                        continue
                    }
                    
                }
                //fmt.Printf("\n")
            }

            //fmt.Printf("\n")
        }

        // check board for bingos

        
        for _, b := range boards {

            for rowI := 0; rowI < 5; rowI++ {
                currRow := b[rowI]

                if currRow == bingo {
                    rowBingo = !rowBingo
                    fmt.Println("BINGO")
                    printBoard(b)
                    winningBoard = b
                    winningNum = chosenNum
                    break
                }

            }

            if rowBingo {
                break
            }

            for colI := 0; colI < 5; colI++ {
                currCol := [5]string{
                    b[0][colI],
                    b[1][colI],
                    b[2][colI],
                    b[3][colI],
                    b[4][colI],
                }

                bingo := [5]string{" *"," *"," *"," *"," *",}

                if currCol == bingo {
                    colBingo = !colBingo
                    fmt.Println("BINGO")
                    printBoard(b)
                    winningBoard = b
                    winningNum = chosenNum
                    break
                }
            }

            if colBingo{
                break
            }

            
            // printBoard(b)
        }

        if rowBingo || colBingo {
            break
        }
    }
    fmt.Println("WINNER")
    printBoard(winningBoard)
    winningSum := 0

    
    for rowI := 0; rowI < 5; rowI++ {
        for colI := 0; colI < 5; colI++ {
            if winningBoard[rowI][colI] != " *" {
                strippedNum := strings.TrimSpace(winningBoard[rowI][colI])
                currNum, err := strconv.Atoi(strippedNum)
                if err != nil {
                   panic(err) 
                }
                winningSum += currNum
            }
        }
    }
    

    // fmt.Println(chosenNums)
    fmt.Println(winningSum)
    fmt.Println(winningNum)

    result <- int64(winningSum * winningNum)

}

func printBoard(b board) {
    for rowI := 0; rowI < 5; rowI++ {
        for colI := 0; colI < 5; colI++ {
            fmt.Printf("%v ", b[rowI][colI])
        }
        fmt.Println()
    }
    fmt.Println()
}

func puzzleB(input []string, result chan int64) {

    result <- 0

}
