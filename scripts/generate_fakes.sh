#!/usr/bin/env bash

set -e

go get github.com/maxbrunsfeld/counterfeiter

pushd src/config_server/store
    counterfeiter config_server/store.IDb
    counterfeiter config_server/store.DbProvider
    counterfeiter config_server/store.IRow
    counterfeiter config_server/store.IRows
    counterfeiter config_server/store.ISql
    counterfeiter config_server/store.Store
    counterfeiter database/sql.Result
popd

pushd src/config_server/types
    counterfeiter config_server/types.CertsLoader
    counterfeiter config_server/types.ValueGenerator
    counterfeiter config_server/types.ValueGeneratorFactory
popd

pushd src/config_server/server
    counterfeiter net/http.Handler
    counterfeiter config_server/server.TokenValidator
popd