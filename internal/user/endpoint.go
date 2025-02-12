package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ismae1-alvarez/rest-api-go/pkg/meta"
)

type (
	Controller func(w http.ResponseWriter, r *http.Request)

	EndPoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Update Controller
		Delete Controller
	}

	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	UpdateReq struct {
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Email     *string `json:"email"`
		Phone     *string `json:"phone"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		Meta   *meta.Meta  `json:"meta,omitempty"`
	}
)

func MakeEndPoints(s Service) EndPoints {

	return EndPoints{
		Create: makeCreateEndPoint(s),
		Get:    makeGetEndPoint(s),
		GetAll: makeGetAllEndPoint(s),
		Update: makeUpdateEndPoint(s),
		Delete: makeDeleteEndPoint(s),
	}
}

func makeCreateEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "Invalid request format"})
			return
		}

		if req.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "First Name is required"})

			return
		}

		if req.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "Last Name is required"})
			return
		}

		user, err := s.Create(req.FirstName, req.LastName, req.Email, req.Phone)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: user})
	}
}

func makeGetAllEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		v := r.URL.Query()

		filters := Filters{
			FirstName: v.Get("first_name"),
			LastName:  v.Get("last_name"),
		}

		limit, _ := strconv.Atoi(v.Get("limit"))
		page, _ := strconv.Atoi(v.Get("page"))

		count, err := s.Count(filters)

		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}

		meta, err := meta.New(page, limit, count)

		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Response{Status: 500, Err: err.Error()})
			return
		}

		users, err := s.GetAll(filters, meta.Offset(), meta.Limit())

		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}

		fmt.Println(users)

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: users, Meta: meta})
	}
}
func makeGetEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		path := mux.Vars(r)

		id := path["id"]

		user, err := s.Get(id)

		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: "user doesn't exist"})
			return
		}

		json.NewEncoder(w).Encode(&Response{Status: 200, Data: user})
	}
}

func makeUpdateEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		var req UpdateReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "Invlid request format"})
			return
		}

		if req.FirstName != nil && *req.FirstName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "first name required"})
			return
		}

		if req.LastName != nil && *req.LastName == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(&Response{Status: 400, Err: "last name required"})

			return
		}

		path := mux.Vars(r)

		id := path["id"]

		if err := s.Update(id, req.FirstName, req.LastName, req.Email, req.Phone); err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: "user doesn't exist"})
			return
		}
		json.NewEncoder(w).Encode(&Response{Status: 200, Data: "success"})

	}
}

func makeDeleteEndPoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		path := mux.Vars(r)

		id := path["id"]

		if err := s.Delete(id); err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(&Response{Status: 404, Err: "user doesn't exist"})

			return
		}
		json.NewEncoder(w).Encode(&Response{Status: 200, Data: "success"})
	}
}
