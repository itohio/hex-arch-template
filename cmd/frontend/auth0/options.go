package auth0

type Options struct {
	domain      string
	clientId    string
	audience    string
	scope       string
	redirectUri string
}

type OptionFunc func(o *Options) error

func WithDomain(domain string) OptionFunc {
	return func(o *Options) error {
		o.domain = domain
		return nil
	}
}
func WithClientID(clientId string) OptionFunc {
	return func(o *Options) error {
		o.clientId = clientId
		return nil
	}
}
func WithAudience(audience string) OptionFunc {
	return func(o *Options) error {
		o.audience = audience
		return nil
	}
}
func WithScope(scope string) OptionFunc {
	return func(o *Options) error {
		o.scope = scope
		return nil
	}
}
func WithRedirectUri(redirectUri string) OptionFunc {
	return func(o *Options) error {
		o.redirectUri = redirectUri
		return nil
	}
}
