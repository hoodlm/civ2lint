package lib

import (
	"go.uber.org/zap"
)

type Config struct {
	Path   string
	Logger *zap.SugaredLogger
}

type Civ2Linter struct {
	Config       Config
	Logger       *zap.SugaredLogger
	Sections     map[string][]string
	SeenSections map[string]bool
	Rules        Civ2Rules
}

type Civ2Rules struct {
	Civilize map[string]Civilize
}

type Civilize struct {
	Name     string
	AiValue  int
	Modifier int
	Preq1    string
	Preq2    string
	Epoch    int
	Category int
}
