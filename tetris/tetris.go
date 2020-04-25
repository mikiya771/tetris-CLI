package tetris

import "fmt"

type Row [10]bool
type Stage [20]Row

func RefreshStage(stage Stage) Stage {
	var ReturnStage Stage
	IndexOfReturnStage := 0
	for _, row := range stage {
		if EvaluateRow(row) == true {
			fmt.Println("nice")
		} else {
			ReturnStage[IndexOfReturnStage] = row
			IndexOfReturnStage++
		}
	}
	return ReturnStage
}
func EvaluateRow(row Row) bool {
	for _, square := range row {
		if square == false {
			return false
		}
	}
	return true
}
