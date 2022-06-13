package main

import (
	"easy-apis/utils"
	"fmt"
	"net/http"
)

func main() {
	//routers.InitAdminRouters()
	ret, _ := utils.GetIP(&http.Request{})
	fmt.Println(ret)
}
