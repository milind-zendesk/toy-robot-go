package table

import "toy-robot-go/model"

func PlaceRobotOnTable(table *model.Table, x_coord int, y_coord int, face string) {
	robot := model.Robot{
		Position: model.Positioning{
			X_Coords: x_coord,
			Y_Coords: y_coord,
		},
		Facing: face,
	}
	table.Robot = &robot
}
