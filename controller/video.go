package controller

import (
	"net/http"
	"strconv"

	"jiminy-app/repository"

	"github.com/labstack/echo"
)

type (
	Video struct {
		Id       int    `json:"id"`
		VideoId  string `json:"video_id"`
		Comment  string `json:"comment"`
		PlayTime string `json:"play_time"`
	}

	PostVideoBody struct {
		Result int `json:"result"`
	}
)

/**
 * Youtube動画全件取得
 */
func GetVideo(c echo.Context) error {
	// Youtube動画検索
	videoRepository := repository.NewMysqlVideoRepository()
	result, err := videoRepository.SelectVideo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, result)
}

/**
 * Youtube動画取得
 */
func GetOneVideo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// Youtube動画検索
	videoRepository := repository.NewMysqlVideoRepository()
	result, err := videoRepository.SelectOneVideo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, result)
}

/**
 * Youtube動画登録
 */
func PostVideo(c echo.Context) error {
	videoRequest := new(Video)
	c.Bind(&videoRequest)

	// Youtube動画登録
	videoRepository := repository.NewMysqlVideoRepository()
	err := videoRepository.InsertVideo(videoRequest.VideoId, videoRequest.Comment, videoRequest.PlayTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// 正常レスポンス返却
	return c.JSON(http.StatusOK, PostVideoBody{Result: 200})
}
