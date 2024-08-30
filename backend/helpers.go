package backend

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// readJSON read json from request body into data. It accepts a sinle JSON of 1MB max size value in the body
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 //maximum allowable bytes is 1MB

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}

	return nil
}

// writeJSON writes arbitray data out as json
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	//add the headers if exists
	if len(headers) > 0 {
		for i, v := range headers[0] {
			w.Header()[i] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
	return nil
}

// badRequest sends a JSON response with the status http.StatusBadRequest, describing the error
func (app *application) badRequest(w http.ResponseWriter, err error) {
	var payload struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	payload.Error = true
	payload.Message = err.Error()
	_ = app.writeJSON(w, http.StatusOK, payload)
}

// // invalidCradentials sends a JSON response for invalid credentials
// func (app *application) invalidCradentials(w http.ResponseWriter) error {
// 	var payload struct {
// 		Error   bool   `json:"error"`
// 		Message string `json:"message"`
// 	}

// 	payload.Error = true
// 	payload.Message = "Invalid authentication credentials"
// 	err := app.writeJSON(w, http.StatusOK, payload)
// 	return err
// }

// func (app *application) passwordMatchers(hashPassword, password string) (bool, error) {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
// 			return false, nil
// 		default:
// 			return false, err
// 		}
// 	}
// 	return true, nil
// }

// // MatchMobileNumberPattern checks if the given number matches the provided regex pattern
// func (app *application) MatchMobileNumberPattern(input, pattern string) bool {
// 	matched, err := regexp.MatchString(pattern, input)
// 	if err != nil {
// 		// Handle error if the regex is invalid
// 		println("Error matching regex:", err)
// 		return false
// 	}
// 	return matched
// }
