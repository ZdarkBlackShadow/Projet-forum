package services

import (
	"database/sql"
	"projet-forum/models/dto"
	"projet-forum/repository"
	"projet-forum/utils"
	"strconv"
)

type HomeServices struct {
	usersRepo   *repository.UsersRepository
	channelRepo *repository.ChannelRepository
	tagRepo     *repository.TagRepository
	messageRepo *repository.MessageRepository
}

func InitHomeServices(db *sql.DB) *HomeServices {
	return &HomeServices{
		usersRepo:   repository.InitUsersRepository(db),
		channelRepo: repository.InitChannelRepository(db),
		tagRepo:     repository.InitTagRepository(db),
		messageRepo: repository.InitMessageRepository(db),
	}
}

func (s *HomeServices) GetUser(token string) (dto.HomeModel, error) {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return dto.HomeModel{}, jwtErr
	}

	user, userErr := s.usersRepo.GetById(userId)
	if userErr != nil {
		return dto.HomeModel{}, userErr
	}

	return dto.HomeModel{
		Userconnected: true,
		ConnectingUser: dto.User{
			Id:      user.UserID,
			Name:    user.Username,
			Bio:     user.Bio,
			ImageID: user.ImageID,
		},
	}, nil
}

func (s *HomeServices) Home(token string) (dto.HomeModel, error) {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return dto.HomeModel{}, jwtErr
	}

	user, userErr := s.usersRepo.GetById(userId)
	if userErr != nil {
		return dto.HomeModel{}, userErr
	}

	publicsChannels, publicsChannelsErr := s.channelRepo.GetAllPublicChannels()
	if publicsChannelsErr != nil {
		return dto.HomeModel{}, publicsChannelsErr
	}

	var publicsChannelsDto []dto.Channel
	for _, channel := range publicsChannels {
		owner, ownerErr := s.usersRepo.GetById(strconv.Itoa(channel.UserID))
		if ownerErr != nil {
			return dto.HomeModel{}, ownerErr
		}

		tags, tagErr := s.tagRepo.GetTagsByChannelId(channel.ChannelID)
		if tagErr != nil {
			return dto.HomeModel{}, tagErr
		}

		messages, messagesErr := s.messageRepo.GetMessagesByChannelID(channel.ChannelID)
		if messagesErr != nil {
			return dto.HomeModel{}, messagesErr
		}

		var messagesDto []dto.Message
		for _, message := range messages {
			creator, creatorErr := s.usersRepo.GetById(strconv.Itoa(message.UserID))
			if creatorErr != nil {
				return dto.HomeModel{}, creatorErr
			}

			messagesDto = append(messagesDto, dto.Message{
				Id:     message.MessageTextID,
				Text:   message.Text,
				Edited: message.Edited,
				Creator: dto.User{
					Id:      creator.UserID,
					Name:    creator.Username,
					Bio:     creator.Bio,
					ImageID: creator.ImageID,
				},
				Image: message.Image,
			})
		}

		publicsChannelsDto = append(publicsChannelsDto, dto.Channel{
			Id:          channel.ChannelID,
			Name:        channel.Name,
			Description: channel.Description,
			IsPrivate:   channel.Private,
			Owner: dto.User{
				Id:      owner.UserID,
				Name:    owner.Username,
				Bio:     owner.Bio,
				ImageID: owner.ImageID,
			},
			Tags: tags,
		})
	}

	return dto.HomeModel{
		Userconnected: true,
		ConnectingUser: dto.User{
			Id:      user.UserID,
			Name:    user.Username,
			Bio:     user.Bio,
			ImageID: user.ImageID,
		},
		PublicCHannel: publicsChannelsDto,
	}, nil
}
