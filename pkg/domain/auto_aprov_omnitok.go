package domain

import (
	"context"
	"time"
)

type Auto_Aprov_OmnitokReq struct {
	ID                      int32       `json:"id" copier:"must"`
	Branch                  string      `gorm:"column:branch;type:VARCHAR;size:10;" json:"branch" copier:"must"`
	IngressIp               string      `json:"ingress_ip" copier:"must"`
	Hostname                string      `json:"hostname" copier:"must"`
	InitialAdminEmail       string      `json:"initial_admin_email" copier:"must"`
	InitialAdminpassword    string      `json:"initial_admin_password" copier:"must"`
	ExternalMongodbUrl      string      `json:"external_mongodb_url" copier:"must"`
	ExternalMongodbOplogUrl string      `json:"external_mongodb_oplog_url" copier:"must"`
	Apply                   bool        `json:"apply" copier:"must"`
	Destroy                 bool        `json:"destroy" copier:"must"`
	Deleted                 bool        `json:"deleted" copier:"must"`
	DateCreate              time.Time   `json:"date_create" copier:"must"`
	DateDeleted             time.Time   `json:"date_deleted" copier:"must"`
	Vars                    interface{} `json:"vars"`
}

type Auto_Aprov_Omnitok struct {
	ID                      int32     `gorm:"primary_key;column:id;autoIncrement;" json:"id"`
	Branch                  string    `gorm:"column:branch;type:VARCHAR;size:10;" json:"branch"`
	IngressIp               string    `gorm:"column:ingress_ip;type:VARCHAR;size:20;" json:"ingress_ip"`
	Hostname                string    `gorm:"column:hostname;type:VARCHAR;size:200;" json:"hostname"`
	InitialAdminEmail       string    `gorm:"column:initial_admin_email;type:VARCHAR;size:250;" json:"initial_admin_email"`
	InitialAdminpassword    string    `gorm:"column:initial_admin_password;type:VARCHAR;size:250;" json:"initial_admin_password"`
	ExternalMongodbUrl      string    `gorm:"column:external_mongodb_url;type:VARCHAR;size:500;" json:"external_mongodb_url"`
	ExternalMongodbOplogUrl string    `gorm:"column:external_mongodb_oplog_url;type:VARCHAR;size:500;" json:"external_mongodb_oplog_url"`
	Apply                   bool      `gorm:"column:apply;type:BOOL;" json:"apply"`
	Destroy                 bool      `gorm:"column:destroy;type:BOOL;" json:"destroy"`
	Deleted                 bool      `gorm:"column:deleted;type:BOOL;" json:"deleted"`
	DateCreate              time.Time `gorm:"column:date_create;autoCreateTime:true;" json:"date_create"`
	DateDeleted             time.Time `gorm:"column:date_deleted;" json:"date_deleted"`
}

// TableName sets the insert table name for this struct type
func (c *Auto_Aprov_Omnitok) TableName() string {
	return "auto_aprov_omnitok"
}

type Auto_Aprov_OmnitokRepository interface {
	FindAll(ctx context.Context) ([]Auto_Aprov_Omnitok, error)
	FindByID(ctx context.Context, id uint) (Auto_Aprov_Omnitok, error)
	FindByServicio(ctx context.Context, idservicio uint) ([]Auto_Aprov_Omnitok, error)
	Save(ctx context.Context, auto_Aprov_Omnitok Auto_Aprov_Omnitok) (Auto_Aprov_Omnitok, error)
	SaveAll(ctx context.Context, auto_Aprov_Omnitok []Auto_Aprov_Omnitok) (string, error)
	Update(ctx context.Context, auto_Aprov_Omnitok Auto_Aprov_Omnitok) (Auto_Aprov_Omnitok, error)
	Delete(ctx context.Context, auto_Aprov_Omnitok Auto_Aprov_Omnitok) error
}
