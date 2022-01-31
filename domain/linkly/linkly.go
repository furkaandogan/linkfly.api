package domain_linkly

import (
	"strings"
	"time"

	"github.com/google/uuid"
	stringrandom "linkfly.api/pkg/stringRandom"
)

type Linkly struct {
	Id          string `bson:"_id"`
	CreatedDate time.Time
	UpdatedDate time.Time
	CreatedUser User
	Hash        string
	Uri         Uri
	Metrics     Metric
}

func NewLinkly(url string, createdUseragent string, createdUserIp string) Linkly {
	uuid, _ := uuid.NewUUID()

	uri := NewUri(url)

	hash := strings.ToLower(stringrandom.RandString(4))

	return Linkly{
		Id:          uuid.String(),
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
		Uri:         uri,
		Hash:        hash,
		Metrics:     Metric{},
	}
}