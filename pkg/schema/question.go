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
	LastEditDate     int64    `json:"last_edit_date"`
	Link             string   `json:"link"`
	Title            string   `json:"title"`
	Body             string   `json:"body"`
	Tags             []string `json:"tags"`
	Owner            Owner    `json:"owner"`
	PostedTimeAgo    int
	UnitPeroid       string
}

type Owner struct {
	Reputation   int    `json:"reputation"`
	UserID       int    `json:"user_id"`
	UserType     string `json:"user_type"`
	ProfileImage string `json:"profile_image"`
	DisplayName  string `json:"display_name"`
	Link         string `json:"link"`
}

type Page struct {
	Index int
}
