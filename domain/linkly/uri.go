package domain_linkly

type Uri struct {
	Host    string
	Path    string
	Query   string
	FullUrl string
}

func NewUri(url string) Uri {
	return Uri{
		FullUrl: url,
	}
}
