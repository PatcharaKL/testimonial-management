package usecases

import (
	"testimonial-management/internal/dto"
	"testimonial-management/internal/entities"
	"testimonial-management/internal/repository"
)

type TestimonialUsecase struct {
	Repo *repository.TestimonialRepository
}

func NewTestimonialUsecase(repo *repository.TestimonialRepository) *TestimonialUsecase {
	return &TestimonialUsecase{Repo: repo}
}

func (u *TestimonialUsecase) CreateTestimonial(req *dto.CreateTestimonialRequest) error {
	t := &entities.Testimonial{
		FullName:    req.FullName,
		Email:       req.Email,
		Role:        req.Role,
		Company:     req.Company,
		Testimonial: req.Testimonial,
		PhotoURL:    req.PhotoURL,
		IsPublic:    req.IsPublic,
	}
	return u.Repo.Create(t)
}

func (u *TestimonialUsecase) GetAllTestimonials() ([]*dto.TestimonialResponse, error) {
	testimonials, err := u.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	var resp []*dto.TestimonialResponse
	for _, t := range testimonials {
		resp = append(resp, &dto.TestimonialResponse{
			ID:          t.ID,
			FullName:    t.FullName,
			Email:       t.Email,
			Role:        t.Role,
			Company:     t.Company,
			Testimonial: t.Testimonial,
			PhotoURL:    t.PhotoURL,
			IsPublic:    t.IsPublic,
			CreatedAt:   t.CreatedAt,
		})
	}
	return resp, nil
}
