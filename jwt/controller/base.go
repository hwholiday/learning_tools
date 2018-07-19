package controller

type BaseController struct {
}

func (c *BaseController) GetJwtKey() string {
	return "holiday"
}
