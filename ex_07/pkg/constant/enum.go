package constant

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyVND Currency = "VND"
)

type WeightUnit string

const (
	WeightUnitGram WeightUnit = "GRAM"
	WeightUnitLBS  WeightUnit = "LBS"
	WeightUnitKg   WeightUnit = "KG"
)

type ProductStatus string

const (
	ProductStatusActive  ProductStatus = "ACTIVE"
	ProductStatusDeleted ProductStatus = "DELETED"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "PENDING"
	PaymentStatusPaid    PaymentStatus = "PAID"
)
