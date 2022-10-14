package domain

import (
	"context"
	"time"
)

type User struct {
	ID uint `gorm:"primary_key;column:us_id;autoIncrement;" json:"us_id"`
	//[ 1] us_usuario                                     VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	UsUsuario string `gorm:"column:us_usuario;type:VARCHAR;size:50;" json:"us_usuario"`
	//[ 2] us_nombre                                      VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	UsNombre string `gorm:"column:us_nombre;type:VARCHAR;size:50;" json:"us_nombre"`
	//[ 3] us_apellido                                    VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	UsApellido string `gorm:"column:us_apellido;type:VARCHAR;size:50;" json:"us_apellido"`
	//[ 4] us_correo                                      VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	UsCorreo string `gorm:"column:us_correo;type:VARCHAR;size:100;" json:"us_correo"`
	//[ 5] us_clave                                       VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	UsClave string `gorm:"column:us_clave;type:VARCHAR;size:50;" json:"us_clave"`
	//[ 6] us_esactivo                                    BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	UsEsactivo bool `gorm:"column:us_esactivo;type:BOOL;" json:"us_esactivo"`
	//[ 7] us_eliminado                                   BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	UsEliminado bool `gorm:"column:us_eliminado;type:BOOL;" json:"us_eliminado"`
	//[ 8] us_fecharegistro                               DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	UsFecharegistro time.Time `gorm:"column:us_fecharegistro;autoCreateTime:true;" json:"us_fecharegistro"`

	Roles []Rol `gorm:"many2many:cx_rol_usuarios;association_jointable_foreignkey:rol_id" json:"us_roles"`
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Tokens struct {
	Token Token `json:"token"`
}
type Token struct {
	Expires_in    int64  `json:"expires_in"`
	Access_token  string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id uint) (User, error)
	FindByIDWithRole(ctx context.Context, id uint) (User, error)
	Save(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, user User) error
	FindByUserName(ctx context.Context, username string) (User, error)
}

// TableName sets the insert table name for this struct type
func (c *User) TableName() string {
	return "usuario"
}

// func (User) BeforeCreate(db *gorm.DB) error {
// 	err := db.SetupJoinTable(&User{}, "Roles", &Rol_Usuario{})
// 	return err
// }

// func (c *User) BeforeSave(db *gorm.DB) error {
// 	c.UsEsactivo = true
// 	return nil
// }

// // BeforeSave invoked before saving, return an error if field is not populated.
// func (c *User) BeforeSave() error {
// 	return nil
// }

// // Prepare invoked before saving, can be used to populate fields etc.
// func (c *User) Prepare() {
// }

// // Validate invoked before performing action, return an error if field is not populated.

// // TableInfo return table meta data
// func (c *User) TableInfo() *TableInfo {
// 	return UserTableInfo
// }
