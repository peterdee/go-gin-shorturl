package link

type CreateLinkPayload struct {
	OriginalURL string `form:"originalURL"`
	Password    string `form:"password"`
}

type DeleteLinkPayload struct {
	Password string `form:"password"`
	ShortID  string `form:"shortID"`
}
