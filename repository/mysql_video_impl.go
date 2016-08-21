package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type (
	Video struct {
		Id   int	`gorm:"primary_key"`
		VideoId string
		Comment string
		PlayTime string
	}
	Model struct {
		gorm.Model
		Videos   []Video
	}
)

type VideoRepository interface {
	SelectVideo(id int) error
	InsertVideo(videoId, comment, playTime string) error
	UpdateVideo(id int, videoId, comment, playTime string) error
	DeleteVideo(id string) error
}

type mysqlVideo struct {
}

func NewMysqlVideoRepository() VideoRepository {
	return &mysqlVideo{}
}

func (v *mysqlVideo) SelectVideo (id int) error {
	var m Model

	err := db.Model(&m).Related(&Video{Id:	id}).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("zoo :%v", m)
	return nil
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