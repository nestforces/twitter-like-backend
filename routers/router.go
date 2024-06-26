package routers

import (
    "twitter-like-backend/handlers"
    "twitter-like-backend/middlewares"

    "github.com/gofiber/fiber/v2"
)


func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

    // Authentication routes
    api.Post("/register", handlers.Register)
    api.Post("/login", handlers.Login)

    // Protected routes
    api.Use(middlewares.AuthRequired())

    // User routes
    api.Patch("/user/change-username", handlers.ChangeUsername)
    api.Patch("/user/update-profile", handlers.UpdateProfile)
	api.Post("/follow/:id", handlers.FollowUser)
    api.Delete("/unfollow/:id", handlers.UnfollowUser)

    // Tweet routes
    api.Post("/tweets", handlers.CreateTweet)
    api.Get("/tweets", handlers.GetTweets)
    api.Post("/tweets/:id/reply", handlers.ReplyToTweet)

    // Like routes
    api.Post("/tweets/:id/like", handlers.LikeTweet)
    api.Delete("/tweets/:id/unlike", handlers.UnlikeTweet)

    // Retweet routes
    api.Post("/tweets/:id/retweet", handlers.Retweet)
    api.Delete("/tweets/:id/unretweet", handlers.RemoveRetweet)

    // Search routes
    api.Get("/search", handlers.SearchTweets)

     // Direct Message routes
     api.Post("/dm", handlers.SendDirectMessage)
     api.Get("/dm/:id", handlers.GetDirectMessages)
}
