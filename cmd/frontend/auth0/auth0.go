package auth0

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

type authState struct {
	ReturnTo string `json:"return_to"`
}

type Context struct {
	options       Options
	token         string
	authenticated bool
	scopes        string
	err           string
	expires       time.Duration
	timestamp     time.Time
	tokenType     string
}

func New(opts ...OptionFunc) (*Context, error) {
	ret := &Context{}
	for _, opt := range opts {
		if err := opt(&ret.options); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (a *Context) Token() string {
	return a.token
}

func (a *Context) Error() string {
	return a.err
}

func (a *Context) Authenticated() bool {
	return a.authenticated
}

func (a *Context) Scopes() []string {
	return strings.Split(a.scopes, ",")
}

func (a *Context) buildQueryUrl(endpoint string, params map[string]string) string {
	query := make([]string, 0, len(params))
	for k, v := range params {
		if v == "" {
			query = append(query, k)
			continue
		}
		query = append(query, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
	}
	return fmt.Sprintf("https://%s%s?%s", a.options.domain, endpoint, strings.Join(query, "&"))
}

func (a *Context) encodeState(s authState) string {
	data, err := json.Marshal(s)
	if err != nil {
		return ""
	}

	return base64.RawStdEncoding.EncodeToString(data)
}

func (a *Context) decodeState(s string) authState {
	var as authState

	str, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		return as
	}

	json.Unmarshal([]byte(str), &as)
	log.Println("decode State: ", as)

	return as
}

func (a *Context) LoginUrl(redirect string) string {
	log.Println("Login: ", redirect, " audience: ", a.options.audience)
	return a.buildQueryUrl(
		"/authorize",
		map[string]string{
			"client_id":    a.options.clientId,
			"audience":     a.options.audience,
			"redirect_uri": a.options.redirectUri,
			// "connection":    "",
			"response_type": "token",
			"state":         a.encodeState(authState{ReturnTo: redirect}),
		},
	)
}

func (a *Context) SignupUrl(redirect string) string {
	return a.buildQueryUrl(
		"/authorize",
		map[string]string{
			"client_id":    a.options.clientId,
			"audience":     a.options.audience,
			"redirect_uri": a.options.redirectUri,
			// "connection":    "",
			"response_type": "token",
			"screen_hint":   "signup",
			"state":         a.encodeState(authState{ReturnTo: redirect}),
		},
	)
}

func (a *Context) LogoutUrl(redirect string) string {
	return a.buildQueryUrl(
		"/v2/logout",
		map[string]string{
			"client_id": a.options.clientId,
			"audience":  a.options.audience,
			"returnTo":  a.options.redirectUri + redirect,
			"state":     a.encodeState(authState{ReturnTo: redirect}),
		},
	)
}

// Handles redirect from auth0
// updates token/authorized/scopes
// returns redirect Uri if authenticated, error otherwise
func (a *Context) Handle(queryStr, hashStr string) (string, error) {
	switch {
	case strings.HasPrefix(hashStr, "#error"):
		return a.handleError(queryStr, hashStr)
	case strings.HasPrefix(hashStr, "#access_token="):
		return a.handleAccessToken(queryStr, hashStr)
	case strings.HasPrefix(hashStr, "#logout"):
		return a.handleLogout()
	}
	a.err = ""

	if _, err := a.handleCookie(queryStr, hashStr); err != nil {
		return "", err
	}

	if _, err := a.handleExpiration(queryStr, hashStr); err != nil {
		return "", err
	}
	return "", nil
}

func readCookie() (c map[string]string) {
	cookiesStr := js.Global().Get("document").Get("cookie").String()
	cookies := strings.Split(cookiesStr, ";")

	c = make(map[string]string, len(cookies))
	for _, kv := range cookies {
		kvArr := strings.SplitN(strings.Trim(kv, " "), "=", 2)
		if len(kvArr) == 2 {
			log.Println(fmt.Sprintf("Found cookie '%s'='%s'", kvArr[0], kvArr[1]))
			qe, err := url.QueryUnescape(kvArr[1])
			if err != nil {
				continue
			}
			c[kvArr[0]] = qe
		}
	}

	return
}

func writeCookie(c map[string]string) {
	for k, v := range c {
		kv := fmt.Sprintf("%s=%s", k, url.QueryEscape(v))
		js.Global().Get("document").Set("cookie", kv)
	}
}

func resetCookie() {
	cookies := readCookie()
	cookies["token"] = ""
	cookies["token_type"] = ""
	cookies["expires"] = ""
	cookies["timestamp"] = ""
	writeCookie(cookies)
}

func (a *Context) handleCookie(queryStr, hashStr string) (string, error) {
	kv := readCookie()

	token, ok1 := kv["token"]
	ttype, ok2 := kv["token_type"]
	timestamp, ok3 := kv["timestamp"]
	expires, ok4 := kv["expires"]

	if !(ok1 && ok2 && ok3 && ok4) {
		return "", nil
	}

	timestampi, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return "", nil
	}
	expiresi, err := strconv.ParseInt(expires, 10, 64)
	if err != nil {
		return "", nil
	}

	expiresIn := time.Second * time.Duration(expiresi)
	ts := time.UnixMilli(timestampi)

	a.authenticated = true
	a.err = ""
	a.token = token
	a.tokenType = ttype
	a.expires = expiresIn
	a.timestamp = ts

	return "", nil
}

func (a *Context) handleExpiration(queryStr, hashStr string) (string, error) {
	if !a.authenticated {
		return "", nil
	}

	if time.Since(a.timestamp) < a.expires {
		return "", nil
	}

	a.handleLogout()
	a.err = "Authentication expired"

	return "/", fmt.Errorf("please login and try again")
}

func (a *Context) handleAccessToken(queryStr, hashStr string) (string, error) {
	params, err := url.ParseQuery(hashStr[1:])
	if err != nil {
		return "", err
	}

	if token, ok := params["access_token"]; ok && len(token) == 1 {
		a.timestamp = time.Now()
		a.authenticated = true
		a.token = token[0]
		a.tokenType = ""
		if tt, ok := params["token_type"]; ok && len(tt) == 1 {
			a.tokenType = tt[0]
		}
		if e, ok := params["expires_in"]; ok && len(e) == 1 {
			ei, err := strconv.Atoi(e[0])
			if err == nil {
				a.expires = time.Second * time.Duration(ei)
			}
		}
	}

	if a.authenticated {
		writeCookie(map[string]string{
			"token":      a.token,
			"token_type": a.tokenType,
			"expires":    fmt.Sprint(int(a.expires.Seconds())),
			"timestamp":  fmt.Sprint(a.timestamp.UnixMilli()),
		})
	} else {
		resetCookie()
	}

	if state, ok := params["state"]; ok && len(state) == 1 {
		return a.decodeState(state[0]).ReturnTo, nil
	}

	return "/", nil
}

func (a *Context) handleLogout() (r string, err error) {
	a.token = ""
	a.authenticated = false
	resetCookie()

	return "/", nil
}

func (a *Context) handleError(queryStr, hashStr string) (r string, err error) {
	a.handleLogout()

	params, err := url.ParseQuery(hashStr[1:])
	if err != nil {
		return "", err
	}

	errStr := ""

	if code, ok := params["error"]; ok {
		errStr = fmt.Sprintf("Error code: %s, ", code)
	}
	if descr, ok := params["error_description"]; ok {
		errStr += fmt.Sprintf("Error: %s", descr)
	}
	a.err = errStr
	return "", fmt.Errorf(errStr)
}
