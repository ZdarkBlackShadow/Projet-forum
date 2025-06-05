package services

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"projet-forum/models/dto"
	"projet-forum/models/entity"
	"projet-forum/repository"
	"projet-forum/utils"
	"strconv"
	"time"
)

type ChannelService struct {
	channelRepo *repository.ChannelRepository
	usersRepo   *repository.UsersRepository
	imageRepo   *repository.ImageRepository
}

func InitChannelServices(db *sql.DB) *ChannelService {
	return &ChannelService{
		channelRepo: repository.InitChannelRepository(db),
		usersRepo:   repository.InitUsersRepository(db),
	}
}

func (s *ChannelService) CreateChannel(channelInfo dto.ChannelCreation, channelImage entity.UserImage, token string) (int, error) {
	newChannel := entity.Channel {
		Name: channelInfo.Name,
		Description: channelInfo.Description,
		Private: channelInfo.Private,
		StateID: 1,
		CreatedAt: time.Now(),
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
			return -1, err
		}
		imageName := channelImage.Handler.Filename
		dstPath := filepath.Join("images/servers", imageName)

		dst, dstErr := os.Create(dstPath)
		if dstErr != nil {
			return -1, dstErr
		}
		defer dst.Close()

		_, copyErr := io.Copy(dst, channelImage.File)
		if copyErr != nil {
			return -1, copyErr
		}

		imageId, imageErr := s.imageRepo.Create(dstPath)
		if imageErr != nil {
			return -1, imageErr
		}
		newChannel.ImageID = imageId
	} else {
		newChannel.ImageID = 14
	}

	return s.channelRepo.CreateChannel(newChannel)
}

func (s *ChannelService) GetChannelById(channelId string, token string) (entity.Channel, error) {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return entity.Channel{}, jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return entity.Channel{}, convErr
	}

	intChannelId, convErr := strconv.Atoi(channelId)
	if convErr != nil {
		return entity.Channel{}, convErr
	}

	canAccess, accesErr := s.channelRepo.VerifyAccess(intChannelId, intUserId)
	if accesErr != nil {
		return entity.Channel{}, accesErr
	}

	if !canAccess {
		return entity.Channel{}, fmt.Errorf("Can't access to this channel")
	}

	return s.channelRepo.GetChannelById(intChannelId)
}