// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateInstructorInput struct {
	Name               string  `json:"name"`
	SyllabicCharacters string  `json:"syllabicCharacters"`
	Biography          *string `json:"biography"`
	PhoneNumber        *string `json:"phoneNumber"`
	Email              *string `json:"email"`
}

type CreateRoomInput struct {
	Name     string `json:"name"`
	Capacity uint   `json:"capacity"`
	StudioID uint   `json:"studioID"`
}

type CreateSchoolInput struct {
	Name string `json:"name"`
}

type CreateStudioInput struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	SchoolID uint   `json:"schoolID"`
}

type CreateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateRoomInput struct {
	ID       uint    `json:"id"`
	Name     *string `json:"name"`
	Capacity uint    `json:"capacity"`
	StudioID *uint   `json:"studioID"`
}

type UpdateSchoolInput struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UpdateStudioInput struct {
	ID       uint    `json:"id"`
	Name     *string `json:"name"`
	Location string  `json:"location"`
	SchoolID *uint   `json:"schoolID"`
}
