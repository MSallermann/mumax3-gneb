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

// CUDA handle for dotproduct2 kernel
var dotproduct2_code cu.Function

// Stores the arguments for dotproduct2 kernel invocation
type dotproduct2_args_t struct {
	arg_dst       unsafe.Pointer
	arg_prefactor float32
	arg_ax        unsafe.Pointer
	arg_ay        unsafe.Pointer
	arg_az        unsafe.Pointer
	arg_bx        unsafe.Pointer
	arg_by        unsafe.Pointer
	arg_bz        unsafe.Pointer
	arg_N         int
	argptr        [9]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for dotproduct2 kernel invocation
var dotproduct2_args dotproduct2_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	dotproduct2_args.argptr[0] = unsafe.Pointer(&dotproduct2_args.arg_dst)
	dotproduct2_args.argptr[1] = unsafe.Pointer(&dotproduct2_args.arg_prefactor)
	dotproduct2_args.argptr[2] = unsafe.Pointer(&dotproduct2_args.arg_ax)
	dotproduct2_args.argptr[3] = unsafe.Pointer(&dotproduct2_args.arg_ay)
	dotproduct2_args.argptr[4] = unsafe.Pointer(&dotproduct2_args.arg_az)
	dotproduct2_args.argptr[5] = unsafe.Pointer(&dotproduct2_args.arg_bx)
	dotproduct2_args.argptr[6] = unsafe.Pointer(&dotproduct2_args.arg_by)
	dotproduct2_args.argptr[7] = unsafe.Pointer(&dotproduct2_args.arg_bz)
	dotproduct2_args.argptr[8] = unsafe.Pointer(&dotproduct2_args.arg_N)
}

// Wrapper for dotproduct2 CUDA kernel, asynchronous.
func k_dotproduct2_async(dst unsafe.Pointer, prefactor float32, ax unsafe.Pointer, ay unsafe.Pointer, az unsafe.Pointer, bx unsafe.Pointer, by unsafe.Pointer, bz unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("dotproduct2")
	}

	dotproduct2_args.Lock()
	defer dotproduct2_args.Unlock()

	if dotproduct2_code == 0 {
		dotproduct2_code = fatbinLoad(dotproduct2_map, "dotproduct2")
	}

	dotproduct2_args.arg_dst = dst
	dotproduct2_args.arg_prefactor = prefactor
	dotproduct2_args.arg_ax = ax
	dotproduct2_args.arg_ay = ay
	dotproduct2_args.arg_az = az
	dotproduct2_args.arg_bx = bx
	dotproduct2_args.arg_by = by
	dotproduct2_args.arg_bz = bz
	dotproduct2_args.arg_N = N

	args := dotproduct2_args.argptr[:]
	cu.LaunchKernel(dotproduct2_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("dotproduct2")
	}
}

// maps compute capability on PTX code for dotproduct2 kernel.
var dotproduct2_map = map[int]string{0: "",
	50: dotproduct2_ptx_50}

// dotproduct2 PTX code for various compute capabilities.
const (
	dotproduct2_ptx_50 = `
.version 7.5
.target sm_50
.address_size 64

	// .globl	dotproduct2

.visible .entry dotproduct2(
	.param .u64 dotproduct2_param_0,
	.param .f32 dotproduct2_param_1,
	.param .u64 dotproduct2_param_2,
	.param .u64 dotproduct2_param_3,
	.param .u64 dotproduct2_param_4,
	.param .u64 dotproduct2_param_5,
	.param .u64 dotproduct2_param_6,
	.param .u64 dotproduct2_param_7,
	.param .u32 dotproduct2_param_8
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<12>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<23>;


	ld.param.u64 	%rd1, [dotproduct2_param_0];
	ld.param.f32 	%f1, [dotproduct2_param_1];
	ld.param.u64 	%rd2, [dotproduct2_param_2];
	ld.param.u64 	%rd3, [dotproduct2_param_3];
	ld.param.u64 	%rd4, [dotproduct2_param_4];
	ld.param.u64 	%rd5, [dotproduct2_param_5];
	ld.param.u64 	%rd6, [dotproduct2_param_6];
	ld.param.u64 	%rd7, [dotproduct2_param_7];
	ld.param.u32 	%r2, [dotproduct2_param_8];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	$L__BB0_2;

	cvta.to.global.u64 	%rd8, %rd2;
	mul.wide.s32 	%rd9, %r1, 4;
	add.s64 	%rd10, %rd8, %rd9;
	cvta.to.global.u64 	%rd11, %rd3;
	add.s64 	%rd12, %rd11, %rd9;
	cvta.to.global.u64 	%rd13, %rd4;
	add.s64 	%rd14, %rd13, %rd9;
	cvta.to.global.u64 	%rd15, %rd5;
	add.s64 	%rd16, %rd15, %rd9;
	cvta.to.global.u64 	%rd17, %rd6;
	add.s64 	%rd18, %rd17, %rd9;
	cvta.to.global.u64 	%rd19, %rd7;
	add.s64 	%rd20, %rd19, %rd9;
	ld.global.nc.f32 	%f2, [%rd16];
	ld.global.nc.f32 	%f3, [%rd10];
	ld.global.nc.f32 	%f4, [%rd18];
	ld.global.nc.f32 	%f5, [%rd12];
	mul.f32 	%f6, %f5, %f4;
	fma.rn.f32 	%f7, %f3, %f2, %f6;
	ld.global.nc.f32 	%f8, [%rd20];
	ld.global.nc.f32 	%f9, [%rd14];
	fma.rn.f32 	%f10, %f9, %f8, %f7;
	mul.f32 	%f11, %f10, %f1;
	cvta.to.global.u64 	%rd21, %rd1;
	add.s64 	%rd22, %rd21, %rd9;
	st.global.f32 	[%rd22], %f11;

$L__BB0_2:
	ret;

}

`
)
