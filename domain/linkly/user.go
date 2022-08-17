package linkly

type User struct {
	Id        string
	UserAgent UserAgent
}

func NewUser(id string, ip string, useragent string) User {
	return User{
		Id: id,
		UserAgent: UserAgent{
			IP:        ip,
			Useragent: useragent,
		},
	}
}
