package model

// UserReceived user received by processor
type UserReceived struct {
	Prefecture string
	City       string
	First_Name string
	Last_Name  string
}

// UserTransformed user transformed by transformer
type UserTransformed struct {
	Prefecture string
	City       string
	Name       string
}
