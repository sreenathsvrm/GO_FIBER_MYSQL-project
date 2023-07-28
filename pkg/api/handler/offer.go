package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/mgo.v2/bson"

	services "github.com/sreenathsvrm/voix-me/pkg/usecase/interface"
)

type OfferCompanyHandler struct {
	offerUseCase services.OfferUseCase
}

func NewOfferCompanyHandler(usecase services.OfferUseCase) *OfferCompanyHandler {
	return &OfferCompanyHandler{
		offerUseCase: usecase,
	}
}

// FindAll godoc
// @summary Get all offer companies
// @description Get all companies
// @tags Offer
// @security ApiKeyAuth
// @Param data_format header string true "Data Format"
// @Param country_name path string  true "Country Name"
// @id FindAll
// @produce json
// @Router /api/offer-company/{country_name} [get]
// @response 200 {object} []domain.OfferCompany "OK"
func (cr *OfferCompanyHandler) FindAll(c *fiber.Ctx) error {
	// expecting data format on ( JSON || BSON )
	dataFormat := c.Get("data_format")

	if dataFormat != "JSON" && dataFormat != "BSON" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "invalid header type for data format",
			"message": "format should be in (JSON || BSON)",
		})
	}

	countryName := c.Params("country_name")

	countries, err := cr.offerUseCase.FindAll(c.Context(), countryName)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Set("Content-Type", "application/json")

	if dataFormat == "JSON" {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"data": countries,
		})
	}

	bsonData, err := bson.Marshal(countries)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to marshal data to BSON",
		})
	}

	c.Set("Content-Type", "application/bson")

	// Respond with the BSON data
	return c.Status(http.StatusOK).Send(bsonData)
}
