package config

// An Env is a possible deployment environment. A local machine
// is represented by None. For this project, we only have local
// and production, but you could imagine adding other environments
// for staging and development purposes.
type Env string

const (
	// Production is the Production environment
	Prod Env = "production"
	// None is the local environment.
	None Env = ""
)

// NewEnvFromString takes a string and asserts that the string
// maps directly to an environment, panicking if not.
func NewEnvFromString(name string) Env {
	var allEnvs = []Env{Prod, None}
	for _, env := range allEnvs {
		if string(env) == name {
			return env
		}
	}
	panic("String provided is not a valid environment.")
}

func (env Env) String() string {
	return string(env)
}
