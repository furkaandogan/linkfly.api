package domain_linkly

type User struct {
	UserAgent UserAgent
}

func NewUser(ip string, useragent string) User {
	return User{
		UserAgent: UserAgent{
			IP:        ip,
			Useragent: useragent,
		},
	}
}
