package linkly

import urlHelpers "net/url"

const (
	hostRegex = "(^(.*:)//)|([A-Za-z0-9\\-\\.]+)(:[0-9]+)?(.*)$"
)

type Uri struct {
	Host    string
	Path    string
	Query   string
	FullUrl string
}

func NewUri(url string) Uri {
	uri, err := urlHelpers.Parse(url)
	if err != nil {
		panic(err)
	}

	return Uri{
		Host:    uri.Host,
		Path:    uri.Path,
		Query:   uri.RawQuery,
		FullUrl: url,
	}
}
