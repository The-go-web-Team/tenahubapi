package main

import (
	// "fmt"
	"net/http"
	// "github.com/tenahubapi/entity"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tenahubapi/delivery/http/handler"

	hcserviceRepository "github.com/tenahubapi/service/repository"
	hcserviceService "github.com/tenahubapi/service/service"

	commentRepository "github.com/tenahubapi/comment/repository"
	commentService "github.com/tenahubapi/comment/service"

	ratingRepository "github.com/tenahubapi/rating/repository"
	ratingService "github.com/tenahubapi/rating/service"

	sesRepository "github.com/tenahubapi/session/repository"
	sesService "github.com/tenahubapi/session/service"

	// serviceRepo "github.com/tenahubapi/service/repository"
	// serviceServ "github.com/tenahubapi/service/service"
	//adminRepo "github.com/tenahubapi/admin/repository"
	//adminServ "github.com/tenahubapi/admin/service"
	//agentRepo "github.com/tenahubapi/agent/repository"
	//agentServ "github.com/tenahubapi/agent/service"

	healthCenterRepo "github.com/tenahubapi/healthcenter/repository"
	healthCenterServ "github.com/tenahubapi/healthcenter/service"
	feedBackRepo "github.com/tenahubapi/comment/repository"
	feedBackServ "github.com/tenahubapi/comment/service"
	"github.com/tenahubapi/user/repository"
	"github.com/tenahubapi/user/service"
	"github.com/julienschmidt/httprouter"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/tenahubapi/entity"
	"fmt"
)


func main()  {
	//dbconn, err := gorm.Open("postgres", "postgres://postgres:0912345678@localhost/tenahub?sslmode=disable")
	//dbconn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/tenahubdb?sslmode=disable")
	dbconn, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	errs := dbconn.CreateTable(&entity.Comment{}, &entity.Rating{},&entity.HealthCenter{}, &entity.Service{},&entity.User{}, &entity.Hcrating{},&entity.Session{}, &entity.UserComment{}).GetErrors()
	fmt.Println(errs)

	if len(errs) > 0 {
		panic(errs)
	}

	userRepo := repository.NewUserGormRepo(dbconn)
	userServ := service.NewUserService(userRepo)
	userHandl := handler.NewUserHander(userServ)

	comRepo := commentRepository.NewCommentGormRepo(dbconn)
	comServ := commentService.NewCommentService(comRepo)
	cmtHandl := handler.NewCommentHandler(comServ)

	ratingRepo := ratingRepository.NewGormRatingRepository(dbconn)
	ratingServ := ratingService.NewHcRatingService(ratingRepo)
	ratingHandl := handler.NewRatingHandler(ratingServ)

	sessionRepo := sesRepository.NewSessionGormRepo(dbconn)
	sessionService := sesService.NewSessionService(sessionRepo)
	sesHandl := handler.NewSessionHandler(sessionService)



	healthCenterRespository := healthCenterRepo.NewHealthCenterGormRepo(dbconn)
	healthCenterService := healthCenterServ.NewHealthCenterService(healthCenterRespository)
	healthCenterHandler := handler.NewHealthCenterHandler(healthCenterService)

	feedBackRepository := feedBackRepo.NewCommentGormRepo(dbconn)
	feedBackService := feedBackServ.NewCommentService(feedBackRepository)
	feedBackHandler := handler.NewCommentHandler(feedBackService)

	serviceRepo := hcserviceRepository.NewServiceGormRepo(dbconn)
	serviceServ := hcserviceService.NewServiceService(serviceRepo)
	serviceHandler := handler.NewServiceHandler(serviceServ)

	router := httprouter.New()

	router.GET("/v1/services/:id", serviceHandler.GetServices)
	router.GET("/v1/pending/services/:id", serviceHandler.GetPendingServices)
	router.PUT("/v1/service/:id", serviceHandler.PutService)
	router.POST("/v1/service", serviceHandler.PostService)
	router.OPTIONS("/v1/service", serviceHandler.PostService)
	router.DELETE("/v1/service/:id", serviceHandler.DeleteService)
	router.GET("/v1/service/:id", serviceHandler.GetSingleService)

	router.GET("/v1/feedback/:id", feedBackHandler.GetComments)
	router.PUT("/v1/feedback/:id", feedBackHandler.PutComment)
	router.POST("/v1/feedback", feedBackHandler.PostComment)
	router.OPTIONS("/v1/feedback", feedBackHandler.PostComment)
	router.DELETE("/v1/feedback/:id", feedBackHandler.DeleteComment)

	router.GET("/v1/users/:id", userHandl.GetSingleUser)
	router.GET("/v1/users/:id/type", userHandl.GetUsers)
	router.POST("/v1/user", userHandl.GetUser)
	router.PUT("/v1/users/:id", userHandl.PutUser)
	router.POST("/v1/users", userHandl.PostUser)
	router.DELETE("/v1/users/:id", userHandl.DeleteUser)


	router.GET("/v1/comments/:id", cmtHandl.GetComments)
	router.GET("/v1/comment/:id", cmtHandl.GetComment)
	router.PUT("/v1/comments/:id", cmtHandl.PutComment)
	router.DELETE("/v1/comments/:id", cmtHandl.DeleteComment)
	router.POST("/v1/comments", cmtHandl.PostComment)
	router.POST("/v1/comments/check", cmtHandl.Check)

	router.GET("/v1/healthcenter/:id", healthCenterHandler.GetSingleHealthCenter)
	router.POST("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	router.GET("/v1/healthcenter/:id/agent", healthCenterHandler.GetHealthCentersByAgentId)
	router.POST("/v1/healthcenter/addhealthcenter", healthCenterHandler.PostHealthCenter)
	router.PUT("/v1/healthcenter/:id", healthCenterHandler.PutHealthCenter)
	router.GET("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	router.GET("/v1/healthcenters", healthCenterHandler.GetHealthCenters)
	router.DELETE("/v1/healthcenter/:id", healthCenterHandler.DeleteHealthCenter)
	router.GET("/v1/healthcenters/search", healthCenterHandler.SearchHealthcenters)
	router.GET("/v1/healthcenters/top/:amount", healthCenterHandler.GetTop)

	router.GET("/v1/rating/:id", ratingHandl.GetRating)
	router.POST("/v1/rating", ratingHandl.PostRating)

	router.GET("/v1/session", sesHandl.GetSession)
	router.POST("/v1/session", sesHandl.PostSession)
	router.DELETE("/v1/session/:uuid", sesHandl.DeleteSession)

	err = http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		panic(err)
	}

	//http.ListenAndServe(":8181", router)
}

