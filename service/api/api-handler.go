package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.loginHandler))
	rt.router.GET("/user/{userId}/stream", rt.wrap(rt.getStream))
	rt.router.PUT("/user/{userId}/username", rt.wrap(rt.setUsername))
	rt.router.GET("/user/{userId}/profile", rt.wrap(rt.getProfile))
	rt.router.GET("/user/{userId}/photos", rt.wrap(rt.postPhoto))
	rt.router.DELETE("/user/{userId}/photos/{photoId}", rt.wrap(rt.deletePhoto))
	rt.router.PUT("/user/{userId}/follow/{followerId}", rt.wrap(rt.addFollower))
	rt.router.DELETE("/user/{userId}/follow/{followerId}", rt.wrap(rt.removeFollow))
	rt.router.PUT("/user/{userId}/ban/{bannedId}", rt.wrap(rt.addBan))
	rt.router.DELETE("/user/{userId}/ban/{bannedId}", rt.wrap(rt.removeBan))
	rt.router.PUT("/user/{userId}/photos/{photoId}/likes/{likeId}", rt.wrap(rt.addLike))
	rt.router.DELETE("/user/{userId}/photos/{photoId}/likes/{likeId}", rt.wrap(rt.removeLike))
	rt.router.POST("/user/{userId}/photos/{photoId}/comments", rt.wrap(rt.addComment))
	rt.router.DELETE("/user/{userId}/photos/{photoId}/comments/{commentId}", rt.wrap(rt.removeComment))
	router.GET("/user/{userId}/stream", rt.getStreamHandler)
	
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
