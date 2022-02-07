package models

type Todo struct {
	Id        int
	Item      string
	Completed int
}

type Web_server interface{
	
	tood *TODO `json:"id"`
}

//whatever changes you made here it will reflect into our local repo after hiting git pull..............
