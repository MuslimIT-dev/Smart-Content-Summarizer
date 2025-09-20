package models

type RequestBody struct {
	Text string `json:"text"`
}

type HFResponse []struct {
	SummaryText string `json:"summary_text"`
}
