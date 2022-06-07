package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// get all peminjaman
func GetPinjamcontrollers(c echo.Context) error {
	var pinjam []models.pinjam
	if err := config.DB.Find(&pinjam).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get all pinjam", pinjam))
}

// get pinjam by id
func GetPinjamcontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	pinjam := models.Pinjam{}
	if err := config.DB.First(&pinjam, id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "pinjam not found",
			})
		}

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success get pinjam", pinjam))
}

// create pinjam by id
func CreatePinjamcontrollers(c echo.Context) error {
	pinjam := models.Pinjam{}
	c.Bind(&pinjam)

	if err := config.DB.Debug().Create(&pinjam).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, helper.BuildResponse("success create new pinjam", pinjam))
}

// delete pinjam by id
func DeletePinjamcontrollers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	pinjam := models.Pinjam{}
	if err := config.DB.Where("id = ?", id).First(&pinjam).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "pinjam not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Where("id = ?", id).Delete(&pinjam).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("pinjam deleted successfully", pinjam))
}

// update pinjam by id
func UpdatePinjamcontrollers(c echo.Context) error {
	id := c.Param("id")
	pinjam := models.pinjam{}

	if err := config.DB.Table("pinjam").First(&pinjam, "id = ?", id).Error; err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "pinjam not found",
			})
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newpinjam := models.Pinjam{}
	c.Bind(&newpinjam)
	fmt.Println("pinjam",pinjam)
	pinjam.Nama = newpinjam.Nama
	pinjam.Jenis = newpinjam.Gender
	pinjam.Total = newpinjam.Major
	if err := config.DB.Table("pinjam").Where("id = ?", id).Debug().Save(&pinjam).Debug().Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.BuildResponse("success update pinjam", pinjam))
}
