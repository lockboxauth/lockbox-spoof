module impractical.co/auth/cmd/spoof

replace impractical.co/auth/sessions v0.0.0 => ../../sessions

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	impractical.co/auth/sessions v0.0.0
)
