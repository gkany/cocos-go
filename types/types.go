package types

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"time"

	sort "github.com/emirpasic/gods/utils"
	"github.com/gkany/cocos-go/util"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

var (
	//	ErrRPCClientNotInitialized      = fmt.Errorf("RPC client is not initialized")
	ErrNotImplemented               = fmt.Errorf("not implemented")
	ErrInvalidInputType             = fmt.Errorf("invalid input type")
	ErrInvalidInputLength           = fmt.Errorf("invalid input length")
	ErrInvalidPublicKey             = fmt.Errorf("invalid PublicKey")
	ErrInvalidAddress               = fmt.Errorf("invalid Address")
	ErrPublicKeyChainPrefixMismatch = fmt.Errorf("PublicKey chain prefix mismatch")
	ErrAddressChainPrefixMismatch   = fmt.Errorf("Address chain prefix mismatch")
	ErrInvalidChecksum              = fmt.Errorf("invalid checksum")
	ErrNoSigningKeyFound            = fmt.Errorf("no signing key found")
	ErrNoVerifyingKeyFound          = fmt.Errorf("no verifying key found")
	ErrInvalidDigestLength          = fmt.Errorf("invalid digest length")
	ErrInvalidPrivateKeyCurve       = fmt.Errorf("invalid PrivateKey curve")
	ErrChainConfigIsUndefined       = fmt.Errorf("chain config is undefined")
)

var (
	EmptyBuffer = []byte{}
	EmptyParams = []interface{}{}
)

type WorkerInitializerType UInt8

const (
	WorkerInitializerTypeRefund WorkerInitializerType = iota
	WorkerInitializerTypeVestingBalance
	WorkerInitializerTypeBurn
)

type CallOrderUpdateExtensionsType UInt8

const (
	CallOrderUpdateExtensionsTypeTargetRatio CallOrderUpdateExtensionsType = iota
)

type AccountCreateExtensionsType UInt8

const (
	AccountCreateExtensionsNullExt AccountCreateExtensionsType = iota
	AccountCreateExtensionsOwnerSpecial
	AccountCreateExtensionsActiveSpecial
	AccountCreateExtensionsBuyback
)

type SpecialAuthorityType UInt8

const (
	SpecialAuthorityTypeNoSpecial SpecialAuthorityType = iota
	SpecialAuthorityTypeTopHolders
)

type VestingPolicyType UInt8

const (
	VestingPolicyTypeLinear VestingPolicyType = iota
	VestingPolicyTypeCCD
)

type AssetType Int8

const (
	AssetTypeUndefined AssetType = -1
	AssetTypeCoreAsset AssetType = iota
	AssetTypeUIA
	AssetTypeSmartCoin
	AssetTypePredictionMarket
)

type OperationType UInt8

const (
	OperationTypeTransfer                              OperationType = iota //0
	OperationTypeLimitOrderCreate                                           //1
	OperationTypeLimitOrderCancel                                           //2
	OperationTypeCallOrderUpdate                                            //3
	OperationTypeFillOrder                                                  //4
	OperationTypeAccountCreate                                              //5
	OperationTypeAccountUpdate                                              //6
	OperationTypeAccountUpgrade                                             //7
	OperationTypeAssetCreate                                                //8
	OperationTypeAssetUpdate                                                //9
	OperationTypeAssetUpdateRestricted                                      //10
	OperationTypeAssetUpdateBitasset                                        //11
	OperationTypeAssetUpdateFeedProducers                                   //12
	OperationTypeAssetIssue                                                 //13
	OperationTypeAssetReserve                                               //14
	OperationTypeAssetSettle                                                //15
	OperationTypeAssetGlobalSettle                                          //16
	OperationTypeAssetPublishFeed                                           //17
	OperationTypeWitnessCreate                                              //18
	OperationTypeWitnessUpdate                                              //19
	OperationTypeProposalCreate                                             //20
	OperationTypeProposalUpdate                                             //21
	OperationTypeProposalDelete                                             //22
	OperationTypeCommitteeMemberCreate                                      //23
	OperationTypeCommitteeMemberUpdate                                      //24
	OperationTypeCommitteeMemberUpdateGlobalParameters                      //25
	OperationTypeVestingBalanceCreate                                       //26
	OperationTypeVestingBalanceWithdraw                                     //27
	OperationTypeWorkerCreate                                               //28
	OperationTypeBalanceClaim                                               //29
	OperationTypeAssetSettleCancel                                          //30
	OperationTypeAssetClaimFees                                             //31
	OperationTypeBidCollateral                                              //32
	OperationTypeExecuteBid                                                 //33
	OperationTypeContractCreate                                             //34
	OperationTypeCallContractFunction                                       //35
	OperationTypeTemporaryAuthorityChange                                   //36
	OperationTypeRegisterNhAssetCreator                                     //37
	OperationTypeCreateWorldView                                            //38
	OperationTypeRelateWorldView                                            //39
	OperationTypeCreateNhAsset                                              //40
	OperationTypeDeleteNhAsset                                              //41
	OperationTypeTransferNhAsset                                            //42
	OperationTypeCreateNhAssetOrder                                         //43
	OperationTypeCancelNhAssetOrder                                         //44
	OperationTypeFillNhAssetOrder                                           //45
	OperationTypeCreateFile                                                 //46
	OperationTypeAddFileRelateAccount                                       //47
	OperationTypeFileSignature                                              //48
	OperationTypeRelateParentFile                                           //49
	OperationTypeReviseContract                                             //50
	OperationTypeCrontabCreate                                              //51
	OperationTypeCrontabCancel                                              //52
	OperationTypeCrontabRecover                                             //53
	OperationTypeUpdateCollateralForGas                                     //54
	OperationTypeAccountAuthentication                                      //55
)

func (p OperationType) OperationName() string {
	return fmt.Sprintf("%sOperation", p.String()[13:])
}

type SpaceType UInt8

const (
	SpaceTypeRelativeProtocol SpaceType = iota
	SpaceTypeProtocol
	SpaceTypeImplementation
	SpaceTypeExtension
	SpaceTypeNHAssetProtocol
	SpaceTypeMarketHistory
)

type ObjectType UInt8

//for SpaceTypeProtocol
const (
	ObjectTypeBase             ObjectType = iota + 1 // 1
	ObjectTypeAccount                                // 2
	ObjectTypeAsset                                  // 3
	ObjectTypeForceSettlement                        // 4
	ObjectTypeCommitteeMember                        // 5
	ObjectTypeWitness                                // 6
	ObjectTypeLimitOrder                             // 7
	ObjectTypeCallOrder                              // 8
	ObjectTypeCustom                                 // 9
	ObjectTypeProposal                               // 10
	ObjectTypeOperationHistory                       // 11
	ObjectTypeCrontab                                // 12
	ObjectTypeVestingBalance                         // 13
	ObjectTypeWorker                                 // 14
	ObjectTypeBalance                                // 15
	ObjectTypeContract                               // 16
	ObjectTypeContractData                           // 17
	ObjectTypeFile                                   // 18
)

// for SpaceTypeImplementation
const (
	ObjectTypeGlobalProperty ObjectType = iota
	ObjectTypeDynamicGlobalProperty
	ObjectTypeContractBinCode
	ObjectTypeAssetDynamicData
	ObjectTypeAssetBitAssetData
	ObjectTypeAccountBalance
	ObjectTypeAccountStatistics
	ObjectTypeTransaction
	ObjectTypeBlockSummary
	ObjectTypeAccountTransactionHistory
	ObjectTypeCollateralBid
	ObjectTypeChainProperty
	ObjectTypeWitnessSchedule
	ObjectTypeBudgetRecord
	ObjectTypeSpecialAuthority
)

// for extension_type_for_nico
const (
	ObjectTypeTemporaryAuthority ObjectType = iota
	ObjectTypeTransactionInBlockInfo
	ObjectTypeAssetRestrictedObject
	ObjectTypeUnsuccessfulCandidates
	ObjectTypeCollateralForGas
)

// for nh object type
const (
	ObjectTypeNhAssetCreator ObjectType = iota // 0
	ObjectTypeWorldView                        // 1
	ObjectTypeNHAsset                          // 2
	ObjectTypeNHAssetOrder                     // 3
)

type AssetPermission Int16

const (
	AssetPermissionChargeMarketFee     AssetPermission = 0x01
	AssetPermissionWhiteList           AssetPermission = 0x02
	AssetPermissionOverrideAuthority   AssetPermission = 0x04
	AssetPermissionTransferRestricted  AssetPermission = 0x08
	AssetPermissionDisableForceSettle  AssetPermission = 0x10
	AssetPermissionGlobalSettle        AssetPermission = 0x20
	AssetPermissionDisableConfidential AssetPermission = 0x40
	AssetPermissionWitnessFedAsset     AssetPermission = 0x80
	AssetPermissionComiteeFedAsset     AssetPermission = 0x100
)

type Rate float64

func (p Rate) Inverse() Rate {
	return 1 / p
}

func (p Rate) Value() float64 {
	return float64(p)
}

func unmarshalUInt(data []byte) (uint64, error) {
	if len(data) == 0 {
		return 0, errors.New("unmarshalUInt: empty input")
	}

	var (
		res uint64
		err error
		len = len(data)
	)

	if data[0] == '"' && data[len-1] == '"' {
		data := data[1 : len-1]
		res, err = strconv.ParseUint(string(data), 10, 64)
		if err != nil {
			return 0, errors.Errorf("unmarshalUInt: unable to parse input %v", data)
		}
	} else if err := ffjson.Unmarshal(data, &res); err != nil {
		return 0, errors.Errorf("unmarshalUInt: unable to unmarshal input %v", data)
	}

	return res, nil
}

func unmarshalInt(data []byte) (int64, error) {
	if len(data) == 0 {
		return 0, errors.New("unmarshalInt: empty input")
	}

	var (
		res int64
		err error
		len = len(data)
	)

	if data[0] == '"' && data[len-1] == '"' {
		data := data[1 : len-1]
		res, err = strconv.ParseInt(string(data), 10, 64)
		if err != nil {
			return 0, errors.Errorf("unmarshalInt: unable to parse input %v", data)
		}
	} else if err := ffjson.Unmarshal(data, &res); err != nil {
		return 0, errors.Errorf("unmarshalInt: unable to unmarshal input %v", data)
	}

	return res, nil
}

func unmarshalFloat(data []byte) (float64, error) {
	if len(data) == 0 {
		return 0, errors.New("unmarshalFloat: empty input")
	}

	var (
		res float64
		err error
		len = len(data)
	)

	if data[0] == '"' && data[len-1] == '"' {
		data := data[1 : len-1]
		res, err = strconv.ParseFloat(string(data), 64)
		if err != nil {
			return 0, errors.Errorf("unmarshalFloat: unable to parse input %v", data)
		}
	} else if err := ffjson.Unmarshal(data, &res); err != nil {
		return 0, errors.Errorf("unmarshalFloat: unable to unmarshal input %v", data)
	}

	return res, nil
}

type UInt uint

func (num *UInt) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt(v)
	return nil
}

func (num UInt) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(uint(num))
}

type UInt8 uint8

func (num *UInt8) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt8(v)
	return nil
}

func (num UInt8) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(uint8(num))
}

type UInt16 uint16

func (num *UInt16) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt16(v)
	return nil
}

func (num UInt16) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(uint16(num))
}

type UInt32 uint32

func (num *UInt32) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt32(v)
	return nil
}

func (num UInt32) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(uint32(num))
}

type UInt64 uint64

func (num *UInt64) UnmarshalJSON(data []byte) error {
	v, err := unmarshalUInt(data)
	if err != nil {
		return err
	}

	*num = UInt64(v)
	return nil
}

func (num UInt64) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(uint64(num))
}

type Int8 int8

func (num *Int8) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int8(v)
	return nil
}

func (num Int8) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(int8(num))
}

type Int16 int16

func (num *Int16) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int16(v)
	return nil
}

func (num Int16) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(int16(num))
}

type Int32 int32

func (num *Int32) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int32(v)
	return nil
}

func (num Int32) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(int32(num))
}

type Int64 int64

func (num *Int64) UnmarshalJSON(data []byte) error {
	v, err := unmarshalInt(data)
	if err != nil {
		return err
	}

	*num = Int64(v)
	return nil
}

func (num Int64) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(int64(num))
}

type Float32 float32

func (num *Float32) UnmarshalJSON(data []byte) error {
	v, err := unmarshalFloat(data)
	if err != nil {
		return err
	}

	*num = Float32(v)
	return nil
}

func (num Float32) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(float32(num))
}

type Float64 float64

func (num *Float64) UnmarshalJSON(data []byte) error {
	v, err := unmarshalFloat(data)
	if err != nil {
		return err
	}

	*num = Float64(v)
	return nil
}

func (num Float64) Marshal(enc *util.TypeEncoder) error {
	return enc.EncodeNumber(float64(num))
}

const TimeFormat = `"2006-01-02T15:04:05"`

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(TimeFormat)), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	tm, err := time.ParseInLocation(TimeFormat, string(data), time.UTC)
	if err != nil {
		return errors.Annotate(err, "ParseInLocation")
	}

	t.Time = tm
	return nil
}

func (t Time) Marshal(enc *util.TypeEncoder) error {
	return enc.Encode(uint32(t.Time.Unix()))
}

func (t Time) Add(dur time.Duration) Time {
	return Time{t.Time.Add(dur)}
}

func (t *Time) FromTime(tm time.Time) {
	t.Time = tm
}

func (t *Time) Set(dur time.Duration) {
	t.Time = time.Now().UTC().Add(dur)
}

func (t Time) IsZero() bool {
	return t.Time.IsZero()
}

type String struct {
	data string
}

func (p *String) SetData(data string) {
	p.data = data
}

func (p String) MarshalJSON() ([]byte, error) {
	return ffjson.Marshal(p.data)
}

func (p String) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeString(p.data); err != nil {
		return errors.Annotate(err, "encode data")
	}
	return nil
}

func (p *String) UnmarshalJSON(data []byte) error {
	if err := ffjson.Unmarshal(data, &p.data); err != nil {
		return errors.Annotate(err, "Unmarshal")
	}
	return nil
}

func (p *String) Unmarshal(dec *util.TypeDecoder) error {
	if err := dec.DecodeString(&p.data); err != nil {
		return errors.Annotate(err, "decode data")
	}
	return nil
}

func (p String) String() string {
	return p.data
}

type Buffers []Buffer
type Buffer []byte

func (p *Buffer) UnmarshalJSON(data []byte) error {
	var b string
	if err := ffjson.Unmarshal(data, &b); err != nil {
		return errors.Annotate(err, "Unmarshal")
	}

	return p.FromString(b)
}

func (p Buffer) Bytes() []byte {
	return p
}

func (p Buffer) Length() int {
	return len(p)
}

func (p Buffer) String() string {
	return hex.EncodeToString(p)
}

func (p *Buffer) FromString(data string) error {
	buf, err := hex.DecodeString(data)
	if err != nil {
		return errors.Annotate(err, "DecodeString")
	}

	*p = buf
	return nil
}

func (p Buffer) MarshalJSON() ([]byte, error) {
	return ffjson.Marshal(p.String())
}

func (p Buffer) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	if err := enc.Encode(p.Bytes()); err != nil {
		return errors.Annotate(err, "encode bytes")
	}

	return nil
}

func (p *Buffer) Unmarshal(dec *util.TypeDecoder) error {
	var len uint64
	if err := dec.DecodeUVarint(&len); err != nil {
		return errors.Annotate(err, "decode length")
	}

	if err := dec.ReadBytes(p, len); err != nil {
		return errors.Annotate(err, "decode bytes")
	}

	return nil
}

//Encrypt AES-encrypts the buffer content
func (p *Buffer) Encrypt(cipherKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return nil, errors.Annotate(err, "NewCipher")
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+p.Length())
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, errors.Annotate(err, "ReadFull")
	}

	cipher.NewCFBEncrypter(block, iv).XORKeyStream(
		ciphertext[aes.BlockSize:],
		p.Bytes(),
	)

	return ciphertext, nil
}

//Decrypt AES decrypts the buffer content
func (p *Buffer) Decrypt(cipherKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return nil, errors.Annotate(err, "NewCipher")
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if byteLen := p.Length(); byteLen < aes.BlockSize {
		return nil, errors.Errorf("invalid cipher size %d", byteLen)
	}

	buf := p.Bytes()
	iv := buf[:aes.BlockSize]
	buf = buf[aes.BlockSize:]

	// XORKeyStream can work in-place if the two arguments are the same.
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(buf, buf)

	return buf, nil
}

type FixedBuffer struct {
	Buffer
}

func (p FixedBuffer) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.Bytes()); err != nil {
		return errors.Annotate(err, "encode bytes")
	}

	return nil
}

func BufferFromString(data string) (b Buffer, err error) {
	b = Buffer{}
	err = b.FromString(data)
	return
}

type Unmarshalable interface {
	UnmarshalJSON(input []byte) error
}
type M map[string]interface{}

//committee member
type SupporterType map[AccountID]AssetAmount

func (p *SupporterType) UnmarshalJSON(data []byte) error {
	var supporters [][]interface{}
	if err := ffjson.Unmarshal(data, &supporters); err != nil {
		return errors.Annotate(err, "unmarshal supporters")
	}

	(*p) = make(map[AccountID]AssetAmount)
	for _, tk := range supporters {
		key, ok := tk[0].(AccountID)
		if !ok {
			return ErrInvalidInputType
		}

		amount, ok := tk[1].(AssetAmount)
		if !ok {
			return ErrInvalidInputType
		}

		(*p)[key] = amount
	}

	return nil
}

func (p SupporterType) MarshalJSON() ([]byte, error) {
	ret := make([]interface{}, 0, len(p))
	for k, v := range p {
		ret = append(ret, []interface{}{k.String(), v})
	}
	return ffjson.Marshal(ret)
}

func (p SupporterType) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	//sort keys
	keys := make([]interface{}, 0, len(p))
	for k := range p {
		keys = append(keys, k)
	}

	var err error
	sort.Sort(keys, func(a, b interface{}) (s int) {
		s, err = publicKeyComparator(a.(*PublicKey), b.(*PublicKey))
		return
	})

	if err != nil {
		return errors.Annotate(err, "Sort")
	}

	for _, k := range keys {
		id := k.(AccountID)
		if err := id.Marshal(enc); err != nil {
			return errors.Annotate(err, "encode PubKey")
		}

		if err := enc.Encode(p[id]); err != nil {
			return errors.Annotate(err, "encode Weight")
		}
	}

	return nil
}

type AccountIDArray []AccountID
type FileIDArray []FileID

type SignaturesType map[AccountID]string

func (p *SignaturesType) UnmarshalJSON(data []byte) error {
	var signatures [][]interface{}
	if err := ffjson.Unmarshal(data, &signatures); err != nil {
		return errors.Annotate(err, "unmarshal signatures")
	}

	(*p) = make(map[AccountID]string)
	for _, tk := range signatures {
		key, ok := tk[0].(AccountID)
		if !ok {
			return ErrInvalidInputType
		}

		value, ok := tk[1].(string)
		if !ok {
			return ErrInvalidInputType
		}

		(*p)[key] = value
	}

	return nil
}

func (p SignaturesType) MarshalJSON() ([]byte, error) {
	ret := make([]interface{}, 0, len(p))
	for k, v := range p {
		ret = append(ret, []interface{}{k.String(), v})
	}
	return ffjson.Marshal(ret)
}

func (p SignaturesType) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	// for _, key := range keys {
	// 	key := key.(AccountID)
	// 	if err := key.Marshal(enc); err != nil {
	// 		return errors.Annotate(err, "encode AccountID")
	// 	}

	// 	if err := enc.Encode(p[key]); err != nil {
	// 		return errors.Annotate(err, "encode value")
	// 	}
	// }

	return nil
}

type ContractIDListType []ContractID
type NHAssetIDListType []NHAssetID
type NHAssetMapType map[ContractID]NHAssetIDListType
type StringMapType map[string]string
type ContractDescribeMapType map[ContractID]StringMapType

type NhAssetLeaseLimitType int // 0 -- black_list, 1 -- white_list
type CharListType []int
