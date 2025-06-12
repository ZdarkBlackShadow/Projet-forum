package mapper

import (
	"projet-forum/models/dto"
	"projet-forum/models/entity"
)

// UserEntityToDTO convertit une entité User en DTO User
func UserEntityToDTO(e entity.User) dto.User {
	return dto.User{
		Id:      e.UserID,
		Name:    e.Username,
		Bio:     e.Bio,
		ImageID: e.ImageID,
	}
}

// ChannelEntityToDTO convertit une entité Channel en DTO Channel
func ChannelEntityToDTO(e entity.Channel, owner dto.User, tags []string) dto.Channel {
	return dto.Channel{
		Id:          e.ChannelID,
		Name:        e.Name,
		Description: e.Description,
		IsPrivate:   e.Private,
		Owner:       owner,
		Tags:        tags,
	}
}

// MessageEntityToDTO convertit une entité Message en DTO Message
func MessageEntityToDTO(e entity.Message, creator dto.User, image dto.Image, reactions []string) dto.Message {
	return dto.Message{
		Id:        e.MessageTextID,
		Text:      e.Text,
		Edited:    e.Edited,
		Creator:   creator,
		Image:     e.Image,
		ImageFile: image,
		Reaction:  reactions,
	}
}

// FriendEntityToDTO convertit une entité Friend en DTO Friend
func FriendEntityToDTO(e entity.User) dto.Friend {
	return dto.Friend{
		Id:      e.UserID,
		Name:    e.Username,
		Bio:     e.Bio,
		ImageID: string(rune(e.ImageID)), // à adapter si nécessaire
	}
}

// FriendRequestEntityToDTO convertit une entité FriendRequest en DTO FriendRequest
func FriendRequestEntityToDTO(e entity.User) dto.FriendRequest {
	return dto.FriendRequest{
		Id:      e.UserID,
		Name:    e.Username,
		Bio:     e.Bio,
		ImageID: string(rune(e.ImageID)), // à adapter si nécessaire
	}
}

// ChannelInvitationEntityToDTO convertit une entité ChannelInvitation en DTO ChannelInvitation
func ChannelInvitationEntityToDTO(e entity.ChannelInvitation) dto.ChannelInvitation {
	return dto.ChannelInvitation{
		ChannelId: e.ChannelID,
		UserId:    e.UserID,
	}
}
