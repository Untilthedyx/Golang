package task

import "gorm.io/gorm"

// 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。
type User struct {
	Id         uint
	PostCounts int
	Posts      []Post `gorm:"foreignKey:UserId"`
}
type Post struct {
	Id            uint
	UserId        uint
	CommentCounts int
	State         string
	Comments      []Comment `gorm:"foreignKey:PostId"`
}
type Comment struct {
	Id     uint
	PostId uint
	Com    string
}

// 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
func FindUserPostComments(db *gorm.DB, id uint, posts []Post, comments []Comment) error {
	if err := db.Model(&User{}).Where("id=?", id).Association("posts").Find(&posts); err != nil {
		return err
	}
	if err := db.Model(&Post{}).Where("user_id=?", id).Association("comments").Find(&comments); err != nil {
		return err
	}
	return nil
}
func FindMaxCommentsPost(db *gorm.DB, post *Post) error {
	err := db.Model(&Post{}).Debug().Order("comment_counts desc").First(post)
	return err.Error
}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (p *Post) AfterCreate(db *gorm.DB) error {
	result := db.Model(&User{}).Where("id=?", p.UserId).Update("post_counts", gorm.Expr("post_counts+?", 1))
	return result.Error
}
func (c *Comment) BeforeDelete(db *gorm.DB) error {
	var post Post
	if err := db.Model(&Post{}).Where("id=?", c.PostId).Find(&post).Error; err != nil {
		return err
	}
	if post.CommentCounts == 0 {
		result := db.Model(&Post{}).Where("id=?", c.PostId).Update("state", "无评论")
		return result.Error
	}
	return nil
}
