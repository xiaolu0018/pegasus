
#//r.GET("/api/activity/voters", ListVotersHandler)
#//r.POST("/api/activity/voter", RegisterVoterHandler)
#//r.POST("/api/activity/voter/:id/vote", VoteHandler)
#ID          string `json:"id"`
#	Name        string `json:"name"`
#	Image       string `json:"image"`
#	Company     string `json:"company"`
#	Mobile      string `json:"mobile"`
#	Declaration string `json:"declaration"`
#	VotedCount  int    `json:"voteCount"`
curl -XPOST 127.0.0.1:9000/api/activity/voter -d '{
    "openid": "1212312312312",
    "name": "潘新元",
    "image": "baidu.com/1.jpg",
    "company": "迪安",
    "mobile": "17744524308",
    "declaration": "123"
}'

curl -XPOST 127.0.0.1:9000/api/activity/voter/17/vote?openid=1212312312312

curl 127.0.0.1:9000/api/activity/voters
curl 192.168.199.198:9000/api/activity/voter?openid=oH4HtwGsY-0JSjhNhJLA7jYYOMsQ

curl 'https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=tAWw6VWFIMkxoCFtQ6RK9J3kpmXet8llCqSyi-AIMlQxfQoZQ-UdteM6zZSk2UVbAsnkWgUMetziOHUxf2kDaiWmzwZshzeKQTmx7pPltU0eJ7H0QgeCoSFoudP8nCGKUURfAIANNK&type=jsapi'
