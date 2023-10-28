package controller

import (
	"app/config"
	"app/helper"
	"app/model/domain"
	"app/model/web"
	"app/utils/request"
	"app/utils/response"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create Schedule
func CreateScheduleController(c echo.Context) error {
	var scheduleRequest web.ScheduleRequest

	if err := c.Bind(&scheduleRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input"))
	}

	schedule := request.ConvertToScheduleRequest(scheduleRequest)

	if err := config.DB.Create(&schedule).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Create Schedule"))
	}

	response := response.ConvertToGetSchedule(schedule)

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success Created Schedule", response))
}

// Get All Schedules
func GetAllSchedulesController(c echo.Context) error {
	var schedules []domain.Schedule

	err := config.DB.Find(&schedules).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Schedules"))
	}

	if len(schedules) == 0 {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("Empty Schedules Data"))
	}

	response := response.ConvertToGetAllSchedules(schedules)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Schedules Data Successfully Retrieved", response))
}

// Get Schedule by ID
func GetScheduleController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var schedule domain.Schedule

	if err := config.DB.First(&schedule, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Schedule"))
	}

	response := response.ConvertToGetSchedule(&schedule)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Schedule Data Successfully Retrieved", response))
}

// Update Schedule by ID
func UpdateScheduleController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var updatedSchedule domain.Schedule

	if err := c.Bind(&updatedSchedule); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Input"))
	}

	var existingSchedule domain.Schedule
	result := config.DB.First(&existingSchedule, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Schedule"))
	}

	config.DB.Model(&existingSchedule).Updates(updatedSchedule)

	response := response.ConvertToGetSchedule(&existingSchedule)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Schedule Data Successfully Updated", response))
}

// Delete Schedule by ID
func DeleteScheduleController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid ID"))
	}

	var existingSchedule domain.Schedule
	result := config.DB.First(&existingSchedule, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Failed to Retrieve Schedule"))
	}

	config.DB.Delete(&existingSchedule)

	return c.JSON(http.StatusOK, helper.SuccessResponse("Schedule Data Successfully Deleted", nil))
}