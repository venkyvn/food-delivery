package skuser

import (
	socketio "github.com/googollee/go-socket.io"
	"go-food-delivery/common"
	"log"
)

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(requester common.Requester) func(conn socketio.Conn, data LocationData) {
	return func(conn socketio.Conn, data LocationData) {
		// location belong to ???

		log.Println("User update location : user id is ", requester.GetUserId(), "at location", data)
	}
}
