package link

type createLinkRequestPayload struct {
	OriginalURL string `form:"originalURL"`
	Password    string `form:"password"`
}

// For Swagger
type createLinkResponsePayload struct {
	ShortID string `json:"shortID"`
}

// Fix for unused struct linter error
var _ = createLinkResponsePayload{}

type deleteLinkRequestPayload struct {
	Password string `form:"password"`
	ShortID  string `form:"shortID"`
}
