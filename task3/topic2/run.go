package topic2

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       int
	Gender    string
	Posts     []Post
	PostCount uint
}

type Post struct {
	gorm.Model
	Title         string
	Body          string
	UserID        uint
	Comments      []Comment
	CommentStatus string
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Post{}).Where("user_id = ?", p.UserID).Count(&count)
	tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", count)
	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Comment{}).Where("post_id = ", c.PostID).Count(&count)
	if count == 0 {
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论")
	}
	return
}

func generateDemoData(db *gorm.DB) {
	users := []User{
		{Name: "张三", Age: 25, Gender: "男"},
		{Name: "李四", Age: 30, Gender: "男"},
		{Name: "王五", Age: 28, Gender: "男"},
		{Name: "赵六", Age: 22, Gender: "女"},
		{Name: "钱七", Age: 35, Gender: "女"},
	}
	db.Create(&users)
	posts := []Post{
		{Title: "我的第一篇博客", Body: "今天开始写博客，记录我的学习历程。", UserID: 1, Comments: []Comment{
			{Content: "非常好", PostID: 1},
			{Content: "太赞了,总结的很到位哦", PostID: 1},
		}},
		{Title: "Go语言学习笔记", Body: "Go语言的并发模型非常强大，值得深入学习。", UserID: 1, Comments: []Comment{}},
		{Title: "GORM使用心得", Body: "GORM是一个优秀的ORM框架，简化了数据库操作。", UserID: 2, Comments: []Comment{}},
		{Title: "数据库设计经验", Body: "良好的数据库设计是系统成功的关键。", UserID: 2, Comments: []Comment{}},
		{Title: "Web开发实践", Body: "最近完成了一个Web项目，收获颇丰。", UserID: 1, Comments: []Comment{}},
		{Title: "编程技巧分享", Body: "分享一些提高编程效率的小技巧。", UserID: 3, Comments: []Comment{}},
		{Title: "项目总结", Body: "这个项目历时三个月，总结一下经验教训。。", UserID: 4, Comments: []Comment{}},
	}
	db.Create(&posts)
}

func Run(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	generateDemoData(db)
}

// Run2 查询某个用户发布的所有文章及其对应的评论信息
func Run2(db *gorm.DB) {
	var user User
	userId := 1
	//嵌套预加载
	db.Preload("Posts.Comments").First(&user, userId)
	fmt.Println(user)
}

// Run3 查询评论数量最多的文章信息
func Run3(db *gorm.DB) {
	var p Post
	db.Select("posts.*").Joins("left join comments on comments.post_id = posts.id").Group("posts.id").Order("count(comments.id) desc").First(&p)
	fmt.Println(p)
}

func Run4(db *gorm.DB) {
	comment := Comment{
		Content: "文笔不错,值得收藏呀!",
		PostID:  2,
	}
	db.Create(&comment)
	db.Delete(&comment)
}
