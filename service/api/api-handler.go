package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(rt.loginHandler))
	rt.router.GET("/user/:userId/stream", rt.wrap(rt.getStreamHandler))
	rt.router.PUT("/user/:userId/username", rt.wrap(rt.setUsernameHandler))
	rt.router.GET("/user/:userId/profile", rt.wrap(rt.getProfileHandler))
	rt.router.POST("/user/:userId/photos/", rt.wrap(rt.uploadPhotoHandler))
	rt.router.DELETE("/user/:userId/photos/:photoId", rt.wrap(rt.deletePhotoHandler))
	rt.router.PUT("/user/:userId/follow/:followerId", rt.wrap(rt.followUserHandler))
	rt.router.DELETE("/user/:userId/follow/:followerId", rt.wrap(rt.unfollowUserHandler))
	rt.router.PUT("/user/:userId/ban/:bannedId", rt.wrap(rt.banUserHandler))
	rt.router.DELETE("/user/:userId/ban/:bannedId", rt.wrap(rt.unbanUserHandler))
	rt.router.PUT("/photos/:photoId/likes/:likeId", rt.wrap(rt.addLikeHandler))
	rt.router.DELETE("/photos/:photoId/likes/:likeId", rt.wrap(rt.removeLikeHandler))
	rt.router.POST("/user/:userId/photos/:photoId/comments/", rt.wrap(rt.addCommentHandler))
	rt.router.DELETE("/user/:userId/photos/:photoId/comments/:commentId", rt.wrap(rt.removeCommentHandler))
	rt.router.GET("/user", rt.wrap(rt.GetAllUsersHandler))
	rt.router.GET("/photos/:photoId/image", rt.wrap(rt.getPhotoHandler))
	rt.router.DELETE("/user/:userId", rt.wrap(rt.DeleteUserHandler))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
