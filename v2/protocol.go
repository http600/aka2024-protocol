package np

type Protocol interface{}

type PlayItem struct {
    PlayLink   string   `json:"play_link,omitempty"`
    TextLink   string   `json:"text_link,omitempty"`
    AvatarLink string   `json:"avatar_link,omitempty"`
    CoverLink  string   `json:"cover_link,omitempty"`
    Title      string   `json:"title,omitempty"`
    Tags       []string `json:"tags,omitempty"`
    ViewCount  int      `json:"view_count,omitempty"`
    AgeInfo    string   `json:"age_info,omitempty"`
}

type TimelineResponse struct {
    Status       int        `json:"status,omitempty"`
    TimelineList []PlayItem `json:"timeline_list,omitempty"`
}

type SyncTimedTextRequest struct {
    ItemKey int `json:"item_key,omitempty"`
}

type QueryTimedTextRequest struct {
    Term string `json:"term,omitempty"`
}

type Query4PlayRequest struct {
    Term string `json:"term,omitempty"`
}
