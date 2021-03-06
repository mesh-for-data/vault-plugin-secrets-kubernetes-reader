package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	kubesecrets "github.com/hashicorp/vault-plugin-kubernetes-secrets"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/plugin"
)

func main() {
	// Boilerplate code to get started.
	// Please see https://www.hashicorp.com/blog/building-a-vault-secure-plugin for more info.
	apiClientMeta := &api.PluginAPIClientMeta{}
	flags := apiClientMeta.FlagSet()
	err := flags.Parse(os.Args[1:])
	if err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("Error in flags.Parse", err)
		os.Exit(1)
	}

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	err = plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: kubesecrets.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	})
	if err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}
