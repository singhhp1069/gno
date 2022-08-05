// Code generated by "stringer -type=Op"; DO NOT EDIT.

package gno

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpInvalid-0]
	_ = x[OpHalt-1]
	_ = x[OpNoop-2]
	_ = x[OpExec-3]
	_ = x[OpPrecall-4]
	_ = x[OpCall-5]
	_ = x[OpCallNativeBody-6]
	_ = x[OpReturn-7]
	_ = x[OpReturnFromBlock-8]
	_ = x[OpReturnToBlock-9]
	_ = x[OpDefer-10]
	_ = x[OpCallDeferNativeBody-11]
	_ = x[OpGo-12]
	_ = x[OpSelect-13]
	_ = x[OpSwitchClause-14]
	_ = x[OpSwitchClauseCase-15]
	_ = x[OpTypeSwitch-16]
	_ = x[OpIfCond-17]
	_ = x[OpPopValue-18]
	_ = x[OpPopResults-19]
	_ = x[OpPopBlock-20]
	_ = x[OpPopFrameAndReset-21]
	_ = x[OpPanic1-22]
	_ = x[OpPanic2-23]
	_ = x[OpUpos-32]
	_ = x[OpUneg-33]
	_ = x[OpUnot-34]
	_ = x[OpUxor-35]
	_ = x[OpUrecv-37]
	_ = x[OpLor-38]
	_ = x[OpLand-39]
	_ = x[OpEql-40]
	_ = x[OpNeq-41]
	_ = x[OpLss-42]
	_ = x[OpLeq-43]
	_ = x[OpGtr-44]
	_ = x[OpGeq-45]
	_ = x[OpAdd-46]
	_ = x[OpSub-47]
	_ = x[OpBor-48]
	_ = x[OpXor-49]
	_ = x[OpMul-50]
	_ = x[OpQuo-51]
	_ = x[OpRem-52]
	_ = x[OpShl-53]
	_ = x[OpShr-54]
	_ = x[OpBand-55]
	_ = x[OpBandn-56]
	_ = x[OpEval-64]
	_ = x[OpBinary1-65]
	_ = x[OpIndex1-66]
	_ = x[OpIndex2-67]
	_ = x[OpSelector-68]
	_ = x[OpSlice-69]
	_ = x[OpStar-70]
	_ = x[OpRef-71]
	_ = x[OpTypeAssert1-72]
	_ = x[OpTypeAssert2-73]
	_ = x[OpStaticTypeOf-74]
	_ = x[OpCompositeLit-75]
	_ = x[OpArrayLit-76]
	_ = x[OpSliceLit-77]
	_ = x[OpSliceLit2-78]
	_ = x[OpMapLit-79]
	_ = x[OpStructLit-80]
	_ = x[OpFuncLit-81]
	_ = x[OpConvert-82]
	_ = x[OpArrayLitGoNative-96]
	_ = x[OpSliceLitGoNative-97]
	_ = x[OpStructLitGoNative-98]
	_ = x[OpCallGoNative-99]
	_ = x[OpFieldType-112]
	_ = x[OpArrayType-113]
	_ = x[OpSliceType-114]
	_ = x[OpPointerType-115]
	_ = x[OpInterfaceType-116]
	_ = x[OpChanType-117]
	_ = x[OpFuncType-118]
	_ = x[OpMapType-119]
	_ = x[OpStructType-120]
	_ = x[OpMaybeNativeType-121]
	_ = x[OpAssign-128]
	_ = x[OpAddAssign-129]
	_ = x[OpSubAssign-130]
	_ = x[OpMulAssign-131]
	_ = x[OpQuoAssign-132]
	_ = x[OpRemAssign-133]
	_ = x[OpBandAssign-134]
	_ = x[OpBandnAssign-135]
	_ = x[OpBorAssign-136]
	_ = x[OpXorAssign-137]
	_ = x[OpShlAssign-138]
	_ = x[OpShrAssign-139]
	_ = x[OpDefine-140]
	_ = x[OpInc-141]
	_ = x[OpDec-142]
	_ = x[OpValueDecl-144]
	_ = x[OpTypeDecl-145]
	_ = x[OpSticky-208]
	_ = x[OpBody-209]
	_ = x[OpForLoop-210]
	_ = x[OpRangeIter-211]
	_ = x[OpRangeIterString-212]
	_ = x[OpRangeIterMap-213]
	_ = x[OpRangeIterArrayPtr-214]
	_ = x[OpReturnCallDefers-215]
}

const (
	_Op_name_0 = "OpInvalidOpHaltOpNoopOpExecOpPrecallOpCallOpCallNativeBodyOpReturnOpReturnFromBlockOpReturnToBlockOpDeferOpCallDeferNativeBodyOpGoOpSelectOpSwitchClauseOpSwitchClauseCaseOpTypeSwitchOpIfCondOpPopValueOpPopResultsOpPopBlockOpPopFrameAndResetOpPanic1OpPanic2"
	_Op_name_1 = "OpUposOpUnegOpUnotOpUxor"
	_Op_name_2 = "OpUrecvOpLorOpLandOpEqlOpNeqOpLssOpLeqOpGtrOpGeqOpAddOpSubOpBorOpXorOpMulOpQuoOpRemOpShlOpShrOpBandOpBandn"
	_Op_name_3 = "OpEvalOpBinary1OpIndex1OpIndex2OpSelectorOpSliceOpStarOpRefOpTypeAssert1OpTypeAssert2OpStaticTypeOfOpCompositeLitOpArrayLitOpSliceLitOpSliceLit2OpMapLitOpStructLitOpFuncLitOpConvert"
	_Op_name_4 = "OpArrayLitGoNativeOpSliceLitGoNativeOpStructLitGoNativeOpCallGoNative"
	_Op_name_5 = "OpFieldTypeOpArrayTypeOpSliceTypeOpPointerTypeOpInterfaceTypeOpChanTypeOpFuncTypeOpMapTypeOpStructTypeOpMaybeNativeType"
	_Op_name_6 = "OpAssignOpAddAssignOpSubAssignOpMulAssignOpQuoAssignOpRemAssignOpBandAssignOpBandnAssignOpBorAssignOpXorAssignOpShlAssignOpShrAssignOpDefineOpIncOpDec"
	_Op_name_7 = "OpValueDeclOpTypeDecl"
	_Op_name_8 = "OpStickyOpBodyOpForLoopOpRangeIterOpRangeIterStringOpRangeIterMapOpRangeIterArrayPtrOpReturnCallDefers"
)

var (
	_Op_index_0 = [...]uint16{0, 9, 15, 21, 27, 36, 42, 58, 66, 83, 98, 105, 126, 130, 138, 152, 170, 182, 190, 200, 212, 222, 240, 248, 256}
	_Op_index_1 = [...]uint8{0, 6, 12, 18, 24}
	_Op_index_2 = [...]uint8{0, 7, 12, 18, 23, 28, 33, 38, 43, 48, 53, 58, 63, 68, 73, 78, 83, 88, 93, 99, 106}
	_Op_index_3 = [...]uint8{0, 6, 15, 23, 31, 41, 48, 54, 59, 72, 85, 99, 113, 123, 133, 144, 152, 163, 172, 181}
	_Op_index_4 = [...]uint8{0, 18, 36, 55, 69}
	_Op_index_5 = [...]uint8{0, 11, 22, 33, 46, 61, 71, 81, 90, 102, 119}
	_Op_index_6 = [...]uint8{0, 8, 19, 30, 41, 52, 63, 75, 88, 99, 110, 121, 132, 140, 145, 150}
	_Op_index_7 = [...]uint8{0, 11, 21}
	_Op_index_8 = [...]uint8{0, 8, 14, 23, 34, 51, 65, 84, 102}
)

func (i Op) String() string {
	switch {
	case i <= 23:
		return _Op_name_0[_Op_index_0[i]:_Op_index_0[i+1]]
	case 32 <= i && i <= 35:
		i -= 32
		return _Op_name_1[_Op_index_1[i]:_Op_index_1[i+1]]
	case 37 <= i && i <= 56:
		i -= 37
		return _Op_name_2[_Op_index_2[i]:_Op_index_2[i+1]]
	case 64 <= i && i <= 82:
		i -= 64
		return _Op_name_3[_Op_index_3[i]:_Op_index_3[i+1]]
	case 96 <= i && i <= 99:
		i -= 96
		return _Op_name_4[_Op_index_4[i]:_Op_index_4[i+1]]
	case 112 <= i && i <= 121:
		i -= 112
		return _Op_name_5[_Op_index_5[i]:_Op_index_5[i+1]]
	case 128 <= i && i <= 142:
		i -= 128
		return _Op_name_6[_Op_index_6[i]:_Op_index_6[i+1]]
	case 144 <= i && i <= 145:
		i -= 144
		return _Op_name_7[_Op_index_7[i]:_Op_index_7[i+1]]
	case 208 <= i && i <= 215:
		i -= 208
		return _Op_name_8[_Op_index_8[i]:_Op_index_8[i+1]]
	default:
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
