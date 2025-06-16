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
func ChannelEntityToDTO(e entity.Channel, owner dto.User, tags []string, messages []dto.Message) dto.Channel {
	return dto.Channel{
		Id:          e.ChannelID,
		Name:        e.Name,
		Description: e.Description,
		IsPrivate:   e.Private,
		Owner:       owner,
		Tags:        tags,
		Messages:    messages,
	}
}

// MessageEntityToDTO convertit une entité Message en DTO Message
func MessageEntityToDTO(e entity.Message, creator dto.User, image dto.Image, reactions []string, upDownVotes []dto.UpDownVote, uservote bool, vote int) dto.Message {
	nbUpVote := 0
	nbDownVote := 0
	for _, vote := range upDownVotes {
		if vote.Vote == 1 {
			nbUpVote++
		} else if vote.Vote == 0 {
			nbDownVote++
		}
	}
	return dto.Message{
		Id:         e.MessageTextID,
		Text:       e.Text,
		Edited:     e.Edited,
		Creator:    creator,
		Image:      e.Image,
		ImageFile:  image,
		Reaction:   reactions,
		UpVotes:    upDownVotes,
		NbUpVote:   nbUpVote,
		NbDownVote: nbDownVote,
		UserVote:   uservote,
		Vote:       vote,
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

func ListOfMessagesEntityToDTO(messages []entity.Message, creators []dto.User, updownList [][]dto.UpDownVote) []dto.Message {
	var messagesDTO []dto.Message
	uservote := false
	vote := 0
	for index, message := range messages {
		for _, updown := range updownList[index] {
			if updown.UserId == creators[index].Id {
				uservote = true
				vote = updown.Vote
			}
		}
		messagesDTO = append(messagesDTO, MessageEntityToDTO(message, creators[index], dto.Image{}, []string{}, updownList[index], uservote, vote))
	}
	return messagesDTO
}

func ListUpDownVoteEntityToDTO(votes []entity.UpDown) []dto.UpDownVote {
	var votesDTO []dto.UpDownVote
	for _, vote := range votes {
		votesDTO = append(
			votesDTO,
			dto.UpDownVote{
				UserId: vote.UserID,
				Vote:   vote.UpDownVoteID,
			},
		)
	}
	return votesDTO
}
