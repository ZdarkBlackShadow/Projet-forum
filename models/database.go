package models

import "time"

type Image struct {
    ImageID int    `gorm:"primaryKey;autoIncrement" json:"image_id"`
    Path    string `gorm:"size:100;not null" json:"path"`
}

type Tag struct {
    TagID     int       `gorm:"primaryKey;autoIncrement" json:"tag_id"`
    Name      string    `gorm:"size:50;unique;not null" json:"name"`
    CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

type State struct {
    StateID int    `gorm:"primaryKey;autoIncrement" json:"state_id"`
    Name    string `gorm:"size:50;unique;not null" json:"name"`
}

type Emoji struct {
    EmojiID int    `gorm:"primaryKey;autoIncrement" json:"emoji_id"`
    Code    string `gorm:"size:50;not null" json:"code"`
}

type Role struct {
    RoleID int    `gorm:"primaryKey;autoIncrement" json:"role_id"`
    Name   string `gorm:"size:50;not null" json:"name"`
}

type Permission struct {
    PermissionID int    `gorm:"primaryKey;autoIncrement" json:"permission_id"`
    Name         string `gorm:"size:50;not null" json:"name"`
}

type UpDownVote struct {
    UpDownVoteID int  `gorm:"primaryKey;autoIncrement" json:"up_down_vote_id"`
    UpDown       bool `gorm:"not null" json:"up_down"`
}

type User struct {
    UserID        int       `gorm:"primaryKey;autoIncrement" json:"user_id"`
    Email         string    `gorm:"size:50;unique;not null" json:"email"`
    Username      string    `gorm:"size:20;unique;not null" json:"username"`
    Password      string    `gorm:"size:300;unique;not null" json:"password"`
    Bio           string    `gorm:"size:500" json:"bio"`
    LastConnection time.Time `gorm:"not null" json:"last_connection"`
    ImageID       int       `gorm:"not null" json:"image_id"`
}

type Message struct {
    MessageTextID int       `gorm:"primaryKey;autoIncrement" json:"message_text_id"`
    Text          string    `gorm:"size:400;not null" json:"text"`
    CreatedAt     time.Time `gorm:"not null" json:"created_at"`
    Edited        bool      `gorm:"not null" json:"edited"`
    DownVote      int       `gorm:"not null" json:"down_vote"`
    UpVote        int       `gorm:"not null" json:"up_vote"`
    Image         bool      `gorm:"not null" json:"image"`
    UserID        int       `gorm:"not null" json:"user_id"`
}

type Channel struct {
    ChannelID   int       `gorm:"primaryKey;autoIncrement" json:"channel_id"`
    Name        string    `gorm:"size:50;unique;not null" json:"name"`
    Description string    `gorm:"size:500;not null" json:"description"`
    CreatedAt   time.Time `gorm:"not null" json:"created_at"`
    Private     bool      `gorm:"not null" json:"private"`
    ImageID     int       `gorm:"not null" json:"image_id"`
    StateID     int       `gorm:"not null" json:"state_id"`
    UserID      int       `gorm:"not null" json:"user_id"`
}

type ChannelTag struct {
    ChannelID int `gorm:"primaryKey" json:"channel_id"`
    TagID     int `gorm:"primaryKey" json:"tag_id"`
}

type MessageImage struct {
    MessageTextID int `gorm:"primaryKey" json:"message_text_id"`
    ImageID       int `gorm:"primaryKey" json:"image_id"`
}

type Reaction struct {
    MessageTextID int `gorm:"primaryKey" json:"message_text_id"`
    EmojiID       int `gorm:"primaryKey" json:"emoji_id"`
}

type Friend struct {
    UserID     int       `gorm:"primaryKey" json:"user_id"`
    UserID1    int       `gorm:"primaryKey" json:"user_id_1"`
    CreatedAt  time.Time `gorm:"not null" json:"created_at"`
}

type UsersWhoCanAccess struct {
    UserID    int `gorm:"primaryKey" json:"user_id"`
    ChannelID int `gorm:"primaryKey" json:"channel_id"`
}

type FriendRequest struct {
    UserID    int       `gorm:"primaryKey" json:"user_id"`
    UserID1   int       `gorm:"primaryKey" json:"user_id_1"`
    CreatedAt time.Time `gorm:"not null" json:"created_at"`
}

type ChannelInvitation struct {
    UserID    int       `gorm:"primaryKey" json:"user_id"`
    UserID1   int       `gorm:"primaryKey" json:"user_id_1"`
    ChannelID string    `gorm:"size:50;not null" json:"channel_id"`
    CreatedAt time.Time `json:"created_at"`
}

type RolePermission struct {
    RoleID       int `gorm:"primaryKey" json:"role_id"`
    PermissionID int `gorm:"primaryKey" json:"permission_id"`
}

type RoleUserChannel struct {
    UserID    int `gorm:"primaryKey" json:"user_id"`
    ChannelID int `gorm:"primaryKey" json:"channel_id"`
    RoleID    int `gorm:"primaryKey" json:"role_id"`
}

type UpDown struct {
    UserID        int `gorm:"primaryKey" json:"user_id"`
    MessageTextID int `gorm:"primaryKey" json:"message_text_id"`
    UpDownVoteID  int `gorm:"primaryKey" json:"up_down_vote_id"`
}
