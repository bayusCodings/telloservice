package middlewares

type validationContextKey string
type authContextKey string

const (
	ValidatedRequestKey validationContextKey = "validatedRequest"
	UserPayloadKey      authContextKey       = "userPayload"
)
