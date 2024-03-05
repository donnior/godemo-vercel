package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/donnior/godemo-vercel/pkg"
)

func ContentHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := pkg.GetContent("https://www.sumingxs.com/xiaoshuo/2/70671/", pkg.SiteConfig{
		Domain:          "www.sumingxs.com",
		ChapterSelector: "#list li",
		ContentSelector: "#c p",
	})
	if err != nil {
		// w.Write("some error")
		fmt.Print("error")
	}
	// w.Write([]byte(res))
	fmt.Print(res)
	json.NewEncoder(w).Encode(res)
}
