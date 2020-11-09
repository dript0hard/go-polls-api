package handlers

import (
	"net/http"
	"time"

	"github.com/dript0hard/pollsapi/database"
	pollserrors "github.com/dript0hard/pollsapi/errors"
	"github.com/dript0hard/pollsapi/models"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

var (
	pollDb, _ = database.OpenDB()
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
	ID         uuid.UUID `json:"id"`
	ChoiceText string    `json:"choice_text"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewChoice(c *models.Choice) *ChoiceSerializer {
	return &ChoiceSerializer{
		ID:         c.ID,
		ChoiceText: c.ChoiceText,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}
}

type PollSerializer struct {
	ID        uuid.UUID           `json:"id"`
	UserID    uuid.UUID           `json:"user_id"`
	Question  string              `json:"question"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	Choices   []*ChoiceSerializer `json:"choices,omitempty"`
}

func NewPoll(p *models.Poll) *PollSerializer {
	return &PollSerializer{
		ID:        p.ID,
		UserID:    p.UserID,
		Question:  p.Question,
		CreatedAt: p.CreatedAt,
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

func createPoll(w http.ResponseWriter, r *http.Request) {
	data := &CreatePollRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, pollserrors.ErrInvalidRequest(err))
		return
	}

	poll := &models.Poll{Question: data.Question, UserID: data.UserId}
	if err := pollDb.Save(poll).Error; err != nil {
		render.Render(w, r, pollserrors.ErrInvalidRequest(err))
		return
	}

	choiceList := []*ChoiceSerializer{}
	for _, choice := range data.Choices {
		c := &models.Choice{ChoiceText: choice, PollID: poll.ID}
		if err := pollDb.Save(c).Error; err != nil {
			render.Render(w, r, pollserrors.ErrInvalidRequest(err))
			return
		}
		choiceList = append(choiceList, NewChoice(c))
	}

	pollResp := &PollResponse{
		Poll: NewPoll(poll),
	}
	pollResp.Poll.Choices = choiceList
	render.Status(r, http.StatusOK)
	render.Render(w, r, pollResp)
	return
}

func listPolls(w http.ResponseWriter, r *http.Request) {

}

func getPollById(w http.ResponseWriter, r *http.Request) {

}

func updatePoll(w http.ResponseWriter, r *http.Request) {

}

func deletePoll(w http.ResponseWriter, r *http.Request) {

}

func listChoices(w http.ResponseWriter, r *http.Request) {

}

func voteChoice(w http.ResponseWriter, r *http.Request) {

}
