package repository

import (
	"testimonial-management/internal/entities"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateTestimonial(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	repo := NewTestimonialRepository(db)
	testimonial := &entities.Testimonial{
		FullName:    "John Doe",
		Email:       "john@example.com",
		Role:        "Developer",
		Company:     "Acme Corp",
		Testimonial: "Great service!",
		PhotoURL:    "http://example.com/photo.jpg",
		IsPublic:    true,
	}

	mock.ExpectExec("INSERT INTO testimonials").
		WithArgs(testimonial.FullName, testimonial.Email, testimonial.Role, testimonial.Company, testimonial.Testimonial, testimonial.PhotoURL, testimonial.IsPublic).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(testimonial)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetAllTestimonials(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	repo := NewTestimonialRepository(db)

	rows := sqlmock.NewRows([]string{"id", "full_name", "email", "role", "company", "testimonial", "photo_url", "is_public", "created_at"}).
		AddRow(1, "John Doe", "john@example.com", "Developer", "Acme Corp", "Great service!", "http://example.com/photo.jpg", true, "2025-09-01T00:00:00Z")

	mock.ExpectQuery("SELECT id, full_name, email, role, company, testimonial, photo_url, is_public, created_at FROM testimonials").
		WillReturnRows(rows)

	testimonials, err := repo.GetAll()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(testimonials) != 1 {
		t.Errorf("expected 1 testimonial, got %d", len(testimonials))
	}
	if testimonials[0].FullName != "John Doe" {
		t.Errorf("expected FullName 'John Doe', got '%s'", testimonials[0].FullName)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
