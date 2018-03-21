package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/net/proxy"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	socks          string
	host           string
	followRedirect bool
)

var cmdRoot = &cobra.Command{
	Use:   "wassit target",
	Short: "wassit is simple request proxy",
	Long: `A fast and simple request http proxy
                with easy to use configuration options
                created by sbani in Go.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Missing target url")
		}

		// Test target url
		targetURL, err := url.Parse(args[0])
		if err != nil {
			return errors.Wrap(err, "Error parsing url")
		}

		// If tor is specified
		if t, terr := cmd.PersistentFlags().GetBool("tor"); terr == nil && t {
			socks = "127.0.0.1:9150"
		}

		fmt.Fprintln(os.Stdout, "Starting reverse host")
		fmt.Fprintf(os.Stdout, "Server running on %s\n", host)
		fmt.Fprintf(os.Stdout, "Pushing request to %s\n", targetURL)

		if socks != "" {
			fmt.Fprintf(os.Stdout, "Using socks proxy %s\n", socks)
		}

		err = RunServer(targetURL)
		fmt.Fprintf(os.Stderr, "Server stopped due to the following error:\n%s", err)

		return nil
	},
}

func init() {
	cmdRoot.PersistentFlags().StringVarP(&host, "listen", "l", ":9001", "Host and port that the wassit server is listening to")
	cmdRoot.PersistentFlags().StringVarP(&socks, "socks5", "s", "", "Use a socks5 socks for connections to the target")
	cmdRoot.PersistentFlags().BoolP("tor", "t", false, "Enable tor socks5 proxy usage. Please don't forget to start tor")
	cmdRoot.PersistentFlags().BoolVarP(&followRedirect, "follow-redirect", "f", false, "Follow the first redirect (if present) and proxies content")
}

func createHTTPTransport() (*http.Transport, error) {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Use socks proxy to tunnel traffic
	if socks != "" {
		dialer, err := proxy.SOCKS5("tcp", socks, nil, &net.Dialer{})
		if err != nil {
			return nil, errors.Wrap(err, "proxy transport")
		}

		tr.Dial = func(network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		}
	}

	return tr, nil
}

func getRequest(method string, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)")
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient
	t, tErr := createHTTPTransport()
	if tErr != nil {
		fmt.Fprint(os.Stderr, tErr.Error())
		return
	}
	client.Transport = t

	return client.Do(req)
}

// RunServer let's you start the server in foreground
func RunServer(target *url.URL) error {
	targetString := strings.TrimRight(target.String(), "/")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := getRequest(r.Method, fmt.Sprintf("%s%s", targetString, r.RequestURI))
		if err != nil {
			return
		}

		// Handle redirects
		if followRedirect && (response.StatusCode == 302 || response.StatusCode == 301) {
			location := response.Header.Get("Location")

			response, err = getRequest(r.Method, location)
			if err != nil {
				return
			}
		}

		w.WriteHeader(response.StatusCode)
		io.Copy(w, response.Body)
	})

	return http.ListenAndServe(host, nil)
}

// main it baby
func main() {
	if err := cmdRoot.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
