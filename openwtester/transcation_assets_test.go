/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package openwtester

import (
	"github.com/blocktree/openwallet/openw"
	"testing"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
)

func testGetAssetsAccountBalance(tm *openw.WalletManager, walletID, accountID string) {
	balance, err := tm.GetAssetsAccountBalance(testApp, walletID, accountID)
	if err != nil {
		log.Error("GetAssetsAccountBalance failed, unexpected error:", err)
		return
	}
	log.Info("balance:", balance)
}

func testGetAssetsAccountTokenBalance(tm *openw.WalletManager, walletID, accountID string, contract openwallet.SmartContract) {
	balance, err := tm.GetAssetsAccountTokenBalance(testApp, walletID, accountID, contract)
	if err != nil {
		log.Error("GetAssetsAccountTokenBalance failed, unexpected error:", err)
		return
	}
	log.Info("token balance:", balance.Balance)
}

func testCreateTransactionStep(tm *openw.WalletManager, walletID, accountID, to, amount, feeRate string, contract *openwallet.SmartContract) (*openwallet.RawTransaction, error) {

	//err := tm.RefreshAssetsAccountBalance(testApp, accountID)
	//if err != nil {
	//	log.Error("RefreshAssetsAccountBalance failed, unexpected error:", err)
	//	return nil, err
	//}

	rawTx, err := tm.CreateTransaction(testApp, walletID, accountID, amount, to, feeRate, "", contract)

	if err != nil {
		log.Error("CreateTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTx, nil
}

func testCreateSummaryTransactionStep(
	tm *openw.WalletManager,
	walletID, accountID, summaryAddress, minTransfer, retainedBalance, feeRate string,
	start, limit int,
	contract *openwallet.SmartContract) ([]*openwallet.RawTransaction, error) {

	rawTxArray, err := tm.CreateSummaryTransaction(testApp, walletID, accountID, summaryAddress, minTransfer,
		retainedBalance, feeRate, start, limit, contract)

	if err != nil {
		log.Error("CreateSummaryTransaction failed, unexpected error:", err)
		return nil, err
	}

	return rawTxArray, nil
}

func testSignTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	_, err := tm.SignTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, "12345678", rawTx)
	if err != nil {
		log.Error("SignTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testVerifyTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	//log.Info("rawTx.Signatures:", rawTx.Signatures)

	_, err := tm.VerifyTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("VerifyTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Infof("rawTx: %+v", rawTx)
	return rawTx, nil
}

func testSubmitTransactionStep(tm *openw.WalletManager, rawTx *openwallet.RawTransaction) (*openwallet.RawTransaction, error) {

	tx, err := tm.SubmitTransaction(testApp, rawTx.Account.WalletID, rawTx.Account.AccountID, rawTx)
	if err != nil {
		log.Error("SubmitTransaction failed, unexpected error:", err)
		return nil, err
	}

	log.Std.Info("tx: %+v", tx)
	log.Info("wxID:", tx.WxID)
	log.Info("txID:", rawTx.TxID)

	return rawTx, nil
}

func TestTransfer_NTN(t *testing.T) {

	addrs := []string{
		//"0x86b887251fd50c1a8777c02857f430ce112dced8",
		//"0xc0bdaae19f31abd801291172d434b366256baa32",
		//"0xb3377952b690187e8b042e30b2a1afc0c6ce7e81",
		//"0x95b2ae4609abec5238b268baadce3feeeb0ae182",
		//"0xfc3462edb9836d815af84daf868df11ea483aad0",
		//"0x0c27e32d7502dad762df0f83d4ecf6e5f8c26636",
		//"0x7a9e5283f2b150716f45240e04a0848c63e77fa4",
		//"0xe48e16b2d32ba89896cc6d8e058feacde09b1ba6",
		//"0xf66c46617094b36b7cae1d97956d5d0e2fd3c824",
		//"0x0191c1a14e16a3f64a6617d9764c5efddbc9176e",

		"0x21d25f3c33208738900a64d4dd45ec4e6c9ef929",
	}

	tm := testInitWalletManager()
	walletID := "Vzx4RZ5ZzCzWjHPx2o7p4hfF1vf85LxMA4"
	accountID := "BQFx3sxavFeHq8u5xVui981zY2eTQaHzvvxeDYHiwFP2"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	for _, to := range addrs {
		rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "0.1", "", nil)
		if err != nil {
			return
		}

		log.Std.Info("rawTx: %+v", rawTx)

		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}
}

func TestTransfer_ERC20(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMTUzB3LWaSKNKEQw9Sn73FjkEoYGHEp4B"
	accountID := "59t47qyjHUMZ6PGAdjkJopE9ffAPUkdUhSinJqcWRYZ1"
	to := "0xd35f9Ea14D063af9B3567064FAB567275b09f03D"

	contract := openwallet.SmartContract{
		Address:  "4092678e4E78230F46A1534C0fbc8fA39780892B",
		Symbol:   "NTN",
		Name:     "OCoin",
		Token:    "OCN",
		Decimals: 18,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTx, err := testCreateTransactionStep(tm, walletID, accountID, to, "12.34567", "", &contract)
	if err != nil {
		return
	}

	_, err = testSignTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testVerifyTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

	_, err = testSubmitTransactionStep(tm, rawTx)
	if err != nil {
		return
	}

}

func TestSummary_NTN(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "Vzx4RZ5ZzCzWjHPx2o7p4hfF1vf85LxMA4"
	accountID := "HbsfbSFUk4Cpjv2zBScWxDSVoJYg1A2t7SZ63348Fs3f"
	summaryAddress := "0x58fb97021ff7ec9fa8c5c88340236dff9f261b61"

	testGetAssetsAccountBalance(tm, walletID, accountID)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, nil)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTx := range rawTxArray {
		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}

}

func TestSummary_ERC20(t *testing.T) {
	tm := testInitWalletManager()
	walletID := "WMTUzB3LWaSKNKEQw9Sn73FjkEoYGHEp4B"
	accountID := "59t47qyjHUMZ6PGAdjkJopE9ffAPUkdUhSinJqcWRYZ1"
	summaryAddress := "0xd35f9Ea14D063af9B3567064FAB567275b09f03D"

	contract := openwallet.SmartContract{
		Address:  "4092678e4E78230F46A1534C0fbc8fA39780892B",
		Symbol:   "NTN",
		Name:     "OCoin",
		Token:    "OCN",
		Decimals: 18,
	}

	testGetAssetsAccountBalance(tm, walletID, accountID)

	testGetAssetsAccountTokenBalance(tm, walletID, accountID, contract)

	rawTxArray, err := testCreateSummaryTransactionStep(tm, walletID, accountID,
		summaryAddress, "", "", "",
		0, 100, &contract)
	if err != nil {
		log.Errorf("CreateSummaryTransaction failed, unexpected error: %v", err)
		return
	}

	//执行汇总交易
	for _, rawTx := range rawTxArray {
		_, err = testSignTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testVerifyTransactionStep(tm, rawTx)
		if err != nil {
			return
		}

		_, err = testSubmitTransactionStep(tm, rawTx)
		if err != nil {
			return
		}
	}

}
