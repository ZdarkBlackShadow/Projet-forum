package controllers

import (
	"html/template"
	"net/http"
	"projet-forum/services"

	"github.com/gorilla/mux"
)

// ImageControllers handles image-related HTTP requests.
type ImageControllers struct {
	service  *services.ImageServices
	template *template.Template
}

// InitImageControllers initializes a new ImageControllers instance.
//
// Parameters:
//   - service: pointer to the ImageServices that contains business logic for image handling.
//   - template: pointer to the template engine (not used directly in this controller).
//
// Returns:
//   - pointer to an initialized ImageControllers instance.
func InitImageControllers(service *services.ImageServices, template *template.Template) *ImageControllers {
	return &ImageControllers{service: service, template: template}
}

// ImageRouter sets up routing for image-related endpoints.
//
// Parameters:
//   - r: the Gorilla Mux router to register image routes with.
func (c *ImageControllers) ImageRouter(r *mux.Router) {
	r.HandleFunc("/image/{filename}", c.GetImageByFilename).Methods("GET")
}

// GetImageByFilename handles the retrieval and serving of an image by its filename.
//
// This endpoint requires a valid "token" cookie for authentication.
// The image data is fetched through the ImageServices layer, and returned with the appropriate content type.
//
// URL Parameters:
//   - filename: the name of the image file to retrieve.
func (c *ImageControllers) GetImageByFilename(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	// Check for authentication token
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Retrieve image data
	data, content, err := c.service.GetImageByName(cookie.Value, filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set content type and write image to response
	w.Header().Set("Content-Type", content)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}