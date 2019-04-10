package schema

import "time"

type Questions struct {
	Items          []Question `json:"items"`
	HasMore        bool       `json:"has_more"`
	QuotaMax       int        `json:"quota_max"`
	QuotaRemaining int        `json:"quota_remaining"`
	Backoff        int        `json:"backoff,omitempty"`
}

type Question struct {
	QuestionId       int   `json:"question_id"`
	ViewCount        int   `json:"view_count"`
	IsAnswered       bool  `json:"is_answered"`
	FavoriteCount    int   `json:"favorite_count"`
	DownVoteCount    int   `json:"down_vote_count"`
	UpVoteCount      int   `json:"up_vote_count"`
	AnswerCount      int   `json:"answer_count"`
	Score            int   `json:"score"`
	LastActivityDate int64 `json:"last_activity_date"`
	IntCreationDate  int64 `json:"creation_date"`
	CreationDate     time.Time
	LastEditDate     int64  `json:"last_edit_date"`
	Link             string `json:"link"`
	Title            string `json:"title"`
	Body             string `json:"body"`
}

type Page struct {
	Index int
}
