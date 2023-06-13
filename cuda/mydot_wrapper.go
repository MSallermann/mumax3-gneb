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

// CUDA handle for mydot kernel
var mydot_code cu.Function

// Stores the arguments for mydot kernel invocation
type mydot_args_t struct {
	arg_sm unsafe.Pointer
	arg_ax unsafe.Pointer
	arg_ay unsafe.Pointer
	arg_az unsafe.Pointer
	arg_bx unsafe.Pointer
	arg_by unsafe.Pointer
	arg_bz unsafe.Pointer
	arg_N  int
	arg_Nz int
	argptr [9]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for mydot kernel invocation
var mydot_args mydot_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	mydot_args.argptr[0] = unsafe.Pointer(&mydot_args.arg_sm)
	mydot_args.argptr[1] = unsafe.Pointer(&mydot_args.arg_ax)
	mydot_args.argptr[2] = unsafe.Pointer(&mydot_args.arg_ay)
	mydot_args.argptr[3] = unsafe.Pointer(&mydot_args.arg_az)
	mydot_args.argptr[4] = unsafe.Pointer(&mydot_args.arg_bx)
	mydot_args.argptr[5] = unsafe.Pointer(&mydot_args.arg_by)
	mydot_args.argptr[6] = unsafe.Pointer(&mydot_args.arg_bz)
	mydot_args.argptr[7] = unsafe.Pointer(&mydot_args.arg_N)
	mydot_args.argptr[8] = unsafe.Pointer(&mydot_args.arg_Nz)
}

// Wrapper for mydot CUDA kernel, asynchronous.
func k_mydot_async(sm unsafe.Pointer, ax unsafe.Pointer, ay unsafe.Pointer, az unsafe.Pointer, bx unsafe.Pointer, by unsafe.Pointer, bz unsafe.Pointer, N int, Nz int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("mydot")
	}

	mydot_args.Lock()
	defer mydot_args.Unlock()

	if mydot_code == 0 {
		mydot_code = fatbinLoad(mydot_map, "mydot")
	}

	mydot_args.arg_sm = sm
	mydot_args.arg_ax = ax
	mydot_args.arg_ay = ay
	mydot_args.arg_az = az
	mydot_args.arg_bx = bx
	mydot_args.arg_by = by
	mydot_args.arg_bz = bz
	mydot_args.arg_N = N
	mydot_args.arg_Nz = Nz

	args := mydot_args.argptr[:]
	cu.LaunchKernel(mydot_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("mydot")
	}
}

// maps compute capability on PTX code for mydot kernel.
var mydot_map = map[int]string{0: "",
	50: mydot_ptx_50}

// mydot PTX code for various compute capabilities.
const (
	mydot_ptx_50 = `
.version 7.5
.target sm_50
.address_size 64

	// .globl	mydot

.visible .entry mydot(
	.param .u64 mydot_param_0,
	.param .u64 mydot_param_1,
	.param .u64 mydot_param_2,
	.param .u64 mydot_param_3,
	.param .u64 mydot_param_4,
	.param .u64 mydot_param_5,
	.param .u64 mydot_param_6,
	.param .u32 mydot_param_7,
	.param .u32 mydot_param_8
)
{
	.reg .pred 	%p<7>;
	.reg .f32 	%f<65>;
	.reg .b32 	%r<26>;
	.reg .b64 	%rd<58>;


	ld.param.u64 	%rd37, [mydot_param_0];
	ld.param.u64 	%rd38, [mydot_param_1];
	ld.param.u64 	%rd39, [mydot_param_2];
	ld.param.u64 	%rd40, [mydot_param_3];
	ld.param.u64 	%rd41, [mydot_param_4];
	ld.param.u64 	%rd42, [mydot_param_5];
	ld.param.u64 	%rd43, [mydot_param_6];
	ld.param.u32 	%r10, [mydot_param_7];
	cvta.to.global.u64 	%rd1, %rd43;
	cvta.to.global.u64 	%rd2, %rd40;
	cvta.to.global.u64 	%rd3, %rd42;
	cvta.to.global.u64 	%rd4, %rd39;
	cvta.to.global.u64 	%rd5, %rd41;
	cvta.to.global.u64 	%rd6, %rd38;
	mov.u32 	%r11, %nctaid.x;
	mov.u32 	%r12, %ctaid.y;
	mov.u32 	%r13, %ctaid.x;
	mad.lo.s32 	%r14, %r12, %r11, %r13;
	mov.u32 	%r15, %ntid.x;
	mul.lo.s32 	%r16, %r14, %r15;
	mov.u32 	%r17, %tid.x;
	neg.s32 	%r18, %r17;
	setp.ne.s32 	%p1, %r16, %r18;
	@%p1 bra 	$L__BB0_9;

	setp.lt.s32 	%p2, %r10, 1;
	mov.f32 	%f64, 0f00000000;
	@%p2 bra 	$L__BB0_8;

	add.s32 	%r20, %r10, -1;
	and.b32  	%r25, %r10, 3;
	setp.lt.u32 	%p3, %r20, 3;
	mov.f32 	%f64, 0f00000000;
	mov.u32 	%r24, 0;
	@%p3 bra 	$L__BB0_5;

	sub.s32 	%r23, %r10, %r25;
	mov.u64 	%rd46, %rd2;
	mov.u64 	%rd47, %rd3;
	mov.u64 	%rd48, %rd4;
	mov.u64 	%rd49, %rd1;
	mov.u64 	%rd50, %rd5;
	mov.u64 	%rd51, %rd6;

$L__BB0_4:
	ld.global.nc.f32 	%f12, [%rd50];
	ld.global.nc.f32 	%f13, [%rd51];
	ld.global.nc.f32 	%f14, [%rd47];
	ld.global.nc.f32 	%f15, [%rd48];
	mul.f32 	%f16, %f15, %f14;
	fma.rn.f32 	%f17, %f13, %f12, %f16;
	ld.global.nc.f32 	%f18, [%rd49];
	ld.global.nc.f32 	%f19, [%rd46];
	fma.rn.f32 	%f20, %f19, %f18, %f17;
	add.f32 	%f21, %f64, %f20;
	ld.global.nc.f32 	%f22, [%rd50+4];
	ld.global.nc.f32 	%f23, [%rd51+4];
	ld.global.nc.f32 	%f24, [%rd47+4];
	ld.global.nc.f32 	%f25, [%rd48+4];
	mul.f32 	%f26, %f25, %f24;
	fma.rn.f32 	%f27, %f23, %f22, %f26;
	ld.global.nc.f32 	%f28, [%rd49+4];
	ld.global.nc.f32 	%f29, [%rd46+4];
	fma.rn.f32 	%f30, %f29, %f28, %f27;
	add.f32 	%f31, %f21, %f30;
	ld.global.nc.f32 	%f32, [%rd50+8];
	ld.global.nc.f32 	%f33, [%rd51+8];
	ld.global.nc.f32 	%f34, [%rd47+8];
	ld.global.nc.f32 	%f35, [%rd48+8];
	mul.f32 	%f36, %f35, %f34;
	fma.rn.f32 	%f37, %f33, %f32, %f36;
	ld.global.nc.f32 	%f38, [%rd49+8];
	ld.global.nc.f32 	%f39, [%rd46+8];
	fma.rn.f32 	%f40, %f39, %f38, %f37;
	add.f32 	%f41, %f31, %f40;
	ld.global.nc.f32 	%f42, [%rd50+12];
	ld.global.nc.f32 	%f43, [%rd51+12];
	ld.global.nc.f32 	%f44, [%rd47+12];
	ld.global.nc.f32 	%f45, [%rd48+12];
	mul.f32 	%f46, %f45, %f44;
	fma.rn.f32 	%f47, %f43, %f42, %f46;
	ld.global.nc.f32 	%f48, [%rd49+12];
	ld.global.nc.f32 	%f49, [%rd46+12];
	fma.rn.f32 	%f50, %f49, %f48, %f47;
	add.f32 	%f64, %f41, %f50;
	add.s32 	%r24, %r24, 4;
	add.s64 	%rd51, %rd51, 16;
	add.s64 	%rd50, %rd50, 16;
	add.s64 	%rd49, %rd49, 16;
	add.s64 	%rd48, %rd48, 16;
	add.s64 	%rd47, %rd47, 16;
	add.s64 	%rd46, %rd46, 16;
	add.s32 	%r23, %r23, -4;
	setp.ne.s32 	%p4, %r23, 0;
	@%p4 bra 	$L__BB0_4;

$L__BB0_5:
	setp.eq.s32 	%p5, %r25, 0;
	@%p5 bra 	$L__BB0_8;

	mul.wide.s32 	%rd44, %r24, 4;
	add.s64 	%rd57, %rd1, %rd44;
	add.s64 	%rd56, %rd2, %rd44;
	add.s64 	%rd55, %rd3, %rd44;
	add.s64 	%rd54, %rd4, %rd44;
	add.s64 	%rd53, %rd5, %rd44;
	add.s64 	%rd52, %rd6, %rd44;

$L__BB0_7:
	.pragma "nounroll";
	ld.global.nc.f32 	%f51, [%rd53];
	ld.global.nc.f32 	%f52, [%rd52];
	ld.global.nc.f32 	%f53, [%rd55];
	ld.global.nc.f32 	%f54, [%rd54];
	mul.f32 	%f55, %f54, %f53;
	fma.rn.f32 	%f56, %f52, %f51, %f55;
	ld.global.nc.f32 	%f57, [%rd57];
	ld.global.nc.f32 	%f58, [%rd56];
	fma.rn.f32 	%f59, %f58, %f57, %f56;
	add.f32 	%f64, %f64, %f59;
	add.s64 	%rd57, %rd57, 4;
	add.s64 	%rd56, %rd56, 4;
	add.s64 	%rd55, %rd55, 4;
	add.s64 	%rd54, %rd54, 4;
	add.s64 	%rd53, %rd53, 4;
	add.s64 	%rd52, %rd52, 4;
	add.s32 	%r25, %r25, -1;
	setp.ne.s32 	%p6, %r25, 0;
	@%p6 bra 	$L__BB0_7;

$L__BB0_8:
	cvta.to.global.u64 	%rd45, %rd37;
	st.global.f32 	[%rd45], %f64;

$L__BB0_9:
	ret;

}

`
)