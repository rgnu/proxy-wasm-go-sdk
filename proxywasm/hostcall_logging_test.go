package proxywasm

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rgnu/proxy-wasm-go-sdk/proxywasm/rawhostcall"
	"github.com/rgnu/proxy-wasm-go-sdk/proxywasm/types"
)

type logHost struct {
	rawhostcall.DefaultProxyWAMSHost
	t           *testing.T
	expMessage  string
	expLogLevel types.LogLevel
}

func (l logHost) ProxyLog(logLevel types.LogLevel, messageData *byte, messageSize int) types.Status {
	actual := RawBytePtrToString(messageData, messageSize)
	assert.Equal(l.t, l.expMessage, actual)
	assert.Equal(l.t, l.expLogLevel, logLevel)
	return types.StatusOK
}

func TestHostCall_Logging(t *testing.T) {
	hostMutex.Lock()
	defer hostMutex.Unlock()

	t.Run("trace", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "trace",
			expLogLevel:          types.LogLevelTrace,
		})
		LogTrace("trace")
	})

	t.Run("tracef", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "trace: log",
			expLogLevel:          types.LogLevelTrace,
		})
		LogTracef("trace: %s", "log")
	})

	t.Run("debug", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "abc",
			expLogLevel:          types.LogLevelDebug,
		})
		LogDebug("abc")
	})

	t.Run("debugf", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "debug: log",
			expLogLevel:          types.LogLevelDebug,
		})
		LogDebugf("debug: %s", "log")
	})

	t.Run("info", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "info",
			expLogLevel:          types.LogLevelInfo,
		})
		LogInfo("info")
	})

	t.Run("infof", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "info: log: 10",
			expLogLevel:          types.LogLevelInfo,
		})
		LogInfof("info: %s: %d", "log", 10)
	})

	t.Run("warn", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "warn",
			expLogLevel:          types.LogLevelWarn,
		})
		LogWarn("warn")
	})

	t.Run("warnf", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "warn: log: 10",
			expLogLevel:          types.LogLevelWarn,
		})
		LogWarnf("warn: %s: %d", "log", 10)
	})

	t.Run("error", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "error",
			expLogLevel:          types.LogLevelError,
		})
		LogError("error")
	})

	t.Run("warnf", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "warn: log: 10",
			expLogLevel:          types.LogLevelWarn,
		})
		LogWarnf("warn: %s: %d", "log", 10)
	})

	t.Run("critical", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "critical error",
			expLogLevel:          types.LogLevelCritical,
		})
		LogCritical("critical error")
	})

	t.Run("criticalf", func(t *testing.T) {
		rawhostcall.RegisterMockWASMHost(logHost{
			DefaultProxyWAMSHost: rawhostcall.DefaultProxyWAMSHost{},
			t:                    t,
			expMessage:           "critical: log: 10",
			expLogLevel:          types.LogLevelCritical,
		})
		LogCriticalf("critical: %s: %d", "log", 10)
	})
}
