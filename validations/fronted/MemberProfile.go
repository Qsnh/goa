package fronted

type MemberProfileValidation struct {
	Company    string `form:"company"`
	Age        int64  `form:"age"`
	Profession string `form:"profession"`
	Website    string `form:"website"`
	Weibo      string `form:"weibo"`
	Wechat     string `form:"wechat"`
}
