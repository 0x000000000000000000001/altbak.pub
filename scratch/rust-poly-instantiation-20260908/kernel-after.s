__ZN13poly_verified21Test_Polymorphism_act28_$u7b$$u7b$closure$u7d$$u7d$23purs_local_2_rec_0_impl17h74268b3ff98dd116E:
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
	mov	x20, x1
	bl	__ZN4core3ptr39drop_in_place$LT$purust_core..Value$GT$17ha2c3ece851164f3bE
	add	x0, x19, x20
	.cfi_def_cfa wsp, 32
	ldp	x29, x30, [sp, #16]
	ldp	x20, x19, [sp], #32
	.cfi_def_cfa_offset 0
	.cfi_restore w30
	.cfi_restore w29
	.cfi_restore w19
	.cfi_restore w20
	ret
	.cfi_endproc
