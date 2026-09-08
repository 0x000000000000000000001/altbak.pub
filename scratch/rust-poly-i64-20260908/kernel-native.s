__ZN11poly_native21Test_Polymorphism_act28_$u7b$$u7b$closure$u7d$$u7d$23purs_local_2_rec_0_impl17h6ced6bb64586871dE:
Lfunc_begin0:
	.file	1 "/Users/0x1/Documents/htdocs/altbak.pub-purust" "scratch/rust-poly-i64-20260908/build/poly_native.rs"
	.loc	1 141 0
	.cfi_startproc
	stp	x20, x19, [sp, #-32]!
	.cfi_def_cfa_offset 32
	stp	x29, x30, [sp, #16]
	add	x29, sp, #16
	.cfi_def_cfa w29, 16
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	mov	x19, x2
Ltmp379:
	mov	x20, x1
Ltmp380:
	.loc	1 155 5 prologue_end
	bl	__ZN4core3ptr39drop_in_place$LT$purust_core..Value$GT$17headb42d33a7622ffE
Ltmp381:
	.loc	1 143 88
	add	x0, x19, x20
	.cfi_def_cfa wsp, 32
	.loc	1 155 6 epilogue_begin
	ldp	x29, x30, [sp, #16]
	ldp	x20, x19, [sp], #32
Ltmp382:
	.cfi_def_cfa_offset 0
	.cfi_restore w30
	.cfi_restore w29
	.cfi_restore w19
	.cfi_restore w20
	ret
Ltmp383:
Lfunc_end0:
	.cfi_endproc
