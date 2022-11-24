package models

type MovieModel struct {
	Code   int64  `json:"code"`
	Status string `json:"status"`
	Data   *Data  `json:"data"`
	Error  error
}
type Data struct {
	Movie *Movie `json:"movie"`
}

type Movie struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Genre    string  `json:"genre"`
	Rating   float64 `json:"rating"`
	Plot     string  `json:"plot"`
	Released bool    `json:"released"`
}
