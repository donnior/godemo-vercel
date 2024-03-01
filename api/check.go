package api

import (
	"fmt"
	"github.com/donnior/godemo-vercel/config"
	"net/http"
)

func Check(w http.ResponseWriter, req *http.Request) {
	err := config.CheckConfig()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "配置成功")
}
