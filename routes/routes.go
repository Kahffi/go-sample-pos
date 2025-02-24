package routes

import (
	"github.com/Kahffi/go-sample-pos/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, productController *controllers.ProductController, receiptController *controllers.ReceiptController,
	employeeController *controllers.EmployeeController, customerController *controllers.CustomerController,
	paymentController *controllers.PaymentController, orderController *controllers.OrderController) {
	//Product Group
	productGroup := app.Group("/products")
	productGroup.Post("/", productController.CreateProduct)
	productGroup.Get("/", productController.GetAllProducts)
	productGroup.Get("/:id", productController.GetProductByID)
	productGroup.Put("/:id", productController.UpdateProduct)
	productGroup.Delete("/:id", productController.DeleteProduct)

	//	Receipt Group
	receiptGroup := app.Group("/receipts")
	receiptGroup.Post("/", receiptController.CreateReceipt)
	receiptGroup.Get("/", receiptController.GetAllReceipts)
	receiptGroup.Get("/:id", receiptController.GetReceiptByID)
	receiptGroup.Put("/:id", receiptController.UpdateReceipt)
	receiptGroup.Delete("/:id", receiptController.DeleteReceipt)

	//	Employee Group
	employeeGroup := app.Group("/employees")
	employeeGroup.Post("/", employeeController.CreateEmployee)
	employeeGroup.Get("/", employeeController.GetAllEmployees)
	employeeGroup.Get("/:id", employeeController.GetEmployeeByID)
	employeeGroup.Put("/:id", employeeController.UpdateEmployee)
	employeeGroup.Delete("/:id", employeeController.DeleteEmployee)

	//	Customer Group
	customerGroup := app.Group("/customers")
	customerGroup.Post("/", customerController.CreateCustomer)
	customerGroup.Get("/", customerController.GetAllCustomers)
	customerGroup.Get("/:id", customerController.GetCustomerByID)
	customerGroup.Put("/:id", customerController.UpdateCustomer)
	customerGroup.Delete("/:id", customerController.DeleteCustomer)

	//	Payment Group
	paymentGroup := app.Group("/payments")
	paymentGroup.Post("/", paymentController.CreatePayment)
	paymentGroup.Get("/", paymentController.GetAllPayments)
	paymentGroup.Get("/:id", paymentController.GetPaymentByID)
	paymentGroup.Put("/:id", paymentController.UpdatePayment)
	paymentGroup.Delete("/:id", paymentController.DeletePayment)

	//	Order Group
	orderGroup := app.Group("/orders")
	orderGroup.Post("/", orderController.CreateOrder)
	orderGroup.Get("/", orderController.GetAllOrders)
	orderGroup.Get("/:id", orderController.GetOrderByID)
	orderGroup.Put("/:id", orderController.UpdateOrder)
	orderGroup.Delete("/:id", orderController.DeleteOrder)
}
