package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"jiminy_app/repository"
)

type (
	Video struct {
		Id int `json:"id"`
		VideoId        string `json:"video_id"`
		Comment string `json:"comment"`
		PlayTime string `json:"play_time"`
	}

	PostVideoBody struct {
		Result        int `json:"result"`
	}
)

/**
 * Youtube動画取得
 */
func GetVideo(c echo.Context) error {
	repository.InitializeMysql()
	defer repository.CloseMysql()

	videoRepository := repository.NewMysqlVideoRepository()
	videoRepository.SelectVideo(1)

	videoBody := Video {
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
	videoRequest := new(Video)
	c.Bind(&videoRequest)

	// MySQLコネクション生成
	repository.InitializeMysql()
	defer repository.CloseMysql()

	// Youtube動画登録
	videoRepository := repository.NewMysqlVideoRepository()
	err := videoRepository.InsertVideo(videoRequest.VideoId,videoRequest.Comment, videoRequest.PlayTime)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// 正常レスポンス返却
	return c.JSON(http.StatusOK, PostVideoBody{ Result:200 })
}