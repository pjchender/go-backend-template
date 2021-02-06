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

type OrganizationHandler struct {
	DB *database.GormDatabase
}

func NewOrganizationHandler(db *database.GormDatabase) *OrganizationHandler {
	return &OrganizationHandler{
		DB: db,
	}
}

func (handler *OrganizationHandler) Get(ctx *gin.Context) {
	var err error
	orgIDStr := ctx.Param("id")
	orgID, err := uuid.Parse(orgIDStr)
	if success := app.SuccessOrAbort(ctx, http.StatusBadRequest, err); !success {
		log.Println("[v1/org] uuid.Parse failed: ", err)
		return
	}

	svc := service.New(ctx, handler.DB)
	org, err := svc.GetOrganizationByID(service.GetOrganizationRequest{ID: orgID})
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[v1/org] svc.GetOrganizationByID failed: ", err)
		return
	}

	ctx.JSON(http.StatusOK, org.ToExternal())
}

func (handler *OrganizationHandler) FirstOrCreateByNISOrgID(ctx *gin.Context) {
	param := service.FirstOrCreateOrganizationByNISOrgIDRequest{}

	svc := service.New(ctx, handler.DB)
	org, err := svc.FirstOrCreateOrganizationByNISOrgID(param)
	if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
		log.Println("[v1/org] svc.FirstOrCreateOrganizationByNISOrgID failed: ", err)
	}

	ctx.JSON(http.StatusOK, org.ToExternal())
}
