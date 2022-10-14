package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	domain "framework-auto-aprov-go/pkg/domain"
	helper "framework-auto-aprov-go/pkg/helpers"
	"framework-auto-aprov-go/pkg/responses"
	services "framework-auto-aprov-go/pkg/services/auto_Aprov_Omnitok/interface"
)

type Auto_Aprov_OmnitokHandler struct {
	auto_Aprov_OmnitokUseCase services.Auto_Aprov_OmnitokUseCase
}

type ResponseAuto_Aprov_Omnitok struct {
	ID                      int32     `json:"pro_usuario_cliente" copier:"must"`
	Branch                  string    `json:"branch" copier:"must"`
	IngressIp               string    `json:"ingress_ip" copier:"must"`
	Hostname                string    `json:"hostname" copier:"must"`
	InitialAdminEmail       string    `json:"initial_admin_email" copier:"must"`
	InitialAdminpassword    string    `json:"initial_admin_password" copier:"must"`
	ExternalMongodbUrl      string    `json:"external_mongodb_url" copier:"must"`
	ExternalMongodbOplogUrl string    `json:"external_mongodb_oplog_url" copier:"must"`
	Apply                   bool      `json:"apply" copier:"must"`
	Destroy                 bool      `json:"destroy" copier:"must"`
	Deleted                 bool      `json:"deleted" copier:"must"`
	DateCreate              time.Time `json:"date_create" copier:"must"`
	DateDeleted             time.Time `json:"date_deleted" copier:"must"`
}

func NewAuto_Aprov_OmnitokHandler(usecase services.Auto_Aprov_OmnitokUseCase) *Auto_Aprov_OmnitokHandler {
	return &Auto_Aprov_OmnitokHandler{
		auto_Aprov_OmnitokUseCase: usecase,
	}
}

// FindAll godoc
// @summary Get all auto_Aprov_Omnitok
// @description Get all auto_Aprov_Omnitok
// @tags auto_Aprov_Omnitok
// @security ApiKeyAuth
// @id FindAll
// @produce json
// @Router /api/auto_Aprov_Omnitok [get]
// @response 200 {object} []Response "OK"
func (cr *Auto_Aprov_OmnitokHandler) FindAll(c *gin.Context) {
	auto_Aprov_Omnitok, err := cr.auto_Aprov_OmnitokUseCase.FindAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	} else {
		response := []ResponseAuto_Aprov_Omnitok{}
		copier.Copy(&response, &auto_Aprov_Omnitok)

		responses.SuccessJSON(c, 200, response)
	}
}

func (cr *Auto_Aprov_OmnitokHandler) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	auto_Aprov_Omnitok, err := cr.auto_Aprov_OmnitokUseCase.FindByID(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := ResponseAuto_Aprov_Omnitok{}
		copier.Copy(&response, &auto_Aprov_Omnitok)

		responses.SuccessJSON(c, 200, response)
	}
}

func (cr *Auto_Aprov_OmnitokHandler) FindByServicio(c *gin.Context) {
	paramsId := c.Param("idservicio")
	idservicio, _ := strconv.Atoi(paramsId)

	auto_Aprov_Omnitok, err := cr.auto_Aprov_OmnitokUseCase.FindByServicio(c.Request.Context(), uint(idservicio))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
	} else {
		response := []ResponseAuto_Aprov_Omnitok{}
		copier.Copy(&response, &auto_Aprov_Omnitok)

		responses.SuccessJSON(c, 200, response)
	}
}

func (cr *Auto_Aprov_OmnitokHandler) Save(c *gin.Context) {
	var auto_Aprov_OmnitokReq domain.Auto_Aprov_OmnitokReq
	var auto_Aprov_Omnitok domain.Auto_Aprov_Omnitok
	//var vars map[string]interface{}

	//if exists set values
	if err := c.BindJSON(&auto_Aprov_OmnitokReq); err != nil {
		responses.ErrorJSON(c, 400, err)
		c.Abort()
		return
	}

	fmt.Printf("vars: %s\n", auto_Aprov_OmnitokReq.Vars)

	//copiar sin vars
	copier.Copy(&auto_Aprov_Omnitok, &auto_Aprov_OmnitokReq)

	auto_Aprov_Omnitok, err := cr.auto_Aprov_OmnitokUseCase.Save(c.Request.Context(), auto_Aprov_Omnitok)

	if err != nil {
		responses.ErrorJSON(c, 200, err)
		c.Abort()
	} else {
		response := ResponseAuto_Aprov_Omnitok{}

		//convert map
		varsMap := auto_Aprov_OmnitokReq.Vars.(map[string]interface{})

		//create yaml

		erryaml := helper.GenerateYaml(c, "values.preproduction.omnitok.yaml", varsMap)
		if erryaml != nil {
			responses.ErrorJSON(c, 400, erryaml.Error())
			fmt.Println("error: " + erryaml.Error())
			c.Abort()
			return
		}
		copier.Copy(&response, &auto_Aprov_Omnitok)
		responses.SuccessJSON(c, 200, response)
	}

}

func (cr *Auto_Aprov_OmnitokHandler) SaveAll(c *gin.Context) {
	var auto_Aprov_Omnitok []domain.Auto_Aprov_Omnitok

	//if exists set values
	if err := c.BindJSON(&auto_Aprov_Omnitok); err != nil {
		responses.ErrorJSON(c, 400, err)
		c.Abort()
		return
	}

	_result, err := cr.auto_Aprov_OmnitokUseCase.SaveAll(c.Request.Context(), auto_Aprov_Omnitok)

	if err != nil {
		responses.ErrorJSON(c, 200, err)
		c.Abort()
	} else {
		// response := ResponseAuto_Aprov_Omnitok{}
		// copier.Copy(&response, &auto_Aprov_Omnitok)

		c.JSON(http.StatusOK, gin.H{"result": _result})
	}

}
func (cr *Auto_Aprov_OmnitokHandler) Update(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	oldauto_Aprov_Omnitok, err := cr.auto_Aprov_OmnitokUseCase.FindByID(c.Request.Context(), uint(id))
	var newauto_Aprov_Omnitok domain.Auto_Aprov_Omnitok
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		//if exists set values
		if err := c.BindJSON(&newauto_Aprov_Omnitok); err != nil {
			responses.ErrorJSON(c, 400, err)
			c.Abort()
			return
		}

		//Set values update
		oldauto_Aprov_Omnitok.Branch = newauto_Aprov_Omnitok.Branch
		oldauto_Aprov_Omnitok.IngressIp = newauto_Aprov_Omnitok.IngressIp
		oldauto_Aprov_Omnitok.Hostname = newauto_Aprov_Omnitok.Hostname
		oldauto_Aprov_Omnitok.InitialAdminEmail = newauto_Aprov_Omnitok.InitialAdminEmail
		oldauto_Aprov_Omnitok.InitialAdminpassword = newauto_Aprov_Omnitok.InitialAdminpassword
		oldauto_Aprov_Omnitok.ExternalMongodbUrl = newauto_Aprov_Omnitok.ExternalMongodbUrl
		oldauto_Aprov_Omnitok.ExternalMongodbOplogUrl = newauto_Aprov_Omnitok.ExternalMongodbOplogUrl
		oldauto_Aprov_Omnitok.Apply = newauto_Aprov_Omnitok.Apply
		oldauto_Aprov_Omnitok.Destroy = newauto_Aprov_Omnitok.Destroy

		// oldauto_Aprov_Omnitok.ProEliminado = newauto_Aprov_Omnitok.ProEliminado

		auto_Aprov_Omnitok, err := cr.auto_Aprov_OmnitokUseCase.Save(c.Request.Context(), oldauto_Aprov_Omnitok)

		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			response := ResponseAuto_Aprov_Omnitok{}
			copier.Copy(&response, &auto_Aprov_Omnitok)

			responses.SuccessJSON(c, 200, response)
		}
	}
}

func (cr *Auto_Aprov_OmnitokHandler) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	ctx := c.Request.Context()
	auto_Aprov_Omnitok, err := cr.auto_Aprov_OmnitokUseCase.FindByID(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if auto_Aprov_Omnitok == (domain.Auto_Aprov_Omnitok{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Auto_Aprov_Omnitok is not booking yet",
		})
		return
	}

	cr.auto_Aprov_OmnitokUseCase.Delete(ctx, auto_Aprov_Omnitok)

	responses.SuccessJSON(c, 200, "Auto_Aprov_Omnitok is deleted successfully")
}
