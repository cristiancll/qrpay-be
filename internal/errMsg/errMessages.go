package errMsg

// Authentication errors
const (
	FailedAuthCheck         = "failed to check authorization"
	FailedUUIDContext       = "failed to get uuid from context"
	TokenNotString          = "token is not a string"
	FailedMetadataContext   = "failed to get metadata from context"
	FailedCookieMetadata    = "failed to get cookie from metadata"
	NoAuthToken             = "no authorization token provided"
	FailedAuthTokenMetadata = "failed to get authorization token from metadata"
	SigningMethodMismatch   = "signing method mismatch"
	KIDMismatch             = "key id mismatch"
	InvalidToken            = "invalid token"
	TokenExpired            = "token expired"
	UserNotVerified         = "user not verified"
	UserDisabled            = "user disabled"
	RoleNotFound            = "role not found"
)

// Input validation errors
const (
	PhoneRequired        = "phone is required"
	PasswordRequired     = "password is required"
	NameRequired         = "name is required"
	UUIDRequired         = "uuid is required"
	CategoryUUIDRequired = "category uuid is required"
	ItemUUIDRequired     = "item uuid is required"
	PriceRequired        = "price is required"
	UserUUIDRequired     = "user uuid is required"
	SaleUnitsRequired    = "sale units are required"
	QuantityRequired     = "quantity is required"
	SKUUUIDRequired      = "sku uuid is required"
	RoleRequired         = "role is required"
	UUIDInvalid          = "uuid is invalid"
	NotAnInteger         = "not an integer"
)

// Auth errors
const (
	FailedCreateAuth = "failed to create auth"
	FailedGetAuth    = "failed to get auth"
)

// User errors
const (
	UserNotAdmin      = "user is not admin"
	UserNotStaff      = "user is not staff"
	IncorrectPassword = "incorrect password"
	UserAlreadyExists = "user already exists"
	FailedCreateUser  = "failed to create user"
	FailedUpdateUser  = "failed to update user"
	FailedGetUser     = "failed to get user"
	FailedGetAllUser  = "failed to get all user"
	FailedCountUser   = "failed to count users"
)

// Category errors
const (
	FailedCreateCategory = "failed to create category"
	FailedUpdateCategory = "failed to update category"
	FailedDeleteCategory = "failed to delete category"
	FailedGetCategory    = "failed to get category"
	FailedGetAllCategory = "failed to get all category"
)

// Item errors
const (
	FailedCreateItem = "failed to create item"
	FailedUpdateItem = "failed to update item"
	FailedDeleteItem = "failed to delete item"
	FailedGetItem    = "failed to get item"
	FailedGetAllItem = "failed to get all item"
)

// SKU errors
const (
	FailedCreateSKU = "failed to create sku"
	FailedUpdateSKU = "failed to update sku"
	FailedDeleteSKU = "failed to delete sku"
	FailedGetSKU    = "failed to get sku"
	FailedGetAllSKU = "failed to get all sku"
)

// Stock errors
const (
	StockAlreadyExists = "stock already exists"
	FailedCreateStock  = "failed to create stock"
	FailedUpdateStock  = "failed to update stock"
	FailedDeleteStock  = "failed to delete stock"
	FailedGetStock     = "failed to get stock"
	FailedGetAllStock  = "failed to get all stock"
	FailedCountStock   = "failed to count stocks"
)

// Sale errors
const (
	FailedCreateSale     = "failed to create sale"
	FailedCreateSaleItem = "failed to create sale item"
	FailedGetAllSaleItem = "failed to get all sale item"
)

// Retrieval errors
const (
	FailedCreateRetrieval = "failed to create retrieval"
	FailedUpdateRetrieval = "failed to update retrieval"
	FailedDeleteRetrieval = "failed to delete retrieval"
	FailedGetRetrieval    = "failed to get retrieval"
	FailedGetAllRetrieval = "failed to get all retrieval"
)

// Operation log errors
const (
	FailedCreateOperationLog = "failed to create operation log"
)

// Other errors
const (
	FailedStringConversion    = "failed to convert to string"
	FailedRoleContext         = "failed to get role from context"
	FailedToMigrateRepository = "failed to migrate repository"
	FailedLogin               = "failed to login"
	FailedHeartbeat           = "failed to heartbeat"
	FailedDeleteCookie        = "failed to delete jwt cookie"
	FailedVerifyToken         = "failed to verify jwt token"
	FailedExecuteTransaction  = "failed to execute transaction"
)
