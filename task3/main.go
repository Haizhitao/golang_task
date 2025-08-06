package main

import (
	"golang_task/task3/topic2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//#SQL语句练习
	//1: 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	// insert into students (name, age, grade) values ('张三', 20, '三年级');
	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	// select * from students where age > 18;
	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	// update students set grade = '四年级' where name = '张三';
	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	// delete from students where age < 15;

	//2: 事务语句
	//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	//start transaction;
	//declare account_a_balance decimal(10, 2);
	//select balance into account_a_balance from accounts where id = A for update;
	//if account_a_balance < 100 then
	//	rollback;
	//else
	//	update accounts set balance = balance - 100 where id = A;
	//	update accounts set balance = balance + 100 where id = B;
	//	insert into transactions (from_account_id, to_account_id, amount) values (A, B, 100);
	//	commit;
	//end if;

	//#Sqlx入门
	//1: 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	//dsn := "root:root@tcp(127.0.0.1:3306)/test?parseTime=true"
	//db, err := sqlx.Connect("mysql", dsn)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	////编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	//topic1.Run(db)
	////编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	//topic1.Run2(db)

	//2.假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	//topic1.Run3(db)

	//#gorm进阶
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//1.假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	//要求 ：
	//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	//编写Go代码，使用Gorm创建这些模型对应的数据库表。
	//topic2.Run(db)

	//2：关联查询
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//topic2.Run2(db)
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	//topic2.Run3(db)

	//3.钩子函数
	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	/*func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
		var count int64
		tx.Model(&Post{}).Where("user_id = ?", p.UserID).Count(&count)
		tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", count)
		return
	}*/
	//topic2.Run(db)
	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	topic2.Run4(db)
}
