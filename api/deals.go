package api

type GetDealsParams struct {
	FilterID      *int   `form:"filter_id" example:"123"`
	IDs           string `form:"ids" example:"1,2,3"`
	OwnerID       *int   `form:"owner_id" example:"456"`
	PersonID      *int   `form:"person_id" example:"789"`
	OrgID         *int   `form:"org_id" example:"101"`
	PipelineID    *int   `form:"pipeline_id" example:"202"`
	StageID       *int   `form:"stage_id" example:"303"`
	Status        string `form:"status" example:"open,winner"`
	UpdatedSince  string `form:"updated_since" example:"2025-01-01T10:20:00Z"`
	UpdatedUntil  string `form:"updated_until" example:"2025-01-02T10:20:00Z"`
	SortBy        string `form:"sort_by" example:"update_time"`
	SortDirection string `form:"sort_direction" example:"desc"`
	IncludeFields string `form:"include_fields" example:"next_activity_id,last_activity_id"`
	CustomFields  string `form:"custom_fields" example:"custom_field_1,custom_field_2"`
	Limit         *int   `form:"limit" example:"100"`
	Cursor        string `form:"cursor" example:"cursor_token"`
}

type GetDealsResponse struct {
	Success        bool        `json:"success" example:"true"`
	Data           interface{} `json:"data"`
	AdditionalData interface{} `json:"additional_data"`
}

type CreateDealParams struct {
	Title             string   `form:"title" binding:"required" example:"Deal Title"`
	OwnerID           *int     `form:"owner_id" example:"1"`
	PersonID          *int     `form:"person_id" example:"1"`
	OrgID             *int     `form:"org_id" example:"1"`
	PipelineID        *int     `form:"pipeline_id" example:"1"`
	StageID           *int     `form:"stage_id" example:"1"`
	Value             *float64 `form:"value" example:"200.5"`
	Currency          string   `form:"currency" example:"USD"`
	AddTime           string   `form:"add_time" example:"2021-01-01T00:00:00Z"`
	UpdateTime        string   `form:"update_time" example:"2021-01-01T00:00:00Z"`
	StageChangeTime   string   `form:"stage_change_time" example:"2021-01-01T00:00:00Z"`
	IsDeleted         *bool    `form:"is_deleted" example:"false"`
	Status            string   `form:"status" example:"open"`
	Probability       *float64 `form:"probability" example:"90"`
	LostReason        string   `form:"lost_reason" example:"Lost Reason"`
	VisibleTo         *int     `form:"visible_to" example:"7"`
	CloseTime         string   `form:"close_time" example:"2021-01-01T00:00:00Z"`
	WonTime           string   `form:"won_time" example:"2021-01-01T00:00:00Z"`
	LostTime          string   `form:"lost_time" example:"2021-01-01T00:00:00Z"`
	ExpectedCloseDate string   `form:"expected_close_date" example:"2021-01-01"`
	LabelIDs          string   `form:"label_ids" example:"1,2,3"`
}

type CreateDealResponse struct {
	Success bool        `json:"success" example:"true"`
	Data    interface{} `json:"data"`
}
