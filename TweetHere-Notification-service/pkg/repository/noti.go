package repository

import (
	"fmt"
	interfaces "tweethere-Notification/pkg/repository/interface"
	"tweethere-Notification/pkg/utils/models"

	"gorm.io/gorm"
)

type notiRepository struct {
	DB *gorm.DB
}

func NewnotiRepository(DB *gorm.DB) interfaces.NotiRepository {
	return &notiRepository{
		DB: DB,
	}
}

func (c *notiRepository) StorenotificationReq(noti models.NotificationReq) error {
	err := c.DB.Exec("INSERT INTO notifications(user_id,sender_id,post_id,message,created_at) VALUES(?,?,?,?,?)", noti.UserID, noti.SenderID, noti.PostID, noti.Message, noti.CreatedAt).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *notiRepository) GetNotification(id int, pag models.Pagination) ([]models.Notification, error) {
	fmt.Println("..................fffffffffffff", id, pag)
	var data []models.Notification
	if pag.Offset <= 0 {
		pag.Offset = 1
	}
	offset := (pag.Offset - 1) * pag.Limit
	err := c.DB.Raw("SELECT sender_id,message, created_at FROM notifications WHERE user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", id, pag.Limit, offset).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
