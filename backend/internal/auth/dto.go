package auth

type UserResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Role      string  `json:"role"`
	ManagerID *string `json:"manager_id"`
}

type MeResponse struct {
	User UserResponse `json:"user"`
}
