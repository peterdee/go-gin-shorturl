package mongodb

type Link struct {
	CreatedAt    int    `bson:"createdAt,omitempty"`
	OriginalURL  string `bson:"originalURL,omitempty"`
	PasswordHash string `bson:"passwordHash,omitempty"`
	ShortID      string `bson:"shortID,omitempty"`
	UpdatedAt    int    `bson:"updatedAt,omitempty"`
}
