package models

type Request struct {
	X   float32 `json: "x"`
	Y   float32 `json: "y"`
	Z   float32 `json: "z"`
	Vel float32 `json: "vel"`
}
