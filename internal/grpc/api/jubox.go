package api

import (
	"context"
	"fmt"
	"github.com/pjchender/go-backend-template/internal/database"
	juboxPB "github.com/pjchender/go-backend-template/internal/grpc/proto/jubox"
	"net/http"
	"strconv"
	"time"
)

type JuboxHandler struct {
	juboxPB.UnimplementedJuboxServer
	DB *database.GormDatabase
}

func NewJuboxHandler(db *database.GormDatabase) *JuboxHandler {
	return &JuboxHandler{
		DB: db,
	}
}

func (handler *JuboxHandler) BedEventUpdate(
	ctx context.Context,
	bedEvent *juboxPB.BedEvent,
) (*juboxPB.UpdateResp, error) {
	var err error

	occurredAt, err := time.Parse(time.RFC3339, bedEvent.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("incorrect format of occurredAt: %v (%w)", bedEvent.Timestamp, err)
	}

	fmt.Println("receive BedEventUpdate ", occurredAt)
	//
	//svc := service.New(ctx, handler.DB)
	//_, err = svc.CreateDetection(service.CreateDetectionRequest{
	//	OccurredAt:  occurredAt,
	//	Event:       bedEvent.EventDetail.Name,
	//	Message:     bedEvent.EventDetail.Message,
	//	BoxDeviceID: bedEvent.DeviceId,
	//})
	//if err != nil {
	//	return nil, fmt.Errorf("svc.CreateDetection failed: %w", err)
	//}

	return &juboxPB.UpdateResp{
		TransactionNo: bedEvent.TransactionNo,
		RespCode:      strconv.Itoa(http.StatusOK),
		RespMessage:   http.StatusText(http.StatusOK),
	}, nil
}
