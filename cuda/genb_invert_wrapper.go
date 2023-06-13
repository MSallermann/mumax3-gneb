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

// CUDA handle for invert kernel
var invert_code cu.Function

// Stores the arguments for invert kernel invocation
type invert_args_t struct {
	arg_dst unsafe.Pointer
	arg_src unsafe.Pointer
	arg_N   int
	argptr  [3]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for invert kernel invocation
var invert_args invert_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	invert_args.argptr[0] = unsafe.Pointer(&invert_args.arg_dst)
	invert_args.argptr[1] = unsafe.Pointer(&invert_args.arg_src)
	invert_args.argptr[2] = unsafe.Pointer(&invert_args.arg_N)
}

// Wrapper for invert CUDA kernel, asynchronous.
func k_invert_async(dst unsafe.Pointer, src unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("invert")
	}

	invert_args.Lock()
	defer invert_args.Unlock()

	if invert_code == 0 {
		invert_code = fatbinLoad(invert_map, "invert")
	}

	invert_args.arg_dst = dst
	invert_args.arg_src = src
	invert_args.arg_N = N

	args := invert_args.argptr[:]
	cu.LaunchKernel(invert_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("invert")
	}
}

// maps compute capability on PTX code for invert kernel.
var invert_map = map[int]string{0: "",
	50: invert_ptx_50}

// invert PTX code for various compute capabilities.
const (
	invert_ptx_50 = `
.version 7.5
.target sm_50
.address_size 64

	// .globl	invert

.visible .entry invert(
	.param .u64 invert_param_0,
	.param .u64 invert_param_1,
	.param .u32 invert_param_2
)
{
	.reg .pred 	%p<3>;
	.reg .f32 	%f<3>;
	.reg .b32 	%r<10>;
	.reg .b64 	%rd<8>;


	ld.param.u64 	%rd2, [invert_param_0];
	ld.param.u64 	%rd3, [invert_param_1];
	ld.param.u32 	%r2, [invert_param_2];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	$L__BB0_4;

	cvta.to.global.u64 	%rd4, %rd3;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f1, [%rd6];
	setp.eq.f32 	%p2, %f1, 0f00000000;
	cvta.to.global.u64 	%rd7, %rd2;
	add.s64 	%rd1, %rd7, %rd5;
	@%p2 bra 	$L__BB0_3;
	bra.uni 	$L__BB0_2;

$L__BB0_3:
	mov.u32 	%r9, 0;
	st.global.u32 	[%rd1], %r9;
	bra.uni 	$L__BB0_4;

$L__BB0_2:
	rcp.rn.f32 	%f2, %f1;
	st.global.f32 	[%rd1], %f2;

$L__BB0_4:
	ret;

}

`
)