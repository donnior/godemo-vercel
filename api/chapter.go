package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/donnior/godemo-vercel/pkg"
)

func ChapterHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := pkg.ListChapter("https://www.sumingxs.com/xiaoshuo/2/", pkg.SiteConfig{
		Domain:          "www.sumingxs.com",
		ChapterSelector: "#list li",
		ContentSelector: "some selector",
	})
	if err != nil {
		// w.Write("some error")
		fmt.Print("error")
	}
	// w.Write([]byte(res))
	fmt.Print(res)
	json.NewEncoder(w).Encode(res)
}
