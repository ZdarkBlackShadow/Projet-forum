package entity

import "time"

type Image struct {
	ImageID int    `json:"image_id"`
	Path    string `json:"path"`
}

type Tag struct {
	TagID     int       `json:"tag_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type State struct {
	StateID int    `json:"state_id"`
	Name    string `json:"name"`
}

type Emoji struct {
	EmojiID int    `json:"emoji_id"`
	Code    string `json:"code"`
}

type Role struct {
	RoleID int    `json:"role_id"`
	Name   string `json:"name"`
}

type Permission struct {
	PermissionID int    `json:"permission_id"`
	Name         string `json:"name"`
}

type UpDownVote struct {
	UpDownVoteID int  `json:"up_down_vote_id"`
	UpDown       bool `json:"up_down"`
}

type User struct {
	UserID         int       `json:"user_id"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Bio            string    `json:"bio"`
	LastConnection time.Time `json:"last_connection"`
	ImageID        int       `json:"image_id"`
	Salt           string    `json:"salt"`
}

type Message struct {
	MessageTextID int       `json:"message_text_id"`
	Text          string    `json:"text"`
	CreatedAt     time.Time `json:"created_at"`
	Edited        bool      `json:"edited"`
	Image         bool      `json:"image"`
	UserID        int       `json:"user_id"`
	Channel_id    int       `json:"channel_id"`
}

type Channel struct {
	ChannelID   int       `json:"channel_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Private     bool      `json:"private"`
	ImageID     int       `json:"image_id"`
	StateID     int       `json:"state_id"`
	UserID      int       `json:"user_id"`
}

type ChannelTag struct {
	ChannelID int `json:"channel_id"`
	TagID     int `json:"tag_id"`
}

type MessageImage struct {
	MessageTextID int `json:"message_text_id"`
	ImageID       int `json:"image_id"`
}

type Reaction struct {
	MessageTextID int `json:"message_text_id"`
	EmojiID       int `json:"emoji_id"`
}

type Friend struct {
	UserID    int       `json:"user_id"`
	UserID1   int       `json:"user_id_1"`
	CreatedAt time.Time `json:"created_at"`
}

type UsersWhoCanAccess struct {
	UserID    int `json:"user_id"`
	ChannelID int `json:"channel_id"`
}

type FriendRequest struct {
	UserIdCreator int       `json:"user_id_creator"`
	UserIdInvite  int       `json:"user_id_invite"`
	CreatedAt     time.Time `json:"created_at"`
}

type ChannelInvitation struct {
	UserID    int       `json:"user_id"`
	UserID1   int       `json:"user_id_1"`
	ChannelID string    `json:"channel_id"`
	CreatedAt time.Time `json:"created_at"`
}

type RolePermission struct {
	RoleID       int `json:"role_id"`
	PermissionID int `json:"permission_id"`
}

type RoleUserChannel struct {
	UserID    int `json:"user_id"`
	ChannelID int `json:"channel_id"`
	RoleID    int `json:"role_id"`
}

type UpDown struct {
	UserID        int `json:"user_id"`
	MessageTextID int `json:"message_text_id"`
	UpDownVoteID  int `json:"up_down_vote_id"`
}
