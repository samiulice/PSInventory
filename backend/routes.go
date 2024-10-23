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
		mux.Post("/mis/get-supplier-id-name-list", app.GetActiveSuppliersIDAndName)
		mux.Post("/mis/get-customer-id-name-list", app.GetActiveCustomersIDAndName)
		mux.Post("/mis/add-supplier", app.AddSupplier)

		//Inventory
		mux.Post("/inventory/add-brand", app.AddBrand)
		mux.Post("/inventory/add-category", app.AddCategory)
		mux.Post("/inventory/add-product", app.AddProduct)
		//memo--products list
		mux.Post("/inventory/memo/get-purchase-product-list", app.FetchPurchaseMemoProductItems)
		mux.Post("/inventory/memo/get-sales-product-list", app.FetchSalesMemoProductItems)
		mux.Post("/inventory/products/get-list-by-id", app.FetchProductItemsbyProductID)
		mux.Post("/inventory/products/search-instock-products-by-serial", app.FetchInstockProductItembySerialNumber) //search in-stock items
		mux.Post("/inventory/products/search-sold-products-by-serial", app.FetchSoldProductItembySerialNumber)       //search sale/sold items
		mux.Post("/inventory/products/search-products-by-serial", app.FetchProductItembySerialNumber)                //search all type items
		//memo--supplier
		mux.Post("/inventory/get-supplier-memo-list", app.GetMemoListBySupplierID)
		//memo--supplier
		mux.Post("/inventory/get-customer-memo-list", app.GetMemoListByCustomerID)
		mux.Post("/inventory/return-product-to-supplier", app.ReturnProductsToSupplier)
		mux.Post("/inventory/restock-product", app.RestockProduct)
		mux.Post("/inventory/sale-products", app.SaleProducts)
		//page details
		mux.Post("/inventory/purchase/get-page-details", app.GetPurchasePageDetails)
		mux.Post("/inventory/sale/get-page-details", app.GetSalePageDetails)
		//warranty
		mux.Post("/inventory/products/warranty/get-history", app.GetClaimWarrantyList)
		mux.Post("/inventory/products/warranty/checkout'", app.CheckoutWarrantyProduct)
		mux.Post("/inventory/products/warranty/checkout/get-list", app.GetClaimWarrantyList)
		mux.Post("/inventory/products/claim-warranty-by-serial-id", app.ClaimWarrantyBySerialID)
		//accounts
		mux.Post("/accounts/receive-collection/get-page-details", app.GetReceiveCollectionPageDetails)

	})
	return mux
}
