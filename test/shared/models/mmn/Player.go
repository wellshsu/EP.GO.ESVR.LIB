package mmn

import (
	"fmt"

	"github.com/hsu2017/EP.GO.ESVR.LIB/core/xutility/xorm"
	"github.com/hsu2017/EP.GO.ESVR.LIB/test/shared/models/mdf"
)

type Player struct {
	xorm.Table `orm:"-" json:"-"`
	ID         int    `orm:"column(id);pk"`
	Account    string `orm:"column(account)"`  // 账号
	Password   string `orm:"column(password)"` // 密码
	Online     int    `orm:"column(online)"`   // 0-不在线 1-在线
	GateUrl    string `orm:"column(gate_url)"` // 网关地址
}

func RegPlayer() *Player {
	this := NewPlayer()
	xorm.RegisterModel(this)
	return this
}

func NewPlayer() *Player {
	this := new(Player)
	this.CTOR(this)
	return this
}

func (this *Player) DBName() string {
	return mdf.TYPE_DB_MAIN
}

func (this *Player) TableName() string {
	return "player"
}

func (this *Player) RedisName() string {
	if this.ID == xorm.ORM_INT_UNDEFINED {
		return fmt.Sprintf("%v_%v", this.DBName(), this.TableName())
	} else {
		return fmt.Sprintf("%v_%v_%v", this.DBName(), this.TableName(), this.ID)
	}
}

func (this *Player) RedisPrefix() string {
	return fmt.Sprintf("%v_%v", this.DBName(), this.TableName())
}
