// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: createnhassetorderoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *CreateNhAssetOrderOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *CreateNhAssetOrderOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "seller":`)

	{

		obj, err = j.Seller.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"otcaccount":`)

	{

		obj, err = j.OTCAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	/* Struct fall back. type=types.AssetAmount kind=struct */
	buf.WriteString(`,"pending_orders_fee":`)
	err = buf.Encode(&j.PendingOrdersFee)
	if err != nil {
		return err
	}
	buf.WriteString(`,"nh_asset":`)

	{

		obj, err = j.NHAsset.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"memo":`)
	fflib.WriteJsonString(buf, string(j.Memo))
	/* Struct fall back. type=types.AssetAmount kind=struct */
	buf.WriteString(`,"price":`)
	err = buf.Encode(&j.Price)
	if err != nil {
		return err
	}
	buf.WriteString(`,"expiration":`)

	{

		obj, err = j.Expiration.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			/* Struct fall back. type=types.AssetAmount kind=struct */
			buf.WriteString(`"fee":`)
			err = buf.Encode(j.Fee)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtCreateNhAssetOrderOperationbase = iota
	ffjtCreateNhAssetOrderOperationnosuchkey

	ffjtCreateNhAssetOrderOperationSeller

	ffjtCreateNhAssetOrderOperationOTCAccount

	ffjtCreateNhAssetOrderOperationPendingOrdersFee

	ffjtCreateNhAssetOrderOperationNHAsset

	ffjtCreateNhAssetOrderOperationMemo

	ffjtCreateNhAssetOrderOperationPrice

	ffjtCreateNhAssetOrderOperationExpiration

	ffjtCreateNhAssetOrderOperationFee
)

var ffjKeyCreateNhAssetOrderOperationSeller = []byte("seller")

var ffjKeyCreateNhAssetOrderOperationOTCAccount = []byte("otcaccount")

var ffjKeyCreateNhAssetOrderOperationPendingOrdersFee = []byte("pending_orders_fee")

var ffjKeyCreateNhAssetOrderOperationNHAsset = []byte("nh_asset")

var ffjKeyCreateNhAssetOrderOperationMemo = []byte("memo")

var ffjKeyCreateNhAssetOrderOperationPrice = []byte("price")

var ffjKeyCreateNhAssetOrderOperationExpiration = []byte("expiration")

var ffjKeyCreateNhAssetOrderOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *CreateNhAssetOrderOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *CreateNhAssetOrderOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtCreateNhAssetOrderOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtCreateNhAssetOrderOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffjKeyCreateNhAssetOrderOperationExpiration, kn) {
						currentKey = ffjtCreateNhAssetOrderOperationExpiration
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyCreateNhAssetOrderOperationFee, kn) {
						currentKey = ffjtCreateNhAssetOrderOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'm':

					if bytes.Equal(ffjKeyCreateNhAssetOrderOperationMemo, kn) {
						currentKey = ffjtCreateNhAssetOrderOperationMemo
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyCreateNhAssetOrderOperationNHAsset, kn) {
						currentKey = ffjtCreateNhAssetOrderOperationNHAsset
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyCreateNhAssetOrderOperationOTCAccount, kn) {
						currentKey = ffjtCreateNhAssetOrderOperationOTCAccount
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyCreateNhAssetOrderOperationPendingOrdersFee, kn) {
						currentKey = ffjtCreateNhAssetOrderOperationPendingOrdersFee
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyCreateNhAssetOrderOperationPrice, kn) {
						currentKey = ffjtCreateNhAssetOrderOperationPrice
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyCreateNhAssetOrderOperationSeller, kn) {
						currentKey = ffjtCreateNhAssetOrderOperationSeller
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyCreateNhAssetOrderOperationFee, kn) {
					currentKey = ffjtCreateNhAssetOrderOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCreateNhAssetOrderOperationExpiration, kn) {
					currentKey = ffjtCreateNhAssetOrderOperationExpiration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCreateNhAssetOrderOperationPrice, kn) {
					currentKey = ffjtCreateNhAssetOrderOperationPrice
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCreateNhAssetOrderOperationMemo, kn) {
					currentKey = ffjtCreateNhAssetOrderOperationMemo
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyCreateNhAssetOrderOperationNHAsset, kn) {
					currentKey = ffjtCreateNhAssetOrderOperationNHAsset
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyCreateNhAssetOrderOperationPendingOrdersFee, kn) {
					currentKey = ffjtCreateNhAssetOrderOperationPendingOrdersFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCreateNhAssetOrderOperationOTCAccount, kn) {
					currentKey = ffjtCreateNhAssetOrderOperationOTCAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyCreateNhAssetOrderOperationSeller, kn) {
					currentKey = ffjtCreateNhAssetOrderOperationSeller
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtCreateNhAssetOrderOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtCreateNhAssetOrderOperationSeller:
					goto handle_Seller

				case ffjtCreateNhAssetOrderOperationOTCAccount:
					goto handle_OTCAccount

				case ffjtCreateNhAssetOrderOperationPendingOrdersFee:
					goto handle_PendingOrdersFee

				case ffjtCreateNhAssetOrderOperationNHAsset:
					goto handle_NHAsset

				case ffjtCreateNhAssetOrderOperationMemo:
					goto handle_Memo

				case ffjtCreateNhAssetOrderOperationPrice:
					goto handle_Price

				case ffjtCreateNhAssetOrderOperationExpiration:
					goto handle_Expiration

				case ffjtCreateNhAssetOrderOperationFee:
					goto handle_Fee

				case ffjtCreateNhAssetOrderOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Seller:

	/* handler: j.Seller type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Seller.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_OTCAccount:

	/* handler: j.OTCAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.OTCAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PendingOrdersFee:

	/* handler: j.PendingOrdersFee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.PendingOrdersFee)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NHAsset:

	/* handler: j.NHAsset type=types.NHAssetID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.NHAsset.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Memo:

	/* handler: j.Memo type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Memo = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Price:

	/* handler: j.Price type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Price)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Expiration:

	/* handler: j.Expiration type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Expiration.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}