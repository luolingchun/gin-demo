/*
@Author  : llc
@Time    : 2020/3/1 13:16
*/

package models

import (
	orm "gin-demo/src/db"
	"github.com/jinzhu/gorm"
	"time"
)

//任务模型
type Job struct {
	ID        uint64    `json:"id" gorm:"primary_key"` //主键
	CreatedAt time.Time `json:"created_at"`            //创建时间
	UpdatedAt time.Time `json:"updated_at"`            //更新时间
	Name      string    `json:"name" form:"name"`      //任务名称
	Result    string    `json:"result" form:"result"`  //任务结果
}

func (j *Job) GetJobs() (jobs []Job, err error) {
	result := orm.DB.Find(&jobs)
	err = result.Error
	return
}

func (j *Job) GetJob() error {
	result := orm.DB.Where("id = ?", j.ID).First(&j)
	return result.Error
}

func (j *Job) CreateJob() error {
	result := orm.DB.Create(&j)
	return result.Error
}

func (j *Job) UpdateJob() error {
	if !ExistRecordById(j.ID) {
		return gorm.ErrRecordNotFound
	}

	result := orm.DB.Model(&j).Updates(&j)
	return result.Error
}

func (j *Job) DeleteJob() error {
	if !ExistRecordById(j.ID) {
		return gorm.ErrRecordNotFound
	}

	result := orm.DB.Delete(&j)
	return result.Error
}

func ExistRecordById(id uint64) bool {
	var job Job
	if result := orm.DB.Where("id = ?", id).First(&job); result.RecordNotFound() {
		return false
	}
	return true
}
