package response

type Response struct {
	Code   int         `json:"code"`
	Errors []string    `json:"errors"`
	Data   interface{} `json:"data"`
}

type DocRegisterResponse200 struct {
	Code   int               `json:"code" example:"200"`
	Errors []string          `json:"errors"`
	Data   map[string]string `json:"data"`
}

type DocRegisterResponse422 struct {
	Code   int               `json:"code" example:"422"`
	Errors []string          `json:"errors" example:"email вже зареєстровано,password обов'язкове поле"`
	Data   map[string]string `json:"data"`
}

type DocRegisterResponse500 struct {
	Code   int               `json:"code" example:"500"`
	Errors []string          `json:"errors" example:"щось пішло не так"`
	Data   map[string]string `json:"data"`
}

type DocLoginResponse200 struct {
	Code   int               `json:"code" example:"200"`
	Errors []string          `json:"errors"`
	Data   map[string]string `json:"data" example:"jwt:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc3OTM5NzcsInN1YiI6NX0.Y90d9seg_kX3zH1JeiTqegMtVAWAqHE74teYF_4Zpxo"`
}

type DocLoginResponse422 struct {
	Code   int               `json:"code" example:"422"`
	Errors []string          `json:"errors" example:"email не знайдено,password обов'язкове поле"`
	Data   map[string]string `json:"data"`
}

type DocLoginResponse500 struct {
	Code   int               `json:"code" example:"500"`
	Errors []string          `json:"errors" example:"щось пішло не так"`
	Data   map[string]string `json:"data"`
}

func NewResponse(code int, errors []string, data interface{}) Response {
	return Response{
		Code:   code,
		Errors: errors,
		Data:   data,
	}
}
