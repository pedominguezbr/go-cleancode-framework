package domain

import (
	"context"
	"time"
)

/*
DB Table Details
-------------------------------------

JSON Sample
-------------------------------------
{    "rol_id": 63,    "rol_nombre": "xHiDZhFKQVQKsFAREDRDyQadH",    "rol_descripcion": "KVurlJlQKJTFgAFsMIxhRfbet",    "rol_esactivo": true,    "rol_eliminado": false,    "rol_fecharegistro": "2043-11-30T10:44:26.891008233-05:00"}



*/

// CXROL struct is a row record of the CX_ROL table in the chambaxtra database
type Rol struct {
	//[ 0] rol_id                                         INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID uint `gorm:"primary_key;column:rol_id;autoIncrement;" json:"rol_id"`
	//[ 1] rol_nombre                                     VARCHAR(20)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 20      default: []
	RolNombre string `gorm:"column:rol_nombre;type:VARCHAR;size:20;" json:"rol_nombre"`
	//[ 2] rol_descripcion                                VARCHAR(100)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	RolDescripcion string `gorm:"column:rol_descripcion;type:VARCHAR;size:100;" json:"rol_descripcion"`
	//[ 3] rol_esactivo                                   BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	RolEsactivo bool `gorm:"column:rol_esactivo;type:BOOL;" json:"rol_esactivo"`
	//[ 4] rol_eliminado                                  BOOL                 null: true   primary: false  isArray: false  auto: false  col: BOOL            len: -1      default: []
	RolEliminado bool `gorm:"column:rol_eliminado;type:BOOL;" json:"rol_eliminado"`
	//[ 5] rol_fecharegistro                              DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	RolFecharegistro time.Time `gorm:"column:rol_fecharegistro;autoCreateTime:true;" json:"rol_fecharegistro"`
}

// TableName sets the insert table name for this struct type
func (c *Rol) TableName() string {
	return "rol"
}

type RolRepository interface {
	FindAll(ctx context.Context) ([]Rol, error)
	FindByID(ctx context.Context, id uint) (Rol, error)
	Save(ctx context.Context, rol Rol) (Rol, error)
	Update(ctx context.Context, rol Rol) (Rol, error)
	Delete(ctx context.Context, rol Rol) error
	DeleteRolUser(ctx context.Context, rol Rol) error
}
