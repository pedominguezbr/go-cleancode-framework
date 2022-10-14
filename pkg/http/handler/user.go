package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	domain "framework-auto-aprov-go/pkg/domain"
	helper "framework-auto-aprov-go/pkg/helpers"
	"framework-auto-aprov-go/pkg/responses"
	servicesrol_Usuario "framework-auto-aprov-go/pkg/services/rol_usuario/interface"
	services "framework-auto-aprov-go/pkg/services/user/interface"
)

type UserHandler struct {
	userUseCase        services.UserUseCase
	rol_UsuarioUseCase servicesrol_Usuario.Rol_UsuarioUseCase
}

type Response struct {
	ID              uint         `json:"us_id" copier:"must"`
	UsCorreo        string       `json:"us_correo" copier:"must"`
	UsUsuario       string       `json:"us_usuario" copier:"must"`
	UsNombre        string       `json:"us_nombre" copier:"must"`
	UsApellido      string       `json:"us_apellido" copier:"must"`
	UsEsactivo      bool         `json:"us_esactivo" copier:"must"`
	UsEliminado     bool         `json:"us_eliminado" copier:"must"`
	UsFecharegistro time.Time    `json:"us_fecharegistro" copier:"must"`
	Roles           []domain.Rol `json:"us_roles" copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase, rol_usuariocase servicesrol_Usuario.Rol_UsuarioUseCase) *UserHandler {
	return &UserHandler{
		userUseCase:        usecase,
		rol_UsuarioUseCase: rol_usuariocase,
	}
}

// FindAll godoc
// @summary Get all users
// @description Get all users
// @tags users
// @security ApiKeyAuth
// @id FindAll
// @produce json
// @Router /api/users [get]
// @response 200 {object} []Response "OK"
func (cr *UserHandler) FindAll(c *gin.Context) {
	users, err := cr.userUseCase.FindAll(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []Response{}
		copier.Copy(&response, &users)

		responses.SuccessJSON(c, 200, response)
	}
}

func (cr *UserHandler) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	user, err := cr.userUseCase.FindByID(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		responses.SuccessJSON(c, 200, response)
	}
}

func (cr *UserHandler) FindByIDWithRole(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	user, err := cr.userUseCase.FindByIDWithRole(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		responses.SuccessJSON(c, 200, response)
	}
}

func (cr *UserHandler) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	ctx := c.Request.Context()
	user, err := cr.userUseCase.FindByIDWithRole(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	//Delete roles asociados primero
	for _, rol := range user.Roles {
		rol_usuarios, err := cr.rol_UsuarioUseCase.FindByIDUserIDRol(ctx, uint(id), rol.ID)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			cr.rol_UsuarioUseCase.DeleteArray(ctx, rol_usuarios)
		}
	}

	cr.userUseCase.Delete(ctx, user)

	c.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully"})
}

func (cr *UserHandler) DeleteRoles(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		responses.ErrorJSON(c, 500, "cannot parse id")
		c.Abort()
		return
	}

	//if exists set values
	var user domain.User
	if err := c.BindJSON(&user); err != nil {
		responses.ErrorJSON(c, 400, err)
		c.Abort()
		return
	}

	ctx := c.Request.Context()

	for _, rol := range user.Roles {
		rol_usuarios, err := cr.rol_UsuarioUseCase.FindByIDUserIDRol(ctx, uint(id), rol.ID)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			cr.rol_UsuarioUseCase.DeleteArray(ctx, rol_usuarios)
		}
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	// if user == (domain.Users{}) {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "User is not booking yet",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "User is deleted successfully"})
}

// SignUp godoc
// @summary SignUp users
// @description  SignUp users
// @tags auth
// @id SignUp
// @produce json
// @Router /api/v1/auth/signUp [post]
// @response 200 {object} []Response "OK"
func (cr *UserHandler) SignUp(c *gin.Context) {
	var user domain.User

	log.Println("Contenido del userUseCase: ", cr.userUseCase)
	log.Println("Contenido del Context: ", c.Request.Body)

	if err := c.BindJSON(&user); err != nil {
		responses.ErrorJSON(c, 400, err)
		c.Abort()
		return
	}

	password := helper.HashPassword(user.UsClave)
	user.UsClave = password

	user, err := cr.userUseCase.Save(c.Request.Context(), user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := Response{}
		copier.Copy(&response, &user)

		responses.SuccessJSON(c, 200, response)
	}
}

// login godoc
// @summary login users
// @description  login users
// @tags auth
// @id login
// @produce json
// @Router /api/v1/auth/login [post]
// @response 200 {object} []Response "OK"
func (cr *UserHandler) LoginHandler(c *gin.Context) {

	var credential domain.Credentials
	var TokenResult domain.Tokens
	var expiresInMilSec int64
	if err := c.BindJSON(&credential); err != nil {
		responses.ErrorJSON(c, 400, err)
		c.Abort()
		return
	}

	user, err := cr.userUseCase.FindByUserName(c.Request.Context(), credential.Username)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	passwordIsValid, msg := helper.VerifyPassword(credential.Password, user.UsClave)
	//defer cancel()
	if !passwordIsValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, expiresInMilSec, err := helper.GenerateAllTokens(user.UsUsuario, user.UsCorreo, user.UsNombre, user.UsApellido, "Admin", user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	TokenResult.Token.Access_token = token
	TokenResult.Token.Refresh_token = refreshToken
	TokenResult.Token.Expires_in = expiresInMilSec

	c.JSON(http.StatusOK, TokenResult)

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"token": gin.H{
	// 			"expires_in":    expiresInMilSec,
	// 			"access_token":  token,
	// 			"refresh_token": refreshToken,
	// 		}})
}

// refresh-token godoc
// @summary refresh-token users
// @description  refresh-token users
// @tags auth
// @id refresh-token
// @produce json
// @Router /api/v1/auth/refresh-token [post]
// @response 200 {object} []Response "OK"
func (cr *UserHandler) RefreshToken(c *gin.Context) {
	var tokenRefresh domain.Tokens
	var Claimsdata *helper.SignedDetails
	var TokenResult domain.Tokens

	if err := c.BindJSON(&tokenRefresh); err != nil {
		responses.ErrorJSON(c, 400, err)
		c.Abort()
		return
	}
	//log.Println("tokenRefresh.Token.Refresh_token:", tokenRefresh.Token.Refresh_token)
	Claimsdata, err := helper.ValidateToken(tokenRefresh.Token.Refresh_token)

	if err != "" {
		log.Println("error:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, error := cr.userUseCase.FindByID(c.Request.Context(), Claimsdata.ID)

	if error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	token, refreshToken, expiresInMilSec, error := helper.GenerateAllTokens(user.UsUsuario, user.UsCorreo, user.UsNombre, user.UsApellido, "Admin", user.ID)

	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": error.Error(),
		})
	}

	TokenResult.Token.Access_token = token
	TokenResult.Token.Refresh_token = refreshToken
	TokenResult.Token.Expires_in = expiresInMilSec

	c.JSON(http.StatusOK, TokenResult)
}
