package main

/*
   "locationInfo": null,
   "authStatus": 0,
   "remarkName": null,
   "avatarUrl": "http://p1.music.126.net/FEckpBBlDjqic0TwsXdHvw==/109951163027286591.jpg",
   "experts": null,
   "expertTags": null,
   "vipType": 0,
   "userId": 512965084,
   "userType": 0,
   "nickname": "夏木秋影"
*/

//User 用户信息
type User struct {
	LocationInfo string `json:"locationInfo"`
	AuthStatus   int    `json:"authStatus"`
	RemarkName   string `json:"remarkName"`
	AvatarURL    string `json:"avatarUrl"`
	Experts      string `json:"experts"`
	ExpertTags   string `json:"expertTags"`
	VipType      string `json:"vipType"`
	UserID       int64  `json:"userId"`
	UserType     int    `json:"userType"`
	Nickname     string `json:"nickname"`
}
