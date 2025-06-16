package services

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"projet-forum/models/dto"
	"projet-forum/models/entity"
	"projet-forum/models/mapper"
	"projet-forum/repository"
	"projet-forum/utils"
	"strconv"
	"time"
)

type ChannelService struct {
	channelRepo *repository.ChannelRepository
	usersRepo   *repository.UsersRepository
	imageRepo   *repository.ImageRepository
	tagRepo     *repository.TagRepository
	updownRepo  *repository.UpDownVoteRepository
	messageRepo *repository.MessageRepository
}

func InitChannelServices(db *sql.DB) *ChannelService {
	return &ChannelService{
		channelRepo: repository.InitChannelRepository(db),
		usersRepo:   repository.InitUsersRepository(db),
		imageRepo:   repository.InitImageRepository(db),
		tagRepo:     repository.InitTagRepository(db),
		updownRepo:  repository.InitUpDownVoteRepository(db),
		messageRepo: repository.InitMessageRepository(db),
	}
}

func (s *ChannelService) CreateChannel(channelInfo dto.ChannelCreation, channelImage entity.UserImage, token string) (int, error) {
	newChannel := entity.Channel{
		Name:        channelInfo.Name,
		Description: channelInfo.Description,
		Private:     channelInfo.Private,
		StateID:     1,
		CreatedAt:   time.Now(),
	}
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return -1, jwtErr
	}

	_, err := s.usersRepo.GetById(userId)
	if err != nil {
		return -1, err
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return -1, convErr
	}

	newChannel.UserID = intUserId

	if channelImage.File != nil && channelImage.Handler != nil {
		err := os.MkdirAll("images/servers", os.ModePerm)
		if err != nil {
			fmt.Println(0)
			return -1, err
		}
		//imageName := channelImage.Handler.Filename
		dstPath := filepath.Join("images/servers", "test.txt")

		dst, dstErr := os.Create(dstPath)
		if dstErr != nil {
			fmt.Println(1)
			stat, _ := dst.Stat()
			fmt.Println(stat.Mode().Perm())
			return -1, dstErr
		}

		defer dst.Close()

		_, copyErr := io.Copy(dst, channelImage.File)
		if copyErr != nil {
			fmt.Println(2)
			return -1, copyErr
		}

		imageId, imageErr := s.imageRepo.Create(dstPath)
		if imageErr != nil {
			fmt.Println(3)
			return -1, imageErr
		}
		newChannel.ImageID = imageId
	} else {
		newChannel.ImageID = 14
	}

	return s.channelRepo.CreateChannel(newChannel)
}

func (s *ChannelService) GetChannelById(channelId string, token string) (dto.Channel, error) {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return dto.Channel{}, jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return dto.Channel{}, convErr
	}

	intChannelId, convErr := strconv.Atoi(channelId)
	if convErr != nil {
		return dto.Channel{}, convErr
	}

	canAccess, accesErr := s.channelRepo.VerifyAccess(intChannelId, intUserId)
	if accesErr != nil {
		return dto.Channel{}, accesErr
	}

	if !canAccess {
		return dto.Channel{}, fmt.Errorf("Can't access to this channel")
	}

	channel, err := s.channelRepo.GetChannelById(intChannelId)
	owner, err := s.usersRepo.GetById(strconv.Itoa(channel.UserID))
	tags, err := s.tagRepo.GetTagsByChannelId(intChannelId)
	messages, err := s.messageRepo.GetMessagesByChannelID(intChannelId)
	var creatorList []dto.User
	var upDownVoteList [][]dto.UpDownVote
	for _, message := range messages {
		creator, err := s.usersRepo.GetById(strconv.Itoa(message.UserID))
		if err != nil {
			return dto.Channel{}, err
		}
		upDownVoteMessage, err := s.updownRepo.GetAllUpDownVoteFromMessageId(message.MessageTextID)
		if err != nil {
			return dto.Channel{}, err
		}
		upDownVoteList = append(upDownVoteList, mapper.ListUpDownVoteEntityToDTO(upDownVoteMessage))
		creatorList = append(creatorList, mapper.UserEntityToDTO(creator))
	}
	dtoMessages := mapper.ListOfMessagesEntityToDTO(messages, creatorList, upDownVoteList)
	channelDto := mapper.ChannelEntityToDTO(channel, mapper.UserEntityToDTO(owner), tags, dtoMessages)
	return channelDto, err
}

func (s *ChannelService) AddTagToChannel(channelId string, tag string, token string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	intChannelId, convErr := strconv.Atoi(channelId)
	if convErr != nil {
		return convErr
	}

	canAccess, accesErr := s.channelRepo.VerifyAccess(intChannelId, intUserId)
	if accesErr != nil {
		return accesErr
	}

	if !canAccess {
		return fmt.Errorf("Can't access to this channel")
	}
	intTagId, reqErr := s.tagRepo.GetTagIdByName(tag)
	if reqErr != nil {
		return reqErr
	}

	err := s.tagRepo.AddTagToChannel(intChannelId, intTagId)
	if err != nil {
		return err
	}
	return nil
}

func (s *ChannelService) RemoveTagFromChannel(channelId string, tags []string, token string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	intChannelId, convErr := strconv.Atoi(channelId)
	if convErr != nil {
		return convErr
	}

	canAccess, accesErr := s.channelRepo.VerifyAccess(intChannelId, intUserId)
	if accesErr != nil {
		return accesErr
	}

	if !canAccess {
		return fmt.Errorf("Can't access to this channel")
	}

	for _, tag := range tags {
		intTagId, reqErr := s.tagRepo.GetTagIdByName(tag)
		if reqErr != nil {
			return reqErr
		}

		err := s.tagRepo.RemoveTagFromChannel(intChannelId, intTagId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ChannelService) DeleteChannel(channelId string, token string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	intChannelId, convErr := strconv.Atoi(channelId)
	if convErr != nil {
		return convErr
	}

	channel, err := s.channelRepo.GetChannelById(intChannelId)
	if err != nil {
		return err
	}

	if channel.UserID != intUserId {
		return fmt.Errorf("Can't delete this channel")
	}

	return s.channelRepo.DeleteChannel(intChannelId)
}

func (s *ChannelService) CreateTag(channelId string, tagName string, token string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intChannleId, convErr := strconv.Atoi(channelId)
	if convErr != nil {
		return convErr
	}

	_, err := s.usersRepo.GetById(userId)
	if err != nil {
		return err
	}

	newTagId, creationErr := s.tagRepo.CreateTag(tagName)
	if creationErr != nil {
		return creationErr
	}

	return s.tagRepo.AddTagToChannel(intChannleId, newTagId)
}

func (s *ChannelService) CreateChannelIvitation(token string, usernameWhoAreInvited string, channelId string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	intChannelId, convErr := strconv.Atoi(channelId)
	if convErr != nil {
		return convErr
	}

	userIdWhoAreInvited, err := s.usersRepo.GetIdByUsername(usernameWhoAreInvited)
	if err != nil {
		return err
	}

	return s.channelRepo.CreateChannelInvitation(intUserId, userIdWhoAreInvited, intChannelId)
}

func (s *ChannelService) GetChannelInvitations(token string) ([]entity.ChannelInvitation, error) {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return nil, jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return nil, convErr
	}

	return s.channelRepo.GetAllChannelInvitations(intUserId)
}

func (s *ChannelService) AcceptChannelInvitation(token string, usernameWhoCreateInvitation string, channelName string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	channel, err := s.channelRepo.GetChannelByName(channelName)

	intInvitationId, reqErr := s.usersRepo.GetIdByUsername(usernameWhoCreateInvitation)
	if reqErr != nil {
		return reqErr
	}

	invitation, err := s.channelRepo.GetChannelInvitation(intInvitationId, intUserId, channel.ChannelID)
	if err != nil {
		return err
	}

	err = s.channelRepo.AddUserToChannel(invitation.UserID1, channel.ChannelID)
	if err != nil {
		return err
	}

	return s.channelRepo.DeleteChannelInvitation(intInvitationId, intUserId, channel.ChannelID)
}

func (s *ChannelService) DeclineChannelInvitation(token string, usernameWhoCreateInvitation string, channelName string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	channel, err := s.channelRepo.GetChannelByName(channelName)

	intInvitationId, reqErr := s.usersRepo.GetIdByUsername(usernameWhoCreateInvitation)
	if reqErr != nil {
		return reqErr
	}

	invitation, err := s.channelRepo.GetChannelInvitation(intInvitationId, intUserId, channel.ChannelID)
	if err != nil {
		return err
	}

	return s.channelRepo.DeleteChannelInvitation(invitation.UserID, intUserId, channel.ChannelID)
}
