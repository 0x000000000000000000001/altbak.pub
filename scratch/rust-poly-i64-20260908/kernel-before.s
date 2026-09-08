__ZN11poly_before21Test_Polymorphism_act28_$u7b$$u7b$closure$u7d$$u7d$23purs_local_2_rec_0_impl17he442e8e87408f71cE:
Lfunc_begin0:
	.file	1 "/Users/0x1/Documents/htdocs/altbak.pub-purust" "scratch/rust-poly-i64-20260908/build/poly_before.rs"
	.loc	1 141 0
	.cfi_startproc
	.cfi_personality 155, _rust_eh_personality
	.cfi_lsda 16, Lexception0
	sub	sp, sp, #112
	.cfi_def_cfa_offset 112
	stp	x24, x23, [sp, #48]
	stp	x22, x21, [sp, #64]
	stp	x20, x19, [sp, #80]
	stp	x29, x30, [sp, #96]
	add	x29, sp, #96
	.cfi_def_cfa w29, 16
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	.cfi_offset w19, -24
	.cfi_offset w20, -32
	.cfi_offset w21, -40
	.cfi_offset w22, -48
	.cfi_offset w23, -56
	.cfi_offset w24, -64
	.cfi_remember_state
	mov	x20, x2
Ltmp433:
	mov	x19, x0
Ltmp434:
	mov	x21, x8
Ltmp435:
	.loc	1 143 88 prologue_end
	cbz	x1, LBB0_7
Ltmp436:
	.loc	1 0 88 is_stmt 0
	mov	x22, x1
Ltmp437:
	mov	x23, #-9223372036854775807
Ltmp438:
LBB0_2:
Ltmp0:
	.loc	1 148 107 is_stmt 1
	add	x8, sp, #24
	mov	x0, x20
	bl	__ZN57_$LT$purust_core..Value$u20$as$u20$core..clone..Clone$GT$5clone17hf4a357cde6f4f741E
Ltmp1:
Ltmp439:
Ltmp3:
	.loc	1 148 116 is_stmt 0
	add	x0, sp, #24
	bl	__ZN11purust_core5Value10unwrap_int17h6ac4fe043bdffe9aE
Ltmp4:
Ltmp440:
	.loc	1 148 92
	add	x8, x0, #1
Ltmp441:
	.file	2 "/Users/0x1/Documents/htdocs/altbak.pub-purust/run/bak/rust/output/purust_output" "purust_core/src/lib.rs"
	.loc	2 5788 42 is_stmt 1
	stp	x23, x8, [sp]
Ltmp442:
Ltmp8:
	.loc	1 148 134
	add	x0, sp, #24
	bl	__ZN4core3ptr39drop_in_place$LT$purust_core..Value$GT$17h61eca746579af08eE
Ltmp443:
Ltmp9:
Ltmp11:
	.loc	1 150 9
	mov	x0, x20
	bl	__ZN4core3ptr39drop_in_place$LT$purust_core..Value$GT$17h61eca746579af08eE
Ltmp12:
Ltmp444:
	ldr	q0, [sp]
	str	q0, [x20]
	ldr	x8, [sp, #16]
	str	x8, [x20, #16]
Ltmp445:
	.loc	1 0 0 is_stmt 0
	sub	x22, x22, #1
Ltmp446:
	cbnz	x22, LBB0_2
Ltmp447:
LBB0_7:
Ltmp16:
	.loc	1 144 87 is_stmt 1
	mov	x8, x21
	mov	x0, x20
	bl	__ZN57_$LT$purust_core..Value$u20$as$u20$core..clone..Clone$GT$5clone17hf4a357cde6f4f741E
Ltmp17:
Ltmp448:
Ltmp21:
	.loc	1 155 5
	mov	x0, x20
	bl	__ZN4core3ptr39drop_in_place$LT$purust_core..Value$GT$17h61eca746579af08eE
Ltmp22:
Ltmp449:
	mov	x0, x19
	bl	__ZN4core3ptr39drop_in_place$LT$purust_core..Value$GT$17h61eca746579af08eE
	.cfi_def_cfa wsp, 112
	.loc	1 155 6 epilogue_begin is_stmt 0
	ldp	x29, x30, [sp, #96]
	ldp	x20, x19, [sp, #80]
Ltmp450:
	ldp	x22, x21, [sp, #64]
	ldp	x24, x23, [sp, #48]
	add	sp, sp, #112
	.cfi_def_cfa_offset 0
	.cfi_restore w30
	.cfi_restore w29
	.cfi_restore w19
	.cfi_restore w20
	.cfi_restore w21
	.cfi_restore w22
	.cfi_restore w23
	.cfi_restore w24
	ret
Ltmp451:
LBB0_10:
	.cfi_restore_state
Ltmp23:
	.loc	1 0 6
	mov	x21, x0
	b	LBB0_19
Ltmp452:
LBB0_11:
Ltmp18:
	b	LBB0_13
Ltmp453:
LBB0_12:
Ltmp2:
LBB0_13:
	mov	x21, x0
	b	LBB0_18
Ltmp454:
LBB0_14:
Ltmp5:
	mov	x21, x0
Ltmp6:
Ltmp455:
	.loc	1 148 134 is_stmt 1
	add	x0, sp, #24
	bl	__ZN4core3ptr39drop_in_place$LT$purust_core..Value$GT$17h61eca746579af08eE
Ltmp7:
	b	LBB0_18
Ltmp456:
LBB0_15:
Ltmp10:
	.loc	1 0 134 is_stmt 0
	mov	x21, x0
	mov	w8, #1
	.loc	1 152 5 is_stmt 1
	tbz	w8, #0, LBB0_18
Ltmp457:
Ltmp14:
	mov	x0, sp
	bl	__ZN4core3ptr39drop_in_place$LT$purust_core..Value$GT$17h61eca746579af08eE
Ltmp15:
	b	LBB0_18
Ltmp458:
LBB0_17:
Ltmp13:
	.loc	1 0 5 is_stmt 0
	mov	x21, x0
Ltmp459:
	.loc	1 150 9 is_stmt 1
	ldr	q0, [sp]
	str	q0, [x20]
	ldr	x8, [sp, #16]
	str	x8, [x20, #16]
Ltmp460:
LBB0_18:
Ltmp19:
	.loc	1 155 5
	mov	x0, x20
	bl	__ZN4core3ptr39drop_in_place$LT$purust_core..Value$GT$17h61eca746579af08eE
Ltmp20:
Ltmp461:
LBB0_19:
Ltmp24:
	mov	x0, x19
	bl	__ZN4core3ptr39drop_in_place$LT$purust_core..Value$GT$17h61eca746579af08eE
Ltmp25:
Ltmp462:
	.loc	1 0 5 is_stmt 0
	mov	x0, x21
	bl	__Unwind_Resume
Ltmp463:
LBB0_21:
Ltmp26:
	.loc	1 141 9 is_stmt 1
	bl	__RNvNtCs6sq8b9ugfBC_4core9panicking16panic_in_cleanup
Ltmp464:
Lfunc_end0:
	.cfi_endproc
