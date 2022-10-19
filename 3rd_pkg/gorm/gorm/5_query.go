package gorm

/*
	获取单条记录的方式有很多种
*/
var fc Character
var tc Character
var lc Character

func GetOnceRec() {
	// 按主键排序获取第一条记录
	DB.First(&fc)
	Println(fc)

	// 不排序获取第一条
	// 一般情况下查询到的数据顺序都是一样的，但也有极少数情况查询到的数据顺序不一样
	DB.Take(&tc)
	Println(tc)

	// 获取最后一条数据，按主键排序
	DB.Last(&lc)
	Println(lc)

	/*
		将获取到的记录填充到集合中
		由于集合本身是不带有任何字段属性的，而且我们也不知道要去查哪个表
		所以需要使用 gorm 提供的 Model 方法来告诉程序如何填充
		反观前面，不管是 First Take Last, 都可以获取到一个对象及他的指针，进而可以得知要去查哪个表
	*/
	res := map[string]interface{}{}
	DB.Model(&Character{}).First(res) // 若获取到的结果不对，可能是 birthday 序列化之后没有办法存进 map. 暂时可以将 birthday 字段注掉
	// 或者指定表名来查询,此时必须改用 Take 方法（这都是文档中写明了的）
	DB.Table("characters").Take(res)
	Println(res)
}

func GetRecWithCondition() {
	// 根据主键条件检索
	//db.First(&user, 10) 	// 后面的参数默认是给主键使用的
	//db.First(&user, "name=?","goku")	// 当然也可以通过指定参数名的方式来附加检索条件
}
