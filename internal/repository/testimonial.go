package repository

import (
	"database/sql"
	"testimonial-management/internal/entities"
)

type TestimonialRepository struct {
	DB *sql.DB
}

func NewTestimonialRepository(db *sql.DB) *TestimonialRepository {
	return &TestimonialRepository{DB: db}
}

func (r *TestimonialRepository) Create(t *entities.Testimonial) error {
	_, err := r.DB.Exec(
		`INSERT INTO testimonials (full_name, email, role, company, testimonial, photo_url, is_public) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		t.FullName, t.Email, t.Role, t.Company, t.Testimonial, t.PhotoURL, t.IsPublic,
	)
	return err
}

func (r *TestimonialRepository) GetAll() ([]*entities.Testimonial, error) {
	rows, err := r.DB.Query(`SELECT id, full_name, email, role, company, testimonial, photo_url, is_public, created_at FROM testimonials`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var testimonials []*entities.Testimonial
	for rows.Next() {
		t := &entities.Testimonial{}
		if err := rows.Scan(&t.ID, &t.FullName, &t.Email, &t.Role, &t.Company, &t.Testimonial, &t.PhotoURL, &t.IsPublic, &t.CreatedAt); err != nil {
			return nil, err
		}
		testimonials = append(testimonials, t)
	}
	return testimonials, nil
}
