package main

import (
	"log"
	"net/http"
	"os"
	"tokbel/config"
	"tokbel/docs"
	"tokbel/handler/category"
	"tokbel/handler/middleware"
	"tokbel/handler/product"
	"tokbel/handler/transaction"
	"tokbel/handler/user"
	"tokbel/repository"
	"tokbel/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Toko Belanja API
//	@version		1.0
//	@description	This API Documentation for Toko Belanja.

//	@BasePath	/api/v1

//	@securityDefinitions.apiKey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description		Bearer token authorization

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load env file")
	}

	docs.SwaggerInfo.Host = os.Getenv("HOST")

	repo := repository.New(config.GetDBConfig())
	userService := services.NewUserService(repo.User)
	categoryService := services.NewCategoryService(repo.Category)
	productService := services.NewProductService(repo.Product, repo.Category)
	trxService := services.NewTransactionService(repo.Locker, repo.Transaction, repo.Product, repo.User, repo.Category)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Toko Belanja is up and running, please check the docs at /swagger/index.html",
		})
	})

	v1 := r.Group("/api/v1")
	{
		userGrp := v1.Group("/users")
		{
			userGrp.POST("/register", user.Register(userService))
			userGrp.POST("/login", user.Login(userService))
			userGrp.Use(middleware.Authenticated()).PATCH("/topup", user.TopUp(userService))
		}

		catGrp := v1.Group("/categories")
		catGrp.Use(middleware.Authenticated(), middleware.AdminOnly(userService))
		{
			catGrp.POST("/", category.Create(categoryService))
			catGrp.GET("/", category.FindAll(categoryService))
			catGrp.PATCH("/:id", category.Update(categoryService))
			catGrp.DELETE("/:id", category.Delete(categoryService))
		}

		productGrp := v1.Group("/products")
		productGrp.Use(middleware.Authenticated(), middleware.AdminOnly(userService))
		{
			productGrp.POST("/", product.Create(productService))
			productGrp.GET("/", product.FindAll(productService))
			productGrp.PUT("/:id", product.Update(productService))
			productGrp.DELETE("/:id", product.Delete(productService))
		}

		trxGrp := v1.Group("/transactions")
		trxGrp.Use(middleware.Authenticated())
		{
			trxGrp.POST("/", transaction.Create(trxService))
			trxGrp.GET("/my-transactions", transaction.FindByUser(trxService))
			trxGrp.Use(middleware.AdminOnly(userService)).GET("/user-transactions", transaction.FindAll(trxService))
		}

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
