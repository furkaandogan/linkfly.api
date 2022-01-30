package application_createlinkly

import (
	"context"

	"github.com/golobby/container/v3"
	linkfly "linkfly.api/domain/linkly"
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
