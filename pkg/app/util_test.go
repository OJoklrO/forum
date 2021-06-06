package app

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	require.Equal(t, TimeFormat(time.Now().Unix()), strconv.Itoa(time.Now().Hour()) + ":" + strconv.Itoa(time.Now().Minute()))
}
