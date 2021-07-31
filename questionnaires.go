package main

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type Questionnaire struct {
    id uuid.UUID `json:"id"`
    name string `json:"name"`
    visibility bool `json:"visiblity"`
    questions string `json:"questions"`
}

var Questionnaires []Questionnaire

func (qnr *Questionnaire) getQuestionnaire(db *sql.DB) error {
    return errors.New("NYI")
}

func (qnr *Questionnaire) getQuestionnaires(db *sql.DB) ([]Questionnaire, error) {
    return nil, errors.New("NYI")
}

func (qnr *Questionnaire) updateQuestionnaire(db *sql.DB) error {
    return errors.New("NYI")
}

func (qnr *Questionnaire) deleteQuestionnaire(db *sql.DB) error {
    return errors.New("NYI")
}
