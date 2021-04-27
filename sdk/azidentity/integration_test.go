package azidentity

import (
	"os"
	"testing"
)

func integrationTestEnvVarGuard(t *testing.T, vl ...string) []string {
	fvl := make([]string, 0, len(vl)+1)
	fvl = append(fvl, "AZURE_GO_SDK_INT_TEST")
	fvl = append(fvl, vl...)
	for idx, v := range fvl {
		ev := os.Getenv(v)
		if ev == "" {
			t.Skipf("Skipping this integration test due to environment variable %q is not set.", v)
		}
		fvl[idx] = ev
	}
	return fvl[1:]
}
