package models

import (
    "github.com/google/uuid"
    "time"
)

type User struct {
    ID        uuid.UUID   `bson:"_id" json:"id" ` // use UUID as the primary key
    Name      string      `bson:"name" json:"name"  binding:"required"`
    UserName  string      `bson:"username" json:"username"  binding:"required"`
    Email     string      `bson:"email" json:"email"  binding:"required"`
    Password  string      `bson:"password" json:"password"  binding:"required"` // don't expose password in JSON
    Bio       string      `bson:"bio,omitempty" json:"bio,omitempty"`
    AvatarURL string      `bson:"avatar_url,omitempty" json:"avatar_url,omitempty"`
    Followers []uuid.UUID `bson:"followers,omitempty" json:"followers,omitempty"`
    Following []uuid.UUID `bson:"following,omitempty" json:"following,omitempty"`
    CreatedAt time.Time   `bson:"created_at,omitempty" json:"created_at,omitempty"`
    UpdatedAt time.Time   `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}


type Post struct {
	ID        uuid.UUID   `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    uuid.UUID   `bson:"user_id" json:"user_id"`
	Caption   string               `bson:"caption,omitempty" json:"caption,omitempty"`
	ImageURL  string               `bson:"image_url" json:"image_url"  binding:"required"`
	Likes     []uuid.UUID `bson:"likes,omitempty" json:"likes,omitempty"`
	Comments  []Comment            `bson:"comments,omitempty" json:"comments,omitempty"`
	CreatedAt time.Time            `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type Comment struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	PostID    uuid.UUID `bson:"post_id" json:"post_id"`
	UserID    uuid.UUID `bson:"user_id" json:"user_id"`
	Text      string    `bson:"text" json:"text"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type Story struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    uuid.UUID `bson:"user_id" json:"user_id"`
	MediaURL  string             `bson:"media_url" json:"media_url"`
	ExpiresAt time.Time          `bson:"expires_at" json:"expires_at"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type Message struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	SenderID  uuid.UUID `bson:"sender_id" json:"sender_id"`
	ReceiverID uuid.UUID `bson:"receiver_id" json:"receiver_id"`
	Text       string             `bson:"text" json:"text"`
	SentAt     time.Time          `bson:"sent_at" json:"sent_at"`
}

type Notification struct {
	ID        uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    uuid.UUID `bson:"user_id" json:"user_id"` // who receives the notification
	Message   string             `bson:"message" json:"message"`
	IsRead    bool               `bson:"is_read" json:"is_read"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
