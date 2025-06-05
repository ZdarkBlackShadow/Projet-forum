package dto

import (
	"projet-forum/models/entity"
)

//dto : data transfert object

type HomeModel struct {
	Userconnected  bool
	ConnectingUser entity.User
}

type ChannelCreation struct {
	Name string
	Description string
	Private bool
}