package main

/*
{
            "user": {
                "locationInfo": null,
                "expertTags": null,
                "remarkName": null,
                "avatarUrl": "http://p1.music.126.net/meuXLG6wmVI_JU_jnUN4Lg==/19162288649180456.jpg",
                "experts": null,
                "authStatus": 0,
                "nickname": "Amloy",
                "userType": 0,
                "vipType": 0,
                "userId": 118770203
            },
            "beReplied": [
                {
                    "user": {
                        "locationInfo": null,
                        "expertTags": null,
                        "remarkName": null,
                        "avatarUrl": "http://p1.music.126.net/0kZUZjoRMRzgWy8GnAUmyg==/18607035278926779.jpg",
                        "experts": null,
                        "authStatus": 0,
                        "nickname": "高高圩糖",
                        "userType": 0,
                        "vipType": 0,
                        "userId": 1295089081
                    },
                    "content": "你们都知道自己离家多少里吗  我是177里",
                    "status": 0
                }
            ],
            "commentId": 644162024,
            "liked": false,
            "likedCount": 0,
            "time": 1513499949880,
            "content": "800多里"
        }
*/

//Comment 评论
type Comment struct {
	User       User      `json:"user"`
	BeReplied  []Comment `json:"beReplied"`
	CommentID  int64     `json:"commentId"`
	Liked      bool      `json:"liked"`
	LikedCount int       `json:"likedCount"`
	Time       int64     `json:"time"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
}

type CommentResult struct {
	IsMusician  bool  `json:"isMusician"`
	UserID      int64 `json:"userId"`
	TopComments string
}
