package appmod

type envType int

const (
	DEV envType = iota // Development
	SND                // Sandbox
	REG                // Regression
	TRN                // Training
	STG                // Staging / Pre-Production
	PRD                // Production
)

type Config struct {
	Env      envType
	Platform string
	//...
}

func (et envType) String() string {
	var value string

	switch et {
	case DEV:
		value = "Development"
	case SND:
		value = "Sandbox"
	case REG:
		value = "Regression"
	case TRN:
		value = "Training"
	case STG:
		value = "Staging"
	case PRD:
		value = "Production"
	}
	return value
}
