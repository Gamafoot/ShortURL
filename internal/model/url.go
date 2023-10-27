package model

import (
	"math/rand"
	"time"
)

type URL struct {
	ID        uint   `gorm:"primary key"`
	Origin    string `gorm:"not null"`
	Short     string `gorm:"unique;not null"`
	CreatedAt time.Time
}

func (URL) TableName() string {
	return "urls"
}

func (this *URL) GenerateShortURL() {
	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 6)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	this.Short = string(b)
}

func CreateURL(url *URL) error {
	return db.Create(url).Error
}

func GetURLByShortURL(shortURL string) (*URL, error) {
	url := &URL{}

	err := db.First(url, "short = ?", shortURL).Error
	if err != nil {
		return nil, err
	}

	return url, nil
}
