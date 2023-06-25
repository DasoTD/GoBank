package api

import (
	"os"
	"testing"
	// "time"

	"github.com/gin-gonic/gin"
	// "github.com/stretchr/testify/require"
	// db "github.com/dasotd/gobank/db/sqlc"
	// "github.com/dasotd/gobank/util"
)

// func newTestServer(t *testing.T, store db.Bank) *Server {
// 	config := util.Config{
// 		TokenSymmetricKey:   util.RandomString(32),
// 		AccessTokenDuration: time.Minute,
// 	}

// 	server, err := NewServer(config, store)
// 	require.NoError(t, err)

// 	return server
// }

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}