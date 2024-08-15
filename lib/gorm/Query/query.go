package curd

import (
	"time"
	conn "yiriGo/lib/gorm/Connection"
	"yiriGo/lib/gorm/model"
	"yiriGo/lib/gorm/utils"
)

/*
获取单条记录的方式有很多种
*/
var fc model.Character
var tc model.Character
var lc model.Character

// 测试gorm查询时，空字段会自动替换成上一个非空字段的问题
type role struct {
	Id         int       `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	ModifyTime time.Time `gorm:"column:modify_time"`
	NickName   string    `gorm:"column:nickname"`
}

func GetOnceRec() {
	// 按主键排序获取第一条记录
	conn.DB.First(&fc)
	utils.Println(fc)

	// 不排序获取第一条
	// 一般情况下查询到的数据顺序都是一样的，但也有极少数情况查询到的数据顺序不一样
	conn.DB.Take(&tc)
	utils.Println(tc)

	// 获取最后一条数据，按主键排序
	conn.DB.Last(&lc)
	utils.Println(lc)

	/*
		将获取到的记录填充到集合中
		由于集合本身是不带有任何字段属性的，而且我们也不知道要去查哪个表
		所以需要使用 curd 提供的 Model 方法来告诉程序如何填充
		反观前面，不管是 First Take Last, 都可以获取到一个对象及他的指针，进而可以得知要去查哪个表
	*/
	res := map[string]interface{}{}
	conn.DB.Model(&model.Character{}).First(res) // 若获取到的结果不对，可能是 birthday 序列化之后没有办法存进 map. 暂时可以将 birthday 字段注掉
	// 或者指定表名来查询,此时必须改用 Take 方法（这都是文档中写明了的）
	conn.DB.Table("characters").Take(res)
	utils.Println(res)
}

func GetRecWithCondition() {
	// 根据主键条件检索
	//db.First(&user, 10) 	// 后面的参数默认是给主键使用的
	var r role
	rec := conn.DB.Table("gormtest").Where("name = ?", "吕秀才").First(&r) // 通过指定参数名的方式来附加检索条件
	utils.Println(rec.Error, "|", rec.RowsAffected, r)                  // 没有记录时，err不为空
}

func GetRecs() {
	//var roles []role
	//conn.DB.Table("gormtest").Find(&roles)
	//fmt.Println(roles[1].ModifyTime)
	//Println(roles)

	var recs []map[string]interface{}
	conn.DB.Table("gormtest").Find(&recs)
	//fmt.Println(recs[1].ModifyTime)
	utils.Println(recs)
}
