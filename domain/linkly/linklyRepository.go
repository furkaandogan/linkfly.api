package domain_linkly

import "context"

type ILinklyRepository interface {
	Insert(context context.Context, linkly *Linkly) (bool, error)
	Update(context context.Context, linkly *Linkly) (bool, error)
	Get(context context.Context, hash string) (*Linkly, error)
	IsExists(context context.Context, hash string) (bool, error)
}
