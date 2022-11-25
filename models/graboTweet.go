package models

import "time"

type GraboTweet struct {
	UserId  string    `bson:"userid" json:"user_id, omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje, omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha, omitempty"`
}
