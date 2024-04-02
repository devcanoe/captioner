package router

func (r *Router) AuthRouter() {
	a := r.router.Group("/auth")
	h := r.auth
	{
		a.POST("/signin", h.Signin).
			POST("/verify-email", h.VerifyEmail).
			POST("/signup", h.Signup).
			POST("/forgot-password", h.ForgotPassword).
			POST("/reset-password", h.ResetPassword)
	}
}
