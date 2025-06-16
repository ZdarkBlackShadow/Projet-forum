package services

import (
	"database/sql"
	"fmt"
	"projet-forum/models/entity"
	"projet-forum/repository"
	"projet-forum/utils"
	"strconv"
	"time"
)

type MessageServices struct {
	messageRepo *repository.MessageRepository
}

func InitMessageServices(db *sql.DB) *MessageServices {
	return &MessageServices{
		messageRepo: repository.InitMessageRepository(db),
	}
}

func (s *MessageServices) CreateMessage(text string, channelId string, token string) (int, error) {
	if len(text) > 400 {
		return -1, fmt.Errorf("Too long message")
	}

	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return -1, jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return -1, convErr
	}

	intChannelId, convErr := strconv.Atoi(channelId)
	if convErr != nil {
		return -1, convErr
	}

	message := entity.Message {
		Text: text,
		CreatedAt: time.Now(),
		UserID: intUserId,
		Edited: false,
		Channel_id: intChannelId,
		Image: false,
	}

	return s.messageRepo.CreateMessage(message)
}

func (s *MessageServices) GetMessageById(id int) (entity.Message, error) {
	return s.messageRepo.GetMessageByID(id)
}

func (s *MessageServices) GetAllMessagesFromAChannel(id int) ([]entity.Message, error) {
	return s.messageRepo.GetMessagesByChannelID(id)
}

func (s *MessageServices) UpdateMessage(newText string, messageId string, token string) error {
	if len(newText) > 400 {
		return fmt.Errorf("Too long message")
	}

	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	intMessageId, convErr := strconv.Atoi(messageId)
	if convErr != nil {
		return convErr
	}

	message, err := s.messageRepo.GetMessageByID(intMessageId)
	if err != nil {
		return err
	}

	if message.UserID != intUserId {
		return fmt.Errorf("This user are not the owner of this message")
	}

	updatedMessage := entity.Message {
		Text: newText,
		Edited: true,
		Image: message.Image,
	}

	return s.messageRepo.UpdateMessage(updatedMessage)
}

func (s *MessageServices) DeleteMessage(messageId string, token string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	intMessageId, convErr := strconv.Atoi(messageId)
	if convErr != nil {
		return convErr
	}

	message, err := s.messageRepo.GetMessageByID(intMessageId)
	if err != nil {
		return err
	}

	if message.UserID != intUserId {
		return fmt.Errorf("This user are not the owner of this message")
	}

	return s.messageRepo.DeleteMessage(intMessageId)
}

func (s *MessageServices) AddUpDownVote(messageId string, token string, voteType string) (entity.Message, error) {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		fmt.Println(jwtErr)
		fmt.Println(token)
		return entity.Message{}, jwtErr
	}

	intVote, convErr := strconv.Atoi(voteType)
	if convErr != nil {
		return entity.Message{}, convErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return entity.Message{}, convErr
	}

	intMessageId, convErr := strconv.Atoi(messageId)
	if convErr != nil {
		return entity.Message{}, convErr
	}

	message, reqErr := s.GetMessageById(intMessageId)
	if reqErr != nil {
		fmt.Println(reqErr)
		return entity.Message{}, reqErr
	}

	err := s.messageRepo.AddUpDownVote(message.MessageTextID, intUserId, intVote)
	if err != nil {
		return entity.Message{}, err
	}
	return message, nil
}

func (s *MessageServices) UpdateUpDownVote(messageId string, token string, newVote string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intVote, convErr := strconv.Atoi(newVote)
	if convErr != nil {
		return convErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	intMessageId, convErr := strconv.Atoi(messageId)
	if convErr != nil {
		return convErr
	}

	message, reqErr := s.messageRepo.GetMessageByID(intMessageId)
	if reqErr != nil {
		return reqErr
	}

	return s.messageRepo.UpdateUpDownVote(message.MessageTextID, intUserId, intVote)
}

