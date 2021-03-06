// This code is available on the terms of the project LICENSE.md file,
// also available online at https://blueoakcouncil.org/license/1.0.0.

package btc

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"decred.org/dcrdex/dex"
	"github.com/btcsuite/btcutil"
	flags "github.com/jessevdk/go-flags"
)

// NetPorts are a set of port to use with the different networks.
type NetPorts struct {
	Mainnet string
	Testnet string
	Simnet  string
}

// RPCPorts are the default BTC ports.
var RPCPorts = NetPorts{
	Mainnet: "8332",
	Testnet: "18332",
	Simnet:  "18443",
}

const (
	defaultHost = "localhost"
)

// Config is a partial bitcoind configuration file, with only the parameters
// needed to initialize the RPC connection. Because we have no need to read
// command line arguments, we can use the IgnoreUnknown flag to allow pulling
// the needed confiruration settings directly from the bitcoin.conf, if the user
// chooses.
type Config struct {
	RPCUser string `long:"rpcuser" description:"JSON-RPC user"`
	RPCPass string `long:"rpcpassword" description:"JSON-RPC password"`
	RPCBind string `long:"rpcbind" description:"RPC address. Can be <addr> or <addr>:<port>, which would override rpcport"`
	RPCPort int    `long:"rpcport" description:"JSON-RPC port"`
}

// LoadConfig loads the configuration settings from the specified filepath.
func LoadConfig(configPath string, name string, network dex.Network, ports NetPorts) (*Config, error) {
	cfg := &Config{}
	// Since we are not reading command-line arguments, and the Config fields
	// share names with the bitcoind configuration options, passing just
	// IgnoreUnknown allows us to have the option to read directly from the
	// bitcoin.conf file.
	parser := flags.NewParser(cfg, flags.IgnoreUnknown)

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("no %q config file found at %s", name, configPath)
	}
	// The config file exists, so attempt to parse it.
	err := flags.NewIniParser(parser).ParseFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error parsing %q ini file %q: %v", name, configPath, err)
	}

	if cfg.RPCUser == "" {
		return nil, fmt.Errorf("no rpcuser set in %q config file", name)
	}
	if cfg.RPCPass == "" {
		return nil, fmt.Errorf("no rpcpassword set in %q config file", name)
	}

	host := defaultHost
	var port string
	switch network {
	case dex.Mainnet:
		port = ports.Mainnet
	case dex.Testnet:
		port = ports.Testnet
	case dex.Regtest:
		port = ports.Simnet
	default:
		return nil, fmt.Errorf("unknown network ID %v", network)
	}

	// RPCPort overrides network default
	if cfg.RPCPort != 0 {
		port = strconv.Itoa(cfg.RPCPort)
	}

	// if RPCBind includes a port, it takes precedence over RPCPort
	if cfg.RPCBind != "" {
		h, p, err := net.SplitHostPort(cfg.RPCBind)
		if err != nil {
			// Will error for i.e. "localhost", but not for "localhost:" or ":1234"
			host = cfg.RPCBind
		} else {
			if h != "" {
				host = h
			}
			if p != "" {
				port = p
			}
		}
	}

	// overwrite rpcbind to use for rpcclient connection
	cfg.RPCBind = net.JoinHostPort(host, port)

	return cfg, nil
}

// SystemConfigPath will return the default config file path for bitcoin-like
// assets.
func SystemConfigPath(asset string) string {
	homeDir := btcutil.AppDataDir(asset, false)
	return filepath.Join(homeDir, fmt.Sprintf("%s.conf", asset))
}
