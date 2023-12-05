package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.loginHandler)
	rt.router.GET("/user/:userId/stream", rt.getStreamHandler)
	rt.router.PUT("/user/:userId/username", rt.setUsernameHandler)
	rt.router.GET("/user/:userId/profile", rt.getProfileHandler)
	rt.router.POST("/user/:userId/photos", rt.uploadPhotoHandler)
	rt.router.DELETE("/user/:userId/photos/:photoId", rt.deletePhotoHandler)
	rt.router.PUT("/user/:userId/follow/:followerId", rt.followUserHandler)
	rt.router.DELETE("/user/:userId/follow/:followerId", rt.unfollowUserHandler)
	rt.router.PUT("/user/:userId/ban/:bannedId", rt.banUserHandler)
	rt.router.DELETE("/user/:userId/ban/:bannedId", rt.unbanUserHandler)
	rt.router.PUT("/photos/:photoId/likes/:likeId", rt.addLikeHandler)
	rt.router.DELETE("/photos/:photoId/likes/:likeId", rt.removeLikeHandler)
	rt.router.POST("/user/:userId/photos/:photoId/comments", rt.addCommentHandler)
	rt.router.DELETE("/user/:userId/photos/:photoId/comments/:commentId", rt.removeCommentHandler)
	
	
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
