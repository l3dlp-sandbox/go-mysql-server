// Copyright 2020-2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/dolthub/vitess/go/mysql"
	"go.opentelemetry.io/otel/trace"

	gms "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/sql"
)

// Server is a MySQL server for SQLe engines.
type Server struct {
	Listener   ProtocolListener
	handler    mysql.Handler
	sessionMgr *SessionManager
	Engine     *gms.Engine
}

// An option to customize the server.
type Option func(e *gms.Engine, sm *SessionManager, handler mysql.Handler) (*gms.Engine, *SessionManager, mysql.Handler)

// Config for the mysql server.
type Config struct {
	// Protocol for the connection.
	Protocol string
	// Address of the server.
	Address string
	// Custom listener for the mysql server. Use this if you don't want ports or unix sockets to be opened automatically.
	// This can be useful in testing by using a pure go net.Conn implementation.
	Listener net.Listener
	// Tracer to use in the server. By default, a noop tracer will be used if
	// no tracer is provided.
	Tracer trace.Tracer
	// Version string to advertise in running server
	Version string
	// ConnReadTimeout is the server's read timeout
	ConnReadTimeout time.Duration
	// ConnWriteTimeout is the server's write timeout
	ConnWriteTimeout time.Duration
	// MaxConnections is the maximum number of simultaneous connections that the server will allow.
	MaxConnections uint64
	// MaxWaitConnections is the maximum number of simultaneous connections that the server will allow to block waiting
	// for a connection before new connections result in immediate rejection.
	MaxWaitConnections uint32
	// MaxWaitConnectionsTimeout is the maximum amount of time that a connection will block waiting for a connection
	MaxWaitConnectionsTimeout time.Duration
	// TLSConfig is the configuration for TLS on this server. If |nil|, TLS is not supported.
	TLSConfig *tls.Config
	// RequestSecureTransport will require incoming connections to be TLS. Requires non-|nil| TLSConfig.
	RequireSecureTransport bool
	// DisableClientMultiStatements will prevent processing of incoming
	// queries as if they contain more than one query. This processing
	// currently works in some simple cases, but breaks in the presence of
	// statements (such as in CREATE TRIGGER queries). Configuring the
	// server to disable processing these is one option for users to get
	// support back for single queries that contain statements, at the cost
	// of not supporting the CLIENT_MULTI_STATEMENTS option on the server.
	DisableClientMultiStatements bool
	// NoDefaults prevents using persisted configuration for new server sessions
	NoDefaults bool
	// Socket is a path to unix socket file
	Socket                   string
	AllowClearTextWithoutTLS bool
	// MaxLoggedQueryLen sets the length at which queries written to the logs are truncated.  A value of 0 will
	// result in no truncation. A value less than 0 will result in the queries being omitted from the logs completely
	MaxLoggedQueryLen int
	// EncodeLoggedQuery determines if logged queries are base64 encoded.
	// If true, queries will be logged as base64 encoded strings.
	// If false (default behavior), queries will be logged as strings, but newlines and tabs will be replaced with spaces.
	EncodeLoggedQuery bool
	// Options gets a chance to visit and mutate the GMS *Engine,
	// *server.SessionManager and the mysql.Handler as the server
	// is being initialized, before the ProtocolListener is
	// constructed.
	Options []Option
	// Used to get the ProtocolListener on server start.
	// If unset, defaults to MySQLProtocolListenerFactory.
	ProtocolListenerFactory ProtocolListenerFunc
}

func (c Config) NewConfig() (Config, error) {
	if _, val, ok := sql.SystemVariables.GetGlobal("max_connections"); ok {
		mc, ok := val.(int64)
		if !ok {
			return Config{}, sql.ErrUnknownSystemVariable.New("max_connections")
		}
		c.MaxConnections = uint64(mc)
	}
	if _, val, ok := sql.SystemVariables.GetGlobal("net_write_timeout"); ok {
		timeout, ok := val.(int64)
		if !ok {
			return Config{}, sql.ErrUnknownSystemVariable.New("net_write_timeout")
		}
		c.ConnWriteTimeout = time.Duration(timeout) * time.Millisecond
	}
	if _, val, ok := sql.SystemVariables.GetGlobal("net_read_timeout"); ok {
		timeout, ok := val.(int64)
		if !ok {
			return Config{}, sql.ErrUnknownSystemVariable.New("net_read_timeout")
		}
		c.ConnReadTimeout = time.Duration(timeout) * time.Millisecond
	}
	return c, nil
}
