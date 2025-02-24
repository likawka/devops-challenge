package api

type GetAllDealsParams struct {
	UserID     *int   `form:"user_id" example:"123"`
	FilterID   *int   `form:"filter_id" example:"456"`
	StageID    *int   `form:"stage_id" example:"2"`
	Status     string `form:"status" example:"open"`
	Start      *int   `form:"start" example:"0"`
	Limit      *int   `form:"limit" example:"100"`
	Sort       string `form:"sort" example:"update_time DESC"`
	OwnedByYou *int   `form:"owned_by_you" example:"1"`
}

type Deal struct {
	ID                    int      `json:"id" example:"1"`
	StageID               int      `json:"stage_id" example:"2"`
	Title                 string   `json:"title" example:"Deal One"`
	Value                 float64  `json:"value" example:"130"`
	Currency              string   `json:"currency" example:"EUR"`
	AddTime               string   `json:"add_time" example:"2019-05-29 04:21:51"`
	UpdateTime            string   `json:"update_time" example:"2019-11-28 16:19:50"`
	StageChangeTime       string   `json:"stage_change_time" example:"2019-11-28 15:41:22"`
	Active                bool     `json:"active" example:"true"`
	Deleted               bool     `json:"deleted" example:"false"`
	Status                string   `json:"status" example:"open"`
	Probability           *int     `json:"probability,omitempty"`
	NextActivityDate      string   `json:"next_activity_date" example:"2019-11-29"`
	NextActivityTime      string   `json:"next_activity_time" example:"11:30:00"`
	NextActivityID        *int     `json:"next_activity_id,omitempty"`
	LastActivityID        *int     `json:"last_activity_id,omitempty"`
	LastActivityDate      *string  `json:"last_activity_date,omitempty"`
	LostReason            *string  `json:"lost_reason,omitempty"`
	VisibleTo             string   `json:"visible_to" example:"1"`
	CloseTime             *string  `json:"close_time,omitempty"`
	PipelineID            int      `json:"pipeline_id" example:"1"`
	WonTime               *string  `json:"won_time,omitempty"`
	FirstWonTime          *string  `json:"first_won_time,omitempty"`
	LostTime              *string  `json:"lost_time,omitempty"`
	ProductsCount         int      `json:"products_count" example:"2"`
	FilesCount            int      `json:"files_count" example:"0"`
	NotesCount            int      `json:"notes_count" example:"2"`
	FollowersCount        int      `json:"followers_count" example:"0"`
	EmailMessagesCount    int      `json:"email_messages_count" example:"4"`
	ActivitiesCount       int      `json:"activities_count" example:"1"`
	DoneActivitiesCount   int      `json:"done_activities_count" example:"0"`
	UndoneActivitiesCount int      `json:"undone_activities_count" example:"1"`
	ParticipantsCount     int      `json:"participants_count" example:"1"`
	ExpectedCloseDate     string   `json:"expected_close_date" example:"2019-06-29"`
	LastIncomingMailTime  *string  `json:"last_incoming_mail_time,omitempty"`
	LastOutgoingMailTime  *string  `json:"last_outgoing_mail_time,omitempty"`
	Label                 string   `json:"label" example:"11"`
	Origin                string   `json:"origin" example:"ManuallyCreated"`
	Channel               *int     `json:"channel,omitempty"`
	ChannelID             *string  `json:"channel_id,omitempty"`
	StageOrderNr          *int     `json:"stage_order_nr,omitempty"`
	MRR                   *float64 `json:"mrr,omitempty"`
	ACVCurrency           string   `json:"acv_currency" example:"EUR"`
	ARRCurrency           string   `json:"arr_currency" example:"EUR"`
	MRRCurrency           string   `json:"mrr_currency" example:"EUR"`
}

type GetAllDealsResponse struct {
	Success bool   `json:"success" example:"true"`
	Data    []Deal `json:"data"`
}
