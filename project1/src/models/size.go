package models

import "strconv"

type Size struct {
	width  int
	height int
}

func GetNewSize(width string, height string) (*Size, string) {
	var size *Size
	var errorMessage string

	// client side does not want to resize image
	if width == "" && height == "" {
		return nil, ""
	}

	// there are both values
	if width != "" && height != "" {
		widthVal, err1 := strconv.Atoi(width)
		heightVal, err2 := strconv.Atoi(width)

		if err1 != nil || err2 != nil {
			errorMessage = "Incorrect width and/or height values."
		} else {
			size = &Size{height: heightVal, width: widthVal}
		}
	} else {
		errorMessage = "Must be set width and height simultaneously or none of them."
	}

	return size, errorMessage
}
