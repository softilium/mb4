package schema

import "entgo.io/ent"

// Ticker holds the schema definition for the Ticker entity.
type Ticker struct {
	ent.Schema
}

// Fields of the Ticker.
func (Ticker) Fields() []ent.Field {
	return nil
}

// Edges of the Ticker.
func (Ticker) Edges() []ent.Edge {
	return nil
}
