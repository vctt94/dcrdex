// This code is available on the terms of the project LICENSE.md file,
// also available online at https://blueoakcouncil.org/license/1.0.0.

package rpcserver

import (
	"errors"
	"fmt"
	"testing"
)

func TestCheckNArgs(t *testing.T) {
	tests := []struct {
		name      string
		have      []string
		wantNArgs []int
		wantErr   bool
	}{{
		name:      "ok exact",
		have:      []string{"1", "2", "3"},
		wantNArgs: []int{3},
		wantErr:   false,
	}, {
		name:      "ok between",
		have:      []string{"1", "2", "3"},
		wantNArgs: []int{2, 4},
		wantErr:   false,
	}, {
		name:      "ok lower",
		have:      []string{"1", "2"},
		wantNArgs: []int{2, 4},
		wantErr:   false,
	}, {
		name:      "ok upper",
		have:      []string{"1", "2", "3", "4"},
		wantNArgs: []int{2, 4},
		wantErr:   false,
	}, {
		name:      "not exact",
		have:      []string{"1", "2", "3"},
		wantNArgs: []int{2},
		wantErr:   true,
	}, {
		name:      "too few",
		have:      []string{"1", "2"},
		wantNArgs: []int{3, 5},
		wantErr:   true,
	}, {
		name:      "too many",
		have:      []string{"1", "2", "3", "4", "5", "6"},
		wantNArgs: []int{2, 5},
		wantErr:   true,
	}}
	for _, test := range tests {
		err := checkNArgs(&RawParams{PWArgs: test.have, Args: test.have}, test.wantNArgs, test.wantNArgs)
		if test.wantErr {
			if err == nil {
				t.Fatalf("expected error for test %s",
					test.name)
			}
			continue
		}
		if err != nil {
			t.Fatalf("unexpected error for test %s: %v",
				test.name, err)
		}
	}
}

func TestParseNewWalletArgs(t *testing.T) {
	paramsWithAssetID := func(id string) *RawParams {
		pwArgs := []string{"password123", "password123"}
		args := []string{
			id,
			"default",
			"/home/wallet.conf",
		}
		return &RawParams{PWArgs: pwArgs, Args: args}
	}
	tests := []struct {
		name    string
		params  *RawParams
		wantErr error
	}{{
		name:   "ok",
		params: paramsWithAssetID("42"),
	}, {
		name:    "assetID is not int",
		params:  paramsWithAssetID("42.1"),
		wantErr: errArgs,
	}}
	for _, test := range tests {
		nwf, err := parseNewWalletArgs(test.params)
		if test.wantErr != nil {
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("unexpected error %v for test %s",
					err, test.name)
			}
			continue
		}
		if nwf.AppPass != test.params.PWArgs[0] {
			t.Fatalf("appPass doesn't match")
		}
		if nwf.WalletPass != test.params.PWArgs[1] {
			t.Fatalf("walletPass doesn't match")
		}
		if fmt.Sprint(nwf.AssetID) != test.params.Args[0] {
			t.Fatalf("assetID doesn't match")
		}
		if nwf.Account != test.params.Args[1] {
			t.Fatalf("account doesn't match")
		}
		if nwf.INIPath != test.params.Args[2] {
			t.Fatalf("inipath doesn't match")
		}
	}
}

func TestParseOpenWalletArgs(t *testing.T) {
	paramsWithAssetID := func(id string) *RawParams {
		pwArgs := []string{"password123"}
		args := []string{id}
		return &RawParams{PWArgs: pwArgs, Args: args}
	}
	tests := []struct {
		name    string
		params  *RawParams
		wantErr error
	}{{
		name:   "ok",
		params: paramsWithAssetID("42"),
	}, {
		name:    "assetID is not int",
		params:  paramsWithAssetID("42.1"),
		wantErr: errArgs,
	}}
	for _, test := range tests {
		owf, err := parseOpenWalletArgs(test.params)
		if test.wantErr != nil {
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("unexpected error %v for test %s",
					err, test.name)
			}
			continue
		}
		if owf.AppPass != test.params.PWArgs[0] {
			t.Fatalf("appPass doesn't match")
		}
		if fmt.Sprint(owf.AssetID) != test.params.Args[0] {
			t.Fatalf("assetID doesn't match")
		}
	}
}

func TestCheckIntArg(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantErr error
	}{{
		name: "ok",
		arg:  "42",
	}, {
		name:    "assetID is not int",
		arg:     "42.1",
		wantErr: errArgs,
	}}
	for _, test := range tests {
		res, err := checkIntArg(test.arg, "name")
		if test.wantErr != nil {
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("unexpected error %v for test %s",
					err, test.name)
			}
			continue
		}
		if fmt.Sprint(res) != test.arg {
			t.Fatalf("strings don't match")
		}
	}
}

func TestCheckBoolArg(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    bool
		wantErr error
	}{{
		name: "ok string lower",
		arg:  "true",
		want: true,
	}, {
		name: "ok string upper",
		arg:  "False",
		want: false,
	}, {
		name: "ok int",
		arg:  "1",
		want: true,
	}, {
		name:    "string but not true or false",
		arg:     "blue",
		wantErr: errArgs,
	}, {
		name:    "int but not 0 or 1",
		arg:     "2",
		wantErr: errArgs,
	}}
	for _, test := range tests {
		res, err := checkBoolArg(test.arg, "name")
		if err != nil {
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("unexpected error %v for test %s",
					err, test.name)
			}
			continue
		}
		if res != test.want {
			t.Fatalf("wanted %v but got %v for test %v", test.want, res, test.name)
		}
	}
}

func TestParseRegisterArgs(t *testing.T) {
	paramsWithFee := func(fee string) *RawParams {
		pwArgs := []string{"password123"}
		args := []string{"dex", fee, "cert"}
		return &RawParams{PWArgs: pwArgs, Args: args}
	}
	tests := []struct {
		name    string
		params  *RawParams
		wantErr error
	}{{
		name:   "ok",
		params: paramsWithFee("1000"),
	}, {
		name:    "fee not int",
		params:  paramsWithFee("1000.0"),
		wantErr: errArgs,
	}}
	for _, test := range tests {
		reg, err := parseRegisterArgs(test.params)
		if test.wantErr != nil {
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("unexpected error %v for test %s",
					err, test.name)
			}
			continue
		}
		if reg.AppPass != test.params.PWArgs[0] {
			t.Fatalf("appPass doesn't match")
		}
		if reg.URL != test.params.Args[0] {
			t.Fatalf("url doesn't match")
		}
		if fmt.Sprint(reg.Fee) != test.params.Args[1] {
			t.Fatalf("fee doesn't match")
		}
		if fmt.Sprint(reg.Cert) != test.params.Args[2] {
			t.Fatalf("cert doesn't match")
		}
	}
}

func TestParseHelpArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    *helpForm
		wantErr error
	}{{
		name: "ok no args",
		want: &helpForm{},
	}, {
		name: "ok help with",
		args: []string{"thing"},
		want: &helpForm{HelpWith: "thing"},
	}, {
		name: "ok help with include passwords",
		args: []string{"thing", "true"},
		want: &helpForm{HelpWith: "thing", IncludePasswords: true},
	}, {
		name:    "include passwords not boolean",
		args:    []string{"thing", "thing2"},
		wantErr: errArgs,
	}}
	for _, test := range tests {
		form, err := parseHelpArgs(&RawParams{Args: test.args})
		if err != nil {
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("unexpected error %v for test %s",
					err, test.name)
			}
			continue
		}
		if len(test.args) > 0 && form.HelpWith != test.args[0] {
			t.Fatalf("helpwith doesn't match")
		}
		if len(test.args) > 1 && fmt.Sprint(form.IncludePasswords) != test.args[1] {
			t.Fatalf("includepasswords doesn't match")
		}
	}
}
