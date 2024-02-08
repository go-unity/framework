package support

const Version string = "v1.0.0"

const (
	EnvRuntime    = "runtime"
	EnvArtisan    = "artisan"
	EnvTest       = "test"
	EnvProduction = "production"
)

var (
	Env                  = EnvRuntime
	EnvPath              = ".env"
	IsKeyGenerateCommand = false
	SecretsName          = ""
	RelativePath         string
	RootPath             string
)
