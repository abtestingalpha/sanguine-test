// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package explorer

import (
	"fmt"
	"io"
	"strconv"
)

type MessageType interface {
	IsMessageType()
}

type AddressChainRanking struct {
	ChainID   *int64   `json:"chainID,omitempty"`
	VolumeUsd *float64 `json:"volumeUsd,omitempty"`
	Rank      *int64   `json:"rank,omitempty"`
}

type AddressDailyCount struct {
	Date  *string `json:"date,omitempty"`
	Count *int64  `json:"count,omitempty"`
}

type AddressData struct {
	BridgeVolume *float64               `json:"bridgeVolume,omitempty"`
	BridgeFees   *float64               `json:"bridgeFees,omitempty"`
	BridgeTxs    *int64                 `json:"bridgeTxs,omitempty"`
	SwapVolume   *float64               `json:"swapVolume,omitempty"`
	SwapFees     *float64               `json:"swapFees,omitempty"`
	SwapTxs      *int64                 `json:"swapTxs,omitempty"`
	Rank         *int64                 `json:"rank,omitempty"`
	EarliestTx   *int64                 `json:"earliestTx,omitempty"`
	ChainRanking []*AddressChainRanking `json:"chainRanking,omitempty"`
	DailyData    []*AddressDailyCount   `json:"dailyData,omitempty"`
}

// AddressRanking gives the amount of transactions that occurred for a specific address across all chains.
type AddressRanking struct {
	Address *string `json:"address,omitempty"`
	Count   *int64  `json:"count,omitempty"`
}

type BlockHeight struct {
	ChainID     *int64        `json:"chainID,omitempty"`
	Type        *ContractType `json:"type,omitempty"`
	BlockNumber *int64        `json:"blockNumber,omitempty"`
}

// BridgeTransaction represents an entire bridge transaction, including both
// to and from transactions. If a `from` transaction does not have a corresponding
// `to` transaction, `pending` will be true.
type BridgeTransaction struct {
	FromInfo    *PartialInfo `json:"fromInfo,omitempty"`
	ToInfo      *PartialInfo `json:"toInfo,omitempty"`
	Kappa       *string      `json:"kappa,omitempty"`
	Pending     *bool        `json:"pending,omitempty"`
	SwapSuccess *bool        `json:"swapSuccess,omitempty"`
}

// BridgeWatcherTx represents a single sided bridge transaction specifically for the bridge watcher.
type BridgeWatcherTx struct {
	BridgeTx    *PartialInfo  `json:"bridgeTx,omitempty"`
	Pending     *bool         `json:"pending,omitempty"`
	Type        *BridgeTxType `json:"type,omitempty"`
	Kappa       *string       `json:"kappa,omitempty"`
	KappaStatus *KappaStatus  `json:"kappaStatus,omitempty"`
}

type ContractQuery struct {
	ChainID int64        `json:"chainID"`
	Type    ContractType `json:"type"`
}

// DateResult is a given statistic for a given date.
type DateResult struct {
	Date  *string  `json:"date,omitempty"`
	Total *float64 `json:"total,omitempty"`
}

// DateResult is a given statistic for a given date.
type DateResultByChain struct {
	Date      *string  `json:"date,omitempty"`
	Ethereum  *float64 `json:"ethereum,omitempty"`
	Optimism  *float64 `json:"optimism,omitempty"`
	Cronos    *float64 `json:"cronos,omitempty"`
	Bsc       *float64 `json:"bsc,omitempty"`
	Polygon   *float64 `json:"polygon,omitempty"`
	Fantom    *float64 `json:"fantom,omitempty"`
	Boba      *float64 `json:"boba,omitempty"`
	Metis     *float64 `json:"metis,omitempty"`
	Moonbeam  *float64 `json:"moonbeam,omitempty"`
	Moonriver *float64 `json:"moonriver,omitempty"`
	Klaytn    *float64 `json:"klaytn,omitempty"`
	Arbitrum  *float64 `json:"arbitrum,omitempty"`
	Avalanche *float64 `json:"avalanche,omitempty"`
	Dfk       *float64 `json:"dfk,omitempty"`
	Aurora    *float64 `json:"aurora,omitempty"`
	Harmony   *float64 `json:"harmony,omitempty"`
	Canto     *float64 `json:"canto,omitempty"`
	Dogechain *float64 `json:"dogechain,omitempty"`
	Base      *float64 `json:"base,omitempty"`
	Blast     *float64 `json:"blast,omitempty"`
	Total     *float64 `json:"total,omitempty"`
}

type HeroType struct {
	Recipient string `json:"recipient"`
	HeroID    string `json:"heroID"`
}

func (HeroType) IsMessageType() {}

// HistoricalResult is a given statistic for dates.
type HistoricalResult struct {
	Total       *float64              `json:"total,omitempty"`
	DateResults []*DateResult         `json:"dateResults,omitempty"`
	Type        *HistoricalResultType `json:"type,omitempty"`
}

type Leaderboard struct {
	Address      *string  `json:"address,omitempty"`
	VolumeUsd    *float64 `json:"volumeUSD,omitempty"`
	Fees         *float64 `json:"fees,omitempty"`
	Txs          *int64   `json:"txs,omitempty"`
	Rank         *int64   `json:"rank,omitempty"`
	AvgVolumeUsd *float64 `json:"avgVolumeUSD,omitempty"`
}

type MessageBusTransaction struct {
	FromInfo  *PartialMessageBusInfo `json:"fromInfo,omitempty"`
	ToInfo    *PartialMessageBusInfo `json:"toInfo,omitempty"`
	Pending   *bool                  `json:"pending,omitempty"`
	MessageID *string                `json:"messageID,omitempty"`
}

// PartialInfo is a transaction that occurred on one chain.
type PartialInfo struct {
	ChainID            *int64   `json:"chainID,omitempty"`
	DestinationChainID *int64   `json:"destinationChainID,omitempty"`
	Address            *string  `json:"address,omitempty"`
	TxnHash            *string  `json:"txnHash,omitempty"`
	Value              *string  `json:"value,omitempty"`
	FormattedValue     *float64 `json:"formattedValue,omitempty"`
	USDValue           *float64 `json:"USDValue,omitempty"`
	TokenAddress       *string  `json:"tokenAddress,omitempty"`
	TokenSymbol        *string  `json:"tokenSymbol,omitempty"`
	BlockNumber        *int64   `json:"blockNumber,omitempty"`
	Time               *int64   `json:"time,omitempty"`
	FormattedTime      *string  `json:"formattedTime,omitempty"`
	FormattedEventType *string  `json:"formattedEventType,omitempty"`
	EventType          *int64   `json:"eventType,omitempty"`
}

type PartialMessageBusInfo struct {
	ChainID              *int64      `json:"chainID,omitempty"`
	ChainName            *string     `json:"chainName,omitempty"`
	DestinationChainID   *int64      `json:"destinationChainID,omitempty"`
	DestinationChainName *string     `json:"destinationChainName,omitempty"`
	ContractAddress      *string     `json:"contractAddress,omitempty"`
	TxnHash              *string     `json:"txnHash,omitempty"`
	Message              *string     `json:"message,omitempty"`
	MessageType          MessageType `json:"messageType,omitempty"`
	BlockNumber          *int64      `json:"blockNumber,omitempty"`
	Time                 *int64      `json:"time,omitempty"`
	FormattedTime        *string     `json:"formattedTime,omitempty"`
	RevertedReason       *string     `json:"revertedReason,omitempty"`
}

type PetType struct {
	Recipient string `json:"recipient"`
	PetID     string `json:"petID"`
	Name      string `json:"name"`
}

func (PetType) IsMessageType() {}

type TearType struct {
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

func (TearType) IsMessageType() {}

// TokenCountResult gives the amount of transactions that occurred for a specific token, separated by chain ID.
type TokenCountResult struct {
	ChainID      *int64  `json:"chainID,omitempty"`
	TokenAddress *string `json:"tokenAddress,omitempty"`
	Count        *int64  `json:"count,omitempty"`
}

// TransactionCountResult gives the amount of transactions that occurred for a specific chain ID.
type TransactionCountResult struct {
	ChainID *int64 `json:"chainID,omitempty"`
	Count   *int64 `json:"count,omitempty"`
}

type UnknownType struct {
	Known bool `json:"known"`
}

func (UnknownType) IsMessageType() {}

// ValueResult is a value result of either USD or numeric value.
type ValueResult struct {
	Value *string `json:"value,omitempty"`
}

type VolumeByChainID struct {
	ChainID *int64   `json:"chainID,omitempty"`
	Total   *float64 `json:"total,omitempty"`
}

type BridgeTxType string

const (
	BridgeTxTypeOrigin      BridgeTxType = "ORIGIN"
	BridgeTxTypeDestination BridgeTxType = "DESTINATION"
)

var AllBridgeTxType = []BridgeTxType{
	BridgeTxTypeOrigin,
	BridgeTxTypeDestination,
}

func (e BridgeTxType) IsValid() bool {
	switch e {
	case BridgeTxTypeOrigin, BridgeTxTypeDestination:
		return true
	}
	return false
}

func (e BridgeTxType) String() string {
	return string(e)
}

func (e *BridgeTxType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BridgeTxType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BridgeTxType", str)
	}
	return nil
}

func (e BridgeTxType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type BridgeType string

const (
	BridgeTypeBridge BridgeType = "BRIDGE"
	BridgeTypeCctp   BridgeType = "CCTP"
)

var AllBridgeType = []BridgeType{
	BridgeTypeBridge,
	BridgeTypeCctp,
}

func (e BridgeType) IsValid() bool {
	switch e {
	case BridgeTypeBridge, BridgeTypeCctp:
		return true
	}
	return false
}

func (e BridgeType) String() string {
	return string(e)
}

func (e *BridgeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BridgeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BridgeType", str)
	}
	return nil
}

func (e BridgeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ContractType string

const (
	ContractTypeBridge ContractType = "BRIDGE"
	ContractTypeCctp   ContractType = "CCTP"
)

var AllContractType = []ContractType{
	ContractTypeBridge,
	ContractTypeCctp,
}

func (e ContractType) IsValid() bool {
	switch e {
	case ContractTypeBridge, ContractTypeCctp:
		return true
	}
	return false
}

func (e ContractType) String() string {
	return string(e)
}

func (e *ContractType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContractType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContractType", str)
	}
	return nil
}

func (e ContractType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DailyStatisticType string

const (
	DailyStatisticTypeVolume       DailyStatisticType = "VOLUME"
	DailyStatisticTypeTransactions DailyStatisticType = "TRANSACTIONS"
	DailyStatisticTypeAddresses    DailyStatisticType = "ADDRESSES"
	DailyStatisticTypeFee          DailyStatisticType = "FEE"
)

var AllDailyStatisticType = []DailyStatisticType{
	DailyStatisticTypeVolume,
	DailyStatisticTypeTransactions,
	DailyStatisticTypeAddresses,
	DailyStatisticTypeFee,
}

func (e DailyStatisticType) IsValid() bool {
	switch e {
	case DailyStatisticTypeVolume, DailyStatisticTypeTransactions, DailyStatisticTypeAddresses, DailyStatisticTypeFee:
		return true
	}
	return false
}

func (e DailyStatisticType) String() string {
	return string(e)
}

func (e *DailyStatisticType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DailyStatisticType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DailyStatisticType", str)
	}
	return nil
}

func (e DailyStatisticType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Direction string

const (
	DirectionIn  Direction = "IN"
	DirectionOut Direction = "OUT"
)

var AllDirection = []Direction{
	DirectionIn,
	DirectionOut,
}

func (e Direction) IsValid() bool {
	switch e {
	case DirectionIn, DirectionOut:
		return true
	}
	return false
}

func (e Direction) String() string {
	return string(e)
}

func (e *Direction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Direction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Direction", str)
	}
	return nil
}

func (e Direction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Duration string

const (
	DurationPastDay     Duration = "PAST_DAY"
	DurationPastMonth   Duration = "PAST_MONTH"
	DurationPast3Months Duration = "PAST_3_MONTHS"
	DurationPast6Months Duration = "PAST_6_MONTHS"
	DurationPastYear    Duration = "PAST_YEAR"
	DurationAllTime     Duration = "ALL_TIME"
)

var AllDuration = []Duration{
	DurationPastDay,
	DurationPastMonth,
	DurationPast3Months,
	DurationPast6Months,
	DurationPastYear,
	DurationAllTime,
}

func (e Duration) IsValid() bool {
	switch e {
	case DurationPastDay, DurationPastMonth, DurationPast3Months, DurationPast6Months, DurationPastYear, DurationAllTime:
		return true
	}
	return false
}

func (e Duration) String() string {
	return string(e)
}

func (e *Duration) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Duration(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Duration", str)
	}
	return nil
}

func (e Duration) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HistoricalResultType string

const (
	HistoricalResultTypeBridgevolume HistoricalResultType = "BRIDGEVOLUME"
	HistoricalResultTypeTransactions HistoricalResultType = "TRANSACTIONS"
	HistoricalResultTypeAddresses    HistoricalResultType = "ADDRESSES"
)

var AllHistoricalResultType = []HistoricalResultType{
	HistoricalResultTypeBridgevolume,
	HistoricalResultTypeTransactions,
	HistoricalResultTypeAddresses,
}

func (e HistoricalResultType) IsValid() bool {
	switch e {
	case HistoricalResultTypeBridgevolume, HistoricalResultTypeTransactions, HistoricalResultTypeAddresses:
		return true
	}
	return false
}

func (e HistoricalResultType) String() string {
	return string(e)
}

func (e *HistoricalResultType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HistoricalResultType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HistoricalResultType", str)
	}
	return nil
}

func (e HistoricalResultType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type KappaStatus string

const (
	KappaStatusExists  KappaStatus = "EXISTS"
	KappaStatusPending KappaStatus = "PENDING"
	KappaStatusUnknown KappaStatus = "UNKNOWN"
)

var AllKappaStatus = []KappaStatus{
	KappaStatusExists,
	KappaStatusPending,
	KappaStatusUnknown,
}

func (e KappaStatus) IsValid() bool {
	switch e {
	case KappaStatusExists, KappaStatusPending, KappaStatusUnknown:
		return true
	}
	return false
}

func (e KappaStatus) String() string {
	return string(e)
}

func (e *KappaStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = KappaStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid KappaStatus", str)
	}
	return nil
}

func (e KappaStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Platform string

const (
	PlatformAll        Platform = "ALL"
	PlatformSwap       Platform = "SWAP"
	PlatformBridge     Platform = "BRIDGE"
	PlatformMessageBus Platform = "MESSAGE_BUS"
)

var AllPlatform = []Platform{
	PlatformAll,
	PlatformSwap,
	PlatformBridge,
	PlatformMessageBus,
}

func (e Platform) IsValid() bool {
	switch e {
	case PlatformAll, PlatformSwap, PlatformBridge, PlatformMessageBus:
		return true
	}
	return false
}

func (e Platform) String() string {
	return string(e)
}

func (e *Platform) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Platform(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Platform", str)
	}
	return nil
}

func (e Platform) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StatisticType string

const (
	StatisticTypeMeanVolumeUsd     StatisticType = "MEAN_VOLUME_USD"
	StatisticTypeMedianVolumeUsd   StatisticType = "MEDIAN_VOLUME_USD"
	StatisticTypeTotalVolumeUsd    StatisticType = "TOTAL_VOLUME_USD"
	StatisticTypeMeanFeeUsd        StatisticType = "MEAN_FEE_USD"
	StatisticTypeMedianFeeUsd      StatisticType = "MEDIAN_FEE_USD"
	StatisticTypeTotalFeeUsd       StatisticType = "TOTAL_FEE_USD"
	StatisticTypeCountTransactions StatisticType = "COUNT_TRANSACTIONS"
	StatisticTypeCountAddresses    StatisticType = "COUNT_ADDRESSES"
)

var AllStatisticType = []StatisticType{
	StatisticTypeMeanVolumeUsd,
	StatisticTypeMedianVolumeUsd,
	StatisticTypeTotalVolumeUsd,
	StatisticTypeMeanFeeUsd,
	StatisticTypeMedianFeeUsd,
	StatisticTypeTotalFeeUsd,
	StatisticTypeCountTransactions,
	StatisticTypeCountAddresses,
}

func (e StatisticType) IsValid() bool {
	switch e {
	case StatisticTypeMeanVolumeUsd, StatisticTypeMedianVolumeUsd, StatisticTypeTotalVolumeUsd, StatisticTypeMeanFeeUsd, StatisticTypeMedianFeeUsd, StatisticTypeTotalFeeUsd, StatisticTypeCountTransactions, StatisticTypeCountAddresses:
		return true
	}
	return false
}

func (e StatisticType) String() string {
	return string(e)
}

func (e *StatisticType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatisticType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatisticType", str)
	}
	return nil
}

func (e StatisticType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
