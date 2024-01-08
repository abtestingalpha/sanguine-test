package relapi

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/service"
)

// Handler is the REST API handler.
type Handler struct {
	db     reldb.Service
	chains map[uint32]*service.Chain
}

// NewHandler creates a new REST API handler.
func NewHandler(db reldb.Service, chains map[uint32]*service.Chain) *Handler {
	return &Handler{
		db:     db, // Store the database connection in the handler
		chains: chains,
	}
}

// GetQuoteRequestStatusByTxID gets the status of a quote request, given a tx id.
func (h *Handler) GetQuoteRequestStatusByTxID(c *gin.Context) {
	txIDStr := c.Query("txID")
	if txIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must specify txID"})
		return
	}

	txIDBytes, err := hexutil.Decode(txIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid txID"})
		return
	}
	var txID [32]byte
	copy(txID[:], txIDBytes)

	quoteRequest, err := h.db.GetQuoteRequestByID(c, txID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := GetQuoteRequestStatusResponse{
		Status: quoteRequest.Status.String(),
		TxID:   hexutil.Encode(quoteRequest.TransactionID[:]),
		TxHash: quoteRequest.DestTxHash.String(),
	}
	c.JSON(http.StatusOK, resp)
}

// GetQuoteRequestStatusByTxHash gets the status of a quote request, given a tx hash.
func (h *Handler) GetQuoteRequestStatusByTxHash(c *gin.Context) {
	txHashStr := c.Query("hash")
	if txHashStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must specify hash"})
		return
	}

	txHash := common.HexToHash(txHashStr)
	quoteRequest, err := h.db.GetQuoteRequestByTxHash(c, txHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := GetQuoteRequestStatusResponse{
		Status: quoteRequest.Status.String(),
		TxID:   hexutil.Encode(quoteRequest.TransactionID[:]),
		TxHash: quoteRequest.DestTxHash.String(),
	}
	c.JSON(http.StatusOK, resp)
}

// PutTxRetry retries a transaction based on tx hash.
func (h *Handler) PutTxRetry(c *gin.Context) {
	txHashStr := c.Query("hash")
	if txHashStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Must specify hash"})
		return
	}

	txHash := common.HexToHash(txHashStr)
	quoteRequest, err := h.db.GetQuoteRequestByTxHash(c, txHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	chainID := quoteRequest.Transaction.DestChainId
	chain, ok := h.chains[uint32(chainID)]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No contract found for chain: %d", chainID)})
		return
	}

	// TODO: this can be deduped with handlers.go code
	gasAmount := big.NewInt(0)
	if quoteRequest.Transaction.SendChainGas {
		gasAmount, err = chain.Bridge.ChainGasAmount(&bind.CallOpts{Context: c})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("could not get chain gas amount: %s", err.Error())})
			return
		}
	}
	nonce, err := chain.SubmitTransaction(c, func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		transactor.Value = core.CopyBigInt(gasAmount)

		tx, err = chain.Bridge.Relay(transactor, quoteRequest.RawRequest)
		if err != nil {
			return nil, fmt.Errorf("could not relay: %w", err)
		}

		return tx, nil
	})

	resp := PutTxRetryResponse{
		TxID:      hexutil.Encode(quoteRequest.TransactionID[:]),
		ChainID:   chainID,
		Nonce:     nonce,
		GasAmount: gasAmount.String(),
	}
	c.JSON(http.StatusOK, resp)
}
