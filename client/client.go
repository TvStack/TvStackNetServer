package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Stream bean
type Channel struct {
	NAME     string     `json:"name"`
	LOGO     string     `json:"logo"`
	CATEGORY string     `json:"category"`
	CoverImg string     `json:"coverimg"`
	LANGUAGE []Language `json:"language"`
}

type Language struct {
	CODE string `json:"code"`
	NAME string `json:"name"`
}

func Get() {

	r, err := http.Get("https://api.github.com/events")

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = r.Body.Close()
	}()
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("%s", body)

}
