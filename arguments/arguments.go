package arguments

type ArgList struct {
	Env          string `arg:"-e" default:"dev"`
	Simulation   string `arg:"positional" default:"carro"`
	Proposal     bool   `arg:"-p" default:"false"`
	ParseResonse bool   `arg:"-r" default:"false"`
}
