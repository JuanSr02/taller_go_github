package api

import (
	"errors"
	"net/http"
	"parte3/internal/user"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	lettersOnlyRegex = regexp.MustCompile(`^[a-zA-Z]+$`)
)

// handler holds the user service and implements HTTP handlers for user CRUD.
type handler struct {
	userService *user.Service
}

// validateUserInput validates the user input for create and update operations
func validateUserInput(name, address, nickname string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(address) == "" {
		return errors.New("address is required")
	}
	if !lettersOnlyRegex.MatchString(name) {
		return errors.New("name must contain only letters")
	}
	if nickname != "" && !lettersOnlyRegex.MatchString(nickname) {
		return errors.New("nickname must contain only letters")
	}
	return nil
}

// handleCreate handles POST /users
func (h *handler) handleCreate(ctx *gin.Context) {
	// request payload
	var req struct {
		Name     string `json:"name"`
		Address  string `json:"address"`
		NickName string `json:"nickname"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	if err := validateUserInput(req.Name, req.Address, req.NickName); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := &user.User{
		Name:     req.Name,
		Address:  req.Address,
		NickName: req.NickName,
	}
	if err := h.userService.Create(u); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, u)
}

// handleRead handles GET /users/:id
func (h *handler) handleRead(ctx *gin.Context) {
	id := ctx.Param("id")

	u, err := h.userService.Get(id)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if user is active
	if u.Status != 1 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// handleUpdate handles PUT /users/:id
func (h *handler) handleUpdate(ctx *gin.Context) {
	id := ctx.Param("id")

	// bind partial update fields
	var fields *user.UpdateFields
	if err := ctx.ShouldBindJSON(&fields); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if at least one field is being updated
	if fields.Name == nil && fields.Address == nil && fields.NickName == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	// Validate fields if they are being updated
	if fields.Name != nil || fields.NickName != nil {
		name := ""
		if fields.Name != nil {
			name = *fields.Name
		}
		nickname := ""
		if fields.NickName != nil {
			nickname = *fields.NickName
		}
		// Address is not required for update
		if err := validateUserInput(name, "dummy address", nickname); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	u, err := h.userService.Update(id, fields)
	if err != nil {
		if errors.Is(err, user.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// handleDelete handles DELETE /users/:id
func (h *handler) handleDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.userService.Delete(id); err != nil {
		if errors.Is(err, user.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}