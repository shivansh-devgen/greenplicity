package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"greenplicity/certificates"
)

// The main func
func main() {
	// create new echo instance
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// create new certificate
	e.POST("/certificates", createCertificate)

	// verify certificate
	e.GET("/certificates/:id/verify", verifyCertificate)

	// start server
	e.Logger.Fatal(e.Start(":1323"))
}

func createCertificate(c echo.Context) error {
	// get certificate data from request body
	producerId := c.FormValue("producerId")
	mwh := c.FormValue("mwh")
	validUntil := c.FormValue("validUntil")
	certificate := c.FormValue("certificate")

	// create new certificate
	newCertificate := certificates.NewEnergyCertificate(producerId, mwh, validUntil, certificate)

	// save certificate to storage
	// code to save the certificate to storage

	// return certificate ID
	return c.JSON(http.StatusCreated, map[string]string{"id": newCertificate.ID})
}

func verifyCertificate(c echo.Context) error {
	// get certificate ID from request path
	certificateId := c.Param("id")

	// get certificate from storage
	// code to get the certificate from storage

	// verify certificate
	if err := certificate.VerifyCertificate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// return certificate data
	return c.JSON(http.StatusOK, certificate)
}

