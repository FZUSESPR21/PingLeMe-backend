//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package serializer

type Status struct {
	Status bool `json:"status"`
}

func BuildStatusResponse(status bool) Response {
	return Response{
		Code: 0,
		Data: Status{Status: status},
	}
}
