package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

var (
	lintExample = `  # linting is what we do
  wman lint`
)

// newLintCmd returns a new initialized instance of the lint sub command
func newLintCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "lint",
		Short:   "Linting exercise",
		Example: lintExample,
		RunE:    lintTest,
	}

	return cmd
}

type T struct {
	Field1 string
	Field2 string
}

type T2 struct {
	Field1 string
	Field2 string
}

func (t T) String() string {
	return fmt.Sprintf("%s.%s", t.Field1, t.Field2)
}

func Check(test string) bool {
	return test != ""
}

func lintTest(cmd *cobra.Command, args []string) error {

	var x T
	y := T2(x)

	fmt.Println(y)

	//Reversable was a good exercise right?
	strs := []string{"kind: Namespace", "kind:Namespace", "kind: Foo", "kind", "kind:  Namespace", "kind: Namespace"}
	newStrs := []string{}

	if len(strs) != 0 {
		fmt.Println("strs is not empty")
	}

	newStrs = append(newStrs, strs...)

	nsRegex := regexp.MustCompile(`kind:\s*Namespace`)
	set := make(map[string]bool)

	for _, str := range newStrs {
		if nsRegex.MatchString(str) {
			fmt.Println(str)
		}
	}

	for key := range set {
		fmt.Println(key)
	}

	return nil
}
