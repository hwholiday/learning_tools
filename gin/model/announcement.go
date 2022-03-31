package model

import (
	"errors"
	"fmt"
	"time"
)

type Announcement struct {
	Id          int    `xorm:"not null pk autoincr INT(11)" json:"id,omitempty"`
	Title       string `xorm:"not null default '' comment('标题') VARCHAR(50)" json:"title,omitempty"`
	Content     string `xorm:"not null comment('内容') TEXT" json:"content,omitempty"`
	Url         string `xorm:"default '' comment('超链接') VARCHAR(150)" json:"url,omitempty"`
	CreateTime  int64  `xorm:"not null default 0 comment('创建时间') BIGINT(20)" json:"create_time,omitempty"`
	UpdateTime  int64  `xorm:"default 0 comment('更新时间') BIGINT(20)" json:"update_time,omitempty"`
	ExpiredTime int64  `xorm:"not null default 0 comment('过期时间(1为永远不过期)') BIGINT(20)" json:"expired_time,omitempty"`
}

func AddAnnouncement(data *Announcement) error {
	ok, err := db.Exist(&Announcement{Title: data.Title})
	if err != nil {
		return err
	}
	if ok {
		return errors.New(MErrExisted)
	}
	data.CreateTime = time.Now().Unix()
	data.UpdateTime = time.Now().Unix()
	_, err = db.Insert(data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAnnouncement(data *Announcement) error {
	id, err := db.Id(data.Id).Delete(&Announcement{})
	if err != nil || id <= 0 {
		return errors.New(MErrDelete)
	}
	return nil
}

func UpdateAnnouncement(data *Announcement) error {
	id, err := db.Id(data.Id).Update(data)
	if err != nil || id <= 0 {
		fmt.Println(err.Error())
		return errors.New(MErrUpdate)
	}
	return nil
}

func GetAnnouncementById(id int) (info Announcement, errs error) {
	ok, err := db.Id(id).Get(&info)
	if !ok || err != nil {
		errs = errors.New(MErrNotFind)
		return
	}
	return
}

func GetAnnouncementAll() (info []Announcement, err error) {
	err = db.Where("expired_time >= ?", time.Now().Unix()).Find(&info)
	return
}
