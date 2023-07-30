package robot

import (
	"testing"
	"toy-robot-go/model"

	"github.com/stretchr/testify/assert"
)

func TestLeft_North(t *testing.T) {
	robot := model.Robot{
		Position: model.Positioning{
			X_Coords: 2,
			Y_Coords: 2,
		},
		Facing: "north",
	}
	Left(&robot)

	exp_robot := model.Robot{
		Position: model.Positioning{
			X_Coords: 2,
			Y_Coords: 2,
		},
		Facing: "east",
	}

	assert.Equal(t, exp_robot, robot)
}

func TestMove(t *testing.T) {
	table := model.Table{
		MaxX: 5,
		MaxY: 5,
		Robot: &model.Robot{
			Position: model.Positioning{
				X_Coords: 2,
				Y_Coords: 2,
			},
			Facing: "north",
		},
	}

	Move(&table)

	exp_table := model.Table{
		MaxX: 5,
		MaxY: 5,
		Robot: &model.Robot{
			Position: model.Positioning{
				X_Coords: 2,
				Y_Coords: 3,
			},
			Facing: "north",
		},
	}
	assert.Equal(t, exp_table, table)
}
