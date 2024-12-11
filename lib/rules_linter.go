package lib

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (cl *Civ2Linter) LintAdvances() error {
	section := "@CIVILIZE"
	lines, ok := cl.Sections[section]
	if !ok {
		message := fmt.Sprintf("section missing: %s", section)
		cl.Logger.Error(message)
		return errors.New(message)
	}

	cl.Rules.Civilize = make(map[string]Civilize, len(lines))
	for i, line := range lines {
		cols := strings.Split(line, ",")
		if len(cols) < 7 {
			return fmt.Errorf("too few columns: %s", line)
		} else if len(cols) > 7 {
			return fmt.Errorf("too many columns: %s", line)
		}
		aiValue, err := strconv.Atoi(strings.TrimSpace(cols[1]))
		if err != nil {
			return err
		}
		modifier, err := strconv.Atoi(strings.TrimSpace(cols[2]))
		if err != nil {
			return err
		}
		epoch, err := strconv.Atoi(strings.TrimSpace(cols[5]))
		if err != nil {
			return err
		}
		category, err := strconv.Atoi(strings.TrimSpace(cols[6]))
		if err != nil {
			return err
		}
		advance := Civilize{
			Name:     strings.TrimSpace(cols[0]),
			AiValue:  aiValue,
			Modifier: modifier,
			Preq1:    strings.TrimSpace(cols[3]),
			Preq2:    strings.TrimSpace(cols[4]),
			Epoch:    epoch,
			Category: category,
		}
		cl.Rules.Civilize[AdvanceCodes[i]] = advance
	}

	return nil
}
