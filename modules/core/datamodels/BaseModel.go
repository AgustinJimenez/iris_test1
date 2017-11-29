package datamodels

import (
	"time"
)

type ModelInterface interface {
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

type Model struct {
	CreatedAt time.Time `xorm:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at" json:"updated_at" form:"updated_at"`
}

func (this *Model) GetCreatedAt() time.Time {
	return this.CreatedAt
}
func (this *Model) GetUpdatedAt() time.Time {
	return this.UpdatedAt
}

func (this *Model) CreatedAtFormated(self ModelInterface) string {

	return self.GetCreatedAt().Format("2006-01-02 15:04:05")
}
func (this *Model) UpdatedAtFormated(self ModelInterface) string {

	return self.GetUpdatedAt().Format("2006-01-02 15:04:05")
}
func (this *Model) TimeToTimestamp(unformated_time time.Time) string {
	return unformated_time.Format("2006-01-02 15:04:05")
}
