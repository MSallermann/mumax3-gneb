package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"sync"
	"unsafe"

	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
)

// CUDA handle for totalforce kernel
var totalforce_code cu.Function

// Stores the arguments for totalforce kernel invocation
type totalforce_args_t struct {
	arg_src     unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	arg_noi     int
	arg_Nz      int
	argptr      [6]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for totalforce kernel invocation
var totalforce_args totalforce_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	totalforce_args.argptr[0] = unsafe.Pointer(&totalforce_args.arg_src)
	totalforce_args.argptr[1] = unsafe.Pointer(&totalforce_args.arg_dst)
	totalforce_args.argptr[2] = unsafe.Pointer(&totalforce_args.arg_initVal)
	totalforce_args.argptr[3] = unsafe.Pointer(&totalforce_args.arg_n)
	totalforce_args.argptr[4] = unsafe.Pointer(&totalforce_args.arg_noi)
	totalforce_args.argptr[5] = unsafe.Pointer(&totalforce_args.arg_Nz)
}

// Wrapper for totalforce CUDA kernel, asynchronous.
func k_totalforce_async(src unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, noi int, Nz int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("totalforce")
	}

	totalforce_args.Lock()
	defer totalforce_args.Unlock()

	if totalforce_code == 0 {
		totalforce_code = fatbinLoad(totalforce_map, "totalforce")
	}

	totalforce_args.arg_src = src
	totalforce_args.arg_dst = dst
	totalforce_args.arg_initVal = initVal
	totalforce_args.arg_n = n
	totalforce_args.arg_noi = noi
	totalforce_args.arg_Nz = Nz

	args := totalforce_args.argptr[:]
	cu.LaunchKernel(totalforce_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("totalforce")
	}
}

// maps compute capability on PTX code for totalforce kernel.
var totalforce_map = map[int]string{0: "",
	50: totalforce_ptx_50}

// totalforce PTX code for various compute capabilities.
const (
	totalforce_ptx_50 = `
.version 7.5
.target sm_50
.address_size 64

	// .globl	totalforce

.visible .entry totalforce(
	.param .u64 totalforce_param_0,
	.param .u64 totalforce_param_1,
	.param .f32 totalforce_param_2,
	.param .u32 totalforce_param_3,
	.param .u32 totalforce_param_4,
	.param .u32 totalforce_param_5
)
{
	.reg .pred 	%p<11>;
	.reg .f32 	%f<48>;
	.reg .b32 	%r<39>;
	.reg .b64 	%rd<17>;
	// demoted variable
	.shared .align 4 .b8 _ZZ10totalforceE5sdata[2048];

	ld.param.u64 	%rd8, [totalforce_param_0];
	ld.param.u64 	%rd7, [totalforce_param_1];
	ld.param.f32 	%f47, [totalforce_param_2];
	ld.param.u32 	%r17, [totalforce_param_3];
	ld.param.u32 	%r18, [totalforce_param_4];
	cvta.to.global.u64 	%rd1, %rd8;
	mov.u32 	%r38, %ntid.x;
	mov.u32 	%r19, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r36, %r19, %r38, %r2;
	mov.u32 	%r20, %nctaid.x;
	mul.lo.s32 	%r4, %r20, %r38;
	setp.ge.s32 	%p1, %r36, %r17;
	@%p1 bra 	$L__BB0_7;

	add.s32 	%r21, %r4, %r17;
	add.s32 	%r22, %r36, %r4;
	not.b32 	%r23, %r22;
	add.s32 	%r24, %r21, %r23;
	div.u32 	%r5, %r24, %r4;
	add.s32 	%r25, %r5, 1;
	and.b32  	%r35, %r25, 3;
	setp.eq.s32 	%p2, %r35, 0;
	@%p2 bra 	$L__BB0_4;

	mul.wide.s32 	%rd9, %r36, 4;
	add.s64 	%rd16, %rd1, %rd9;
	mul.wide.s32 	%rd3, %r4, 4;

$L__BB0_3:
	.pragma "nounroll";
	ld.global.nc.f32 	%f10, [%rd16];
	add.f32 	%f47, %f47, %f10;
	add.s32 	%r36, %r36, %r4;
	add.s64 	%rd16, %rd16, %rd3;
	add.s32 	%r35, %r35, -1;
	setp.ne.s32 	%p3, %r35, 0;
	@%p3 bra 	$L__BB0_3;

$L__BB0_4:
	setp.lt.u32 	%p4, %r5, 3;
	@%p4 bra 	$L__BB0_7;

	mul.wide.s32 	%rd6, %r4, 4;

$L__BB0_6:
	mul.wide.s32 	%rd10, %r36, 4;
	add.s64 	%rd11, %rd1, %rd10;
	ld.global.nc.f32 	%f11, [%rd11];
	add.f32 	%f12, %f47, %f11;
	add.s64 	%rd12, %rd11, %rd6;
	ld.global.nc.f32 	%f13, [%rd12];
	add.f32 	%f14, %f12, %f13;
	add.s32 	%r26, %r36, %r4;
	add.s32 	%r27, %r26, %r4;
	add.s64 	%rd13, %rd12, %rd6;
	ld.global.nc.f32 	%f15, [%rd13];
	add.f32 	%f16, %f14, %f15;
	add.s32 	%r28, %r27, %r4;
	add.s64 	%rd14, %rd13, %rd6;
	ld.global.nc.f32 	%f17, [%rd14];
	add.f32 	%f47, %f16, %f17;
	add.s32 	%r36, %r28, %r4;
	setp.lt.s32 	%p5, %r36, %r17;
	@%p5 bra 	$L__BB0_6;

$L__BB0_7:
	shl.b32 	%r29, %r2, 2;
	mov.u32 	%r30, _ZZ10totalforceE5sdata;
	add.s32 	%r14, %r30, %r29;
	st.shared.f32 	[%r14], %f47;
	bar.sync 	0;
	setp.lt.u32 	%p6, %r38, 66;
	@%p6 bra 	$L__BB0_12;

$L__BB0_9:
	shr.u32 	%r16, %r38, 1;
	setp.ge.u32 	%p7, %r2, %r16;
	@%p7 bra 	$L__BB0_11;

	ld.shared.f32 	%f18, [%r14];
	shl.b32 	%r31, %r16, 2;
	add.s32 	%r32, %r14, %r31;
	ld.shared.f32 	%f19, [%r32];
	add.f32 	%f20, %f18, %f19;
	st.shared.f32 	[%r14], %f20;

$L__BB0_11:
	bar.sync 	0;
	setp.gt.u32 	%p8, %r38, 131;
	mov.u32 	%r38, %r16;
	@%p8 bra 	$L__BB0_9;

$L__BB0_12:
	setp.gt.s32 	%p9, %r2, 31;
	@%p9 bra 	$L__BB0_14;

	ld.volatile.shared.f32 	%f21, [%r14];
	ld.volatile.shared.f32 	%f22, [%r14+128];
	add.f32 	%f23, %f21, %f22;
	st.volatile.shared.f32 	[%r14], %f23;
	ld.volatile.shared.f32 	%f24, [%r14+64];
	ld.volatile.shared.f32 	%f25, [%r14];
	add.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r14], %f26;
	ld.volatile.shared.f32 	%f27, [%r14+32];
	ld.volatile.shared.f32 	%f28, [%r14];
	add.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r14], %f29;
	ld.volatile.shared.f32 	%f30, [%r14+16];
	ld.volatile.shared.f32 	%f31, [%r14];
	add.f32 	%f32, %f31, %f30;
	st.volatile.shared.f32 	[%r14], %f32;
	ld.volatile.shared.f32 	%f33, [%r14+8];
	ld.volatile.shared.f32 	%f34, [%r14];
	add.f32 	%f35, %f34, %f33;
	st.volatile.shared.f32 	[%r14], %f35;
	ld.volatile.shared.f32 	%f36, [%r14+4];
	ld.volatile.shared.f32 	%f37, [%r14];
	add.f32 	%f38, %f37, %f36;
	st.volatile.shared.f32 	[%r14], %f38;

$L__BB0_14:
	setp.ne.s32 	%p10, %r2, 0;
	@%p10 bra 	$L__BB0_16;

	ld.shared.f32 	%f39, [_ZZ10totalforceE5sdata];
	mul.lo.s32 	%r33, %r18, %r17;
	cvt.rn.f32.s32 	%f40, %r33;
	div.rn.f32 	%f41, %f39, %f40;
	cvta.to.global.u64 	%rd15, %rd7;
	atom.global.add.f32 	%f42, [%rd15], %f41;

$L__BB0_16:
	ret;

}

`
)