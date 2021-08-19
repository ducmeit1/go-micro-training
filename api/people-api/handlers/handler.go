package handlers

import (
	"gin-training/api/people-api/requests"
	"gin-training/api/people-api/responses"
	"gin-training/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PeopleHandler interface {
	CreatePeople(c *gin.Context)
}

type peopleHandler struct {
	peopleClient pb.FPTPeopleClient
}

func NewPeopleHandler(peopleClient pb.FPTPeopleClient) PeopleHandler {
	return &peopleHandler{
		peopleClient: peopleClient,
	}
}

func (h *peopleHandler) CreatePeople(c *gin.Context) {
	req := requests.CreatePeopleRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.People{
		Slut:     req.Slut,
		Name:     req.Name,
		Age:      req.Age,
		Address:  req.Address,
		Contacts: []*pb.Contact{},
	}

	for _, v := range req.Contacts {
		pReq.Contacts = append(pReq.Contacts, &pb.Contact{
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email,
			Fax:         v.Fax,
		})
	}

	pRes, err := h.peopleClient.CreatePeople(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &responses.PeopleResponse{
		ID:       pRes.Id,
		Name:     pRes.Name,
		Address:  pRes.Address,
		Age:      pRes.Age,
		Slut:     pRes.Slut,
		Contacts: make([]*responses.PeopleContact, 0),
	}

	for _, v := range pRes.Contacts {
		dto.Contacts = append(dto.Contacts, &responses.PeopleContact{
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email,
			Fax:         v.Fax,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}
