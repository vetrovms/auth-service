package response

// Response Структура відповіді.
type Response struct {
	Code   int         `json:"code"`
	Errors []string    `json:"errors"`
	Data   interface{} `json:"data"`
}

// DocRegisterResponse200 Приклад відповіді для генерації документації.
type DocRegisterResponse200 struct {
	Code   int               `json:"code" example:"200"`
	Errors []string          `json:"errors"`
	Data   map[string]string `json:"data"`
}

// DocRegisterResponse422 Приклад відповіді для генерації документації.
type DocRegisterResponse422 struct {
	Code   int               `json:"code" example:"422"`
	Errors []string          `json:"errors" example:"email вже зареєстровано,password обов'язкове поле"`
	Data   map[string]string `json:"data"`
}

// DocRegisterResponse500 Приклад відповіді для генерації документації.
type DocRegisterResponse500 struct {
	Code   int               `json:"code" example:"500"`
	Errors []string          `json:"errors" example:"щось пішло не так"`
	Data   map[string]string `json:"data"`
}

// DocLoginResponse200 Приклад відповіді для генерації документації.
type DocLoginResponse200 struct {
	Code   int               `json:"code" example:"200"`
	Errors []string          `json:"errors"`
	Data   map[string]string `json:"data" example:"jwt:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc3OTM5NzcsInN1YiI6NX0.Y90d9seg_kX3zH1JeiTqegMtVAWAqHE74teYF_4Zpxo"`
}

// DocLoginResponse422 Приклад відповіді для генерації документації.
type DocLoginResponse422 struct {
	Code   int               `json:"code" example:"422"`
	Errors []string          `json:"errors" example:"email не знайдено,password обов'язкове поле"`
	Data   map[string]string `json:"data"`
}

// DocLoginResponse500 Приклад відповіді для генерації документації.
type DocLoginResponse500 struct {
	Code   int               `json:"code" example:"500"`
	Errors []string          `json:"errors" example:"щось пішло не так"`
	Data   map[string]string `json:"data"`
}

// NewResonse Конструктор повертає структуру відповіді.
func NewResponse(code int, errors []string, data interface{}) Response {
	return Response{
		Code:   code,
		Errors: errors,
		Data:   data,
	}
}
