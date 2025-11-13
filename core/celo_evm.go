package core

import (
	"github.com/bestyourwallet/celo-geth/common"
	"github.com/bestyourwallet/celo-geth/contracts"
	"github.com/bestyourwallet/celo-geth/core/types"
	"github.com/bestyourwallet/celo-geth/core/vm"
	"github.com/bestyourwallet/celo-geth/log"
	"github.com/bestyourwallet/celo-geth/params"
)

func GetFeeCurrencyContext(header *types.Header, config *params.ChainConfig, statedb vm.StateDB) *common.FeeCurrencyContext {
	if !config.IsCel2(header.Time) {
		return &common.FeeCurrencyContext{}
	}

	caller := &contracts.CeloBackend{ChainConfig: config, State: statedb}

	feeCurrencyContext, err := contracts.GetFeeCurrencyContext(caller)
	if err != nil {
		log.Error("Error fetching exchange rates!", "err", err)
	}
	return &feeCurrencyContext
}

func GetExchangeRates(header *types.Header, config *params.ChainConfig, statedb vm.StateDB) common.ExchangeRates {
	if !config.IsCel2(header.Time) {
		return nil
	}

	caller := &contracts.CeloBackend{ChainConfig: config, State: statedb}

	feeCurrencyContext, err := contracts.GetFeeCurrencyContext(caller)
	if err != nil {
		log.Error("Error fetching exchange rates!", "err", err)
	}
	return feeCurrencyContext.ExchangeRates
}
