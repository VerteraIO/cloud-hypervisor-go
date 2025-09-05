// Package unixhttp provides utilities for communicating with Cloud Hypervisor
// over Unix domain sockets using HTTP.
package unixhttp

import (
	"context"
	"net"
	"net/http"
	"time"
)

// NewClient returns an *http.Client configured to communicate over a Unix domain socket.
// Use this with the generated client like:
//
//	httpc := unixhttp.NewClient("/run/cloud-hypervisor.sock")
//	client, err := chclient.NewClient(
//		chclient.WithServer("http://unix/api/v1"),
//		chclient.WithHTTPClient(httpc),
//	)
//
// The server URL host ("unix") is ignored by the transport since all requests
// are routed to the specified socket path.
func NewClient(socketPath string) *http.Client {
	dialer := &net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   10,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			// Ignore the network and addr parameters and always dial the Unix socket
			return dialer.DialContext(ctx, "unix", socketPath)
		},
	}

	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
}

// NewClientWithTimeout returns a Unix socket HTTP client with a custom timeout.
func NewClientWithTimeout(socketPath string, timeout time.Duration) *http.Client {
	client := NewClient(socketPath)
	client.Timeout = timeout
	return client
}
