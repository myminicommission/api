package utils

// ContextKey defines a type for context keys shared in the app
type ContextKey string

// ServerConfig defines the configuration for the server
type ServerConfig struct {
	Port          string
	SessionSecret string
	JWT           JWTConfig
	GraphQL       GQLConfig
	Database      DBConfig
	AuthProviders []AuthProvider
}

//JWTConfig defines the options for JWT tokens
type JWTConfig struct {
	Secret    string
	Algorithm string
}

// GQLConfig defines the configuration for the GQL Server
type GQLConfig struct {
	Path                string
	PlaygroundPath      string
	IsPlaygroundEnabled bool
}

// DBConfig defines the configuration for the DB config
type DBConfig struct {
	Dialect     string
	DSN         string
	LogMode     bool
	AutoMigrate bool
}

// AuthProvider defines the configuration for the Goth config
type AuthProvider struct {
	Provider  string
	ClientKey string
	Secret    string
	Domain    string // If needed, like with auth0
	Scopes    []string
}
