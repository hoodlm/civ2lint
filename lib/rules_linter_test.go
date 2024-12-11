package lib_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"

	"github.com/leonsp/civ2lint/lib"
)

var sugar *zap.SugaredLogger

var _ = Describe("RuleLinter", func() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar = logger.Sugar()

	Describe("Linting advances", func() {
		It("Detects partial no", func() {
			cl := lib.Civ2Linter{
				Logger: sugar,
				Sections: map[string][]string{
					"@CIVILIZE": {
						"Advanced Flight,    4,-2,  nil, no, 3, 4",
					},
				},
			}

			Expect(cl.LintAdvances()).To(HaveOccurred())
		})
		It("Detects self-reference", func() {
			cl := lib.Civ2Linter{
				Logger: sugar,
				Sections: map[string][]string{
					"@CIVILIZE": {
						"Advanced Flight,    4,-2,  AFl, nil, 3, 4",
					},
				},
			}

			Expect(cl.LintAdvances()).To(HaveOccurred())
		})
		It("Detects loops", func() {
			cl := lib.Civ2Linter{
				Logger: sugar,
				Sections: map[string][]string{
					"@CIVILIZE": {
						"Advanced Flight,    4,-2,  Alp, nil, 3, 4",
						"Alphabet,           5, 1,  AFl, nil, 0, 3",
					},
				},
			}

			Expect(cl.LintAdvances()).To(HaveOccurred())
		})
	})
})
