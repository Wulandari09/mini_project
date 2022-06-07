package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all petugas
func GetPetugasscontrollers(c echo.Context) error {
	var petugas []models.petugas
	if err := config.DB.Find(&petugas).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all petugas", petugas))
}

// get petugas by id
func GetPetugascontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	petugas := models.Petugas{}
	if err := config.DB.First(&petugas, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "petugas not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get petugas", petugas))
}

// create petugas by id
func CreatePetugascontrollers(c echo.Context) error {
	petugas := models.Petugas{}
	c.Bind(&petugas)

	if err := config.DB.Debug().Create(&petugas).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new petugas", petugas))
}

// delete petugas by id
func DeletePetugascontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	petugas := models.Petugas{}
	if err := config.DB.Where("id = ?", id).First(&petugas).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "petugas not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Where("id = ?", id).Delete(&petugas).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("petugas deleted successfully", petugas))
}

// update petugas by id
func UpdatePetugascontrollers(c echo.Context) error {
	id := c.Param("id")
	petugas := models.Petugas{}

	if err := config.DB.Table("petugas").First(&petugas, "id = ?", id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "petugas not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newpetugas := models.Petugas{}
	c.Bind(&newpetugas)
	fmt.Println("petugas", v)
	petugas.Nama = newpetugas.Nama
	petugas.Gender = newpetugas.Gender
	petugas.Major = newpetugas.Major
	if err := config.DB.Table("petugas").Where("id = ?", id).Debug().Save(&petugas).Debug().Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update petugas", petugas))
}
