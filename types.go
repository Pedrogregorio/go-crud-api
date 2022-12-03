package main

type User struct {
	ID    int    `json:"user_id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type UserRequest struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type ValidateAttributes struct {
	Invalid bool
	Errors  []string
}

type BaseRespose struct {
	Status   int      `json:"status"`
	Response []User   `json:"data"`
	Message  string   `json:"message"`
	Errors   []string `json:"errors"`
}

type BaseResposeWithUsers struct {
	Status   int    `json:"status"`
	Response []User `json:"data"`
}

type BaseResposeWithUser struct {
	Status   int  `json:"status"`
	Response User `json:"data"`
}

type BaseResposeWithError struct {
	Status int      `json:"status"`
	Errors []string `json:"errors"`
}
