package http

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"gitlab.com/tirava/shop/internal/models"
)

const (
	appVersion      = "0.4.3"
	requestID       = "requestID"
	badJSON         = "Bad JSON"
	errJSONResponse = "Error JSON response"
	badUserID       = "Bad user ID"
	userNotFound    = "User not found"
	errorCreateUser = "Error create user"
	errorGetUser    = "Error get user"
	errorDeleteUser = "Error delete user"
	errorUpdateUser = "Error update user"
)

func (s Server) root(w http.ResponseWriter, _ *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		s.logger.Error().Msg(err.Error())
	}

	if _, err := w.Write([]byte("Welcome to our shop!\nMy pod hostname: " + host + "\n")); err != nil {
		s.logger.Error().Msg(err.Error())
	}
}

func (s Server) health(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, render.M{"status": "OK"})
}

func (s Server) version(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, render.M{"version": appVersion})
}

func (s Server) sendUser(w http.ResponseWriter, user *models.User, rid string) {
	if err := json.NewEncoder(w).Encode(user); err != nil {
		s.logger.Error().Str(requestID, rid).Msg(err.Error())
		s.error.send(w, http.StatusInternalServerError, errJSONResponse)
	}
}

func (s Server) sendUsers(w http.ResponseWriter, users []models.User, rid string) {
	if err := json.NewEncoder(w).Encode(users); err != nil {
		s.logger.Error().Str(requestID, rid).Msg(err.Error())
		s.error.send(w, http.StatusInternalServerError, errJSONResponse)
	}
}

func (s Server) getUserID(w http.ResponseWriter, r *http.Request, rid string) (int, error) {
	uid, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		s.logger.Error().Str(requestID, rid).Msg(err.Error())
		s.error.send(w, http.StatusBadRequest, badUserID)

		return 0, err
	}

	return uid, nil
}

func (s Server) newUser(w http.ResponseWriter, r *http.Request) {
	rid := middleware.GetReqID(r.Context())
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		s.logger.Error().Str(requestID, rid).Msg(err.Error())
		s.error.send(w, http.StatusBadRequest, badJSON)

		return
	}

	db := s.db.Create(user)
	if db.Error != nil {
		s.logger.Error().Str(requestID, rid).Msg(db.Error.Error())
		s.error.send(w, http.StatusInternalServerError, errorCreateUser)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	s.logger.Info().Str(requestID, rid).Msgf("created user with id: %d", user.ID)
	s.sendUser(w, user, rid)
}

func (s Server) getUsers(w http.ResponseWriter, r *http.Request) {
	rid := middleware.GetReqID(r.Context())
	users := make([]models.User, 0)

	db := s.db.Find(&users)
	if db.Error != nil {
		s.logger.Error().Str(requestID, rid).Msg(db.Error.Error())
		s.error.send(w, http.StatusInternalServerError, errorGetUser)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	s.logger.Info().Str(requestID, rid).Msg("got all users")
	s.sendUsers(w, users, rid)
}

func (s Server) getUser(w http.ResponseWriter, r *http.Request) {
	rid := middleware.GetReqID(r.Context())

	uid, err := s.getUserID(w, r, rid)
	if err != nil {
		return
	}

	user := &models.User{}

	db := s.db.First(&user, uid)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		s.logger.Error().Str(requestID, rid).Msg(db.Error.Error())
		s.error.send(w, http.StatusInternalServerError, errorGetUser)

		return
	}

	if user.ID == 0 {
		s.error.send(w, http.StatusNotFound, userNotFound)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	s.logger.Info().Str(requestID, rid).Msgf("got user with id: %d", user.ID)
	s.sendUser(w, user, rid)
}

func (s Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	rid := middleware.GetReqID(r.Context())

	uid, err := s.getUserID(w, r, rid)
	if err != nil {
		return
	}

	db := s.db.Delete(&models.User{}, uint(uid))
	if db.Error != nil {
		s.logger.Error().Str(requestID, rid).Msg(db.Error.Error())
		s.error.send(w, http.StatusInternalServerError, errorDeleteUser)

		return
	}

	if db.RowsAffected == 0 {
		s.error.send(w, http.StatusNotFound, userNotFound)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	s.logger.Info().Str(requestID, rid).Msgf("deleted user with id: %d", uid)
}

func (s Server) updateUser(w http.ResponseWriter, r *http.Request) {
	rid := middleware.GetReqID(r.Context())
	user := &models.User{}

	uid, err := s.getUserID(w, r, rid)
	if err != nil {
		return
	}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		s.logger.Error().Str(requestID, rid).Msg(err.Error())
		s.error.send(w, http.StatusBadRequest, badJSON)

		return
	}

	find := &models.User{}

	db := s.db.First(find, uid)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		s.logger.Error().Str(requestID, rid).Msg(db.Error.Error())
		s.error.send(w, http.StatusInternalServerError, errorGetUser)

		return
	}

	if db.Error == gorm.ErrRecordNotFound {
		s.error.send(w, http.StatusNotFound, userNotFound)

		return
	}

	user.ID = uint(uid)
	user.CreatedAt = find.CreatedAt

	db = s.db.Save(user)
	if db.Error != nil {
		s.logger.Error().Str(requestID, rid).Msg(db.Error.Error())
		s.error.send(w, http.StatusInternalServerError, errorUpdateUser)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	s.logger.Info().Str(requestID, rid).Msgf("updated user with id: %d", uid)
	s.sendUser(w, user, rid)
}
