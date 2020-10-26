package models

import(
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

// BaseSchema : base struct for db structs
type BaseSchema struct {
	ID        uuid.UUID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
    BaseSchema
    Username string
    Email    string
    Password string
}

type Poll struct {
    BaseSchema
    Question string
    UserID   uuid.UUID
    User     User
}

type Choice struct {
    BaseSchema
    ChoiceText string
    PollID     uuid.UUID
    Poll       Poll
}

type Vote struct {
    ChoiceID uuid.UUID
    Choice   Choice
    PollID   uuid.UUID
    Poll     Poll
    UserID   uuid.UUID
    User     User
}

func (b *BaseSchema) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
    return nil
}

