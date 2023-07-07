package arbgasinfo

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../internal/nitro-contracts/src/precompiles/ArbGasInfo.sol --pkg arbgasinfo --sol-version 0.8.17 --filename arbgasinfo
// here we generate some interfaces we use in for our mocks. TODO this should be automated in abigen for all contracts + be condensed
//go:generate go run github.com/vburenin/ifacemaker -f arbgasinfo.abigen.go -s ArbGasInfoCaller -i IArbGasInfoCaller -p arbgasinfo -o icaller_generated.go -c "autogenerated file"
//go:generate go run github.com/vektra/mockery/v2 --name IArbGasInfoCaller --output ./mocks --case=underscore
