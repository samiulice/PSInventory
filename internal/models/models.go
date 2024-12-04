package models

import (
	"time"
)

//MM-PR-randomAlphanumeriac(6)+CurrentIndexOfPurchaseHistory Table //Purchase
//MM-SL-randomAlphanumeriac(6)+CurrentIndexOfSalesHistory Table //Sale
//MM-WC-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Warranty Claimed
//MM-WCO-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Warranty Checked Out
//MM-WD-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Warranty Delivered
//MM-RR-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Repair Receive
//MM-RC-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Repair Completed
//MM-RD-randomAlphanumeriac(6)+CurrentIndexOfWarrantyHistory Table //Repair Delivered

const (
	// AustraliaRegex matches Australian mobile numbers with or without country code (+61)
	AustraliaRegex = `^(\+?61|0)4\d{8}$`

	// BangladeshRegex matches Bangladeshi mobile numbers with or without country code (+880)
	BangladeshRegex = `^\+?(880)?1[3-9]\d{8}$`

	// CanadaRegex matches Canadian phone numbers in various formats
	CanadaRegex = `^(\+?1)?[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}$`

	// FranceRegex matches French phone numbers with or without country code (+33)
	FranceRegex = `^(?:(?:\+|00)33|0)\s*[1-9](?:[\s.-]*\d{2}){4}$`

	// GermanyRegex matches German phone numbers with or without country code (+49)
	GermanyRegex = `^(\+?49|0)(\d{3,4})?[ -]?(\d{3,4})?[ -]?(\d{4,6})$`

	// IndiaRegex matches Indian mobile numbers with or without country code (+91)
	IndiaRegex = `^\+?(91)?\d{10}$`

	// JapanRegex matches Japanese phone numbers with or without country code (+81)
	JapanRegex = `^\+?81[-.\s]?\d{1,4}[-.\s]?\d{1,4}[-.\s]?\d{4}$`

	// PakistanRegex matches Pakistani mobile numbers with or without country code (+92)
	PakistanRegex = `^\+?(92)?\d{10}$`

	// SriLankaRegex matches Sri Lankan mobile numbers with or without country code (+94)
	SriLankaRegex = `^\+?(94)?\d{9}$`

	// UKRegex matches UK phone numbers including landline, mobile, and toll-free numbers
	UKRegex = `^(?:(?:\+44\s?|0)(?:\d{5}\s?\d{5}|\d{4}\s?\d{4}\s?\d{4}|\d{3}\s?\d{3}\s?\d{4}|\d{2}\s?\d{4}\s?\d{4}|\d{4}\s?\d{4}|\d{4}\s?\d{3})|\d{5}\s?\d{4}\s?\d{4}|0800\s?\d{3}\s?\d{4})$`

	// USRegex matches US phone numbers in various formats
	USRegex = `^\+?1?[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}$`
)

// Brand is the type for Brand
type Brand struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Category is the type for Category
type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Employee is the type for employees
type Employee struct {
	ID            int    `json:"id"`
	AccountCode   string `json:"account_code"`
	AccountName   string `json:"account_name"`
	ContactPerson string `json:"contact_person"`
	// Gender                string         `json:"gender"`
	// DateOfBith            time.Time      `json:"date_of_birth"`
	// Experties        string         `json:"exparties"`
	Division        string    `json:"division"`
	District        string    `json:"district"`
	Upazila         string    `json:"upazila"`
	Area            string    `json:"area"`
	Mobile          string    `json:"mobile"`
	DueAmount       int       `json:"due_amount"`
	WhatsappAccount string    `json:"whatsapp_account"`
	Email           string    `json:"email"`
	ImageLink       string    `json:"image_link"`     //username_profile_id_yy-mm-dd_hh-mm-ss.jpg
	AccountStatus   bool      `json:"account_status"` //Active = true, Inactive = false
	MonthlySalary   int       `json:"monthly_salary"`
	OpeningBalance  int       `json:"opening_balance"`
	CVLink          string    `json:"cv_link"`
	NIDLink         string    `json:"nid_link"`
	JoiningDate     time.Time `json:"joining_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
type StakeHolder struct {
	ID            int    `json:"id"`
	AccountType   string `json:"account_type"`
	AccountCode   string `json:"account_code"`
	AccountName   string `json:"account_name"`
	ContactPerson string `json:"contact_person"`
	// Gender                string         `json:"gender"`
	// DateOfBith            time.Time      `json:"date_of_birth"`
	// Experties        string         `json:"exparties"`
	Division        string    `json:"division"`
	District        string    `json:"district"`
	Upazila         string    `json:"upazila"`
	Area            string    `json:"area"`
	Mobile          string    `json:"mobile"`
	DueAmount       int       `json:"due_amount"`
	WhatsappAccount string    `json:"whatsapp_account"`
	Email           string    `json:"email"`
	ImageLink       string    `json:"image_link"`     //username_profile_id_yy-mm-dd_hh-mm-ss.jpg
	AccountStatus   bool      `json:"account_status"` //Active = true, Inactive = false
	MonthlySalary   int       `json:"monthly_salary"`
	OpeningBalance  int       `json:"opening_balance"`
	CVLink          string    `json:"cv_link"`
	NIDLink         string    `json:"nid_link"`
	JoiningDate     time.Time `json:"joining_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Customer is the type for customers
type Customer struct {
	ID              int       `json:"id"`
	AccountCode     string    `json:"account_code"`
	AccountName     string    `json:"account_name"`
	ContactPerson   string    `json:"contact_person"`
	Division        string    `json:"division"`
	District        string    `json:"district"`
	Upazila         string    `json:"upazila"`
	Area            string    `json:"area"`
	Mobile          string    `json:"mobile"`
	DueAmount       int       `json:"due_amount"`
	Email           string    `json:"email"`
	WhatsappAccount string    `json:"whatsapp_account"`
	AccountStatus   bool      `json:"account_status"` //Active = true, Inactive = false
	Discount        int       `json:"discount"`
	OpeningBalance  int       `json:"opening_balance"`
	JoiningDate     time.Time `json:"joining_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Supplier is the type for suppliers
type Supplier struct {
	ID              int       `json:"id"`
	AccountCode     string    `json:"account_code"`
	AccountName     string    `json:"account_name"`
	ContactPerson   string    `json:"contact_person"`
	Division        string    `json:"division"`
	District        string    `json:"district"`
	Upazila         string    `json:"upazila"`
	Area            string    `json:"area"`
	Mobile          string    `json:"mobile"`
	DueAmount       int       `json:"due_amount"`
	Email           string    `json:"email"`
	WhatsappAccount string    `json:"whatsapp_account"`
	AccountStatus   bool      `json:"account_status"` //Active = true, Inactive = false
	Discount        int       `json:"discount"`
	JoiningDate     time.Time `json:"joining_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// HeadAccount is the type for HeadAccount
type HeadAccount struct {
	ID             int       `json:"id"`
	AccountCode    string    `json:"account_code"`
	AccountName    string    `json:"account_name"`
	AccountType    string    `json:"account_type"`
	AccountStatus  bool      `json:"account_status"`
	CurrentBalance int       `json:"current_balance"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Product is the type for products
type Product struct {
	ID                int                `json:"id"`
	ProductCode       string             `json:"product_code"`
	ProductName       string             `json:"product_name"`
	Description       string             `json:"product_description"`
	ProductStatus     bool               `json:"product_status"`
	QuantityPurchased int                `json:"quantity_purchased"`
	PurchaseCost      int                `json:"purchase_cost"`
	QuantitySold      int                `json:"quantity_sold"`
	SoldPrice         int                `json:"sold_price"`
	QuantityInStock   int                `json:"quantity_in_stock"`
	CategoryID        int                `json:"category_id"`
	BrandID           int                `json:"brand_id"`
	StockAlertLevel   int                `json:"stock_alert_level"`
	CreatedAt         time.Time          `json:"created_at"`
	UpdatedAt         time.Time          `json:"updated_at"`
	Category          Category           `json:"category"`
	Brand             Brand              `json:"brand"`
	ProductMetadata   []*ProductMetadata `json:"product_metadata"`
}

// ProductShortInfo stores short information about product
type ProductShortInfo struct {
	ProductID     int      `json:"product_id"`
	ProductName   string   `json:"product_name"`
	CategoryName  string   `json:"category_name"`
	BrandName     string   `json:"brand_name"`
	SerialNumbers []string `json:"serial_numbers"`
	Warranty      int      `json:"warranty"`
	Quantity      int      `json:"quantity"`
	MRP           int      `json:"mrp"`
	SubTotal      int      `json:"sub_total"`
}

// SalesInvoice stores sales data for invoice generation
type SalesInvoice struct {
	CustomerInfo    Customer            `json:"customer_info"`
	HeadAccountInfo HeadAccount         `json:"head_account_info"`
	ProductItems    []*ProductShortInfo `json:"selected_items"`
	SaleDate        string              `json:"sale_date"`
	ChalanNo        string              `json:"chalan_no"`
	MemoNo          string              `json:"memo_no"`
	Note            string              `json:"note"`
	BillAmount      int                 `json:"bill_amount"`
	Discount        int                 `json:"discount"`
	TotalAmount     int                 `json:"total_amount"`
	PaidAmount      int                 `json:"paid_amount"`
	GrossProfit     int                 `json:"gross_profit"`
}

// ProductMetadata holds products meta data
type ProductMetadata struct {
	ID                     int       `json:"id"`
	SerialNumber           string    `json:"serial_number"`
	ProductID              int       `json:"product_id"`
	PurchaseHistoryID      int       `json:"purchase_history_id"`
	SalesHistoryID         int       `json:"sales_history_id"`
	Status                 string    `json:"status,"`
	WarrantyPeriod         int       `json:"warranty"`
	WarrantyStatus         string    `json:"warranty_status"`
	LatesWarrantyHistoryID int       `json:"latest_warranty_history_id"`
	WarrantyHistoryIDs     string    `json:"warranty_history_ids"`
	MaxRetailPrice         int       `json:"max_retail_price"` //max_retail_price = total_amount+profit/quantity_purchased
	PurchaseRate           int       `json:"purchase_rate"`    //purchase_rate = total_amount/quantity_purchased
	SaleRate               int       `json:"sale_rate"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type PurchasedProduct struct {
	Product        Product  `json:"product_info"`
	ProductSerialNo  []string `json:"products_serial_no"`
	Quantity       int      `json:"quantity"`
	MaxRetailPrice int      `json:"max_retail_price"`
	PurchaseRate   int      `json:"purchase_rate"`
	ShippingCost   int      `json:"shipping_cost"`
	WarrantyPeriod int      `json:"warranty"`
}

type PurchasePayload struct{
	ID                int         `json:"id"`
	PurchaseDate      string      `json:"purchase_date"`
	Supplier          Supplier    `json:"supplier_info"`
	HeadAccount       HeadAccount `json:"head_account_info"`
	ChalanNO          string      `json:"chalan_no"`
	MemoNo            string      `json:"memo_no"` //MM-P-randomAlphanumeriac(6)CurrentIndexOfPurchaseHistoryDB
	Note              string      `json:"note"`
	BillAmount        int         `json:"bill_amount"`
	Discount          int         `json:"discount"`
	TotalAmount       int         `json:"total_amount"`
	PaidAmount        int         `json:"paid_amount"`
	PurchasedProduct PurchasedProduct `json:"purchased_product"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

// Purchase is the type for purchase info
type Purchase struct {
	ID                int         `json:"id"`
	PurchaseDate      string      `json:"purchase_date"`
	Supplier          Supplier    `json:"supplier_info"`
	Product           Product     `json:"product_info"`
	Quantity          int         `json:"quantity"`
	QuantityPurchased int         `json:"quantity_purchased"`
	QuantitySold      int         `json:"quantity_sold"`
	ProductsSerialNo  []string    `json:"products_serial_no"`
	HeadAccount       HeadAccount `json:"head_account_info"`
	ChalanNO          string      `json:"chalan_no"`
	MemoNo            string      `json:"memo_no"` //MM-P-randomAlphanumeriac(6)CurrentIndexOfPurchaseHistoryDB
	Note              string      `json:"note"`
	MaxRetailPrice    int         `json:"max_retail_price"`
	PurchaseRate      int         `json:"purchase_rate"`
	WarrantyPeriod    int         `json:"warranty"`
	ShippingCost      int         `json:"shipping_cost"`
	BillAmount        int         `json:"bill_amount"`
	Discount          int         `json:"discount"`
	TotalAmount       int         `json:"total_amount"`
	PaidAmount        int         `json:"paid_amount"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

type SelectedItems struct {
	ProductID     int      `json:"product_id"`
	SerialNumbers []string `json:"serialNumbers"`
}

// Sale is the type for purchase info
type Sale struct {
	ID               int              `json:"id"`
	SaleDate         string           `json:"sale_date"`
	CustomerID       int              `json:"customer_id"`
	ProductsSerialNo []string         `json:"products_serial_no"`
	SelectedItems    []*SelectedItems `json:"selected_items"`
	AccountID        int              `json:"account_id"`
	ChalanNO         string           `json:"chalan_no"`
	MemoNo           string           `json:"memo_no"`
	Note             string           `json:"note"`
	BillAmount       int              `json:"bill_amount"`
	Discount         int              `json:"discount"`
	TotalAmount      int              `json:"total_amount"`
	PaidAmount       int              `json:"paid_amount"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
	Customer         Customer         `json:"customer_info"`
}

type Warranty struct {
	ID               int       `json:"id"`
	Status           string    `json:"status"`
	ProductSerialID  int       `json:"product_serial_id"`
	PreviousSerialNo string    `json:"previous_serial_no"`
	NewSerialNo      string    `json:"new_serial_no"`
	MemoNo           string    `json:"memo_no"`
	ContactNumber    string    `json:"contact_number"`
	RequestedDate    string    `json:"requested_date"`
	ReportedProblem  string    `json:"reported_problem"`
	ReceivedBy       string    `json:"received_by"`
	CheckoutDate     string    `json:"checkout_date"`
	DeliveryDate     string    `json:"delivery_date"`
	DeliveredBy      string    `json:"delivered_by"`
	Comment          string    `json:"comment"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ProductInfo      Product   `json:"product_info"`
	SalesInfo        Sale      `json:"sales_info"`
	CustomerInfo     Customer  `json:"customer_info"`
	SupplierInfo     Supplier  `json:"supplier_info"`
}

// Service is the type for services
type Service struct {
	//service_code,service_name,service_description,min_charge,discount,status,created_at,updated_at)
	ID            int       `json:"id"`
	ServiceCode   string    `json:"service_code"`
	ServiceName   string    `json:"service_name"`
	Description   string    `json:"service_description"`
	ServiceStatus string    `json:"service_status"`
	BaseFee       int       `json:"base_fee"`
	TrackRecord   int       `json:"track_record"`
	Discount      int       `json:"discount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Reception is the definition for Receive & Collection info
type Reception struct {
	VoucherNo          string       `json:"voucher_no"`
	ReceivedDate       string       `json:"received_date"`
	ReceivedAmount     int          `json:"received_amount"`
	SourceAccount      *Customer    `json:"source_account"`
	DestinationAccount *HeadAccount `json:"destination_account"`
	ChequeNo           string       `json:"cheque_no"`
	Carrier            string       `json:"carrier_info"`
	Description        string       `json:"description"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}

// Payment is the definition for payment info
type Payment struct {
	VoucherNo          string       `json:"voucher_no"`
	PaymentDate        string       `json:"payment_date"`
	PaidAmount         int          `json:"paid_amount"`
	SourceAccount      *HeadAccount `json:"source_account"`
	DestinationAccount *Supplier    `json:"destination_account"`
	ChequeNo           string       `json:"cheque_no"`
	Carrier            string       `json:"carrier_info"`
	Description        string       `json:"description"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}

// AmountTransfer is the definition for amount transfer info
type AmountTransfer struct {
	VoucherNo          string       `json:"voucher_no"`
	TransactionDate    string       `json:"transaction_date"`
	TransferAmount     int          `json:"transfer_amount"`
	SourceAccount      *HeadAccount `json:"source_account"`
	DestinationAccount *HeadAccount `json:"destination_account"`
	ChequeNo           string       `json:"cheque_no"`
	Carrier            string       `json:"carrier_info"`
	Description        string       `json:"description"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}

// FundAcquisition is the definition for Fund Acquisition info
type FundAcquisition struct {
	VoucherNo          string       `json:"voucher_no"`
	TransactionDate    string       `json:"transaction_date"`
	TransferAmount     int          `json:"transfer_amount"`
	SourceAccount      *StakeHolder `json:"source_account"`
	DestinationAccount *HeadAccount `json:"destination_account"`
	ChequeNo           string       `json:"cheque_no"`
	Carrier            string       `json:"carrier_info"`
	Description        string       `json:"description"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}

// AmountPayable is the definition for amount payable info
type AmountPayable struct {
	VoucherNo     string      `json:"voucher_no"`
	Date          string      `json:"date"`
	PayableAmount int         `json:"payable_amount"`
	Reason        string      `json:"reason"`
	AccountType   string      `json:"account_type"`
	AccountID     int         `json:"account_id"`
	AccountName   string      `json:"account_name"`
	HeadAccount   HeadAccount `json:"head_account"`
	ChequeNo      string      `json:"cheque_no"`
	Carrier       string      `json:"carrier_info"`
	Description   string      `json:"description"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

// AmountReceivable is the definition for amount receivable info
type AmountReceivable struct {
	VoucherNo        string      `json:"voucher_no"`
	Date             string      `json:"date"`
	ReceivableAmount int         `json:"receivable_amount"`
	Reason           string      `json:"reason"`
	AccountType      string      `json:"account_type"`
	AccountID        int         `json:"account_id"`
	AccountName      string      `json:"account_name"`
	HeadAccount      HeadAccount `json:"head_account"`
	ChequeNo         string      `json:"cheque_no"`
	Carrier          string      `json:"carrier_info"`
	Description      string      `json:"description"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}

// ExpenseList is the definition for expense type list
type ExpenseType struct {
	ID           int       `json:"id"`
	ExpenseName  string    `json:"expense_name"`
	TotalExpense string    `json:"total_expense"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Expense is the definition for expense info
type Expense struct {
	ID            int          `json:"id"`
	VoucherNo     string       `json:"voucher_no"`
	ExpenseType   ExpenseType  `json:"expense_type"`
	ExpenseDate   string       `json:"expense_date"`
	ExpenseAmount int          `json:"expense_amount"`
	SourceAccount *HeadAccount `json:"source_account"`
	ChequeNo      string       `json:"cheque_no"`
	Carrier       string       `json:"carrier_info"`
	Description   string       `json:"description"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

// IncomeStatement is the definition for income statement
type IncomeStatement struct {
	//Revenue Section
	GrossSales    int `json:"gross_sales"`
	SalesReturn   int `json:"sales_return"`
	SalesDiscount int `json:"sales_discount"`
	//Cost of Goods Sold
	GoodsPurchased      int `json:"goods_purchased"`
	PurchaseReturn      int `json:"purchase_return"`
	PurchaseDiscount    int `json:"purchase_discount"`
	TotalAvailableGoods int `json:"total_available_goods"`
	//Non Operating Sources
	NonOperatingIncome  int `json:"non_operating_income"`
	NonOperatingExpense int `json:"non_operating_expense"`

	//Expenses
	ExpenseSection struct {
		Rent                      int `json:"rent"`
		Utilities                 int `json:"utilities"`
		SalariesAndWages          int `json:"salaries_and_wages"`
		AdvertisingAndPromotions  int `json:"advertising_and_promotions"`
		MaintenanceAndRepairs     int `json:"maintenance_and_repairs"`
		OfficeSupplies            int `json:"office_supplies"`
		Insurance                 int `json:"insurance"`
		DeliveryAndFreightCharges int `json:"delivery_and_freight_charges"`
		Depreciation              int `json:"depreciation"`
		TaxesAndLicenses          int `json:"taxes_and_licenses"`
		InventoryCosts            int `json:"inventory_costs"`
		OfficeExpense             int `json:"office_expense"`
		TravelExpense             int `json:"travel_expense"`
		TrainingAndDevelopment    int `json:"training_and_development"`
		BankChargesAndFees        int `json:"bank_charges_and_fees"`
		InterestOnLoans           int `json:"interest_on_loans"`
		SoftwareAndSubscriptions  int `json:"software_and_subscriptions"`
		SecurityCosts             int `json:"security_costs"`
		WasteDisposal             int `json:"waste_disposal"`
		Miscellaneous             int `json:"other"`
	} `json:"expense_section"`
}

//Transaction is the definition for the financial transaction

type Transaction struct {
	ID                     int       `json:"id"`
	VoucherNo              string    `json:"voucher_no"`
	TransactionType        string    `json:"transaction_type"`
	SourceAccountName      string    `json:"source_account_name"`
	SourceType             string    `json:"source_type"`
	SourceID               int       `json:"source_id"`
	DestinationType        string    `json:"destination_type"`
	DestinationAccountName string    `json:"destination_account_name"`
	DestinationID          int       `json:"destination_id"`
	Amount                 int       `json:"amount"`
	CurrentBalance         int       `json:"current_balance"`
	TransactionDate        time.Time `json:"transaction_date"`
	Description            string    `json:"description"`
	ChequeNo               string    `json:"cheque_no"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type TopSheet struct {
	ID                    int       `json:"id"`
	SheetDate             time.Time `json:"sheet_date"`
	TotalPurchases        int       `json:"total_purchases"`
	PurchasesDiscount     int       `json:"purchases_discount"`
	TotalPayments         int       `json:"total_payments"`
	TotalPurchaseReturns  int       `json:"total_purchase_returns"`
	TotalSales            int       `json:"total_sales"`
	SalesDiscount         int       `json:"sales_discount"`
	TotalReceivedPayments int       `json:"total_received_payments"`
	TotalSaleReturns      int       `json:"total_sale_returns"`
	TotalExpenses         int       `json:"total_expenses"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
type TrialBalance struct {
	CapitalInvestment          int       `json:"capital_investment"`
	CashBankAccounts           int       `json:"cash_bank_accounts"`
	CurrentAssets              int       `json:"current_assets"`
	CurrentLiabilities         int       `json:"current_liabilities"`
	CustomerAccountsReceivable int       `json:"customer_accounts_receivable"`
	SupplierAccountsPayable    int       `json:"supplier_accounts_payable"`
	EmployeeAccountsPayable    int       `json:"employee_accounts_payable"`
	ExpenseAccounts            int       `json:"expense_accounts"`
	FixedAssets                int       `json:"fixed_assets"`
	LoanAccounts               int       `json:"loan_accounts"`
	RevenueAccounts            int       `json:"revenue_accounts"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}
type BalanceSheet struct {
	CashBankAccounts           int       `json:"cash_bank_accounts"`
	CurrentAssets              int       `json:"current_assets"`
	CustomerAccountsReceivable int       `json:"customer_accounts_receivable"`
	SupplierAccountsReceivable int       `json:"supplier_accounts_receivable"`
	CustomerAccountsPayable    int       `json:"customer_accounts_payable"`
	SupplierAccountsPayable    int       `json:"supplier_accounts_payable"`
	CapitalInvestment          int       `json:"capital_investment"`
	LoanAccounts               int       `json:"loan_accounts"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}
type CompanyProfile struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Slogan          string    `json:"slogan"`
	Mobile          string    `json:"mobile"`
	WhatsappAccount string    `json:"whatsapp"`
	Telephone       string    `json:"tel"`
	Email           string    `json:"email"`
	Division        string    `json:"division"`
	District        string    `json:"district"`
	Upazila         string    `json:"upazila"`
	Area            string    `json:"area"`
	PostalCode      string    `json:"postal_code"`
	LogoLink        string    `json:"logo_link"`
	OpeningDate     time.Time `json:"opening_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
