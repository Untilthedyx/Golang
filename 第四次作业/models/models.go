package models

type User struct {
	UserId   int       `gorm:"primaryKey;column:user_id"`
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Posts    []Post    `gorm:"foreignKey:UserId"`
	Comments []Comment `gorm:"foreignKey:UserId"`
}

type Post struct {
	PostId    int    `gorm:"primaryKey;column:post_id"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	UserId    int
	CreatedAt string
	UpdatedAt string
	Comments  []Comment `gorm:"foreignKey:PostId"`
}

type Comment struct {
	CommentId int    `gorm:"primaryKey;column:comment_id"`
	Content   string `gorm:"not null"`
	UserId    int
	PostId    int
	CreatedAt string
}
