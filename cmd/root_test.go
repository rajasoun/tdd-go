package cmd

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("GO_ENV", "test")
}

func TestNewRootCmd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		testCaseName string
		input        string
		want         string
	}{
		{"Check with No Parameter should prompt help", "", "Usage"},
		{"Check with Wrong Parameter shows nothing", "dummy", ""},
		{"Check with -v Parameter shows version", "-v", "0.1v"},
		{"Check with -h Parameter shows help", "-h", "Usage"},
	}
	for _, tc := range tests {
		t.Run(tc.testCaseName, func(t *testing.T) {
			outputBuffer := bytes.NewBufferString("")
			rootCmd.SetOut(outputBuffer)
			//set the args only when input parameter is not equal to empty string
			if tc.input != "" {
				rootCmd.SetArgs([]string{tc.input})
			}
			err := rootCmd.Execute()
			if err != nil {
				assert.Error(t, err)
			}
			out, err := ioutil.ReadAll(outputBuffer)
			if err != nil {
				assert.Error(t, err)
			}
			got := string(out)
			assert.Contains(t, got, tc.want, "expected \"%s\" got \"%s\"", tc.want, got)

		})
	}
}

func TestExecute(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"Execute Method with Empty Input", "", "Usage", false},
		{"Execute Method with Non Empty Input", "dummy", "", false},
	}
	for _, tc := range tests {
		tc := tc // rebind tc into this lexical scope
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			rootCmd.SetArgs([]string{tc.input})
			if err := Execute(); (err != nil) != tc.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
