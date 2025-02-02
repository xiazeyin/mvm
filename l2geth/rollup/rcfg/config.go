package rcfg

import (
	"math/big"
	"os"
	"strconv"
	"github.com/ethereum-optimism/optimism/l2geth/common"
)

// UsingOVM is used to enable or disable functionality necessary for the OVM.
var (
	UsingOVM               bool
	SuicideForkNumber      uint64
	PeerHealthCheckSeconds int64
	ChainID                uint64
)

var (
	// l2GasPriceSlot refers to the storage slot that the L2 gas price is stored
	// in in the OVM_GasPriceOracle predeploy
	L2GasPriceSlot = common.BigToHash(big.NewInt(1))
	// l1GasPriceSlot refers to the storage slot that the L1 gas price is stored
	// in in the OVM_GasPriceOracle predeploy
	L1GasPriceSlot = common.BigToHash(big.NewInt(2))
	// l2GasPriceOracleOwnerSlot refers to the storage slot that the owner of
	// the OVM_GasPriceOracle is stored in
	L2GasPriceOracleOwnerSlot = common.BigToHash(big.NewInt(0))
	// l2GasPriceOracleAddress is the address of the OVM_GasPriceOracle
	// predeploy
	L2GasPriceOracleAddress = common.HexToAddress("0x420000000000000000000000000000000000000F")
	// OverheadSlot refers to the storage slot in the OVM_GasPriceOracle that
	// holds the per transaction overhead. This is added to the L1 cost portion
	// of the fee
	OverheadSlot = common.BigToHash(big.NewInt(3))
	// ScalarSlot refers to the storage slot in the OVM_GasPriceOracle that
	// holds the transaction fee scalar. This value is scaled upwards by
	// the number of decimals
	ScalarSlot = common.BigToHash(big.NewInt(4))
	// DecimalsSlot refers to the storage slot in the OVM_GasPriceOracle that
	// holds the number of decimals in the fee scalar
	DecimalsSlot = common.BigToHash(big.NewInt(5))
)

func init() {
	UsingOVM = os.Getenv("USING_OVM") == "true"

	suicideForkNumber := os.Getenv("EMERGENCY_FORK020222_NUMBER")
	if suicideForkNumber == "" {
		SuicideForkNumber = ^uint64(0)
	} else {
		parsed, err := strconv.ParseUint(suicideForkNumber, 0, 64)
		if err != nil {
			panic(err)
		}
		SuicideForkNumber = parsed
	}

	peerHealthCheck := os.Getenv("PEER_HEALTH_CHECK")
	if peerHealthCheck == "" {
		PeerHealthCheckSeconds = ^int64(0)
	} else {
		parsed, err := strconv.ParseInt(peerHealthCheck, 10, 64)
		if err != nil {
			panic(err)
		}
		PeerHealthCheckSeconds = parsed
	}

	envChainID := os.Getenv("CHAIN_ID")
	if envChainID == "" {
		ChainID = ^uint64(0)
	} else {
		parsed, err := strconv.ParseUint(envChainID, 0, 64)
		if err != nil {
			panic(err)
		}
		ChainID = parsed
	}
}
