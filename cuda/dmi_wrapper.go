package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var adddmi_code cu.Function

type adddmi_args struct {
	arg_Hx unsafe.Pointer
	arg_Hy unsafe.Pointer
	arg_Hz unsafe.Pointer
	arg_mx unsafe.Pointer
	arg_my unsafe.Pointer
	arg_mz unsafe.Pointer
	arg_Dx float32
	arg_Dy float32
	arg_Dz float32
	arg_N0 int
	arg_N1 int
	arg_N2 int
	argptr [12]unsafe.Pointer
}

// Wrapper for adddmi CUDA kernel, asynchronous.
func k_adddmi_async(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Dx float32, Dy float32, Dz float32, N0 int, N1 int, N2 int, cfg *config, str cu.Stream) {
	if adddmi_code == 0 {
		adddmi_code = fatbinLoad(adddmi_map, "adddmi")
	}

	var a adddmi_args

	a.arg_Hx = Hx
	a.argptr[0] = unsafe.Pointer(&a.arg_Hx)
	a.arg_Hy = Hy
	a.argptr[1] = unsafe.Pointer(&a.arg_Hy)
	a.arg_Hz = Hz
	a.argptr[2] = unsafe.Pointer(&a.arg_Hz)
	a.arg_mx = mx
	a.argptr[3] = unsafe.Pointer(&a.arg_mx)
	a.arg_my = my
	a.argptr[4] = unsafe.Pointer(&a.arg_my)
	a.arg_mz = mz
	a.argptr[5] = unsafe.Pointer(&a.arg_mz)
	a.arg_Dx = Dx
	a.argptr[6] = unsafe.Pointer(&a.arg_Dx)
	a.arg_Dy = Dy
	a.argptr[7] = unsafe.Pointer(&a.arg_Dy)
	a.arg_Dz = Dz
	a.argptr[8] = unsafe.Pointer(&a.arg_Dz)
	a.arg_N0 = N0
	a.argptr[9] = unsafe.Pointer(&a.arg_N0)
	a.arg_N1 = N1
	a.argptr[10] = unsafe.Pointer(&a.arg_N1)
	a.arg_N2 = N2
	a.argptr[11] = unsafe.Pointer(&a.arg_N2)

	args := a.argptr[:]
	cu.LaunchKernel(adddmi_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for adddmi CUDA kernel, synchronized.
func k_adddmi(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Dx float32, Dy float32, Dz float32, N0 int, N1 int, N2 int, cfg *config) {
	str := stream()
	k_adddmi_async(Hx, Hy, Hz, mx, my, mz, Dx, Dy, Dz, N0, N1, N2, cfg, str)
	syncAndRecycle(str)
}

var adddmi_map = map[int]string{0: "",
	20: adddmi_ptx_20,
	30: adddmi_ptx_30,
	35: adddmi_ptx_35}

const (
	adddmi_ptx_20 = `
.version 3.1
.target sm_20
.address_size 64


.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .u32 adddmi_param_9,
	.param .u32 adddmi_param_10,
	.param .u32 adddmi_param_11
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<66>;
	.reg .f32 	%f<24>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd3, [adddmi_param_0];
	ld.param.u64 	%rd4, [adddmi_param_1];
	ld.param.u64 	%rd5, [adddmi_param_2];
	ld.param.u64 	%rd7, [adddmi_param_3];
	ld.param.u64 	%rd8, [adddmi_param_4];
	ld.param.u64 	%rd6, [adddmi_param_5];
	ld.param.f32 	%f1, [adddmi_param_7];
	ld.param.f32 	%f2, [adddmi_param_8];
	ld.param.u32 	%r22, [adddmi_param_9];
	ld.param.u32 	%r23, [adddmi_param_10];
	ld.param.u32 	%r24, [adddmi_param_11];
	cvta.to.global.u64 	%rd1, %rd7;
	cvta.to.global.u64 	%rd2, %rd8;
	.loc 2 13 1
	mov.u32 	%r1, %ntid.x;
	mov.u32 	%r2, %ctaid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 2 14 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 16 1
	setp.lt.s32 	%p1, %r8, %r24;
	setp.lt.s32 	%p2, %r4, %r23;
	and.pred  	%p3, %p2, %p1;
	.loc 2 20 1
	setp.gt.s32 	%p4, %r22, 0;
	.loc 2 16 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_3;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 25 1
	add.s32 	%r9, %r22, -1;
	add.s32 	%r26, %r23, -1;
	mov.u32 	%r25, 0;
	.loc 3 238 5
	max.s32 	%r27, %r4, %r25;
	.loc 3 210 5
	min.s32 	%r10, %r27, %r26;
	.loc 2 25 1
	add.s32 	%r28, %r8, 1;
	.loc 3 238 5
	max.s32 	%r29, %r28, %r25;
	.loc 2 25 1
	add.s32 	%r30, %r24, -1;
	.loc 3 210 5
	min.s32 	%r11, %r29, %r30;
	.loc 2 25 1
	add.s32 	%r31, %r8, -1;
	.loc 3 238 5
	max.s32 	%r32, %r31, %r25;
	.loc 3 210 5
	min.s32 	%r12, %r32, %r30;
	.loc 2 26 1
	add.s32 	%r33, %r4, 1;
	.loc 3 238 5
	max.s32 	%r34, %r33, %r25;
	.loc 3 210 5
	min.s32 	%r13, %r34, %r26;
	.loc 3 238 5
	max.s32 	%r35, %r8, %r25;
	.loc 3 210 5
	min.s32 	%r14, %r35, %r30;
	.loc 2 26 1
	add.s32 	%r36, %r4, -1;
	.loc 3 238 5
	max.s32 	%r37, %r36, %r25;
	.loc 3 210 5
	min.s32 	%r15, %r37, %r26;
	.loc 2 20 1
	mad.lo.s32 	%r64, %r24, %r4, %r8;
	mul.lo.s32 	%r17, %r24, %r23;
	cvta.to.global.u64 	%rd9, %rd3;
	cvta.to.global.u64 	%rd12, %rd4;
	cvta.to.global.u64 	%rd14, %rd5;
	cvta.to.global.u64 	%rd16, %rd6;
	mov.u32 	%r65, %r25;

BB0_2:
	.loc 2 23 1
	mov.u32 	%r19, %r65;
	mul.wide.s32 	%rd10, %r64, 4;
	add.s64 	%rd11, %rd9, %rd10;
	add.s64 	%rd13, %rd12, %rd10;
	add.s64 	%rd15, %rd14, %rd10;
	.loc 3 238 5
	max.s32 	%r41, %r19, %r25;
	.loc 3 210 5
	min.s32 	%r42, %r41, %r9;
	.loc 2 25 1
	mad.lo.s32 	%r43, %r42, %r23, %r10;
	mad.lo.s32 	%r44, %r43, %r24, %r11;
	mul.wide.s32 	%rd17, %r44, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mad.lo.s32 	%r45, %r43, %r24, %r12;
	mul.wide.s32 	%rd19, %r45, 4;
	add.s64 	%rd20, %rd16, %rd19;
	ld.global.f32 	%f3, [%rd20];
	ld.global.f32 	%f4, [%rd18];
	sub.f32 	%f5, %f4, %f3;
	.loc 2 23 1
	ld.global.f32 	%f6, [%rd11];
	.loc 2 25 1
	fma.rn.f32 	%f7, %f5, %f2, %f6;
	.loc 2 26 1
	mad.lo.s32 	%r49, %r42, %r23, %r13;
	mad.lo.s32 	%r50, %r49, %r24, %r14;
	mul.wide.s32 	%rd21, %r50, 4;
	add.s64 	%rd22, %rd2, %rd21;
	mad.lo.s32 	%r51, %r42, %r23, %r15;
	mad.lo.s32 	%r52, %r51, %r24, %r14;
	mul.wide.s32 	%rd23, %r52, 4;
	add.s64 	%rd24, %rd2, %rd23;
	ld.global.f32 	%f8, [%rd24];
	ld.global.f32 	%f9, [%rd22];
	sub.f32 	%f10, %f9, %f8;
	fma.rn.f32 	%f11, %f10, %f1, %f7;
	.loc 2 28 1
	add.s64 	%rd25, %rd1, %rd21;
	add.s64 	%rd26, %rd1, %rd23;
	ld.global.f32 	%f12, [%rd26];
	ld.global.f32 	%f13, [%rd25];
	sub.f32 	%f14, %f13, %f12;
	mul.f32 	%f15, %f14, %f1;
	.loc 2 23 1
	ld.global.f32 	%f16, [%rd13];
	.loc 2 28 1
	sub.f32 	%f17, %f16, %f15;
	.loc 2 29 1
	add.s64 	%rd27, %rd1, %rd17;
	add.s64 	%rd28, %rd1, %rd19;
	ld.global.f32 	%f18, [%rd28];
	ld.global.f32 	%f19, [%rd27];
	sub.f32 	%f20, %f19, %f18;
	mul.f32 	%f21, %f20, %f2;
	.loc 2 23 1
	ld.global.f32 	%f22, [%rd15];
	.loc 2 29 1
	sub.f32 	%f23, %f22, %f21;
	.loc 2 32 1
	st.global.f32 	[%rd11], %f11;
	.loc 2 33 1
	st.global.f32 	[%rd13], %f17;
	.loc 2 34 1
	st.global.f32 	[%rd15], %f23;
	.loc 2 20 1
	add.s32 	%r64, %r64, %r17;
	.loc 2 20 18
	add.s32 	%r21, %r19, 1;
	.loc 2 20 1
	setp.lt.s32 	%p6, %r21, %r22;
	mov.u32 	%r65, %r21;
	@%p6 bra 	BB0_2;

BB0_3:
	.loc 2 36 2
	ret;
}


`
	adddmi_ptx_30 = `
.version 3.1
.target sm_30
.address_size 64


.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .u32 adddmi_param_9,
	.param .u32 adddmi_param_10,
	.param .u32 adddmi_param_11
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<66>;
	.reg .f32 	%f<24>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd3, [adddmi_param_0];
	ld.param.u64 	%rd4, [adddmi_param_1];
	ld.param.u64 	%rd5, [adddmi_param_2];
	ld.param.u64 	%rd7, [adddmi_param_3];
	ld.param.u64 	%rd8, [adddmi_param_4];
	ld.param.u64 	%rd6, [adddmi_param_5];
	ld.param.f32 	%f1, [adddmi_param_7];
	ld.param.f32 	%f2, [adddmi_param_8];
	ld.param.u32 	%r22, [adddmi_param_9];
	ld.param.u32 	%r23, [adddmi_param_10];
	ld.param.u32 	%r24, [adddmi_param_11];
	cvta.to.global.u64 	%rd1, %rd7;
	cvta.to.global.u64 	%rd2, %rd8;
	.loc 2 13 1
	mov.u32 	%r1, %ntid.x;
	mov.u32 	%r2, %ctaid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 2 14 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 2 16 1
	setp.lt.s32 	%p1, %r8, %r24;
	setp.lt.s32 	%p2, %r4, %r23;
	and.pred  	%p3, %p2, %p1;
	.loc 2 20 1
	setp.gt.s32 	%p4, %r22, 0;
	.loc 2 16 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_3;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 25 1
	add.s32 	%r9, %r22, -1;
	add.s32 	%r26, %r23, -1;
	mov.u32 	%r25, 0;
	.loc 3 238 5
	max.s32 	%r27, %r4, %r25;
	.loc 3 210 5
	min.s32 	%r10, %r27, %r26;
	.loc 2 25 1
	add.s32 	%r28, %r8, 1;
	.loc 3 238 5
	max.s32 	%r29, %r28, %r25;
	.loc 2 25 1
	add.s32 	%r30, %r24, -1;
	.loc 3 210 5
	min.s32 	%r11, %r29, %r30;
	.loc 2 25 1
	add.s32 	%r31, %r8, -1;
	.loc 3 238 5
	max.s32 	%r32, %r31, %r25;
	.loc 3 210 5
	min.s32 	%r12, %r32, %r30;
	.loc 2 26 1
	add.s32 	%r33, %r4, 1;
	.loc 3 238 5
	max.s32 	%r34, %r33, %r25;
	.loc 3 210 5
	min.s32 	%r13, %r34, %r26;
	.loc 3 238 5
	max.s32 	%r35, %r8, %r25;
	.loc 3 210 5
	min.s32 	%r14, %r35, %r30;
	.loc 2 26 1
	add.s32 	%r36, %r4, -1;
	.loc 3 238 5
	max.s32 	%r37, %r36, %r25;
	.loc 3 210 5
	min.s32 	%r15, %r37, %r26;
	.loc 2 20 1
	mad.lo.s32 	%r64, %r24, %r4, %r8;
	mul.lo.s32 	%r17, %r24, %r23;
	cvta.to.global.u64 	%rd9, %rd3;
	cvta.to.global.u64 	%rd12, %rd4;
	cvta.to.global.u64 	%rd14, %rd5;
	cvta.to.global.u64 	%rd16, %rd6;
	mov.u32 	%r65, %r25;

BB0_2:
	.loc 2 23 1
	mov.u32 	%r19, %r65;
	mul.wide.s32 	%rd10, %r64, 4;
	add.s64 	%rd11, %rd9, %rd10;
	add.s64 	%rd13, %rd12, %rd10;
	add.s64 	%rd15, %rd14, %rd10;
	.loc 3 238 5
	max.s32 	%r41, %r19, %r25;
	.loc 3 210 5
	min.s32 	%r42, %r41, %r9;
	.loc 2 25 1
	mad.lo.s32 	%r43, %r42, %r23, %r10;
	mad.lo.s32 	%r44, %r43, %r24, %r11;
	mul.wide.s32 	%rd17, %r44, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mad.lo.s32 	%r45, %r43, %r24, %r12;
	mul.wide.s32 	%rd19, %r45, 4;
	add.s64 	%rd20, %rd16, %rd19;
	ld.global.f32 	%f3, [%rd20];
	ld.global.f32 	%f4, [%rd18];
	sub.f32 	%f5, %f4, %f3;
	.loc 2 23 1
	ld.global.f32 	%f6, [%rd11];
	.loc 2 25 1
	fma.rn.f32 	%f7, %f5, %f2, %f6;
	.loc 2 26 1
	mad.lo.s32 	%r49, %r42, %r23, %r13;
	mad.lo.s32 	%r50, %r49, %r24, %r14;
	mul.wide.s32 	%rd21, %r50, 4;
	add.s64 	%rd22, %rd2, %rd21;
	mad.lo.s32 	%r51, %r42, %r23, %r15;
	mad.lo.s32 	%r52, %r51, %r24, %r14;
	mul.wide.s32 	%rd23, %r52, 4;
	add.s64 	%rd24, %rd2, %rd23;
	ld.global.f32 	%f8, [%rd24];
	ld.global.f32 	%f9, [%rd22];
	sub.f32 	%f10, %f9, %f8;
	fma.rn.f32 	%f11, %f10, %f1, %f7;
	.loc 2 28 1
	add.s64 	%rd25, %rd1, %rd21;
	add.s64 	%rd26, %rd1, %rd23;
	ld.global.f32 	%f12, [%rd26];
	ld.global.f32 	%f13, [%rd25];
	sub.f32 	%f14, %f13, %f12;
	mul.f32 	%f15, %f14, %f1;
	.loc 2 23 1
	ld.global.f32 	%f16, [%rd13];
	.loc 2 28 1
	sub.f32 	%f17, %f16, %f15;
	.loc 2 29 1
	add.s64 	%rd27, %rd1, %rd17;
	add.s64 	%rd28, %rd1, %rd19;
	ld.global.f32 	%f18, [%rd28];
	ld.global.f32 	%f19, [%rd27];
	sub.f32 	%f20, %f19, %f18;
	mul.f32 	%f21, %f20, %f2;
	.loc 2 23 1
	ld.global.f32 	%f22, [%rd15];
	.loc 2 29 1
	sub.f32 	%f23, %f22, %f21;
	.loc 2 32 1
	st.global.f32 	[%rd11], %f11;
	.loc 2 33 1
	st.global.f32 	[%rd13], %f17;
	.loc 2 34 1
	st.global.f32 	[%rd15], %f23;
	.loc 2 20 1
	add.s32 	%r64, %r64, %r17;
	.loc 2 20 18
	add.s32 	%r21, %r19, 1;
	.loc 2 20 1
	setp.lt.s32 	%p6, %r21, %r22;
	mov.u32 	%r65, %r21;
	@%p6 bra 	BB0_2;

BB0_3:
	.loc 2 36 2
	ret;
}


`
	adddmi_ptx_35 = `
.version 3.1
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry adddmi(
	.param .u64 adddmi_param_0,
	.param .u64 adddmi_param_1,
	.param .u64 adddmi_param_2,
	.param .u64 adddmi_param_3,
	.param .u64 adddmi_param_4,
	.param .u64 adddmi_param_5,
	.param .f32 adddmi_param_6,
	.param .f32 adddmi_param_7,
	.param .f32 adddmi_param_8,
	.param .u32 adddmi_param_9,
	.param .u32 adddmi_param_10,
	.param .u32 adddmi_param_11
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<58>;
	.reg .f32 	%f<24>;
	.reg .s64 	%rd<29>;


	ld.param.u64 	%rd3, [adddmi_param_0];
	ld.param.u64 	%rd4, [adddmi_param_1];
	ld.param.u64 	%rd5, [adddmi_param_2];
	ld.param.u64 	%rd7, [adddmi_param_3];
	ld.param.u64 	%rd8, [adddmi_param_4];
	ld.param.u64 	%rd6, [adddmi_param_5];
	ld.param.f32 	%f1, [adddmi_param_7];
	ld.param.f32 	%f2, [adddmi_param_8];
	ld.param.u32 	%r22, [adddmi_param_9];
	ld.param.u32 	%r23, [adddmi_param_10];
	ld.param.u32 	%r24, [adddmi_param_11];
	cvta.to.global.u64 	%rd1, %rd7;
	cvta.to.global.u64 	%rd2, %rd8;
	.loc 3 13 1
	mov.u32 	%r1, %ntid.x;
	mov.u32 	%r2, %ctaid.x;
	mov.u32 	%r3, %tid.x;
	mad.lo.s32 	%r4, %r1, %r2, %r3;
	.loc 3 14 1
	mov.u32 	%r5, %ntid.y;
	mov.u32 	%r6, %ctaid.y;
	mov.u32 	%r7, %tid.y;
	mad.lo.s32 	%r8, %r5, %r6, %r7;
	.loc 3 16 1
	setp.lt.s32 	%p1, %r8, %r24;
	setp.lt.s32 	%p2, %r4, %r23;
	and.pred  	%p3, %p2, %p1;
	.loc 3 20 1
	setp.gt.s32 	%p4, %r22, 0;
	.loc 3 16 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB2_3;
	bra.uni 	BB2_1;

BB2_1:
	.loc 3 25 1
	add.s32 	%r9, %r22, -1;
	add.s32 	%r26, %r23, -1;
	mov.u32 	%r25, 0;
	.loc 4 238 5
	max.s32 	%r27, %r4, %r25;
	.loc 4 210 5
	min.s32 	%r10, %r27, %r26;
	.loc 3 25 1
	add.s32 	%r28, %r8, 1;
	.loc 4 238 5
	max.s32 	%r29, %r28, %r25;
	.loc 3 25 1
	add.s32 	%r30, %r24, -1;
	.loc 4 210 5
	min.s32 	%r11, %r29, %r30;
	.loc 3 25 1
	add.s32 	%r31, %r8, -1;
	.loc 4 238 5
	max.s32 	%r32, %r31, %r25;
	.loc 4 210 5
	min.s32 	%r12, %r32, %r30;
	.loc 3 26 1
	add.s32 	%r33, %r4, 1;
	.loc 4 238 5
	max.s32 	%r34, %r33, %r25;
	.loc 4 210 5
	min.s32 	%r13, %r34, %r26;
	.loc 4 238 5
	max.s32 	%r35, %r8, %r25;
	.loc 4 210 5
	min.s32 	%r14, %r35, %r30;
	.loc 3 26 1
	add.s32 	%r36, %r4, -1;
	.loc 4 238 5
	max.s32 	%r37, %r36, %r25;
	.loc 4 210 5
	min.s32 	%r15, %r37, %r26;
	.loc 3 20 1
	mad.lo.s32 	%r56, %r24, %r4, %r8;
	mul.lo.s32 	%r17, %r24, %r23;
	cvta.to.global.u64 	%rd9, %rd3;
	cvta.to.global.u64 	%rd12, %rd4;
	cvta.to.global.u64 	%rd14, %rd5;
	cvta.to.global.u64 	%rd16, %rd6;
	mov.u32 	%r57, %r25;

BB2_2:
	.loc 3 23 1
	mov.u32 	%r19, %r57;
	mul.wide.s32 	%rd10, %r56, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.f32 	%f3, [%rd11];
	add.s64 	%rd13, %rd12, %rd10;
	ld.global.f32 	%f4, [%rd13];
	add.s64 	%rd15, %rd14, %rd10;
	ld.global.f32 	%f5, [%rd15];
	.loc 4 238 5
	max.s32 	%r44, %r19, %r25;
	.loc 4 210 5
	min.s32 	%r45, %r44, %r9;
	.loc 3 25 1
	mad.lo.s32 	%r46, %r45, %r23, %r10;
	mad.lo.s32 	%r47, %r46, %r24, %r11;
	mul.wide.s32 	%rd17, %r47, 4;
	add.s64 	%rd18, %rd16, %rd17;
	ld.global.nc.f32 	%f6, [%rd18];
	mad.lo.s32 	%r48, %r46, %r24, %r12;
	mul.wide.s32 	%rd19, %r48, 4;
	add.s64 	%rd20, %rd16, %rd19;
	ld.global.nc.f32 	%f7, [%rd20];
	sub.f32 	%f8, %f6, %f7;
	fma.rn.f32 	%f9, %f8, %f2, %f3;
	.loc 3 26 1
	mad.lo.s32 	%r49, %r45, %r23, %r13;
	mad.lo.s32 	%r50, %r49, %r24, %r14;
	mul.wide.s32 	%rd21, %r50, 4;
	add.s64 	%rd22, %rd2, %rd21;
	ld.global.nc.f32 	%f10, [%rd22];
	mad.lo.s32 	%r51, %r45, %r23, %r15;
	mad.lo.s32 	%r52, %r51, %r24, %r14;
	mul.wide.s32 	%rd23, %r52, 4;
	add.s64 	%rd24, %rd2, %rd23;
	ld.global.nc.f32 	%f11, [%rd24];
	sub.f32 	%f12, %f10, %f11;
	fma.rn.f32 	%f13, %f12, %f1, %f9;
	.loc 3 28 1
	add.s64 	%rd25, %rd1, %rd21;
	ld.global.nc.f32 	%f14, [%rd25];
	add.s64 	%rd26, %rd1, %rd23;
	ld.global.nc.f32 	%f15, [%rd26];
	sub.f32 	%f16, %f14, %f15;
	mul.f32 	%f17, %f16, %f1;
	sub.f32 	%f18, %f4, %f17;
	.loc 3 29 1
	add.s64 	%rd27, %rd1, %rd17;
	ld.global.nc.f32 	%f19, [%rd27];
	add.s64 	%rd28, %rd1, %rd19;
	ld.global.nc.f32 	%f20, [%rd28];
	sub.f32 	%f21, %f19, %f20;
	mul.f32 	%f22, %f21, %f2;
	sub.f32 	%f23, %f5, %f22;
	.loc 3 32 1
	st.global.f32 	[%rd11], %f13;
	.loc 3 33 1
	st.global.f32 	[%rd13], %f18;
	.loc 3 34 1
	st.global.f32 	[%rd15], %f23;
	.loc 3 20 1
	add.s32 	%r56, %r56, %r17;
	.loc 3 20 18
	add.s32 	%r21, %r19, 1;
	.loc 3 20 1
	setp.lt.s32 	%p6, %r21, %r22;
	mov.u32 	%r57, %r21;
	@%p6 bra 	BB2_2;

BB2_3:
	.loc 3 36 2
	ret;
}


`
)
