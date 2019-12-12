// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: proposaldeleteoperation.go

package operations

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gkany/graphSDK/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *ProposalDeleteOperation) MarshalJSON() ([]byte, error) {
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
func (j *ProposalDeleteOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"fee_paying_account":`)

	{

		obj, err = j.FeePayingAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"proposal":`)

	{

		obj, err = j.Proposal.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	if j.UsingOwnerAuthority {
		buf.WriteString(`,"using_owner_authority":true`)
	} else {
		buf.WriteString(`,"using_owner_authority":false`)
	}
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			buf.WriteString(`"fee":`)

			{

				err = j.Fee.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtProposalDeleteOperationbase = iota
	ffjtProposalDeleteOperationnosuchkey

	ffjtProposalDeleteOperationExtensions

	ffjtProposalDeleteOperationFeePayingAccount

	ffjtProposalDeleteOperationProposal

	ffjtProposalDeleteOperationUsingOwnerAuthority

	ffjtProposalDeleteOperationFee
)

var ffjKeyProposalDeleteOperationExtensions = []byte("extensions")

var ffjKeyProposalDeleteOperationFeePayingAccount = []byte("fee_paying_account")

var ffjKeyProposalDeleteOperationProposal = []byte("proposal")

var ffjKeyProposalDeleteOperationUsingOwnerAuthority = []byte("using_owner_authority")

var ffjKeyProposalDeleteOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *ProposalDeleteOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *ProposalDeleteOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtProposalDeleteOperationbase
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
				currentKey = ffjtProposalDeleteOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffjKeyProposalDeleteOperationExtensions, kn) {
						currentKey = ffjtProposalDeleteOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyProposalDeleteOperationFeePayingAccount, kn) {
						currentKey = ffjtProposalDeleteOperationFeePayingAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyProposalDeleteOperationFee, kn) {
						currentKey = ffjtProposalDeleteOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyProposalDeleteOperationProposal, kn) {
						currentKey = ffjtProposalDeleteOperationProposal
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffjKeyProposalDeleteOperationUsingOwnerAuthority, kn) {
						currentKey = ffjtProposalDeleteOperationUsingOwnerAuthority
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyProposalDeleteOperationFee, kn) {
					currentKey = ffjtProposalDeleteOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalDeleteOperationUsingOwnerAuthority, kn) {
					currentKey = ffjtProposalDeleteOperationUsingOwnerAuthority
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalDeleteOperationProposal, kn) {
					currentKey = ffjtProposalDeleteOperationProposal
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyProposalDeleteOperationFeePayingAccount, kn) {
					currentKey = ffjtProposalDeleteOperationFeePayingAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProposalDeleteOperationExtensions, kn) {
					currentKey = ffjtProposalDeleteOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtProposalDeleteOperationnosuchkey
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

				case ffjtProposalDeleteOperationExtensions:
					goto handle_Extensions

				case ffjtProposalDeleteOperationFeePayingAccount:
					goto handle_FeePayingAccount

				case ffjtProposalDeleteOperationProposal:
					goto handle_Proposal

				case ffjtProposalDeleteOperationUsingOwnerAuthority:
					goto handle_UsingOwnerAuthority

				case ffjtProposalDeleteOperationFee:
					goto handle_Fee

				case ffjtProposalDeleteOperationnosuchkey:
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

handle_Extensions:

	/* handler: j.Extensions type=types.Extensions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Extensions.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FeePayingAccount:

	/* handler: j.FeePayingAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.FeePayingAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Proposal:

	/* handler: j.Proposal type=types.ProposalID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Proposal.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_UsingOwnerAuthority:

	/* handler: j.UsingOwnerAuthority type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.UsingOwnerAuthority = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.UsingOwnerAuthority = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.Fee = nil

		} else {

			if j.Fee == nil {
				j.Fee = new(types.AssetAmount)
			}

			err = j.Fee.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
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
