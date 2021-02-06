package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/pjchender/go-backend-template/internal/database"
	"github.com/pjchender/go-backend-template/internal/service"
	"github.com/pjchender/go-backend-template/pkg/app"
	"net/http"
)

type UserHandler struct {
	DB *database.GormDatabase
}

func NewUserHandler(db *database.GormDatabase) *UserHandler {
	return &UserHandler{DB: db}
}

func (d *UserHandler) Me(ctx *gin.Context) {
	var err error

	user, err := app.ParseUser(ctx)
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[user] app.ParseUser failed: ", err)
		return
	}

	ctx.JSON(http.StatusOK, user.ToExternal())
}

func (d *UserHandler) Get(ctx *gin.Context) {
	var err error
	userIDStr := ctx.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if success := app.SuccessOrAbort(ctx, http.StatusBadRequest, err); !success {
		log.Println("[user] uuid.Parse failed: ", err)
		return
	}

	svc := service.New(ctx, d.DB)
	user, err := svc.GetUserByID(service.GetUserRequest{ID: userID})
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[user] svc.GetUserByID failed: ", err)
		return
	}

	ctx.JSON(http.StatusOK, user.ToExternal())
}
