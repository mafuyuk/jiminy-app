package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	historyBody struct {
		Number int `json:"number"`
		Url        string `json:"url"`
		SendUser string `json:"send_user"`
		TotalLike int `json:"total_like"`
	}

	dummyBody struct {
		Result        int `json:"result"`
	}


)

// 投稿履歴取得
func GetHistory(c echo.Context) error {
	var list []historyBody

	list = append(list, historyBody{
		Number: 1,
		Url: "3exsRhw3xt8",
		SendUser: "kamono",
		TotalLike: 5,
	})
	list = append(list, historyBody{
		Number: 2,
		Url: "STg4Ya8bEFo",
		SendUser: "nakamura",
		TotalLike: 15,
	})

	return c.JSON(http.StatusOK, list)
}

// ダミーPOST
func PostDummy(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		dummyBody{
			Result: 201,
		})
}