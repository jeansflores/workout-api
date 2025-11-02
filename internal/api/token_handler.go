package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jeansflores/workout-api/internal/store"
	"github.com/jeansflores/workout-api/internal/tokens"
	"github.com/jeansflores/workout-api/internal/utils"
)

type TokenHandler struct {
	tokenStore store.TokenStore
	userStore  store.UserStore
	logger     *log.Logger
}

type CreateTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewTokenHandler(tokenStore store.TokenStore, userStore store.UserStore, logger *log.Logger) *TokenHandler {
	return &TokenHandler{
		tokenStore: tokenStore,
		userStore:  userStore,
		logger:     logger,
	}
}

func (h *TokenHandler) HandleCreateToken(w http.ResponseWriter, r *http.Request) {
	var req CreateTokenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Printf("ERROR: CreateTokenRequest: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid request payload"})
		return
	}

	user, err := h.userStore.GetUserByUserName(req.Username)
	if err != nil || user == nil {
		h.logger.Printf("ERROR: GetUserByUserName: %v", err)
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid credentials"})
		return
	}

	passwordsDoMatch, err := user.PasswordHash.Matches(req.Password)
	if err != nil {
		h.logger.Printf("ERROR: Passwords do not match: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal server error"})
		return
	}

	if !passwordsDoMatch {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid credentials"})
		return
	}

	token, err := h.tokenStore.CreateNewToken(user.ID, 24*time.Hour, tokens.ScopeAuth)
	if err != nil {
		h.logger.Printf("ERROR: CreateNewToken: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"auth_token": token})
}
