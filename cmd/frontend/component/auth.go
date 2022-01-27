package component

import (
	"log"
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

type AuthPre struct {
	vecty.Core
	state  *GlobalState
	apiUrl string

	Error       error
	RedirectUri string
}

type AuthPost struct {
	vecty.Core
	Pre *AuthPre
}

func getpathname() string {
	return js.Global().Get("location").Get("pathname").String()
}

func locationorigin() string {
	return js.Global().Get("location").Get("origin").String()
}

func NewAuth(s *GlobalState, ApiUri string) (*AuthPre, *AuthPost) {
	pre := &AuthPre{state: s, apiUrl: ApiUri}
	post := &AuthPost{Pre: pre}

	switch {
	case s.Auth == nil:
		panic("Auth must be set")
	case ApiUri == "":
		panic("ApiUrl must be set")
	}

	queryStr := js.Global().Get("location").Get("search").String()
	hashStr := js.Global().Get("location").Get("hash").String()
	pre.RedirectUri, pre.Error = s.Auth.Handle(queryStr, hashStr)
	s.API.Init(ApiUri, s.Auth.Token())

	return pre, post
}

func (p *AuthPre) Render() vecty.ComponentOrHTML {
	return elem.Div()
}

func (p *AuthPost) Render() vecty.ComponentOrHTML {
	switch {
	case p.Pre == nil:
		panic("AuthPost: Auth must be set")
	}

	log.Println("Error: ", p.Pre.Error)
	log.Println("Redirect: ", p.Pre.RedirectUri)
	if p.Pre.Error == nil && p.Pre.RedirectUri != "" {
		router.Redirect(p.Pre.RedirectUri)
	} else if p.Pre.Error != nil {
		router.Redirect("/error")
	}

	return elem.Div()
}
