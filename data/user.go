package data

type User struct {
	ID           string   `json:"id" bson:"_id"`
	Username     string   `json:"username" bson:"username"`
	PasswordHash string   `json:"password" bson:"password"`
	Notes        []string `json:"notes" bson:"notes"`
	Timestamp    int64    `json:"timestamp" bson:"timestamp"`
}
