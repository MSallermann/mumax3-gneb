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

// CUDA handle for myzero kernel
var myzero_code cu.Function

// Stores the arguments for myzero kernel invocation
type myzero_args_t struct {
	arg_ax unsafe.Pointer
	arg_ay unsafe.Pointer
	arg_az unsafe.Pointer
	arg_Nx int
	arg_Ny int
	arg_Nz int
	argptr [6]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for myzero kernel invocation
var myzero_args myzero_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	myzero_args.argptr[0] = unsafe.Pointer(&myzero_args.arg_ax)
	myzero_args.argptr[1] = unsafe.Pointer(&myzero_args.arg_ay)
	myzero_args.argptr[2] = unsafe.Pointer(&myzero_args.arg_az)
	myzero_args.argptr[3] = unsafe.Pointer(&myzero_args.arg_Nx)
	myzero_args.argptr[4] = unsafe.Pointer(&myzero_args.arg_Ny)
	myzero_args.argptr[5] = unsafe.Pointer(&myzero_args.arg_Nz)
}

// Wrapper for myzero CUDA kernel, asynchronous.
func k_myzero_async(ax unsafe.Pointer, ay unsafe.Pointer, az unsafe.Pointer, Nx int, Ny int, Nz int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("myzero")
	}

	myzero_args.Lock()
	defer myzero_args.Unlock()

	if myzero_code == 0 {
		myzero_code = fatbinLoad(myzero_map, "myzero")
	}

	myzero_args.arg_ax = ax
	myzero_args.arg_ay = ay
	myzero_args.arg_az = az
	myzero_args.arg_Nx = Nx
	myzero_args.arg_Ny = Ny
	myzero_args.arg_Nz = Nz

	args := myzero_args.argptr[:]
	cu.LaunchKernel(myzero_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("myzero")
	}
}

// maps compute capability on PTX code for myzero kernel.
var myzero_map = map[int]string{0: "",
	50: myzero_ptx_50}

// myzero PTX code for various compute capabilities.
const (
	myzero_ptx_50 = `
.version 7.5
.target sm_50
.address_size 64

	// .globl	myzero

.visible .entry myzero(
	.param .u64 myzero_param_0,
	.param .u64 myzero_param_1,
	.param .u64 myzero_param_2,
	.param .u32 myzero_param_3,
	.param .u32 myzero_param_4,
	.param .u32 myzero_param_5
)
{
	.reg .pred 	%p<3>;
	.reg .b32 	%r<20>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [myzero_param_0];
	ld.param.u64 	%rd2, [myzero_param_1];
	ld.param.u64 	%rd3, [myzero_param_2];
	ld.param.u32 	%r3, [myzero_param_3];
	ld.param.u32 	%r4, [myzero_param_4];
	ld.param.u32 	%r5, [myzero_param_5];
	mov.u32 	%r6, %nctaid.x;
	mov.u32 	%r7, %ctaid.y;
	mov.u32 	%r8, %ctaid.x;
	mad.lo.s32 	%r9, %r7, %r6, %r8;
	mov.u32 	%r10, %ntid.x;
	mov.u32 	%r11, %tid.x;
	mad.lo.s32 	%r1, %r9, %r10, %r11;
	mul.lo.s32 	%r2, %r4, %r3;
	mul.lo.s32 	%r12, %r2, %r5;
	setp.ge.s32 	%p1, %r1, %r12;
	@%p1 bra 	$L__BB0_3;

	div.s32 	%r13, %r1, %r2;
	mul.hi.s32 	%r14, %r13, 1431655766;
	shr.u32 	%r15, %r14, 31;
	add.s32 	%r16, %r14, %r15;
	mul.lo.s32 	%r17, %r16, 3;
	sub.s32 	%r18, %r13, %r17;
	setp.eq.s32 	%p2, %r18, 0;
	@%p2 bra 	$L__BB0_3;

	cvta.to.global.u64 	%rd4, %rd1;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	mov.u32 	%r19, 0;
	st.global.u32 	[%rd6], %r19;
	cvta.to.global.u64 	%rd7, %rd2;
	add.s64 	%rd8, %rd7, %rd5;
	st.global.u32 	[%rd8], %r19;
	cvta.to.global.u64 	%rd9, %rd3;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.u32 	[%rd10], %r19;

$L__BB0_3:
	ret;

}

`
)