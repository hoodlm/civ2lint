package lib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

func New(c Config, l *zap.SugaredLogger) Civ2Linter {
	cl := Civ2Linter{
		Config:       c,
		Logger:       l,
		Sections:     make(map[string][]string, 25),
		SeenSections: make(map[string]bool, 25),
	}
	return cl
}

func (cl *Civ2Linter) Lint() error {
	var err error
	err = cl.parseFile("rules.txt")
	if err != nil {
		cl.Logger.Error("parsing failed:", zap.Error(err))
		return err
	}

	fmt.Println("Seen sections", cl.SeenSections)

	err = cl.LintAdvances()
	if err != nil {
		cl.Logger.Error("linting advances failed:", zap.Error(err))
		return err
	}

	fmt.Println(cl.Rules.Civilize)

	return nil
}

func (cl *Civ2Linter) parseFile(filename string) error {
	filePath := filepath.Join(cl.Config.Path, filename)
	_, err := os.Stat(filePath)
	if err != nil {
		cl.Logger.Error(fmt.Sprintf("%s does not exist:", filename), zap.Error(err))
		return err
	}

	readFile, err := os.Open(filePath)
	if err != nil {
		cl.Logger.Error(fmt.Sprintf("could not open %s:", filename), zap.Error(err))
		return err
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	currentSection := ""
	for fileScanner.Scan() {
		line := strings.TrimSpace(
			strings.Split(
				fileScanner.Text(), ";")[0])
		// cl.Logger.Info(line)
		if strings.HasPrefix(line, ";") {
			continue
		} else if strings.HasPrefix(line, "@") {
			cl.Logger.Infof("new section: %s", line)
			currentSection = line
			cl.SeenSections[currentSection] = true
			cl.Sections[currentSection] = []string{}
		} else if len(line) == 0 {
			continue
		} else if currentSection == "" {
			return fmt.Errorf("reached content before any section: %s", line)
		} else {
			cl.Sections[currentSection] = append(cl.Sections[currentSection], line)
		}
	}
	return nil
}
