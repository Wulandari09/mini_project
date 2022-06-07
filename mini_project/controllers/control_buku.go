package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all buku
func GetBukuscontrollers(c echo.Context) error {
	var buku []models.buku
	if err := config.DB.Find(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all buku", buku))
}

// get buku by id
func GetBukucontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	buku := models.Buku{}
	if err := config.DB.First(&buku, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "buku not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get buku", buku))
}

// create buku by id
func CreateBukucontrollers(c echo.Context) error {
	buku := models.Buku{}
	c.Bind(&buku)

	if err := config.DB.Debug().Create(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new buku", buku))
}

// delete buku by id
func DeleteBukucontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	buku := models.Buku{}
	if err := config.DB.Where("id = ?", id).First(&buku).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "buku not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Where("id = ?", id).Delete(&buku).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("buku deleted successfully", buku))
}

// update anggota by id
func UpdateBukucontrollers(c echo.Context) error {
	id := c.Param("id")
	buku := models.Buku{}

	if err := config.DB.Table("buku").First(&buku, "id = ?", id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "anggota not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newbuku := models.Book{}
	c.Bind(&newbuku)
	fmt.Println("buku",buku)
	buku.Nama = newbuku.Nama
	buku.Jenis = newbuku.Gender
	buku.Stock = newbuku.Major
	if err := config.DB.Table("buku").Where("id = ?", id).Debug().Save(&buku).Debug().Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update buku", buku))
}
