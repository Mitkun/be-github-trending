package req

type ReqUpdateUser struct {
	FullName string `json:"full_name,omitempty" validate:"required"` // tags
	Email    string `json:"email,omitempty" validate:"required,email"`
}
