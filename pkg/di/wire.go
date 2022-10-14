//go:build wireinject
// +build wireinject

package di

import (
	config "framework-auto-aprov-go/pkg/config"
	db "framework-auto-aprov-go/pkg/db"
	http "framework-auto-aprov-go/pkg/http"
	handler "framework-auto-aprov-go/pkg/http/handler"
	auto_Aprov_Omnitokrepository "framework-auto-aprov-go/pkg/services/auto_Aprov_Omnitok/repository"
	auto_Aprov_Omnitokusecase "framework-auto-aprov-go/pkg/services/auto_Aprov_Omnitok/usecase"
	rolrepository "framework-auto-aprov-go/pkg/services/rol/repository"
	rolusecase "framework-auto-aprov-go/pkg/services/rol/usecase"
	rol_usuariorepository "framework-auto-aprov-go/pkg/services/rol_usuario/repository"
	rol_usuariousecase "framework-auto-aprov-go/pkg/services/rol_usuario/usecase"
	userrepository "framework-auto-aprov-go/pkg/services/user/repository"
	userusecase "framework-auto-aprov-go/pkg/services/user/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		userrepository.NewUserRepository,
		userusecase.NewUserUseCase,
		handler.NewUserHandler,
		rolrepository.NewRolRepository,
		rolusecase.NewRolUseCase,
		handler.NewRolHandler,
		rol_usuariorepository.NewRol_UsuarioRepository,
		rol_usuariousecase.NewRol_UsuarioUseCase,
		auto_Aprov_Omnitokrepository.NewAuto_Aprov_OmnitokRepository,
		auto_Aprov_Omnitokusecase.NewAuto_Aprov_OmnitokUseCase,
		handler.NewAuto_Aprov_OmnitokHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
