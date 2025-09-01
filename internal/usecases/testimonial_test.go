package usecases

import (
	"testimonial-management/internal/dto"
	"testimonial-management/internal/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateTestimonialUsecase(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewTestimonialRepository(db)
	usecase := NewTestimonialUsecase(repo)

	req := &dto.CreateTestimonialRequest{
		FullName:    "Jane Doe",
		Email:       "jane@example.com",
		Role:        "Designer",
		Company:     "Beta Corp",
		Testimonial: "Awesome experience!",
		PhotoURL:    "http://example.com/photo2.jpg",
		IsPublic:    false,
	}

	mock.ExpectExec("INSERT INTO testimonials").
		WithArgs(req.FullName, req.Email, req.Role, req.Company, req.Testimonial, req.PhotoURL, req.IsPublic).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = usecase.CreateTestimonial(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetAllTestimonialsUsecase(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	defer db.Close()

	repo := repository.NewTestimonialRepository(db)
	usecase := NewTestimonialUsecase(repo)

	rows := sqlmock.NewRows([]string{"id", "full_name", "email", "role", "company", "testimonial", "photo_url", "is_public", "created_at"}).
		AddRow(2, "Jane Doe", "jane@example.com", "Designer", "Beta Corp", "Awesome experience!", "http://example.com/photo2.jpg", false, "2025-09-01T00:00:00Z")

	mock.ExpectQuery("SELECT id, full_name, email, role, company, testimonial, photo_url, is_public, created_at FROM testimonials").
		WillReturnRows(rows)

	resp, err := usecase.GetAllTestimonials()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(resp) != 1 {
		t.Errorf("expected 1 testimonial, got %d", len(resp))
	}
	if resp[0].FullName != "Jane Doe" {
		t.Errorf("expected FullName 'Jane Doe', got '%s'", resp[0].FullName)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
