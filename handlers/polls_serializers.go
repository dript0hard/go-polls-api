package handlers

import (
	"net/http"
	"time"

	"github.com/dript0hard/pollsapi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreatePollRequest struct {
	Question string    `json:"question"`
	Choices  []string  `json:"choices"`
	UserId   uuid.UUID `json:"user_id"`
}

func (p *CreatePollRequest) Bind(r *http.Request) error {
	return nil
}

type ChoiceSerializer struct {
	PollID     uuid.UUID      `json:"poll_id"`
	ID         uuid.UUID      `json:"id"`
	ChoiceText string         `json:"choice_text"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,-"`
}

type ListChoiceResponse struct {
	Choices []*ChoiceSerializer `json:"choices"`
}

func (p *ListChoiceResponse) Render(
                            w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewChoice(c *models.Choice) *ChoiceSerializer {
	return &ChoiceSerializer{
		ID:         c.ID,
		ChoiceText: c.ChoiceText,
		CreatedAt:  c.CreatedAt,
		DeletedAt:  c.DeletedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

type PollSerializer struct {
	ID        uuid.UUID           `json:"id"`
	UserID    uuid.UUID           `json:"user_id"`
	Question  string              `json:"question"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
    Choices   []*ChoiceSerializer `json:"choices ,omitempty" gorm:"-"`
    DeletedAt gorm.DeletedAt      `json:"deleted_at ,-"`
}

func NewPoll(p *models.Poll) *PollSerializer {
	return &PollSerializer{
		ID:        p.ID,
		UserID:    p.UserID,
		Question:  p.Question,
		CreatedAt: p.CreatedAt,
		DeletedAt: p.DeletedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (p *PollSerializer) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

type PollResponse struct {
	Poll *PollSerializer `json:"poll"`
}

func (p *PollResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type ListPollsResponse struct {
	Polls []*PollSerializer `json:"polls"`
}

func (p *ListPollsResponse) Render(
                                w http.ResponseWriter, r *http.Request) error {
	return nil
}

