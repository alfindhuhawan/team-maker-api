package util

import "team-maker-api/shared/model/vo"

func CreateMediumSize(dimension vo.Size3D) (int, int, int) {
	length := 30
	if dimension[0] > 0 {
		length = dimension[0]
	}

	width := 22
	if dimension[1] > 0 {
		width = dimension[1]
	}

	height := 12
	if dimension[2] > 0 {
		height = dimension[2]
	}

	return length, width, height
}
