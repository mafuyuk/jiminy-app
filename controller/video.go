package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"jiminy_app/repository"
)

type (
	VideoBody struct {
		Id int `json:"id"`
		VideoId        string `json:"video_id"`
		Comment string `json:"comment"`
		PlayTime string `json:"play_time"`
	}

	PostVideoRequest struct {
		VideoId	string	`json:"video_id"`
		Comment string	`json:"comment"`
		PlayTime string	`json:"play_time"`
	}
)

/**
 * Youtube動画取得
 */
func GetVideo(c echo.Context) error {
	repository.Initialize()
	defer repository.Close()

	videoRepository := repository.NewMysqlVideoRepository()
	videoRepository.SelectVideo(1)

	videoBody := VideoBody {
		Id: 1,
		VideoId: "STg4Ya8bEFo",
		Comment: "hello",
		PlayTime: "03:00",
	}

	return c.JSON(http.StatusOK, videoBody)
}

/**
 * Youtube動画登録
 */
func PostVideo(c echo.Context) error {
	postVideo := new(PostVideoRequest)
	c.Bind(&postVideo)

	repository.Initialize()
	defer repository.Close()

	videoRepository := repository.NewMysqlVideoRepository()
	videoRepository.InsertVideo(postVideo.VideoId,postVideo.Comment, postVideo.PlayTime)


	return c.JSON(http.StatusOK, "ok")
}