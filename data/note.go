package data

type Note struct {
	ID        string `json:"id" bson:"_id"`
	Owner     string `json:"owner" bson:"owner"`
	Title     string `json:"title" bson:"title"`
	Content   string `json:"content" bson:"content"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}
