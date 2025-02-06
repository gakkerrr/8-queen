package main

import (
	"fmt"
	"math/rand"
)

type ChessBoard [8][8]int

var (
	numbersOfQueens      int
	chessBoard           ChessBoard
	unattackedPostitions [8][8]bool
	curNumbersOfQueens   int
)

func main() {

	// количество необходимых ферзей на доске
	numbersOfQueens = 8

	// бесконечный цикл пока не найдется решение и он не прервется
	for {
		curNumbersOfQueens = 0
		// с целью spicy outcome первый ферзь добавляется случайно
		setQueenAtRandomCoord(&chessBoard)
		if queens(&chessBoard) {
			printBoard(&chessBoard)
			break
		}
	}
}

func printBoard(board *ChessBoard) {
	for _, row := range board {
		for _, col := range row {
			var printable string
			if col == 2 {
				printable = "👑"
			} else if col == 1 {
				printable = "🟥"
			} else {
				printable = "⬜"
			}
			fmt.Printf("%v", printable)
		}
		fmt.Println()
	}
	fmt.Println("===============================================================")
}

// Принцип работы функции - сначала вся доска обнуляется, кроме ферзей, затем для каждого ферзя на доске отмечаем клетки на траектории его атаки
func checkZonesOfQueen(board *ChessBoard) {

	// Обнуление всех клеток кроме ферзей
	for rowId, row := range board {
		for colId, col := range row {
			if col != 2 {
				board[rowId][colId] = 0
			}
		}
	}

	for rowId, row := range board {
		for colId, col := range row {
			if col == 2 { // Если это ферзь
				// Помечаем атакующие клетки для текущего ферзя
				for i := 0; i < 8; i++ {
					// Горизонтальная линия
					if i != colId {
						board[rowId][i] = 1
					}
					// Вертикальная линия
					if i != rowId {
						board[i][colId] = 1
					}
					// Диагонали
					if rowId+i < 8 && colId+i < 8 && (i != 0) {
						board[rowId+i][colId+i] = 1
					}
					if rowId-i >= 0 && colId-i >= 0 && (i != 0) {
						board[rowId-i][colId-i] = 1
					}
					if rowId+i < 8 && colId-i >= 0 && (i != 0) {
						board[rowId+i][colId-i] = 1
					}
					if rowId-i >= 0 && colId+i < 8 && (i != 0) {
						board[rowId-i][colId+i] = 1
					}
				}
			}
		}
	}
}

func setQueenAtRandomCoord(board *ChessBoard) {
	randX := rand.Intn(8)
	randY := rand.Intn(8)
	board[randX][randY] = 2
	checkZonesOfQueen(board)
	curNumbersOfQueens++
}

func queens(board *ChessBoard) bool {
	if curNumbersOfQueens == numbersOfQueens {
		return true
	}
	unattackedPostitions = updateUnattackedPostitions(board)

	for rowId, row := range unattackedPostitions {
		for colId, col := range row {
			if col {
				placeQueen(board, rowId, colId)
				checkZonesOfQueen(board)
				unattackedPostitions = updateUnattackedPostitions(board)
				if queens(board) {
					return true
				}
				removeQueen(board, rowId, colId)
				checkZonesOfQueen(board)
				unattackedPostitions = updateUnattackedPostitions(board)
			}
		}
	}
	return false
}

func placeQueen(board *ChessBoard, x, y int) {
	board[x][y] = 2
	curNumbersOfQueens++
}

func removeQueen(board *ChessBoard, x, y int) {
	board[x][y] = 0
	curNumbersOfQueens--
}

// Возвращаем двухмерный булевый массив, помечающий неатакованные клетки
func updateUnattackedPostitions(board *ChessBoard) [8][8]bool {
	var unattackedPostitions [8][8]bool
	for rowId, row := range board {
		for colId, col := range row {
			if col == 0 {
				unattackedPostitions[rowId][colId] = true
			}
		}
	}
	return unattackedPostitions
}
