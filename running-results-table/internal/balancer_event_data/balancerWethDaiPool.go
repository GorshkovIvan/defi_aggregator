// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancerEvents

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BalancerEventsABI is the input ABI used to generate the binding from.
const BalancerEventsABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LOG_CALL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_EXIT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"name\":\"LOG_JOIN\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denorm\",\"type\":\"uint256\"}],\"name\":\"bind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcInGivenOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcOutGivenIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolInGivenSingleOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolOutGivenSingleIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleInGivenPoolOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleOutGivenPoolIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"}],\"name\":\"exitPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPoolAmountIn\",\"type\":\"uint256\"}],\"name\":\"exitswapExternAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"exitswapPoolAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFinalTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getNormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPriceSansFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"gulp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"isBound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPublicSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"}],\"name\":\"joinPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPoolAmountOut\",\"type\":\"uint256\"}],\"name\":\"joinswapExternAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"joinswapPoolAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denorm\",\"type\":\"uint256\"}],\"name\":\"rebind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"public_\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceAfter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceAfter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unbind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BalancerEvents is an auto generated Go binding around an Ethereum contract.
type BalancerEvents struct {
	BalancerEventsCaller     // Read-only binding to the contract
	BalancerEventsTransactor // Write-only binding to the contract
	BalancerEventsFilterer   // Log filterer for contract events
}

// BalancerEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerEventsSession struct {
	Contract     *BalancerEvents   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerEventsCallerSession struct {
	Contract *BalancerEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BalancerEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerEventsTransactorSession struct {
	Contract     *BalancerEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BalancerEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerEventsRaw struct {
	Contract *BalancerEvents // Generic contract binding to access the raw methods on
}

// BalancerEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerEventsCallerRaw struct {
	Contract *BalancerEventsCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerEventsTransactorRaw struct {
	Contract *BalancerEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerEvents creates a new instance of BalancerEvents, bound to a specific deployed contract.
func NewBalancerEvents(address common.Address, backend bind.ContractBackend) (*BalancerEvents, error) {
	contract, err := bindBalancerEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerEvents{BalancerEventsCaller: BalancerEventsCaller{contract: contract}, BalancerEventsTransactor: BalancerEventsTransactor{contract: contract}, BalancerEventsFilterer: BalancerEventsFilterer{contract: contract}}, nil
}

// NewBalancerEventsCaller creates a new read-only instance of BalancerEvents, bound to a specific deployed contract.
func NewBalancerEventsCaller(address common.Address, caller bind.ContractCaller) (*BalancerEventsCaller, error) {
	contract, err := bindBalancerEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerEventsCaller{contract: contract}, nil
}

// NewBalancerEventsTransactor creates a new write-only instance of BalancerEvents, bound to a specific deployed contract.
func NewBalancerEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerEventsTransactor, error) {
	contract, err := bindBalancerEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerEventsTransactor{contract: contract}, nil
}

// NewBalancerEventsFilterer creates a new log filterer instance of BalancerEvents, bound to a specific deployed contract.
func NewBalancerEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerEventsFilterer, error) {
	contract, err := bindBalancerEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerEventsFilterer{contract: contract}, nil
}

// bindBalancerEvents binds a generic wrapper to an already deployed contract.
func bindBalancerEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerEvents *BalancerEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerEvents.Contract.BalancerEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerEvents *BalancerEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerEvents.Contract.BalancerEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerEvents *BalancerEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerEvents.Contract.BalancerEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerEvents *BalancerEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerEvents *BalancerEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerEvents *BalancerEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerEvents.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) BONE() (*big.Int, error) {
	return _BalancerEvents.Contract.BONE(&_BalancerEvents.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) BONE() (*big.Int, error) {
	return _BalancerEvents.Contract.BONE(&_BalancerEvents.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) BPOWPRECISION() (*big.Int, error) {
	return _BalancerEvents.Contract.BPOWPRECISION(&_BalancerEvents.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _BalancerEvents.Contract.BPOWPRECISION(&_BalancerEvents.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) EXITFEE() (*big.Int, error) {
	return _BalancerEvents.Contract.EXITFEE(&_BalancerEvents.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) EXITFEE() (*big.Int, error) {
	return _BalancerEvents.Contract.EXITFEE(&_BalancerEvents.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BalancerEvents.Contract.INITPOOLSUPPLY(&_BalancerEvents.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BalancerEvents.Contract.INITPOOLSUPPLY(&_BalancerEvents.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXBOUNDTOKENS(&_BalancerEvents.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXBOUNDTOKENS(&_BalancerEvents.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MAXBPOWBASE() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXBPOWBASE(&_BalancerEvents.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXBPOWBASE(&_BalancerEvents.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MAXFEE() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXFEE(&_BalancerEvents.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MAXFEE() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXFEE(&_BalancerEvents.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MAXINRATIO() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXINRATIO(&_BalancerEvents.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MAXINRATIO() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXINRATIO(&_BalancerEvents.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MAXOUTRATIO() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXOUTRATIO(&_BalancerEvents.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXOUTRATIO(&_BalancerEvents.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXTOTALWEIGHT(&_BalancerEvents.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXTOTALWEIGHT(&_BalancerEvents.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MAXWEIGHT() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXWEIGHT(&_BalancerEvents.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _BalancerEvents.Contract.MAXWEIGHT(&_BalancerEvents.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MINBALANCE() (*big.Int, error) {
	return _BalancerEvents.Contract.MINBALANCE(&_BalancerEvents.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MINBALANCE() (*big.Int, error) {
	return _BalancerEvents.Contract.MINBALANCE(&_BalancerEvents.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BalancerEvents.Contract.MINBOUNDTOKENS(&_BalancerEvents.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BalancerEvents.Contract.MINBOUNDTOKENS(&_BalancerEvents.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MINBPOWBASE() (*big.Int, error) {
	return _BalancerEvents.Contract.MINBPOWBASE(&_BalancerEvents.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _BalancerEvents.Contract.MINBPOWBASE(&_BalancerEvents.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MINFEE() (*big.Int, error) {
	return _BalancerEvents.Contract.MINFEE(&_BalancerEvents.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MINFEE() (*big.Int, error) {
	return _BalancerEvents.Contract.MINFEE(&_BalancerEvents.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) MINWEIGHT() (*big.Int, error) {
	return _BalancerEvents.Contract.MINWEIGHT(&_BalancerEvents.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) MINWEIGHT() (*big.Int, error) {
	return _BalancerEvents.Contract.MINWEIGHT(&_BalancerEvents.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) Allowance(opts *bind.CallOpts, src common.Address, dst common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "allowance", src, dst)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.Allowance(&_BalancerEvents.CallOpts, src, dst)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.Allowance(&_BalancerEvents.CallOpts, src, dst)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "balanceOf", whom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.BalanceOf(&_BalancerEvents.CallOpts, whom)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.BalanceOf(&_BalancerEvents.CallOpts, whom)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsCaller) CalcInGivenOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "calcInGivenOut", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcInGivenOut(&_BalancerEvents.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsCallerSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcInGivenOut(&_BalancerEvents.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsCaller) CalcOutGivenIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "calcOutGivenIn", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcOutGivenIn(&_BalancerEvents.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsCallerSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcOutGivenIn(&_BalancerEvents.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BalancerEvents *BalancerEventsCaller) CalcPoolInGivenSingleOut(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "calcPoolInGivenSingleOut", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BalancerEvents *BalancerEventsSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcPoolInGivenSingleOut(&_BalancerEvents.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BalancerEvents *BalancerEventsCallerSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcPoolInGivenSingleOut(&_BalancerEvents.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BalancerEvents *BalancerEventsCaller) CalcPoolOutGivenSingleIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "calcPoolOutGivenSingleIn", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BalancerEvents *BalancerEventsSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcPoolOutGivenSingleIn(&_BalancerEvents.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BalancerEvents *BalancerEventsCallerSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcPoolOutGivenSingleIn(&_BalancerEvents.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsCaller) CalcSingleInGivenPoolOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "calcSingleInGivenPoolOut", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcSingleInGivenPoolOut(&_BalancerEvents.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsCallerSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcSingleInGivenPoolOut(&_BalancerEvents.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsCaller) CalcSingleOutGivenPoolIn(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "calcSingleOutGivenPoolIn", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcSingleOutGivenPoolIn(&_BalancerEvents.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsCallerSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcSingleOutGivenPoolIn(&_BalancerEvents.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BalancerEvents *BalancerEventsCaller) CalcSpotPrice(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "calcSpotPrice", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BalancerEvents *BalancerEventsSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcSpotPrice(&_BalancerEvents.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BalancerEvents *BalancerEventsCallerSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerEvents.Contract.CalcSpotPrice(&_BalancerEvents.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerEvents *BalancerEventsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerEvents *BalancerEventsSession) Decimals() (uint8, error) {
	return _BalancerEvents.Contract.Decimals(&_BalancerEvents.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerEvents *BalancerEventsCallerSession) Decimals() (uint8, error) {
	return _BalancerEvents.Contract.Decimals(&_BalancerEvents.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.GetBalance(&_BalancerEvents.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.GetBalance(&_BalancerEvents.CallOpts, token)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BalancerEvents *BalancerEventsCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BalancerEvents *BalancerEventsSession) GetColor() ([32]byte, error) {
	return _BalancerEvents.Contract.GetColor(&_BalancerEvents.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BalancerEvents *BalancerEventsCallerSession) GetColor() ([32]byte, error) {
	return _BalancerEvents.Contract.GetColor(&_BalancerEvents.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_BalancerEvents *BalancerEventsCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_BalancerEvents *BalancerEventsSession) GetController() (common.Address, error) {
	return _BalancerEvents.Contract.GetController(&_BalancerEvents.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_BalancerEvents *BalancerEventsCallerSession) GetController() (common.Address, error) {
	return _BalancerEvents.Contract.GetController(&_BalancerEvents.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_BalancerEvents *BalancerEventsCaller) GetCurrentTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getCurrentTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_BalancerEvents *BalancerEventsSession) GetCurrentTokens() ([]common.Address, error) {
	return _BalancerEvents.Contract.GetCurrentTokens(&_BalancerEvents.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_BalancerEvents *BalancerEventsCallerSession) GetCurrentTokens() ([]common.Address, error) {
	return _BalancerEvents.Contract.GetCurrentTokens(&_BalancerEvents.CallOpts)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) GetDenormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getDenormalizedWeight", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.GetDenormalizedWeight(&_BalancerEvents.CallOpts, token)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.GetDenormalizedWeight(&_BalancerEvents.CallOpts, token)
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_BalancerEvents *BalancerEventsCaller) GetFinalTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getFinalTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_BalancerEvents *BalancerEventsSession) GetFinalTokens() ([]common.Address, error) {
	return _BalancerEvents.Contract.GetFinalTokens(&_BalancerEvents.CallOpts)
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_BalancerEvents *BalancerEventsCallerSession) GetFinalTokens() ([]common.Address, error) {
	return _BalancerEvents.Contract.GetFinalTokens(&_BalancerEvents.CallOpts)
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) GetNormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getNormalizedWeight", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) GetNormalizedWeight(token common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.GetNormalizedWeight(&_BalancerEvents.CallOpts, token)
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) GetNormalizedWeight(token common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.GetNormalizedWeight(&_BalancerEvents.CallOpts, token)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) GetNumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getNumTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) GetNumTokens() (*big.Int, error) {
	return _BalancerEvents.Contract.GetNumTokens(&_BalancerEvents.CallOpts)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) GetNumTokens() (*big.Int, error) {
	return _BalancerEvents.Contract.GetNumTokens(&_BalancerEvents.CallOpts)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerEvents *BalancerEventsCaller) GetSpotPrice(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getSpotPrice", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerEvents *BalancerEventsSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.GetSpotPrice(&_BalancerEvents.CallOpts, tokenIn, tokenOut)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerEvents *BalancerEventsCallerSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.GetSpotPrice(&_BalancerEvents.CallOpts, tokenIn, tokenOut)
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerEvents *BalancerEventsCaller) GetSpotPriceSansFee(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getSpotPriceSansFee", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerEvents *BalancerEventsSession) GetSpotPriceSansFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.GetSpotPriceSansFee(&_BalancerEvents.CallOpts, tokenIn, tokenOut)
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerEvents *BalancerEventsCallerSession) GetSpotPriceSansFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BalancerEvents.Contract.GetSpotPriceSansFee(&_BalancerEvents.CallOpts, tokenIn, tokenOut)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) GetSwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getSwapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) GetSwapFee() (*big.Int, error) {
	return _BalancerEvents.Contract.GetSwapFee(&_BalancerEvents.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) GetSwapFee() (*big.Int, error) {
	return _BalancerEvents.Contract.GetSwapFee(&_BalancerEvents.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) GetTotalDenormalizedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "getTotalDenormalizedWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _BalancerEvents.Contract.GetTotalDenormalizedWeight(&_BalancerEvents.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _BalancerEvents.Contract.GetTotalDenormalizedWeight(&_BalancerEvents.CallOpts)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_BalancerEvents *BalancerEventsCaller) IsBound(opts *bind.CallOpts, t common.Address) (bool, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "isBound", t)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_BalancerEvents *BalancerEventsSession) IsBound(t common.Address) (bool, error) {
	return _BalancerEvents.Contract.IsBound(&_BalancerEvents.CallOpts, t)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_BalancerEvents *BalancerEventsCallerSession) IsBound(t common.Address) (bool, error) {
	return _BalancerEvents.Contract.IsBound(&_BalancerEvents.CallOpts, t)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_BalancerEvents *BalancerEventsCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_BalancerEvents *BalancerEventsSession) IsFinalized() (bool, error) {
	return _BalancerEvents.Contract.IsFinalized(&_BalancerEvents.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_BalancerEvents *BalancerEventsCallerSession) IsFinalized() (bool, error) {
	return _BalancerEvents.Contract.IsFinalized(&_BalancerEvents.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_BalancerEvents *BalancerEventsCaller) IsPublicSwap(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "isPublicSwap")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_BalancerEvents *BalancerEventsSession) IsPublicSwap() (bool, error) {
	return _BalancerEvents.Contract.IsPublicSwap(&_BalancerEvents.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_BalancerEvents *BalancerEventsCallerSession) IsPublicSwap() (bool, error) {
	return _BalancerEvents.Contract.IsPublicSwap(&_BalancerEvents.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerEvents *BalancerEventsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerEvents *BalancerEventsSession) Name() (string, error) {
	return _BalancerEvents.Contract.Name(&_BalancerEvents.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerEvents *BalancerEventsCallerSession) Name() (string, error) {
	return _BalancerEvents.Contract.Name(&_BalancerEvents.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerEvents *BalancerEventsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerEvents *BalancerEventsSession) Symbol() (string, error) {
	return _BalancerEvents.Contract.Symbol(&_BalancerEvents.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerEvents *BalancerEventsCallerSession) Symbol() (string, error) {
	return _BalancerEvents.Contract.Symbol(&_BalancerEvents.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerEvents *BalancerEventsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerEvents.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerEvents *BalancerEventsSession) TotalSupply() (*big.Int, error) {
	return _BalancerEvents.Contract.TotalSupply(&_BalancerEvents.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerEvents *BalancerEventsCallerSession) TotalSupply() (*big.Int, error) {
	return _BalancerEvents.Contract.TotalSupply(&_BalancerEvents.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsTransactor) Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "approve", dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Approve(&_BalancerEvents.TransactOpts, dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsTransactorSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Approve(&_BalancerEvents.TransactOpts, dst, amt)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerEvents *BalancerEventsTransactor) Bind(opts *bind.TransactOpts, token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "bind", token, balance, denorm)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerEvents *BalancerEventsSession) Bind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Bind(&_BalancerEvents.TransactOpts, token, balance, denorm)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerEvents *BalancerEventsTransactorSession) Bind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Bind(&_BalancerEvents.TransactOpts, token, balance, denorm)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsTransactor) DecreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "decreaseApproval", dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.DecreaseApproval(&_BalancerEvents.TransactOpts, dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsTransactorSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.DecreaseApproval(&_BalancerEvents.TransactOpts, dst, amt)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_BalancerEvents *BalancerEventsTransactor) ExitPool(opts *bind.TransactOpts, poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "exitPool", poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_BalancerEvents *BalancerEventsSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.ExitPool(&_BalancerEvents.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_BalancerEvents *BalancerEventsTransactorSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.ExitPool(&_BalancerEvents.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_BalancerEvents *BalancerEventsTransactor) ExitswapExternAmountOut(opts *bind.TransactOpts, tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "exitswapExternAmountOut", tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_BalancerEvents *BalancerEventsSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.ExitswapExternAmountOut(&_BalancerEvents.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_BalancerEvents *BalancerEventsTransactorSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.ExitswapExternAmountOut(&_BalancerEvents.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsTransactor) ExitswapPoolAmountIn(opts *bind.TransactOpts, tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "exitswapPoolAmountIn", tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.ExitswapPoolAmountIn(&_BalancerEvents.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsTransactorSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.ExitswapPoolAmountIn(&_BalancerEvents.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BalancerEvents *BalancerEventsTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BalancerEvents *BalancerEventsSession) Finalize() (*types.Transaction, error) {
	return _BalancerEvents.Contract.Finalize(&_BalancerEvents.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BalancerEvents *BalancerEventsTransactorSession) Finalize() (*types.Transaction, error) {
	return _BalancerEvents.Contract.Finalize(&_BalancerEvents.TransactOpts)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_BalancerEvents *BalancerEventsTransactor) Gulp(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "gulp", token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_BalancerEvents *BalancerEventsSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Gulp(&_BalancerEvents.TransactOpts, token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_BalancerEvents *BalancerEventsTransactorSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Gulp(&_BalancerEvents.TransactOpts, token)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsTransactor) IncreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "increaseApproval", dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.IncreaseApproval(&_BalancerEvents.TransactOpts, dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsTransactorSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.IncreaseApproval(&_BalancerEvents.TransactOpts, dst, amt)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_BalancerEvents *BalancerEventsTransactor) JoinPool(opts *bind.TransactOpts, poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "joinPool", poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_BalancerEvents *BalancerEventsSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.JoinPool(&_BalancerEvents.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_BalancerEvents *BalancerEventsTransactorSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.JoinPool(&_BalancerEvents.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_BalancerEvents *BalancerEventsTransactor) JoinswapExternAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "joinswapExternAmountIn", tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_BalancerEvents *BalancerEventsSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.JoinswapExternAmountIn(&_BalancerEvents.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_BalancerEvents *BalancerEventsTransactorSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.JoinswapExternAmountIn(&_BalancerEvents.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsTransactor) JoinswapPoolAmountOut(opts *bind.TransactOpts, tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "joinswapPoolAmountOut", tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.JoinswapPoolAmountOut(&_BalancerEvents.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsTransactorSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.JoinswapPoolAmountOut(&_BalancerEvents.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerEvents *BalancerEventsTransactor) Rebind(opts *bind.TransactOpts, token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "rebind", token, balance, denorm)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerEvents *BalancerEventsSession) Rebind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Rebind(&_BalancerEvents.TransactOpts, token, balance, denorm)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerEvents *BalancerEventsTransactorSession) Rebind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Rebind(&_BalancerEvents.TransactOpts, token, balance, denorm)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_BalancerEvents *BalancerEventsTransactor) SetController(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "setController", manager)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_BalancerEvents *BalancerEventsSession) SetController(manager common.Address) (*types.Transaction, error) {
	return _BalancerEvents.Contract.SetController(&_BalancerEvents.TransactOpts, manager)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_BalancerEvents *BalancerEventsTransactorSession) SetController(manager common.Address) (*types.Transaction, error) {
	return _BalancerEvents.Contract.SetController(&_BalancerEvents.TransactOpts, manager)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_BalancerEvents *BalancerEventsTransactor) SetPublicSwap(opts *bind.TransactOpts, public_ bool) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "setPublicSwap", public_)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_BalancerEvents *BalancerEventsSession) SetPublicSwap(public_ bool) (*types.Transaction, error) {
	return _BalancerEvents.Contract.SetPublicSwap(&_BalancerEvents.TransactOpts, public_)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_BalancerEvents *BalancerEventsTransactorSession) SetPublicSwap(public_ bool) (*types.Transaction, error) {
	return _BalancerEvents.Contract.SetPublicSwap(&_BalancerEvents.TransactOpts, public_)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_BalancerEvents *BalancerEventsTransactor) SetSwapFee(opts *bind.TransactOpts, swapFee *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "setSwapFee", swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_BalancerEvents *BalancerEventsSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.SetSwapFee(&_BalancerEvents.TransactOpts, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_BalancerEvents *BalancerEventsTransactorSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.SetSwapFee(&_BalancerEvents.TransactOpts, swapFee)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_BalancerEvents *BalancerEventsTransactor) SwapExactAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "swapExactAmountIn", tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_BalancerEvents *BalancerEventsSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.SwapExactAmountIn(&_BalancerEvents.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_BalancerEvents *BalancerEventsTransactorSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.SwapExactAmountIn(&_BalancerEvents.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_BalancerEvents *BalancerEventsTransactor) SwapExactAmountOut(opts *bind.TransactOpts, tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "swapExactAmountOut", tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_BalancerEvents *BalancerEventsSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.SwapExactAmountOut(&_BalancerEvents.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_BalancerEvents *BalancerEventsTransactorSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.SwapExactAmountOut(&_BalancerEvents.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "transfer", dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Transfer(&_BalancerEvents.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsTransactorSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Transfer(&_BalancerEvents.TransactOpts, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "transferFrom", src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.TransferFrom(&_BalancerEvents.TransactOpts, src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BalancerEvents *BalancerEventsTransactorSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerEvents.Contract.TransferFrom(&_BalancerEvents.TransactOpts, src, dst, amt)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_BalancerEvents *BalancerEventsTransactor) Unbind(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BalancerEvents.contract.Transact(opts, "unbind", token)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_BalancerEvents *BalancerEventsSession) Unbind(token common.Address) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Unbind(&_BalancerEvents.TransactOpts, token)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_BalancerEvents *BalancerEventsTransactorSession) Unbind(token common.Address) (*types.Transaction, error) {
	return _BalancerEvents.Contract.Unbind(&_BalancerEvents.TransactOpts, token)
}

// BalancerEventsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BalancerEvents contract.
type BalancerEventsApprovalIterator struct {
	Event *BalancerEventsApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BalancerEventsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerEventsApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BalancerEventsApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BalancerEventsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerEventsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerEventsApproval represents a Approval event raised by the BalancerEvents contract.
type BalancerEventsApproval struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BalancerEvents *BalancerEventsFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BalancerEventsApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BalancerEvents.contract.FilterLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BalancerEventsApprovalIterator{contract: _BalancerEvents.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BalancerEvents *BalancerEventsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BalancerEventsApproval, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BalancerEvents.contract.WatchLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerEventsApproval)
				if err := _BalancerEvents.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BalancerEvents *BalancerEventsFilterer) ParseApproval(log types.Log) (*BalancerEventsApproval, error) {
	event := new(BalancerEventsApproval)
	if err := _BalancerEvents.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerEventsLOGEXITIterator is returned from FilterLOGEXIT and is used to iterate over the raw logs and unpacked data for LOGEXIT events raised by the BalancerEvents contract.
type BalancerEventsLOGEXITIterator struct {
	Event *BalancerEventsLOGEXIT // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BalancerEventsLOGEXITIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerEventsLOGEXIT)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BalancerEventsLOGEXIT)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BalancerEventsLOGEXITIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerEventsLOGEXITIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerEventsLOGEXIT represents a LOGEXIT event raised by the BalancerEvents contract.
type BalancerEventsLOGEXIT struct {
	Caller         common.Address
	TokenOut       common.Address
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGEXIT is a free log retrieval operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsFilterer) FilterLOGEXIT(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address) (*BalancerEventsLOGEXITIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BalancerEvents.contract.FilterLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BalancerEventsLOGEXITIterator{contract: _BalancerEvents.contract, event: "LOG_EXIT", logs: logs, sub: sub}, nil
}

// WatchLOGEXIT is a free log subscription operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsFilterer) WatchLOGEXIT(opts *bind.WatchOpts, sink chan<- *BalancerEventsLOGEXIT, caller []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BalancerEvents.contract.WatchLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerEventsLOGEXIT)
				if err := _BalancerEvents.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLOGEXIT is a log parse operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsFilterer) ParseLOGEXIT(log types.Log) (*BalancerEventsLOGEXIT, error) {
	event := new(BalancerEventsLOGEXIT)
	if err := _BalancerEvents.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerEventsLOGJOINIterator is returned from FilterLOGJOIN and is used to iterate over the raw logs and unpacked data for LOGJOIN events raised by the BalancerEvents contract.
type BalancerEventsLOGJOINIterator struct {
	Event *BalancerEventsLOGJOIN // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BalancerEventsLOGJOINIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerEventsLOGJOIN)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BalancerEventsLOGJOIN)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BalancerEventsLOGJOINIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerEventsLOGJOINIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerEventsLOGJOIN represents a LOGJOIN event raised by the BalancerEvents contract.
type BalancerEventsLOGJOIN struct {
	Caller        common.Address
	TokenIn       common.Address
	TokenAmountIn *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGJOIN is a free log retrieval operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsFilterer) FilterLOGJOIN(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address) (*BalancerEventsLOGJOINIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _BalancerEvents.contract.FilterLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &BalancerEventsLOGJOINIterator{contract: _BalancerEvents.contract, event: "LOG_JOIN", logs: logs, sub: sub}, nil
}

// WatchLOGJOIN is a free log subscription operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsFilterer) WatchLOGJOIN(opts *bind.WatchOpts, sink chan<- *BalancerEventsLOGJOIN, caller []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _BalancerEvents.contract.WatchLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerEventsLOGJOIN)
				if err := _BalancerEvents.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLOGJOIN is a log parse operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_BalancerEvents *BalancerEventsFilterer) ParseLOGJOIN(log types.Log) (*BalancerEventsLOGJOIN, error) {
	event := new(BalancerEventsLOGJOIN)
	if err := _BalancerEvents.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerEventsLOGSWAPIterator is returned from FilterLOGSWAP and is used to iterate over the raw logs and unpacked data for LOGSWAP events raised by the BalancerEvents contract.
type BalancerEventsLOGSWAPIterator struct {
	Event *BalancerEventsLOGSWAP // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BalancerEventsLOGSWAPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerEventsLOGSWAP)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BalancerEventsLOGSWAP)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BalancerEventsLOGSWAPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerEventsLOGSWAPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerEventsLOGSWAP represents a LOGSWAP event raised by the BalancerEvents contract.
type BalancerEventsLOGSWAP struct {
	Caller         common.Address
	TokenIn        common.Address
	TokenOut       common.Address
	TokenAmountIn  *big.Int
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGSWAP is a free log retrieval operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsFilterer) FilterLOGSWAP(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*BalancerEventsLOGSWAPIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BalancerEvents.contract.FilterLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BalancerEventsLOGSWAPIterator{contract: _BalancerEvents.contract, event: "LOG_SWAP", logs: logs, sub: sub}, nil
}

// WatchLOGSWAP is a free log subscription operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsFilterer) WatchLOGSWAP(opts *bind.WatchOpts, sink chan<- *BalancerEventsLOGSWAP, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BalancerEvents.contract.WatchLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerEventsLOGSWAP)
				if err := _BalancerEvents.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLOGSWAP is a log parse operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_BalancerEvents *BalancerEventsFilterer) ParseLOGSWAP(log types.Log) (*BalancerEventsLOGSWAP, error) {
	event := new(BalancerEventsLOGSWAP)
	if err := _BalancerEvents.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerEventsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BalancerEvents contract.
type BalancerEventsTransferIterator struct {
	Event *BalancerEventsTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BalancerEventsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerEventsTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BalancerEventsTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BalancerEventsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerEventsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerEventsTransfer represents a Transfer event raised by the BalancerEvents contract.
type BalancerEventsTransfer struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BalancerEvents *BalancerEventsFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BalancerEventsTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BalancerEvents.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BalancerEventsTransferIterator{contract: _BalancerEvents.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BalancerEvents *BalancerEventsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BalancerEventsTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BalancerEvents.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerEventsTransfer)
				if err := _BalancerEvents.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BalancerEvents *BalancerEventsFilterer) ParseTransfer(log types.Log) (*BalancerEventsTransfer, error) {
	event := new(BalancerEventsTransfer)
	if err := _BalancerEvents.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
