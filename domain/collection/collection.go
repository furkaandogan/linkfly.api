package collection

import (
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	Id          string `bson:"_id"`
	CreatedDate time.Time
	UpdatedDate time.Time
	CreatedUser User
	Title       string
}

func(self *Collection) AddLinkly()

func NewCollection(title string, userId string, createdUseragent string, createdUserIp string) Collection {
	uuid, _ := uuid.NewUUID()

	createdUser := NewUser(userId, createdUserIp, createdUseragent)

	return Collection{
		Id:          uuid.String(),
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		CreatedUser: createdUser,
		Title:       title,
	}
}
