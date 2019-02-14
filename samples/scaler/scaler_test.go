package scaler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type scalerSuite struct {
	suite.Suite
	scaler scaler
}

func (s *scalerSuite) TestScaler() {
	r := require.New(s.T())
	initial := fmt.Sprintf("%#v", s.scaler)

	s.scaler.scale(2.0)
	scaled := fmt.Sprintf("%#v", s.scaler)

	s.scaler.scale(0.5)
	scaledBack := fmt.Sprintf("%#v", s.scaler)

	fmt.Printf("Initial: %s\nScaled:  %s\nBack:    %s\n", initial, scaled, scaledBack)
	r.Equal(initial, scaledBack)
	r.NotEqual(initial, scaled)
}

func TestStorer(t *testing.T) {
	suite.Run(t, &scalerSuite{scaler: &circle{2}})
	suite.Run(t, &scalerSuite{scaler: &rect{1, 1}})
}
