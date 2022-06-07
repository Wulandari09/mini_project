package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all pengembalian
func GetPengembaliancontrollers(c echo.Context) error {
	var pengembalian []models.pengembalian
	if err := config.DB.Find(&pengembalian).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all pengembalian", pengembalian))
}

// get pengembalian by id
func GetPengembaliancontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	pengembalian := models.Pengembalian{}
	if err := config.DB.First(&pengembalian, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "pengembalian not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get pengembalian", pengembalian))
}

// create pengembalian by id
func CreatePengembaliancontrollers(c echo.Context) error {
	pengembalian := models.Pengembalian{}
	c.Bind(&pengembalian)

	if err := config.DB.Debug().Create(&pengembalian).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new pengembalian", pengembalian))
}

// delete pengembalian by id
func DeletePinjamcontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	pengembalian := models.Pengembalian{}
	if err := config.DB.Where("id = ?", id).First(&pengembalian).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "pengembalian not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Where("id = ?", id).Delete(&pengembalian).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("pengembalian deleted successfully", pengembalian))
}

	newpengembalian := models.Pengembalian{}
	c.Bind(&newpengembalian)
	fmt.Println("pengembalian",pengembalian)
	pengembalian.Nama = newpengembalian.Nama
	pengembalian.Jenis = newpengembalian.Jenis
	pengembalian.Total = newpengembalian.Total
	if err := config.DB.Table("pengembalian").Where("id = ?", id).Debug().Save(&pengembalian).Debug().Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update pengembalian", pengembalian))
}
