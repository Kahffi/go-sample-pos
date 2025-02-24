package main

import (
	"fmt"

	"github.com/Kahffi/go-sample-pos/controllers"
	"github.com/Kahffi/go-sample-pos/database"
	"github.com/Kahffi/go-sample-pos/repositories"
	"github.com/Kahffi/go-sample-pos/routes"
	"github.com/Kahffi/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect to database
	database.ConnectDatabase()

	// Dependency Injection
	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	orderRepo := repositories.NewOrderRepository(database.DB)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	receiptRepo := repositories.NewReceiptRepository(database.DB)
	receiptService := services.NewReceiptService(receiptRepo)
	receiptController := controllers.NewReceiptController(receiptService)

	employeeRepo := repositories.NewEmployeeRepository(database.DB)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeService)

	paymentRepo := repositories.NewPaymentRepository(database.DB)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentController := controllers.NewPaymentController(paymentService)

	customerRepo := repositories.NewCustomerRepository(database.DB)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	// Setup Routes
	routes.SetupRoutes(app, productController, receiptController, employeeController, customerController, paymentController, orderController)

	// Start server
	err := app.Listen(":8081")
	if err != nil {
		fmt.Println(err)
	}
}
