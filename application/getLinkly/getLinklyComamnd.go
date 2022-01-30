package application_getlinkly

import (
	"context"

	"github.com/golobby/container/v3"
	linkfly "linkfly.api/domain/linkly"
)

type GetLinklyCommand struct {
	IP        string
	Hash      string
	Useragent string
}

func (command *GetLinklyCommand) Execute(context context.Context) (GetLinklyCommandResult, error) {

	var linklyRepository linkfly.ILinklyRepository

	if err := container.Resolve(&linklyRepository); err != nil {
		panic(err)
	}
	var linkly *linkfly.Linkly
	var err error

	if linkly, err = linklyRepository.Get(context, command.Hash); err != nil {
		panic(err)
	}

	return GetLinklyCommandResult{
		Linkly: linkly,
	}, nil
}
