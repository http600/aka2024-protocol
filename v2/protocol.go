package v2

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
