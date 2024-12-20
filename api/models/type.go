package models

import "github.com/lib/pq"

// MoveType represents a type in the application. It contains properties such as GenericId, Name, DoubleDamage,
// HalfDamage, and NoDamage. The GenericId field is the unique identifier of the type. The Name field represents
// the name of the type. The DoubleDamage field contains a list of type IDs that are super effective against
// this type. The HalfDamage field contains a list of type IDs that are not very effective against this type.
// The NoDamage field contains a list of type IDs that have no effect on this type.
type MoveType struct {
	GenericId `gorm:"primary_key;column:id" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	//TODO: convert to manytomany table somehow
	DoubleDamageTo   pq.Int64Array `gorm:"type:integer[]" json:"double_damage_to,omitempty"`
	HalfDamageTo     pq.Int64Array `gorm:"type:integer[]"  json:"half_damage_to,omitempty"`
	NoDamageTo       pq.Int64Array `gorm:"type:integer[]"  json:"no_damage_to,omitempty"`
	DoubleDamageFrom pq.Int64Array `gorm:"type:integer[]" json:"double_damage_from,omitempty"`
	HalfDamageFrom   pq.Int64Array `gorm:"type:integer[]"  json:"half_damage_from,omitempty"`
	NoDamageFrom     pq.Int64Array `gorm:"type:integer[]"  json:"no_damage_from,omitempty"`
	ImgUrl           string        `json:"img_url,omitempty"`
}
