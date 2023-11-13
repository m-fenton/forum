package main

import (
	f "forum/services"
	"net/http"
)

// github client ID
// 4764bca26a3fa9ab0e32

//github secret
// 4b97506ea58726f63c065ce8ed60b35c1fcbf532

// Client ID
// 623060904480-iudd85gmtab7k4sr1utdo92e5lpfsuqd.apps.googleusercontent.com

// Client secret
// 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

// Creation date
// 28 June 2023 at 14:21:45 GMT+1

func main() {

	http.HandleFunc("/", f.ServePage)
	http.HandleFunc("/registerauth", f.RegisterUserAuth)
	http.HandleFunc("/loginauth", f.LogInUserAuth)
	http.HandleFunc("/google-login", f.InitiateGoogleAuth)
	http.HandleFunc("/github-login", f.InitiateGithubAuth)
	http.HandleFunc("/google-callback", f.GoogleHandler)
	http.HandleFunc("/github-callback", f.GithubHandler)
	http.HandleFunc("/logout", f.LogOut)
	http.HandleFunc("/upload", f.Authorize(f.CreateUserPost))
	http.HandleFunc("/uploadComment", f.Authorize(f.CreateUserComment))
	http.HandleFunc("/likeOrDislike", f.HandleLikesOrDislikes)
	http.HandleFunc("/myActivity", f.Authorize(f.MyActivity))
	http.HandleFunc("/category", f.CategoryPage)
	f.StartServer()

}
