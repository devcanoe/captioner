package router

func (r *Router) Merouter() {
	me := r.router.Group("/me")

	{
		me.GET("/profile", r.user.HealthCheck)
	}
}
