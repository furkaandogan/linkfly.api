package application_createlinkly

import (
	"context"

	linkfly "api.linkfly.com/domain/linkly"
	"github.com/golobby/container/v3"
)

type CreateLinklyWithFocusElementCommand struct {
	IP        string
	Url       string
	XPath     string
	Useragent string
}

func (command *CreateLinklyWithFocusElementCommand) Execute(context context.Context) (CreateLinklyWithFocusElementCommandResult, error) {

	linkly := linkfly.NewLinklyWithFocusElement(command.Url, command.XPath, command.Useragent, command.IP)

	var linklyRepository linkfly.ILinklyRepository

	if err := container.Resolve(&linklyRepository); err != nil {
		panic(err)
	}

	if _, err := linklyRepository.Insert(context, &linkly); err != nil {
		panic(err)
	}

	return CreateLinklyWithFocusElementCommandResult{
		Linkly: linkly,
	}, nil
}
