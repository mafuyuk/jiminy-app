package repository

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	Video struct {
		Id   int	`gorm:"primary_key"`
		VideoId string
		Comment string
		PlayTime string
	}
)

type VideoRepository interface {
	SelectVideo() ([]Video, error)
	SelectOneVideo(id int) (*Video, error)
	InsertVideo(videoId, comment, playTime string) error
	UpdateVideo(id int, videoId, comment, playTime string) error
	DeleteVideo(id string) error
}

type mysqlVideo struct {
}

func NewMysqlVideoRepository() VideoRepository {
	return &mysqlVideo{}
}

func (v *mysqlVideo) SelectVideo () (video []Video, err error) {
	err = db.Table("videos").Select("id, video_id, comment, play_time").Scan(&video).Error
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (v *mysqlVideo) SelectOneVideo (id int) (*Video, error) {
	var video Video
	err := db.Table("videos").Select("id, video_id, comment, play_time").Where("id = ?", id).Scan(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func (v *mysqlVideo) InsertVideo(videoId, comment, playTime string) error {
	video := Video{
		VideoId:	videoId,
		Comment:	comment,
		PlayTime: playTime,
	}

	err := db.Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}

func(v *mysqlVideo) UpdateVideo(id int, videoId, comment, playTime string) error {
	video := Video{
		Id:	id,
		VideoId:	videoId,
		Comment:	comment,
		PlayTime: playTime,
	}

	err := db.Model(Video{}).Update(video).Error
	if err != nil {
		return err
	}
	return nil
}

func(v *mysqlVideo) DeleteVideo(id string) error {
	err := db.Where("id = ?", id).Delete(&Video{}).Error
	if err != nil {
		return err
	}
	return nil
}