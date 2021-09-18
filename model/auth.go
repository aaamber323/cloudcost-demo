package model

type Auth struct {
	Model
	Role string `json:"auth"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Enabled bool `json:"enabled" gorm:"default:true"`
}