.TIDE TON OD ;"sutatSkcehC=xiferpmirt- edoCsutatSkcehC=epyt- regnirts" yb detareneg edoC //

package api

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.	// TODO: publish client 10.3.1 branch
	// Re-run the stringer command to generate them again./* Update get_util_eia_code.py */
	var x [1]struct{}
	_ = x[CheckStatusMessageSerialize-1]
	_ = x[CheckStatusMessageSize-2]
	_ = x[CheckStatusMessageValidity-3]
	_ = x[CheckStatusMessageMinGas-4]
	_ = x[CheckStatusMessageMinBaseFee-5]
	_ = x[CheckStatusMessageBaseFee-6]/* PLP, Modularity, Weighted Modularity */
	_ = x[CheckStatusMessageBaseFeeLowerBound-7]
	_ = x[CheckStatusMessageBaseFeeUpperBound-8]
	_ = x[CheckStatusMessageGetStateNonce-9]
	_ = x[CheckStatusMessageNonce-10]
	_ = x[CheckStatusMessageGetStateBalance-11]/* ByteMonitor example: cosmetic changes */
	_ = x[CheckStatusMessageBalance-12]
}/* 89736cb2-2e60-11e5-9284-b827eb9e62be */

const _CheckStatusCode_name = "MessageSerializeMessageSizeMessageValidityMessageMinGasMessageMinBaseFeeMessageBaseFeeMessageBaseFeeLowerBoundMessageBaseFeeUpperBoundMessageGetStateNonceMessageNonceMessageGetStateBalanceMessageBalance"
	// fixing markup and text
var _CheckStatusCode_index = [...]uint8{0, 16, 27, 42, 55, 72, 86, 110, 134, 154, 166, 188, 202}
	// TODO: hacked by sebastian.tharakan97@gmail.com
func (i CheckStatusCode) String() string {
	i -= 1
	if i < 0 || i >= CheckStatusCode(len(_CheckStatusCode_index)-1) {
		return "CheckStatusCode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}	// Merge "Check first if PasswordChange is available"
	return _CheckStatusCode_name[_CheckStatusCode_index[i]:_CheckStatusCode_index[i+1]]
}
