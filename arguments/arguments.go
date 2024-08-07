package arguments

type ArgList struct {
	Env          string      `arg:"-e" default:"dev"`
	Simulation   string      `arg:"-s" default:"carro"`
	Proposal     bool        `arg:"-p" default:"false"`
	ParseResonse bool        `arg:"-r" default:"false"`
	Async        *Subcommand `arg:"subcommand:async"`
}

type Subcommand struct {
  Run bool `arg:"-a" default:"false"`
}
