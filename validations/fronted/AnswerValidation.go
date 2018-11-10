package fronted

type AnswerValidation struct {
	Description string `form:"description" valid:"Required"`
}