package main

// Import Packages
import (
	"fmt"
	"net/http"

	// Import DisGOAuth
	disgoauth "github.com/ruxwez/disgoauth"
)

// Main function
func main() {
	// Establish a new discord client
	var dc *disgoauth.Client = disgoauth.Init(&disgoauth.Client{
		ClientID:     "883006609280864257",
		ClientSecret: "X-9n0rEBywVu1KKKOQSHskRQM7L8UlOV",
		RedirectURI:  "http://localhost:8000/redirect",
		Scopes:       []string{disgoauth.ScopeIdentify},
	})

	////////////////////////////////////////////////////////////////////////
	//
	// Home Page Handler
	//
	// It is suggested to put this in it's own function,
	// I only did it like for the showcase.
	//
	////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Send the user to the discord authentication
		// website. This is where they authorize access.
		//
		// The third parameter in the RedirectHandler is the
		// state. If you're storing a state, PLEASE base64 encode
		// it beforehand!
		dc.RedirectHandler(w, r, "") // w: http.ResponseWriter, r: *http.Request, state: string
	})

	////////////////////////////////////////////////////////////////////////
	//
	// The OAuth URL Redirect Uri
	//
	// It is suggested to put this in it's own function,
	// I only did it like for the showcase.
	//
	////////////////////////////////////////////////////////////////////////
	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		// Define Variables
		var (
			// Get the code from the redirect parameters (&code=...)
			codeFromURLParamaters = r.URL.Query()["code"][0]

			// Get the access token using the above codeFromURLParamaters
			accessToken, _ = dc.GetOnlyAccessToken(codeFromURLParamaters)

			// Get the authorized user's data using the above accessToken
			userData, _ = disgoauth.GetUserData(accessToken)
		)
		// Print the user data map
		fmt.Fprint(w, userData)
	})

	// Listen and Serve to the incoming http requests
	http.ListenAndServe(":8000", nil)
}
