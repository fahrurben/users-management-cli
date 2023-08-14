package internal

type User struct {
	Id       string   `json:"_id"`
	Index    int      `json:"index"`
	Guid     string   `json:"guid"`
	IsActive bool     `json:"isActive"`
	Balance  string   `json:"balance"`
	Tags     []string `json:"tags"`
	Friends  []Friend `json:"friends"`
}

type Friend struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
