package neo4j

import (
	"testing"

	"github.com/researchbot/server/internal/health"
)

var _ health.Checker = (*Client)(nil)

func TestClientImplementsChecker(t *testing.T) {
	var _ health.Checker = (*Client)(nil)
}
