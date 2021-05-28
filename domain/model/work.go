package model
type Workers struct {
	ID int64 `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name string `json:"name"`
	Nums string `json:"nums" gorm:"unique_index"`//工号
	Sex string `json:"sex"` //性别
	Level int64 `json:"level"`//等级
	Score int64 `json:"score"`//信誉分
	Place string `json:"place"`//住址
	Telephone string `json:"telephone"`//电话
	Mail string `json:"mail"`
	Description string `json:"description"`//补充描述
	ISWork bool `json:"is_work"`//是否在职
	Username string `json:"user_name" gorm:"not_null;unique_index"`
	Password string `json:"password"`
}