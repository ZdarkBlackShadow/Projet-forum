package dto

import (
	"mime/multipart"
)

type HomeModel struct {
	Userconnected  bool
	ConnectingUser User
	PublicCHannel  []Channel
}

type ChannelCreation struct {
	Name        string
	Description string
	Private     bool
}

type Channel struct {
	Id          int
	Name        string
	Description string
	IsPrivate   bool
	Owner       User
	Messages    []Message
	Tags        []string
}

type User struct {
	Id               int
	Name             string
	Bio              string
	ImageID          int
	Friends          []Friend
	FriendsRequests  []FriendRequest
	Channels         []Channel
	ChannlesRequests []ChannelInvitation
}

type Message struct {
	Id        int
	Text      string
	Edited    bool
	Creator   User
	Image     bool
	ImageFile Image
	Reaction  []string
}

type Image struct {
	File    multipart.File
	Handler *multipart.FileHeader
}

type Friend struct {
	Id      int
	Name    string
	Bio     string
	ImageID string
}

type FriendRequest struct {
	Id      int
	Name    string
	Bio     string
	ImageID string
}

type ChannelInvitation struct {
	ChannelId string
	UserId    int
}

type UserInformation struct {
	Email    string
	Username string
	Bio      string
	Image    string
}
