package domain

import (
	"context"
	"time"
)

/*

JSON Sample
-------------------------------------
{    "ru_usuarioid": 20,    "ru_rol_Usuarioid": 36,    "ru_esactivo": false,    "ru_fecharegistro": "2237-10-01T09:10:21.50787503-05:00"}


*/

// CXROLUSUARIO struct is a row record of the CX_ROL_USUARIO table in the chambaxtra database
type Rol_Usuario struct {
	//[ 0] ru_usuarioid                                   INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	UserID uint `gorm:"primary_key;column:user_id;autoIncrement:false" json:"user_id"`
	//[ 1] ru_rolid                                       INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	RolID      uint `gorm:"primary_key;column:rol_id;autoIncrement:false" json:"rol_id"`
	RuEsactivo bool `gorm:"column:ru_esactivo;type:BOOL;default:true" json:"ru_esactivo"`

	//[ 3] ru_fecharegistro                               DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	RuFecharegistro time.Time `gorm:"column:ru_fecharegistro;autoCreateTime:true;" json:"ru_fecharegistro"`
}

// TableName sets the insert table name for this struct type
func (c *Rol_Usuario) TableName() string {
	return "rol_usuarios"
}

type Rol_UsuarioRepository interface {
	FindAll(ctx context.Context) ([]Rol_Usuario, error)
	FindByID(ctx context.Context, id uint) (Rol_Usuario, error)
	FindByIDUserIDRol(ctx context.Context, idUser uint, idRol uint) ([]Rol_Usuario, error)
	Save(ctx context.Context, rol_Usuario Rol_Usuario) (Rol_Usuario, error)
	Update(ctx context.Context, rol_Usuario Rol_Usuario) (Rol_Usuario, error)
	Delete(ctx context.Context, rol_Usuario Rol_Usuario) error
	DeleteArray(ctx context.Context, rol_Usuario []Rol_Usuario) error
}
