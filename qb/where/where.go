package where

import "github.com/scylladb/gocqlx/qb"

// BuildStms execute Build function over all the given Where statements on the slice and return a slice of gocqlx qb.Cmp
// or an error.
func BuildStms(stms []Stm) ([]qb.Cmp, error) {
	cmps := make([]qb.Cmp, 0)

	for _, stm := range stms {
		cmp, err := stm.Build()
		if err != nil {
			return nil, err
		}

		cmps = append(cmps, cmp)
	}

	return cmps, nil
}
