package configs

// Servers consts

const (
	ADDR_SERVER = "10.12.16.52:8080"
	MONGO_HOST  = "mongodb://localhost:27017"
)

//Router consts

const (
	USER_PATH    = "/user/"
	TABLE_PATH   = "/table/"
	TABLE_ID     = "{master}/"
	PRODUCT_PATH = "/product/"
)

//Database consts

const (
	CLIENT_DATABASE    = "project"
	USER_COLLECTION    = "users"
	TABLE_COLLECTION   = "tables"
	PRODUCT_COLLECTION = "products"
)

// ERROR consts

const (
	STRUCTURE_ERROR    = "[ERROR] Invalid Structure: "
	RESPONSE_STRUCTURE = `{"error": "structure error"}`

	MARSHAL_ERROR    = "[ERROR] Failed to Marshal: "
	RESPONSE_MARSHAL = `{"error": "marshal failed"}`

	UNMARSHAL_ERROR    = "[ERROR] Failed to Unmarshal: "
	RESPONSE_UNMARSHAL = `{"error": "unmarshal failed"}`

	COOKIE_ERROR    = "[ERROR] Cannot find cookie: "
	RESPONSE_COOKIE = `{"error": "not cookie"}`

	VALIDATION_ERROR    = "[ERROR] Validation error: "
	RESPONSE_VALIDATION = `{"error": "unvalid requisition"}`

	RESPONSE_DATABASE = `{"error": "internal error"}`
)

// Token Consts

const (
	COOKIE_NAME    = "Token"
	LOGIN_ERROR    = "[ERROR] Failed on login user "
	RESPONSE_LOGIN = `{"error": "failed login"}`
	TOKEN_ERROR    = `{"error": "failed token"}`
)
