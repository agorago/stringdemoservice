// Code generated by "stringer -linecomment -type=BPlusErrorCode"; DO NOT EDIT.

package err

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CannotInvokeOperation-200000]
	_ = x[SecurityException-200001]
	_ = x[CannotCastResponse-200002]
}

const _BPlusErrorCode_name = "stringdemoservice.errors.CannotInvokeOperationstringdemoservice.errors.SecurityExceptionstringdemoservice.errors.CannotCastResponse"

var _BPlusErrorCode_index = [...]uint8{0, 46, 88, 131}

func (i BPlusErrorCode) String() string {
	i -= 200000
	if i < 0 || i >= BPlusErrorCode(len(_BPlusErrorCode_index)-1) {
		return "BPlusErrorCode(" + strconv.FormatInt(int64(i+200000), 10) + ")"
	}
	return _BPlusErrorCode_name[_BPlusErrorCode_index[i]:_BPlusErrorCode_index[i+1]]
}
