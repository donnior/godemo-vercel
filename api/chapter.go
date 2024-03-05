package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/donnior/godemo-vercel/internal"
)

func Chapter(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := internal.ListChapter("some book link", internal.SiteConfig{
		Domain:          "some domain",
		ChapterSelector: "some selector",
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
