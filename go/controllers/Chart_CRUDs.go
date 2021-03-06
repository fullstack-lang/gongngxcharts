// generated by stacks/gong/go/models/controller_file.go
package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/fullstack-lang/gongngxcharts/go/models"
	"github.com/fullstack-lang/gongngxcharts/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Chart__dummysDeclaration__ models.Chart
var __Chart_time__dummyDeclaration time.Duration

// An ChartID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getChart updateChart deleteChart
type ChartID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ChartInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postChart updateChart
type ChartInput struct {
	// The Chart to submit or modify
	// in: body
	Chart *orm.ChartAPI
}

// GetCharts
//
// swagger:route GET /charts charts getCharts
//
// Get all charts
//
// Responses:
//    default: genericError
//        200: chartDBsResponse
func GetCharts(c *gin.Context) {
	db := orm.BackRepo.BackRepoChart.GetDB()

	// source slice
	var chartDBs []orm.ChartDB
	query := db.Find(&chartDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	chartAPIs := make([]orm.ChartAPI, 0)

	// for each chart, update fields from the database nullable fields
	for idx := range chartDBs {
		chartDB := &chartDBs[idx]
		_ = chartDB
		var chartAPI orm.ChartAPI

		// insertion point for updating fields
		chartAPI.ID = chartDB.ID
		chartDB.CopyBasicFieldsToChart(&chartAPI.Chart)
		chartAPI.ChartPointersEnconding = chartDB.ChartPointersEnconding
		chartAPIs = append(chartAPIs, chartAPI)
	}

	c.JSON(http.StatusOK, chartAPIs)
}

// PostChart
//
// swagger:route POST /charts charts postChart
//
// Creates a chart
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: chartDBResponse
func PostChart(c *gin.Context) {
	db := orm.BackRepo.BackRepoChart.GetDB()

	// Validate input
	var input orm.ChartAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create chart
	chartDB := orm.ChartDB{}
	chartDB.ChartPointersEnconding = input.ChartPointersEnconding
	chartDB.CopyBasicFieldsFromChart(&input.Chart)

	query := db.Create(&chartDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, chartDB)
}

// GetChart
//
// swagger:route GET /charts/{ID} charts getChart
//
// Gets the details for a chart.
//
// Responses:
//    default: genericError
//        200: chartDBResponse
func GetChart(c *gin.Context) {
	db := orm.BackRepo.BackRepoChart.GetDB()

	// Get chartDB in DB
	var chartDB orm.ChartDB
	if err := db.First(&chartDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var chartAPI orm.ChartAPI
	chartAPI.ID = chartDB.ID
	chartAPI.ChartPointersEnconding = chartDB.ChartPointersEnconding
	chartDB.CopyBasicFieldsToChart(&chartAPI.Chart)

	c.JSON(http.StatusOK, chartAPI)
}

// UpdateChart
//
// swagger:route PATCH /charts/{ID} charts updateChart
//
// Update a chart
//
// Responses:
//    default: genericError
//        200: chartDBResponse
func UpdateChart(c *gin.Context) {
	db := orm.BackRepo.BackRepoChart.GetDB()

	// Get model if exist
	var chartDB orm.ChartDB

	// fetch the chart
	query := db.First(&chartDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.ChartAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	chartDB.CopyBasicFieldsFromChart(&input.Chart)
	chartDB.ChartPointersEnconding = input.ChartPointersEnconding

	query = db.Model(&chartDB).Updates(chartDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the chartDB
	c.JSON(http.StatusOK, chartDB)
}

// DeleteChart
//
// swagger:route DELETE /charts/{ID} charts deleteChart
//
// Delete a chart
//
// Responses:
//    default: genericError
func DeleteChart(c *gin.Context) {
	db := orm.BackRepo.BackRepoChart.GetDB()

	// Get model if exist
	var chartDB orm.ChartDB
	if err := db.First(&chartDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&chartDB)

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	orm.BackRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})
}
