// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: signedtransaction.go

package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *ProcessedTransaction) MarshalJSON() ([]byte, error) {
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
func (j *ProcessedTransaction) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"Operationresults":`)
	if j.Operationresults != nil {
		buf.WriteString(`[`)
		for i, v := range j.Operationresults {
			if i != 0 {
				buf.WriteString(`,`)
			}
			/* Interface types must use runtime reflection. type=types.OperationResult kind=interface */
			err = buf.Encode(v)
			if err != nil {
				return err
			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte(',')
	if len(j.AgreedTask) != 0 {
		buf.WriteString(`"agreed_task":`)
		if j.AgreedTask != nil {
			buf.WriteString(`[`)
			for i, v := range j.AgreedTask {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"signatures":`)
	if j.Signatures != nil {
		buf.WriteString(`[`)
		for i, v := range j.Signatures {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"ref_block_num":`)
	fflib.FormatBits2(buf, uint64(j.RefBlockNum), 10, false)
	buf.WriteString(`,"ref_block_prefix":`)
	fflib.FormatBits2(buf, uint64(j.RefBlockPrefix), 10, false)
	buf.WriteString(`,"expiration":`)

	{

		obj, err = j.Expiration.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"operations":`)

	{

		obj, err = j.Operations.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtProcessedTransactionbase = iota
	ffjtProcessedTransactionnosuchkey

	ffjtProcessedTransactionOperationresults

	ffjtProcessedTransactionAgreedTask

	ffjtProcessedTransactionSignatures

	ffjtProcessedTransactionRefBlockNum

	ffjtProcessedTransactionRefBlockPrefix

	ffjtProcessedTransactionExpiration

	ffjtProcessedTransactionOperations

	ffjtProcessedTransactionExtensions
)

var ffjKeyProcessedTransactionOperationresults = []byte("Operationresults")

var ffjKeyProcessedTransactionAgreedTask = []byte("agreed_task")

var ffjKeyProcessedTransactionSignatures = []byte("signatures")

var ffjKeyProcessedTransactionRefBlockNum = []byte("ref_block_num")

var ffjKeyProcessedTransactionRefBlockPrefix = []byte("ref_block_prefix")

var ffjKeyProcessedTransactionExpiration = []byte("expiration")

var ffjKeyProcessedTransactionOperations = []byte("operations")

var ffjKeyProcessedTransactionExtensions = []byte("extensions")

// UnmarshalJSON umarshall json - template of ffjson
func (j *ProcessedTransaction) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *ProcessedTransaction) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtProcessedTransactionbase
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
				currentKey = ffjtProcessedTransactionnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'O':

					if bytes.Equal(ffjKeyProcessedTransactionOperationresults, kn) {
						currentKey = ffjtProcessedTransactionOperationresults
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'a':

					if bytes.Equal(ffjKeyProcessedTransactionAgreedTask, kn) {
						currentKey = ffjtProcessedTransactionAgreedTask
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyProcessedTransactionExpiration, kn) {
						currentKey = ffjtProcessedTransactionExpiration
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyProcessedTransactionExtensions, kn) {
						currentKey = ffjtProcessedTransactionExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyProcessedTransactionOperations, kn) {
						currentKey = ffjtProcessedTransactionOperations
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyProcessedTransactionRefBlockNum, kn) {
						currentKey = ffjtProcessedTransactionRefBlockNum
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyProcessedTransactionRefBlockPrefix, kn) {
						currentKey = ffjtProcessedTransactionRefBlockPrefix
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyProcessedTransactionSignatures, kn) {
						currentKey = ffjtProcessedTransactionSignatures
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyProcessedTransactionExtensions, kn) {
					currentKey = ffjtProcessedTransactionExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProcessedTransactionOperations, kn) {
					currentKey = ffjtProcessedTransactionOperations
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyProcessedTransactionExpiration, kn) {
					currentKey = ffjtProcessedTransactionExpiration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProcessedTransactionRefBlockPrefix, kn) {
					currentKey = ffjtProcessedTransactionRefBlockPrefix
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProcessedTransactionRefBlockNum, kn) {
					currentKey = ffjtProcessedTransactionRefBlockNum
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProcessedTransactionSignatures, kn) {
					currentKey = ffjtProcessedTransactionSignatures
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProcessedTransactionAgreedTask, kn) {
					currentKey = ffjtProcessedTransactionAgreedTask
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyProcessedTransactionOperationresults, kn) {
					currentKey = ffjtProcessedTransactionOperationresults
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtProcessedTransactionnosuchkey
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

				case ffjtProcessedTransactionOperationresults:
					goto handle_Operationresults

				case ffjtProcessedTransactionAgreedTask:
					goto handle_AgreedTask

				case ffjtProcessedTransactionSignatures:
					goto handle_Signatures

				case ffjtProcessedTransactionRefBlockNum:
					goto handle_RefBlockNum

				case ffjtProcessedTransactionRefBlockPrefix:
					goto handle_RefBlockPrefix

				case ffjtProcessedTransactionExpiration:
					goto handle_Expiration

				case ffjtProcessedTransactionOperations:
					goto handle_Operations

				case ffjtProcessedTransactionExtensions:
					goto handle_Extensions

				case ffjtProcessedTransactionnosuchkey:
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

handle_Operationresults:

	/* handler: j.Operationresults type=types.OperationResults kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for OperationResults", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Operationresults = nil
		} else {

			j.Operationresults = []OperationResult{}

			wantVal := true

			for {

				var tmpJOperationresults OperationResult

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJOperationresults type=types.OperationResult kind=interface quoted=false*/

				{
					/* Falling back. type=types.OperationResult kind=interface */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJOperationresults)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.Operationresults = append(j.Operationresults, tmpJOperationresults)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AgreedTask:

	/* handler: j.AgreedTask type=types.AgreedTaskPair kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AgreedTaskPair", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.AgreedTask = nil
		} else {

			j.AgreedTask = []string{}

			wantVal := true

			for {

				var tmpJAgreedTask string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJAgreedTask type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJAgreedTask = string(string(outBuf))

					}
				}

				j.AgreedTask = append(j.AgreedTask, tmpJAgreedTask)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Signatures:

	/* handler: j.Signatures type=types.Signatures kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Signatures", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Signatures = nil
		} else {

			j.Signatures = []Buffer{}

			wantVal := true

			for {

				var tmpJSignatures Buffer

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJSignatures type=types.Buffer kind=slice quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJSignatures.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.Signatures = append(j.Signatures, tmpJSignatures)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RefBlockNum:

	/* handler: j.RefBlockNum type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.RefBlockNum.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RefBlockPrefix:

	/* handler: j.RefBlockPrefix type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.RefBlockPrefix.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
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

handle_Operations:

	/* handler: j.Operations type=types.Operations kind=slice quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Operations.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

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

// MarshalJSON marshal bytes to json - template
func (j *SignedTransaction) MarshalJSON() ([]byte, error) {
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
func (j *SignedTransaction) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteByte('{')
	if len(j.AgreedTask) != 0 {
		buf.WriteString(`"agreed_task":`)
		if j.AgreedTask != nil {
			buf.WriteString(`[`)
			for i, v := range j.AgreedTask {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	buf.WriteString(`"signatures":`)
	if j.Signatures != nil {
		buf.WriteString(`[`)
		for i, v := range j.Signatures {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"ref_block_num":`)
	fflib.FormatBits2(buf, uint64(j.RefBlockNum), 10, false)
	buf.WriteString(`,"ref_block_prefix":`)
	fflib.FormatBits2(buf, uint64(j.RefBlockPrefix), 10, false)
	buf.WriteString(`,"expiration":`)

	{

		obj, err = j.Expiration.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"operations":`)

	{

		obj, err = j.Operations.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtSignedTransactionbase = iota
	ffjtSignedTransactionnosuchkey

	ffjtSignedTransactionAgreedTask

	ffjtSignedTransactionSignatures

	ffjtSignedTransactionRefBlockNum

	ffjtSignedTransactionRefBlockPrefix

	ffjtSignedTransactionExpiration

	ffjtSignedTransactionOperations

	ffjtSignedTransactionExtensions
)

var ffjKeySignedTransactionAgreedTask = []byte("agreed_task")

var ffjKeySignedTransactionSignatures = []byte("signatures")

var ffjKeySignedTransactionRefBlockNum = []byte("ref_block_num")

var ffjKeySignedTransactionRefBlockPrefix = []byte("ref_block_prefix")

var ffjKeySignedTransactionExpiration = []byte("expiration")

var ffjKeySignedTransactionOperations = []byte("operations")

var ffjKeySignedTransactionExtensions = []byte("extensions")

// UnmarshalJSON umarshall json - template of ffjson
func (j *SignedTransaction) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *SignedTransaction) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtSignedTransactionbase
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
				currentKey = ffjtSignedTransactionnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeySignedTransactionAgreedTask, kn) {
						currentKey = ffjtSignedTransactionAgreedTask
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeySignedTransactionExpiration, kn) {
						currentKey = ffjtSignedTransactionExpiration
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeySignedTransactionExtensions, kn) {
						currentKey = ffjtSignedTransactionExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeySignedTransactionOperations, kn) {
						currentKey = ffjtSignedTransactionOperations
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeySignedTransactionRefBlockNum, kn) {
						currentKey = ffjtSignedTransactionRefBlockNum
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeySignedTransactionRefBlockPrefix, kn) {
						currentKey = ffjtSignedTransactionRefBlockPrefix
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeySignedTransactionSignatures, kn) {
						currentKey = ffjtSignedTransactionSignatures
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeySignedTransactionExtensions, kn) {
					currentKey = ffjtSignedTransactionExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeySignedTransactionOperations, kn) {
					currentKey = ffjtSignedTransactionOperations
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeySignedTransactionExpiration, kn) {
					currentKey = ffjtSignedTransactionExpiration
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeySignedTransactionRefBlockPrefix, kn) {
					currentKey = ffjtSignedTransactionRefBlockPrefix
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeySignedTransactionRefBlockNum, kn) {
					currentKey = ffjtSignedTransactionRefBlockNum
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeySignedTransactionSignatures, kn) {
					currentKey = ffjtSignedTransactionSignatures
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeySignedTransactionAgreedTask, kn) {
					currentKey = ffjtSignedTransactionAgreedTask
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtSignedTransactionnosuchkey
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

				case ffjtSignedTransactionAgreedTask:
					goto handle_AgreedTask

				case ffjtSignedTransactionSignatures:
					goto handle_Signatures

				case ffjtSignedTransactionRefBlockNum:
					goto handle_RefBlockNum

				case ffjtSignedTransactionRefBlockPrefix:
					goto handle_RefBlockPrefix

				case ffjtSignedTransactionExpiration:
					goto handle_Expiration

				case ffjtSignedTransactionOperations:
					goto handle_Operations

				case ffjtSignedTransactionExtensions:
					goto handle_Extensions

				case ffjtSignedTransactionnosuchkey:
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

handle_AgreedTask:

	/* handler: j.AgreedTask type=types.AgreedTaskPair kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AgreedTaskPair", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.AgreedTask = nil
		} else {

			j.AgreedTask = []string{}

			wantVal := true

			for {

				var tmpJAgreedTask string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJAgreedTask type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJAgreedTask = string(string(outBuf))

					}
				}

				j.AgreedTask = append(j.AgreedTask, tmpJAgreedTask)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Signatures:

	/* handler: j.Signatures type=types.Signatures kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for Signatures", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Signatures = nil
		} else {

			j.Signatures = []Buffer{}

			wantVal := true

			for {

				var tmpJSignatures Buffer

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJSignatures type=types.Buffer kind=slice quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJSignatures.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.Signatures = append(j.Signatures, tmpJSignatures)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RefBlockNum:

	/* handler: j.RefBlockNum type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.RefBlockNum.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RefBlockPrefix:

	/* handler: j.RefBlockPrefix type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.RefBlockPrefix.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
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

handle_Operations:

	/* handler: j.Operations type=types.Operations kind=slice quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Operations.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

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
