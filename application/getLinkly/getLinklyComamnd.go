package application_getlinkly

import (
	"context"

	linkfly "api.linkfly.com/domain/linkly"
	"github.com/golobby/container/v3"
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

	linkly.Visited()

	if result, err := linklyRepository.Update(context, linkly); err != nil || !result {
		panic(err)
	}

	return GetLinklyCommandResult{
		Linkly: linkly,
	}, nil
}
