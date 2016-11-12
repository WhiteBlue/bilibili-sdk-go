package client

type BCli struct {
	Rank    RankService
	Bangumi BangumiService
	Video   VideoService
	Special SpecialService
	User    UserService
	Others  OthersService
}

func NewClient(appkey, secret string) *BCli {
	params := BaseParam{
		Appkey: appkey,
		Secret: secret,
	}

	client := NewHttpClient()

	base := BaseService{params, client}
	return &BCli{
		Rank:    RankService{base},
		Bangumi: BangumiService{base},
		Video:   VideoService{base},
		Special: SpecialService{base},
		User:    UserService{base},
		Others:  OthersService{base},
	}
}
