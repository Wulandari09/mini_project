package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all anggota
func GetAnggotascontrollers(c echo.Context) error {
	var anggota []models.angggota
	if err := config.DB.Find(&anggota).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all anggota", anggota))
}

// get anggota by id
func GetAnggotacontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	anggota := models.Anggota{}
	if err := config.DB.First(&anggota, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "anggota not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get anggota", anggota))
}

// create anggota by id
func CreateAnggotacontrollers(c echo.Context) error {
	anggota := models.Anggota{}
	c.Bind(&anggota)

	if err := config.DB.Debug().Create(&anggota).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new anggota", anggota))
}

// delete anggota by id
func DeleteAnggotacontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	anggota := models.Anggota{}
	if err := config.DB.Where("id = ?", id).First(&anggota).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "anggota not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Where("id = ?", id).Delete(&anggota).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("anggota deleted successfully", anggota))
}

// update anggota by id
func UpdateAnggotacontrollers(c echo.Context) error {
	id := c.Param("id")
	anggota := models.Anggota{}

	if err := config.DB.Table("anggota").First(&anggota, "id = ?", id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "anggota not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newanggota := models.Anggota{}
	c.Bind(&newanggota)
	fmt.Println("anggota", anggota)
	anggota.Nama = newanggota.Nama
	anggota.Gender = newanggota.Gender
	anggota.Major = newanggota.Major
	if err := config.DB.Table("anggota").Where("id = ?", id).Debug().Save(&anggota).Debug().Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update anggota", anggota))
}
