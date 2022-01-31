package application_createlinkly

import (
	"context"

	linkfly "api.linkfly.com/domain/linkly"
	"github.com/golobby/container/v3"
)

type CreateLinklyCommand struct {
	IP        string
	Url       string
	Useragent string
}

func (command *CreateLinklyCommand) Execute(context context.Context) (CreateLinklyCommandResult, error) {

	linkly := linkfly.NewLinkly(command.Url, command.Useragent, command.IP)

	var linklyRepository linkfly.ILinklyRepository

	if err := container.Resolve(&linklyRepository); err != nil {
		panic(err)
	}

	if _, err := linklyRepository.Insert(context, &linkly); err != nil {
		panic(err)
	}

	return CreateLinklyCommandResult{
		Linkly: linkly,
	}, nil
}
