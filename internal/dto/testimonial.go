package dto

type CreateTestimonialRequest struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Company     string `json:"company"`
	Testimonial string `json:"testimonial"`
	PhotoURL    string `json:"photo_url"`
	IsPublic    bool   `json:"is_public"`
}

type TestimonialResponse struct {
	ID          int    `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Company     string `json:"company"`
	Testimonial string `json:"testimonial"`
	PhotoURL    string `json:"photo_url"`
	IsPublic    bool   `json:"is_public"`
	CreatedAt   string `json:"created_at"`
}
