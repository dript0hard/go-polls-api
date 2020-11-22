package handlers

import (
	"errors"
	"net/http"

	"github.com/dript0hard/pollsapi/database"
	pollserrors "github.com/dript0hard/pollsapi/errors"
	"github.com/dript0hard/pollsapi/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	pollDb, _ = database.OpenDB()
)

func createPoll(w http.ResponseWriter, r *http.Request) {
	data := &CreatePollRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, pollserrors.ErrInvalidRequest(err))
		return
	}

	// Create the poll.
	poll := &models.Poll{Question: data.Question, UserID: data.UserId}
	if err := pollDb.Save(poll).Error; err != nil {
		render.Render(w, r, pollserrors.ErrInvalidRequest(err))
		return
	}

	// Create and add all the ids of choices for that poll.
	choiceList := []*ChoiceSerializer{}
	for _, choice := range data.Choices {
		c := &models.Choice{ChoiceText: choice, PollID: poll.ID}
		if err := pollDb.Save(c).Error; err != nil {
			render.Render(w, r, pollserrors.ErrInvalidRequest(err))
			return
		}
		choiceList = append(choiceList, NewChoice(c))
	}

	//Done
	pollResp := &PollResponse{
		Poll: NewPoll(poll),
	}
	pollResp.Poll.Choices = choiceList
	render.Status(r, http.StatusOK)
	render.Render(w, r, pollResp)
	return
}

func listPolls(w http.ResponseWriter, r *http.Request) {
    polls := &ListPollsResponse{}
    pollDb.Table("polls").Find(&polls.Polls)

	render.Status(r, http.StatusOK)
	render.Render(w, r, polls)
	return
}

func getPollById(w http.ResponseWriter, r *http.Request) {
	pollID := chi.URLParamFromCtx(r.Context(), "pollId")
	poll := &models.Poll{}
	choices := []*ChoiceSerializer{}

	if err := pollDb.Table("polls").Where("id = ?", pollID).
		First(poll).Error; err != nil {
		render.Render(w, r, pollserrors.ErrPollDoesNotExist)
		return
	}

	pollDb.Table("choices").Where("poll_id = ?", pollID).Find(&choices)

	pollSerializer := NewPoll(poll)
	pollSerializer.Choices = choices

	render.Status(r, http.StatusOK)
	render.Render(w, r, pollSerializer)
	return
}

func updatePoll(w http.ResponseWriter, r *http.Request) {

}

func deletePoll(w http.ResponseWriter, r *http.Request) {
	pollID := chi.URLParamFromCtx(r.Context(), "pollId")
	pid, err := uuid.Parse(pollID)
	if err != nil {
		render.Render(w, r, pollserrors.ErrPollDoesNotExist)
		return
	}

    poll := &models.Poll{}
	poll.ID = pid
	if err := pollDb.Delete(poll).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			render.Render(w, r, pollserrors.ErrPollDoesNotExist)
			return
		}
		render.Render(w, r, pollserrors.ErrInternalServerErr(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.Render(w, r, NewPoll(poll))
	return
}

func listChoices(w http.ResponseWriter, r *http.Request) {
	pollID := chi.URLParamFromCtx(r.Context(), "pollId")
	pid, err := uuid.Parse(pollID)
	if err != nil {
		render.Render(w, r, pollserrors.ErrPollDoesNotExist)
		return
	}

    choices := &ListChoiceResponse{}
	pollDb.Table("choices").Where("poll_id = ?", pid).Find(&choices.Choices)
	render.Status(r, http.StatusOK)
	render.Render(w, r, choices)
	return
}

func voteChoice(w http.ResponseWriter, r *http.Request) {
	choiceId, err := uuid.Parse(chi.URLParamFromCtx(r.Context(), "choiceId"))
	if err != nil {
		render.Render(w, r, pollserrors.ErrChoiceDoesNotExist)
		return
	}

	pollId, err := uuid.Parse(chi.URLParamFromCtx(r.Context(), "pollId"))
	if err != nil {
		render.Render(w, r, pollserrors.ErrPollDoesNotExist)
		return
	}

    //(TODO: dript0hard)
    // Hardcode the users UUID for now later will be taken from token.
    uid, _ := uuid.Parse("430fe4d2-affd-485e-be5a-eee4a552eae0")
    vote := &models.Vote{
        PollID: pollId,
        ChoiceID: choiceId,
        UserID: uid,
    }

    if err := pollDb.Save(vote).Error; err != nil {
        render.Render(w, r, pollserrors.ErrInvalidRequest(err))
        return
    }

	render.Status(r, http.StatusOK)
	return
}
