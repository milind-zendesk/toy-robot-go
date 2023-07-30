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
