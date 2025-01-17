package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	db "mwserver/db/sqlc"
	"mwserver/schemas"
	"mwserver/services"
	"mwserver/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MetricController struct {
	db  *db.Queries
	ctx context.Context
}

func NewMetricController(db *db.Queries, ctx context.Context) *MetricController {
	return &MetricController{db, ctx}
}

// Create Metric  handler
// @Summary Create a new metric
// @Description
// @Tags metric
// @ID create-metric
// @Accept  json
// @Produce  json
// @Param request body schemas.CreateMetricPayload true "query params"
// @Success 200 {object} schemas.MetricResponse
// @Router /metrics [post]
func (cc *MetricController) CreateMetric(ctx *gin.Context) {
	var payload *schemas.CreateMetricPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	now := time.Now()

	parsedTime, err := time.Parse(util.DEFAULT_STRING_LAYOUT, payload.DoneDate)
	args := &db.CreateMetricParams{
		Description:      payload.Description,
		IsDone:           payload.IsDone,
		DoneDate:         sql.NullTime{Time: parsedTime, Valid: err != nil},
		MetricEstimation: int32(payload.MetricEstimation),
		WayUuid:          uuid.MustParse(payload.WayUuid),
		UpdatedAt:        now,
	}

	metric, err := cc.db.CreateMetric(ctx, *args)
	util.HandleErrorGin(ctx, err)
	err = services.UpdateWayIsCompletedStatus(cc.db, ctx, metric.WayUuid)
	util.HandleErrorGin(ctx, err)

	ctx.JSON(http.StatusOK, metric)
}

// Update Metric handler
// @Summary Update metric by UUID
// @Description
// @Tags metric
// @ID update-metric
// @Accept  json
// @Produce  json
// @Param request body schemas.UpdateMetricPayload true "query params"
// @Param metricId path string true "metric UUID"
// @Success 200 {object} schemas.MetricResponse
// @Router /metrics/{metricId} [patch]
func (cc *MetricController) UpdateMetric(ctx *gin.Context) {
	var payload *schemas.UpdateMetricPayload
	metricId := ctx.Param("metricId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	now := time.Now()
	var parsedDoneDate time.Time
	var isDoneDateExistInPayload = payload.DoneDate != ""
	parsedDoneDate, err := time.Parse(util.DEFAULT_STRING_LAYOUT, payload.DoneDate)

	args := &db.UpdateMetricParams{
		Uuid:             uuid.MustParse(metricId),
		UpdatedAt:        sql.NullTime{Time: now, Valid: true},
		Description:      sql.NullString{String: payload.Description, Valid: payload.Description != ""},
		IsDone:           sql.NullBool{Bool: payload.IsDone, Valid: true},
		DoneDate:         sql.NullTime{Time: parsedDoneDate, Valid: isDoneDateExistInPayload && err == nil},
		MetricEstimation: sql.NullInt32{Int32: int32(payload.MetricEstimation), Valid: payload.MetricEstimation != 0},
	}

	metric, err := cc.db.UpdateMetric(ctx, *args)
	util.HandleErrorGin(ctx, err)
	err = services.UpdateWayIsCompletedStatus(cc.db, ctx, metric.WayUuid)
	util.HandleErrorGin(ctx, err)

	response := schemas.MetricResponse{
		Uuid:             metric.Uuid.String(),
		CreatedAt:        metric.CreatedAt.Format(util.DEFAULT_STRING_LAYOUT),
		UpdatedAt:        metric.UpdatedAt.Format(util.DEFAULT_STRING_LAYOUT),
		Description:      metric.Description,
		IsDone:           metric.IsDone,
		DoneDate:         util.MarshalNullTime(metric.DoneDate),
		MetricEstimation: metric.MetricEstimation,
	}

	ctx.JSON(http.StatusOK, response)
}

// Deleting Metric handlers
// @Summary Delete metric by UUID
// @Description
// @Tags metric
// @ID delete-metric
// @Accept  json
// @Produce  json
// @Param metricId path string true "metric ID"
// @Success 200
// @Router /metrics/{metricId} [delete]
func (cc *MetricController) DeleteMetricById(ctx *gin.Context) {
	metricId := ctx.Param("metricId")

	removedMetric, err := cc.db.DeleteMetric(ctx, uuid.MustParse(metricId))
	util.HandleErrorGin(ctx, err)
	err = services.UpdateWayIsCompletedStatus(cc.db, ctx, removedMetric.WayUuid)
	util.HandleErrorGin(ctx, err)

	ctx.JSON(http.StatusNoContent, gin.H{"status": "successfully deleted"})

}
