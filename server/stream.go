package server

import (
	"github.com/gin-gonic/gin"
)

// InitStream init stream
func InitStream() {

	router := gin.Default()
	var stream Stream
	streams, _ := stream.Streams()
	router.GET("/getstream", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"code":    1,
			"message": "",
			"data":    streams,
		})
	})
	router.Run(":8080")
}

//Stream bean
type Stream struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	CoverImg  string `json:"coverimg"`
	VideoURL  string `json:"videourl"`
	StreamURL string `json:"streamurl"`
	VideoType string `json:"videotype"`
}

//Streams return stream list
func (stream *Stream) Streams() (streams []Stream, err error) {
	currentStreams := make([]Stream, 10)
	for i := 0; i < 10; i++ {
		currentStreams[i] = Stream{ID: string(i),
			Title:     "CCTV1",
			Desc:      "test desc",
			CoverImg:  "https://upload.wikimedia.org/wikipedia/zh/6/65/CCTV-1_Logo.png",
			VideoURL:  "",
			StreamURL: "http://117.148.187.37/PLTV/88888888/224/3221226400/index.m3u8", VideoType: "mp4"}
	}
	return currentStreams, nil
}
