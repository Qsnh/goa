package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:DashboardController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:DashboardController"],
		beego.ControllerComments{
			Method: "MemberQuestions",
			Router: `/user/:user_id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:DashboardController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:DashboardController"],
		beego.ControllerComments{
			Method: "MemberAnswers",
			Router: `/user/:user_id/answers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:IndexController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:IndexController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:IndexController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:IndexController"],
		beego.ControllerComments{
			Method: "CaptchaShow",
			Router: `/captcha`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/member`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "SendActiveMail",
			Router: `/member/active/mail/send`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "SendActiveMailHandler",
			Router: `/member/active/mail/send`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "Answers",
			Router: `/member/answers`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "ChangeAvatar",
			Router: `/member/avatar`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "ChangeAvatarHandler",
			Router: `/member/avatar`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/member/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "ChangePassword",
			Router: `/member/password`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "ChangePasswordHandler",
			Router: `/member/password`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "Profile",
			Router: `/member/profile`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "SaveProfileHandler",
			Router: `/member/profile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:MemberController"],
		beego.ControllerComments{
			Method: "Questions",
			Router: `/member/questions`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"],
		beego.ControllerComments{
			Method: "AnswerHandler",
			Router: `/member/questions/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/member/questions/:question_id/edit`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/member/questions/:question_id/edit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/member/questions/create`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"],
		beego.ControllerComments{
			Method: "Store",
			Router: `/member/questions/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:QuestionController"],
		beego.ControllerComments{
			Method: "Show",
			Router: `/questions/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UploadController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UploadController"],
		beego.ControllerComments{
			Method: "Image",
			Router: `/member/upload/image`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"],
		beego.ControllerComments{
			Method: "LoginHandler",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"],
		beego.ControllerComments{
			Method: "FindPassword",
			Router: `/password/find`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"],
		beego.ControllerComments{
			Method: "FindPasswordHandler",
			Router: `/password/find`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"],
		beego.ControllerComments{
			Method: "PasswordReset",
			Router: `/password/reset`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"],
		beego.ControllerComments{
			Method: "PasswordResetHandler",
			Router: `/password/reset`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"],
		beego.ControllerComments{
			Method: "RegisterHandler",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"] = append(beego.GlobalControllerRouter["github.com/Qsnh/goa/controllers/frontend:UserController"],
		beego.ControllerComments{
			Method: "ActiveHandler",
			Router: `/user/active`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
