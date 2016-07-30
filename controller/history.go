package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	history struct {
		Number int `json:"number"`
		Url        string `json:"url"`
		SendUser string `json:"send_user"`
		TotalLike int `json:"total_like"`
	}

	response struct {
		Header *header `json:"header"`
		Body *body `json:"body"`
	}

	header struct {
		ResultCode string `json:"result_code"`
		ResponseId string `json:"response_id"`
	}

	body struct {
		HistoryList []history `json:"history_list"`
	}


)

// 投稿履歴取得
func GetHistory(c echo.Context) error {
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

	body := new(body)
	body.HistoryList = historyList

	header := new(header)
	header.ResultCode = "200"
	header.ResponseId = "responseId00001"

	response := response{
		Header: header,
		Body: body,
	}

	return c.JSON(http.StatusOK, response)
}

// ダミーPOST
func PostDummy(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello post")
}