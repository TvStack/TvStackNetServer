package api

import "github.com/gin-gonic/gin"

func InitStreamApi() {

	router := gin.Default()
	router.GET("/api/streamget", getStream)
	router.Run(":8080")

}

func getStream(c *gin.Context) {

	allStreams := []Stream{{ID: 0, Title: "SpaceX", Desc: "", CoverImg: "", VideoURL: "", StreamURL: "", VideoType: ""},
		{ID: 0, Title: "SpaceX", Desc: "", CoverImg: "", VideoURL: "", StreamURL: "", VideoType: ""}}
	c.IndentedJSON(200, allStreams)
}
c
type Stream struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	CoverImg  string `json:"coverimg"`
	VideoURL  string `json:"videourl"`
	StreamURL string `json:"streamurl"`
	VideoType string `json:"videotype"`
}
