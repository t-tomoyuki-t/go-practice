package object

type Todo struct {
	Id    int    `uri:"id" json:"id"`
	Title string `json:"title"`
}
