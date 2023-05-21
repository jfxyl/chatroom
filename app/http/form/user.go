package form

type RegisterForm struct {
	Name            string `form:"name" json:"name" binding:"required,min=2,max=10"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=10"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password"`
	Gender          string `form:"gender" json:"gender" binding:"required,oneof=0 1 2"`
	Birthday        string `form:"birthday" json:"birthday" binding:"required"`
	Avatar          string `form:"avatar" json:"avatar" binding:"required"`
}
