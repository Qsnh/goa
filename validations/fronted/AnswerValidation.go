package validations

type AnswerValidation struct {
	Description string `form:"description" valid:"Required"`
}