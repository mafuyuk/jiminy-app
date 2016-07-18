package main

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/labstack/echo/engine/fasthttp"
)

type (
	history struct {
		Number int `json:"number"`
		Url        string `json:"url"`
		SendUser string `json:"send_user"`
		TotalLike int `json:"tatl_like"`
	}

	responseHistoryList struct {
		HistoryList []history `json:"history_list"`
	}

)

func main() {
	router := Init()
	router.Run(fasthttp.New(":8888"))
}

func Init() *echo.Echo {
	e := echo.New()

	// Routes
	v1 := e.Group("/jiminy/v1")
	{
		v1.GET("/history/list/:size", getHistoryList)
	}
	return e
}

// 投稿履歴取得
func getHistoryList(c echo.Context) error {
	var historyList []history

	historyList = append(historyList, history{
		Number: 1,
		Url: "3exsRhw3xt8",
		SendUser: "kamono",
		TotalLike: 5,
	})
	historyList = append(historyList, history{
		Number: 2,
		Url: "STg4Ya8bEFo",
		SendUser: "nakamura",
		TotalLike: 15,
	})

	response := new(responseHistoryList)
	response.HistoryList = historyList

	return c.JSON(http.StatusOK, response)
}
