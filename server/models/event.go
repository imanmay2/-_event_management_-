package models

type Event struct {
	ID          string   `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	OrganizerID string   `json:"organizer_id" bson:"organizer_id"`
	Attendees   []string `json:"attendees" bson:"attendees"`
}
