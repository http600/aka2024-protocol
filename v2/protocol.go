package np

type Protocol interface{}

type PlayItem struct {
    PlayURL   string   `json:"play_url,omitempty"`
    TextURL   string   `json:"text_url,omitempty"`
    AvatarURL string   `json:"avatar_url,omitempty"`
    Title     string   `json:"title,omitempty"`
    Tags      []string `json:"tags,omitempty"`
    ViewCount int      `json:"view_count,omitempty"`
    AgeInfo   string   `json:"age_info,omitempty"`
}

type TimelineResponse struct {
    Status       int        `json:"status,omitempty"`
    TimelineList []PlayItem `json:"timeline_list,omitempty"`
}
