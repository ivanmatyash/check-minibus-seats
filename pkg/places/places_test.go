package places

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateDate(t *testing.T) {
	type testCase struct {
		name string
		in   string
		err  string
	}

	cases := []testCase{
		{
			name: "invalid date (1.2.3)",
			in:   "1.2.3",
			err:  "Error: Date '1.2.3' doesn't match dd.mm.yyyy format.",
		},
		{
			name: "valid date (01.01.2000)",
			in:   "01.01.2000",
			err:  "",
		},
		{
			name: "invalid date format (1.1.2000)",
			in:   "1.1.2000",
			err:  "Error: Date '1.1.2000' doesn't match dd.mm.yyyy format.",
		},
		{
			name: "invalid date format (15/10/2013)",
			in:   "15/10/2013",
			err:  "Error: Date '15/10/2013' doesn't match dd.mm.yyyy format.",
		},
		{
			name: "invalid date format (15-10-2013)",
			in:   "15-10-2013",
			err:  "Error: Date '15-10-2013' doesn't match dd.mm.yyyy format.",
		},
		{
			name: "invalid date (32.10.2005)",
			in:   "32.10.2005",
			err:  "Error: Date '32.10.2005' doesn't match dd.mm.yyyy format.",
		},
		{
			name: "invalid date (30.15.2005)",
			in:   "30.15.2005",
			err:  "Error: Date '30.15.2005' doesn't match dd.mm.yyyy format.",
		},
	}

	for _, c := range cases {
		t := t // Using the variable on range scope `t` in function literal
		t.Run(c.name, func(t *testing.T) {
			err := ValidateDate(&c.in)
			if err == nil && c.err != "" {
				t.Errorf("ValidateDate() returned (nil) error, but expected (%s).", c.err)
			} else if err != nil {
				assert.EqualError(t, err, c.err)
			}
		})

	}

}
