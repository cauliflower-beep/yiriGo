package curd

type ArticleTag struct {
	article_id  int64
	tag_id      int64
	status      int64
	create_time int64
}

func (ArticleTag) TableName() string {
	return "t_article_tag"
}

// 这里不能定义成一个指针类型，否则一切改动都会影响数据本身
var c = Character{
	Name:     "conan",
	Age:      16,
	Fight:    48,
	Birthday: 800434227,
	Pal:      "小兰",
	Enemy: struct {
		String string
		Valid  bool
	}{String: "黑衣组织", Valid: true},
	Hobby: []string{"足球", "推理"},
	JobInfo1: Job{
		Title:    "小学生",
		Location: "帝丹小学",
	},
	JobInfo2: Job{
		Title:    "侦探",
		Location: "毛利侦探事务所",
	},
}

// 创建记录的几种方式：

// CreateRec 直接传入一个对象
func CreateRec() {
	c1 := c
	res := DB.Create(&c1)
	Println(res.RowsAffected, res.Error, c) // res.RowsAffected 返回找到的记录数，相当于 `len(res)`
}

// SelectCreate 指定几个字段创建
func SelectCreate() {
	c2 := c
	// 只会创建选中的字段 以及 curd 默认追踪的字段（例如 updateAt）. 未选中的字段为空
	DB.Select("Name", "Age").Create(&c2)
}

// OmitCreate 忽略某些字段，其余字段增加
func OmitCreate() {
	c3 := c
	DB.Omit("Pal").Create(&c3)
}

// BatchCreate 批量 insert. 传入一个对象数组（切片）
func BatchCreate() {
	var cs = []Character{
		{
			Name:     "Naruto",
			Age:      17,
			Birthday: 800434227,
			JobInfo1: Job{
				Title:    "人柱力",
				Location: "木叶村",
			},
			JobInfo2: Job{
				Title:    "火影",
				Location: "木叶村",
			},
		},
		{
			Name:     "Goku",
			Age:      28,
			Birthday: 800434227,
			JobInfo1: Job{
				Title:    "武道家",
				Location: "包子山",
			},
			JobInfo2: Job{
				Title:    "赛亚人",
				Location: "贝吉塔行星",
			},
		},
		{
			Name:     "lufy",
			Age:      23,
			Birthday: 800434227,
			JobInfo1: Job{
				Title:    "海贼",
				Location: "伟大航路",
			},
			JobInfo2: Job{
				Title:    "乔伊波伊",
				Location: "伟大航路",
			},
		},
	}
	// 可以直接将上述切片传递给 Create 方法
	DB.Create(&cs)

	// 也可以使用 CreateInBatches 创建，并指定创建数量
	//DB.CreateInBatches(cs, 10)
}
