package main

import (
	"fmt"
	models "golang-teko/model"
	"math"
)

type ListSolution struct {
	roomId int
	chair  [][]int
}

var listSolutionForRoom []ListSolution

func choiceChair(row int, column int, chair [][]int, ticket models.Ticket, distance int, roomId int) {

	var checkDone bool
	checkDone = true
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if math.Abs(float64(ticket.Row-i))+math.Abs(float64(ticket.Column-j)) >= float64(distance) {
				checkDone = false
				break
			}
		}
	}
	if checkDone {
		listSolutionForRoomObject := ListSolution{roomId, chair}
		listSolutionForRoom = append(listSolutionForRoom, listSolutionForRoomObject)
	} else {

		for i := 0; i < row; i++ {
			for j := 0; j < column; j++ {
				if math.Abs(float64(ticket.Row-i))+math.Abs(float64(ticket.Column-j)) < float64(distance) {
					chair[i][j] = 2
				}
			}
		}
		for i := 0; i < row; i++ {
			for j := 0; j < column; j++ {
				if math.Abs(float64(ticket.Row-i))+math.Abs(float64(ticket.Column-j)) == float64(distance) && chair[i][j] != 2 {
					chair[i][j] = 1
					ticketNew := models.Ticket{i, j, true, roomId}
					choiceChair(row, column, chair, ticketNew, distance, roomId)
					chair[i][j] = 0
				}
			}
		}
	}

}

func solution(row int, column int, distance int, roomId int) int {
	var chair [][]int
	//danh dau vi tri 0,0 dc chon dau tien
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			chair[i][j] = 0
		}
	}
	chair[0][0] = 1
	ticket := models.Ticket{0, 0, true, roomId}
	choiceChair(row, column, chair, ticket, distance, roomId)
	best := 0
	var position int
	position = -1
	for i := 0; i < len(listSolutionForRoom); i++ {
		if listSolutionForRoom[i].roomId == roomId {
			cal := 0
			for k := 0; k < len(chair); k++ {
				for h := 0; h < len(chair[k]); h++ {
					if chair[k][h] == 1 {
						cal = cal + 1
					}
				}
			}
			if cal > best {
				best = cal
				position = i
			}
		}
	}
	return position
}

func main() {
	fmt.Println("Moi Ban nhap thong tin.")
	fmt.Println("Moi ban nhap room Id:")
	var roomId, row, column, distance int
	_, err := fmt.Scanf("%d", &roomId)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Moi Ban nhap so hang")
	_, err = fmt.Scanf("%d", &row)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Moi Ban nhap so cot")
	_, err = fmt.Scanf("%d", &column)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Moi Ban nhap khoang cach")
	_, err = fmt.Scanf("%d", &column)
	if err != nil {
		fmt.Println(err.Error())
	}
	var result int
	result = solution(row, column, distance, roomId)
	fmt.Println("Ket qua ghe ngoi so 1 la noi ban co the ngoi")
	for i := 0; i < len(listSolutionForRoom[result].chair); i++ {
		for j := 0; j < len(listSolutionForRoom[result].chair[i]); j++ {
			fmt.Printf("%d ", listSolutionForRoom[result].chair[i][j])
		}
		fmt.Printf("\n")
	}

}
