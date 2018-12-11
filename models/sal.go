package models

import "time"

// Fighter Represents an S.A.L. Fighter
type Fighter struct {
	ID       int64
	Nickname string
}

//Match Hold Match Meta data
type Match struct {
	ID          int64
	Contestants int
	Date        time.Time
}

// MatchResult Holds match results for a Fighter
type MatchResult struct {
	ID             int64
	FighterID      int64
	KOs            int
	Falls          int
	DamageTaken    int
	DamageRecieved int
	Place          int
}
