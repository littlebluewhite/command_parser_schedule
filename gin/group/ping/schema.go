package ping

type SwaggerPing struct {
	Example string `json:"example" binging:"required" example:"asdfasdf"`
}

type SwaggerListPing struct {
	Name string `json:"name" binging:"required" example:"wilson"`
	Age  int    `json:"age" binging:"required" example:"20"`
}
