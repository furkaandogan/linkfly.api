package linkly

import (
	"strings"
	"time"

	stringrandom "api.linkfly.com/pkg/stringRandom"
	"github.com/google/uuid"
)

type Linkly struct {
	Id           string `bson:"_id"`
	CreatedDate  time.Time
	UpdatedDate  time.Time
	CreatedUser  User
	Hash         string
	Uri          Uri
	Metrics      Metric
	FocusElement *FocusElement
}

func (self *Linkly) Visited() *Linkly {
	self.Metrics.Visited += 1
	return self
}

func NewLinkly(url string, userId string, createdUseragent string, createdUserIp string) Linkly {
	uuid, _ := uuid.NewUUID()

	uri := NewUri(url)

	hash := strings.ToLower(stringrandom.RandString(4))

	createdUser := NewUser(userId, createdUserIp, createdUseragent)

	return Linkly{
		Id:          uuid.String(),
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		CreatedUser: createdUser,
		Uri:         uri,
		Hash:        hash,
		Metrics:     Metric{},
	}
}

func NewLinklyWithFocusElement(url string, xpath string, userId string, createdUseragent string, createdUserIp string) Linkly {
	linkly := NewLinkly(url, userId, createdUseragent, createdUserIp)
	linkly.FocusElement = NewFocusElement(xpath)
	return linkly
}
