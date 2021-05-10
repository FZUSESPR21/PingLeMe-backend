//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package util

import "strconv"

type InterfaceTypeErr struct {
	Name string
}

type RecordAlreadyExistErr struct {
	Row int
}

func (err *InterfaceTypeErr) Error() string {
	return "unknown interface " + err.Name + " ."
}

func (err *RecordAlreadyExistErr) Error() string {
	return "record already exist in " + "row :" + strconv.Itoa(err.Row+2) + "."
}
