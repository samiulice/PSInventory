package backend

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//Secure routes
	mux.Route("/api", func(mux chi.Router) {
		// mux.Use(app.AuthAdmin)

		//HR Managenment
		mux.Post("/hr/view-employee/{type}", app.GetEmployees)
		mux.Post("/hr/add-employee", app.AddEmployee)
		//MIS
		mux.Post("/mis/view-customer/{type}", app.GetCustomers)
		mux.Post("/mis/add-customer", app.AddCustomer)
		mux.Post("/mis/view-supplier/{type}", app.GetSuppliers)
		mux.Post("/mis/get-supplier-id-name", app.GetSuppliersIDAndName)
		mux.Post("/mis/add-supplier", app.AddSupplier)

		//Inventory
		mux.Post("/inventory/add-brand", app.AddBrand)
		mux.Post("/inventory/add-category", app.AddCategory)
		mux.Post("/inventory/add-product", app.AddProduct)
		mux.Post("/inventory/memo/get-product-list", app.FetchMemoProductItems)
		mux.Post("/inventory/get-memo-list", app.GetMemoListBySupplierID)
		mux.Post("/inventory/return-product-to-supplier", app.ReturnProductsToSupplier)
		mux.Post("/inventory/restock-product", app.RestockProduct)
		mux.Post("/inventory/purchase/getPageDetails", app.GetPurchasePageDetails)
		// mux.Post("/inventory/purchase-return/getPageDetails", app.GetPurchaseReturnPageDetails)
	})
	return mux
}
