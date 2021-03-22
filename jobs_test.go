package p4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type jobsInput struct {
	args []string
	res  []map[interface{}]interface{}
}

type jobsTest struct {
	input jobsInput
	want  []Job
}

var jobsTests = []jobsTest{
	{
		input: jobsInput{},
		want:  []Job{},
	},
	{
		// Single user protection
		input: jobsInput{
			args: []string{},
			res: []map[interface{}]interface{}{{
				"Job":    "job000001",
				"Change": "1",
				"Date":   "1612571080",
				"User":   "perforce",
				"Client": "p4_ws",
				"Status": "closed",
				"code":   "stat",
			}},
		},
		want: []Job{
			{
				Job:    "job000001",
				Change: "1",
				Date:   "1612571080",
				User:   "perforce",
				Client: "p4_ws",
				Status: "closed",
				Code:   "stat",
			},
		},
	},
	{
		// Single user protection
		input: jobsInput{
			args: []string{},
			res: []map[interface{}]interface{}{
				{
					"Job":    "job000001",
					"Change": "1",
					"Date":   "1612571080",
					"User":   "perforce",
					"Client": "p4_ws",
					"Status": "closed",
					"code":   "stat",
				},
				{
					"Job":    "job000002",
					"Change": "1",
					"Date":   "1612571081",
					"User":   "perforce",
					"Client": "p4_ws",
					"Status": "closed",
					"code":   "stat",
				},
			},
		},
		want: []Job{
			{
				Job:    "job000001",
				Change: "1",
				Date:   "1612571080",
				User:   "perforce",
				Client: "p4_ws",
				Status: "closed",
				Code:   "stat",
			},
			{
				Job:    "job000002",
				Change: "1",
				Date:   "1612571081",
				User:   "perforce",
				Client: "p4_ws",
				Status: "closed",
				Code:   "stat",
			},
		},
	},
}

func TestJobs(t *testing.T) {
	for _, tst := range jobsTests {
		fp4 := FakeP4Runner{}
		fp4.On("Run", []string{"jobs"}).Return(tst.input.res, nil)
		fs, err := RunJobs(&fp4, tst.input.args)
		assert.Nil(t, err)
		assert.Equal(t, tst.want, fs)
	}
}
