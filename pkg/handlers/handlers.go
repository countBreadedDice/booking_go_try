package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"github.com/countBreadedDice/booking_go_try/models"
	"github.com/countBreadedDice/booking_go_try/pkg/config"
	"github.com/countBreadedDice/booking_go_try/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct{ App *config.AppConfig }

// NewRepo creates the repository
func NewRepo(a *config.AppConfig) *Repository { return &Repository{App: a} }

// NewHandlers sets the repository for th repo
func NewHandlers(r *Repository) { Repo = r }

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic here
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

// Sum is the sum page handler
func (m *Repository) Sum(w http.ResponseWriter, r *http.Request) {
	sum := addValue(2, 2)
	fmt.Fprintf(w, fmt.Sprintf(" This is the about page. 2 +2 is = to %d", sum))
}
func (m *Repository) Divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32 = 100.0, 0.0
	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is equal to %f", x, y, f))
}

// addValue adds two integers
func addValue(x, y int) int {
	sum := x + y
	return sum
}

// Divides divdes two floats

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}
