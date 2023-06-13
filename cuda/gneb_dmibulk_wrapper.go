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

// CUDA handle for gneb_adddmibulk kernel
var gneb_adddmibulk_code cu.Function

// Stores the arguments for gneb_adddmibulk kernel invocation
type gneb_adddmibulk_args_t struct {
	arg_Hx      unsafe.Pointer
	arg_Hy      unsafe.Pointer
	arg_Hz      unsafe.Pointer
	arg_mx      unsafe.Pointer
	arg_my      unsafe.Pointer
	arg_mz      unsafe.Pointer
	arg_Ms_     unsafe.Pointer
	arg_Ms_mul  float32
	arg_aLUT2d  unsafe.Pointer
	arg_DLUT2d  unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_cx      float32
	arg_cy      float32
	arg_cz      float32
	arg_Nx      int
	arg_Ny      int
	arg_Nz      int
	arg_noi     int
	arg_PBC     byte
	arg_OpenBC  byte
	arg_GNEB    byte
	argptr      [21]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for gneb_adddmibulk kernel invocation
var gneb_adddmibulk_args gneb_adddmibulk_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	gneb_adddmibulk_args.argptr[0] = unsafe.Pointer(&gneb_adddmibulk_args.arg_Hx)
	gneb_adddmibulk_args.argptr[1] = unsafe.Pointer(&gneb_adddmibulk_args.arg_Hy)
	gneb_adddmibulk_args.argptr[2] = unsafe.Pointer(&gneb_adddmibulk_args.arg_Hz)
	gneb_adddmibulk_args.argptr[3] = unsafe.Pointer(&gneb_adddmibulk_args.arg_mx)
	gneb_adddmibulk_args.argptr[4] = unsafe.Pointer(&gneb_adddmibulk_args.arg_my)
	gneb_adddmibulk_args.argptr[5] = unsafe.Pointer(&gneb_adddmibulk_args.arg_mz)
	gneb_adddmibulk_args.argptr[6] = unsafe.Pointer(&gneb_adddmibulk_args.arg_Ms_)
	gneb_adddmibulk_args.argptr[7] = unsafe.Pointer(&gneb_adddmibulk_args.arg_Ms_mul)
	gneb_adddmibulk_args.argptr[8] = unsafe.Pointer(&gneb_adddmibulk_args.arg_aLUT2d)
	gneb_adddmibulk_args.argptr[9] = unsafe.Pointer(&gneb_adddmibulk_args.arg_DLUT2d)
	gneb_adddmibulk_args.argptr[10] = unsafe.Pointer(&gneb_adddmibulk_args.arg_regions)
	gneb_adddmibulk_args.argptr[11] = unsafe.Pointer(&gneb_adddmibulk_args.arg_cx)
	gneb_adddmibulk_args.argptr[12] = unsafe.Pointer(&gneb_adddmibulk_args.arg_cy)
	gneb_adddmibulk_args.argptr[13] = unsafe.Pointer(&gneb_adddmibulk_args.arg_cz)
	gneb_adddmibulk_args.argptr[14] = unsafe.Pointer(&gneb_adddmibulk_args.arg_Nx)
	gneb_adddmibulk_args.argptr[15] = unsafe.Pointer(&gneb_adddmibulk_args.arg_Ny)
	gneb_adddmibulk_args.argptr[16] = unsafe.Pointer(&gneb_adddmibulk_args.arg_Nz)
	gneb_adddmibulk_args.argptr[17] = unsafe.Pointer(&gneb_adddmibulk_args.arg_noi)
	gneb_adddmibulk_args.argptr[18] = unsafe.Pointer(&gneb_adddmibulk_args.arg_PBC)
	gneb_adddmibulk_args.argptr[19] = unsafe.Pointer(&gneb_adddmibulk_args.arg_OpenBC)
	gneb_adddmibulk_args.argptr[20] = unsafe.Pointer(&gneb_adddmibulk_args.arg_GNEB)
}

// Wrapper for gneb_adddmibulk CUDA kernel, asynchronous.
func k_gneb_adddmibulk_async(Hx unsafe.Pointer, Hy unsafe.Pointer, Hz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Ms_ unsafe.Pointer, Ms_mul float32, aLUT2d unsafe.Pointer, DLUT2d unsafe.Pointer, regions unsafe.Pointer, cx float32, cy float32, cz float32, Nx int, Ny int, Nz int, noi int, PBC byte, OpenBC byte, GNEB byte, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("gneb_adddmibulk")
	}

	gneb_adddmibulk_args.Lock()
	defer gneb_adddmibulk_args.Unlock()

	if gneb_adddmibulk_code == 0 {
		gneb_adddmibulk_code = fatbinLoad(gneb_adddmibulk_map, "gneb_adddmibulk")
	}

	gneb_adddmibulk_args.arg_Hx = Hx
	gneb_adddmibulk_args.arg_Hy = Hy
	gneb_adddmibulk_args.arg_Hz = Hz
	gneb_adddmibulk_args.arg_mx = mx
	gneb_adddmibulk_args.arg_my = my
	gneb_adddmibulk_args.arg_mz = mz
	gneb_adddmibulk_args.arg_Ms_ = Ms_
	gneb_adddmibulk_args.arg_Ms_mul = Ms_mul
	gneb_adddmibulk_args.arg_aLUT2d = aLUT2d
	gneb_adddmibulk_args.arg_DLUT2d = DLUT2d
	gneb_adddmibulk_args.arg_regions = regions
	gneb_adddmibulk_args.arg_cx = cx
	gneb_adddmibulk_args.arg_cy = cy
	gneb_adddmibulk_args.arg_cz = cz
	gneb_adddmibulk_args.arg_Nx = Nx
	gneb_adddmibulk_args.arg_Ny = Ny
	gneb_adddmibulk_args.arg_Nz = Nz
	gneb_adddmibulk_args.arg_noi = noi
	gneb_adddmibulk_args.arg_PBC = PBC
	gneb_adddmibulk_args.arg_OpenBC = OpenBC
	gneb_adddmibulk_args.arg_GNEB = GNEB

	args := gneb_adddmibulk_args.argptr[:]
	cu.LaunchKernel(gneb_adddmibulk_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("gneb_adddmibulk")
	}
}

// maps compute capability on PTX code for gneb_adddmibulk kernel.
var gneb_adddmibulk_map = map[int]string{0: "",
	50: gneb_adddmibulk_ptx_50}

// gneb_adddmibulk PTX code for various compute capabilities.
const (
	gneb_adddmibulk_ptx_50 = `
.version 7.5
.target sm_50
.address_size 64

	// .globl	gneb_adddmibulk

.visible .entry gneb_adddmibulk(
	.param .u64 gneb_adddmibulk_param_0,
	.param .u64 gneb_adddmibulk_param_1,
	.param .u64 gneb_adddmibulk_param_2,
	.param .u64 gneb_adddmibulk_param_3,
	.param .u64 gneb_adddmibulk_param_4,
	.param .u64 gneb_adddmibulk_param_5,
	.param .u64 gneb_adddmibulk_param_6,
	.param .f32 gneb_adddmibulk_param_7,
	.param .u64 gneb_adddmibulk_param_8,
	.param .u64 gneb_adddmibulk_param_9,
	.param .u64 gneb_adddmibulk_param_10,
	.param .f32 gneb_adddmibulk_param_11,
	.param .f32 gneb_adddmibulk_param_12,
	.param .f32 gneb_adddmibulk_param_13,
	.param .u32 gneb_adddmibulk_param_14,
	.param .u32 gneb_adddmibulk_param_15,
	.param .u32 gneb_adddmibulk_param_16,
	.param .u32 gneb_adddmibulk_param_17,
	.param .u8 gneb_adddmibulk_param_18,
	.param .u8 gneb_adddmibulk_param_19,
	.param .u8 gneb_adddmibulk_param_20
)
{
	.reg .pred 	%p<94>;
	.reg .b16 	%rs<67>;
	.reg .f32 	%f<497>;
	.reg .b32 	%r<166>;
	.reg .b64 	%rd<105>;


	ld.param.u8 	%rs24, [gneb_adddmibulk_param_20];
	ld.param.u8 	%rs23, [gneb_adddmibulk_param_19];
	ld.param.u8 	%rs22, [gneb_adddmibulk_param_18];
	ld.param.u64 	%rd7, [gneb_adddmibulk_param_0];
	ld.param.u64 	%rd8, [gneb_adddmibulk_param_1];
	ld.param.u64 	%rd9, [gneb_adddmibulk_param_2];
	ld.param.u64 	%rd11, [gneb_adddmibulk_param_3];
	ld.param.u64 	%rd12, [gneb_adddmibulk_param_4];
	ld.param.u64 	%rd13, [gneb_adddmibulk_param_5];
	ld.param.u64 	%rd10, [gneb_adddmibulk_param_6];
	ld.param.f32 	%f495, [gneb_adddmibulk_param_7];
	ld.param.u64 	%rd14, [gneb_adddmibulk_param_8];
	ld.param.u64 	%rd15, [gneb_adddmibulk_param_9];
	ld.param.u64 	%rd16, [gneb_adddmibulk_param_10];
	ld.param.f32 	%f252, [gneb_adddmibulk_param_11];
	ld.param.f32 	%f253, [gneb_adddmibulk_param_12];
	ld.param.f32 	%f254, [gneb_adddmibulk_param_13];
	ld.param.u32 	%r55, [gneb_adddmibulk_param_14];
	ld.param.u32 	%r56, [gneb_adddmibulk_param_15];
	ld.param.u32 	%r57, [gneb_adddmibulk_param_16];
	ld.param.u32 	%r58, [gneb_adddmibulk_param_17];
	cvta.to.global.u64 	%rd1, %rd15;
	cvta.to.global.u64 	%rd2, %rd14;
	cvta.to.global.u64 	%rd3, %rd16;
	cvta.to.global.u64 	%rd4, %rd13;
	cvta.to.global.u64 	%rd5, %rd12;
	cvta.to.global.u64 	%rd6, %rd11;
	mov.u32 	%r59, %ntid.x;
	mov.u32 	%r60, %ctaid.x;
	mov.u32 	%r61, %tid.x;
	mad.lo.s32 	%r1, %r60, %r59, %r61;
	mov.u32 	%r62, %ntid.y;
	mov.u32 	%r63, %ctaid.y;
	mov.u32 	%r64, %tid.y;
	mad.lo.s32 	%r2, %r63, %r62, %r64;
	mov.u32 	%r65, %ntid.z;
	mov.u32 	%r66, %ctaid.z;
	mov.u32 	%r67, %tid.z;
	mad.lo.s32 	%r3, %r66, %r65, %r67;
	setp.ge.s32 	%p1, %r1, %r55;
	setp.ge.s32 	%p2, %r2, %r56;
	or.pred  	%p3, %p1, %p2;
	setp.ge.s32 	%p4, %r3, %r57;
	or.pred  	%p5, %p3, %p4;
	@%p5 bra 	$L__BB0_120;

	mul.lo.s32 	%r4, %r3, %r56;
	add.s32 	%r68, %r4, %r2;
	mul.lo.s32 	%r5, %r68, %r55;
	add.s32 	%r6, %r5, %r1;
	mul.wide.s32 	%rd17, %r6, 4;
	add.s64 	%rd18, %rd6, %rd17;
	cvt.s64.s32 	%rd19, %r6;
	add.s64 	%rd20, %rd5, %rd17;
	add.s64 	%rd21, %rd4, %rd17;
	add.s64 	%rd22, %rd3, %rd19;
	ld.global.nc.u8 	%rs1, [%rd22];
	ld.global.nc.f32 	%f1, [%rd18];
	ld.global.nc.f32 	%f5, [%rd20];
	mul.f32 	%f257, %f5, %f5;
	fma.rn.f32 	%f258, %f1, %f1, %f257;
	ld.global.nc.f32 	%f6, [%rd21];
	fma.rn.f32 	%f259, %f6, %f6, %f258;
	setp.eq.f32 	%p6, %f259, 0f00000000;
	@%p6 bra 	$L__BB0_120;

	and.b16  	%rs2, %rs22, 1;
	setp.eq.s16 	%p7, %rs2, 0;
	add.s32 	%r7, %r1, -1;
	@%p7 bra 	$L__BB0_4;
	bra.uni 	$L__BB0_3;

$L__BB0_4:
	max.s32 	%r156, %r7, 0;
	bra.uni 	$L__BB0_5;

$L__BB0_3:
	rem.s32 	%r69, %r7, %r55;
	add.s32 	%r70, %r69, %r55;
	rem.s32 	%r156, %r70, %r55;

$L__BB0_5:
	add.s32 	%r11, %r156, %r5;
	setp.lt.s32 	%p9, %r1, 1;
	mov.f32 	%f421, 0f00000000;
	and.pred  	%p10, %p9, %p7;
	mov.f32 	%f420, %f421;
	mov.f32 	%f419, %f421;
	@%p10 bra 	$L__BB0_7;

	mul.wide.s32 	%rd23, %r11, 4;
	add.s64 	%rd24, %rd6, %rd23;
	ld.global.nc.f32 	%f419, [%rd24];
	add.s64 	%rd25, %rd5, %rd23;
	ld.global.nc.f32 	%f420, [%rd25];
	add.s64 	%rd26, %rd4, %rd23;
	ld.global.nc.f32 	%f421, [%rd26];

$L__BB0_7:
	mul.f32 	%f263, %f419, %f419;
	fma.rn.f32 	%f264, %f420, %f420, %f263;
	fma.rn.f32 	%f13, %f421, %f421, %f264;
	setp.eq.f32 	%p11, %f13, 0f00000000;
	mov.u16 	%rs59, %rs1;
	@%p11 bra 	$L__BB0_9;

	cvt.s64.s32 	%rd27, %r11;
	add.s64 	%rd28, %rd3, %rd27;
	ld.global.nc.u8 	%rs59, [%rd28];

$L__BB0_9:
	min.u16 	%rs27, %rs59, %rs1;
	cvt.u32.u16 	%r71, %rs27;
	max.u16 	%rs28, %rs59, %rs1;
	cvt.u32.u16 	%r72, %rs28;
	add.s32 	%r73, %r72, 1;
	mul.lo.s32 	%r74, %r73, %r72;
	shr.u32 	%r75, %r74, 1;
	add.s32 	%r76, %r75, %r71;
	mul.wide.s32 	%rd29, %r76, 4;
	add.s64 	%rd30, %rd2, %rd29;
	add.s64 	%rd31, %rd1, %rd29;
	ld.global.nc.f32 	%f14, [%rd31];
	ld.global.nc.f32 	%f15, [%rd30];
	setp.leu.f32 	%p12, %f15, 0f00000000;
	mov.f32 	%f418, %f14;
	@%p12 bra 	$L__BB0_11;

	add.f32 	%f265, %f15, %f15;
	div.rn.f32 	%f418, %f14, %f265;

$L__BB0_11:
	setp.ne.s16 	%p13, %rs23, 0;
	mov.f32 	%f432, 0f00000000;
	and.pred  	%p15, %p13, %p11;
	mov.f32 	%f433, %f432;
	mov.f32 	%f434, %f432;
	@%p15 bra 	$L__BB0_15;

	setp.neu.f32 	%p16, %f13, 0f00000000;
	@%p16 bra 	$L__BB0_14;

	mul.f32 	%f269, %f418, %f252;
	fma.rn.f32 	%f420, %f6, %f269, %f5;
	mul.f32 	%f270, %f5, %f269;
	sub.f32 	%f421, %f6, %f270;
	mov.f32 	%f419, %f1;

$L__BB0_14:
	mul.f32 	%f271, %f252, %f252;
	add.f32 	%f272, %f15, %f15;
	div.rn.f32 	%f273, %f272, %f271;
	sub.f32 	%f274, %f419, %f1;
	sub.f32 	%f275, %f420, %f5;
	sub.f32 	%f276, %f421, %f6;
	fma.rn.f32 	%f432, %f273, %f274, 0f00000000;
	fma.rn.f32 	%f277, %f273, %f275, 0f00000000;
	fma.rn.f32 	%f278, %f273, %f276, 0f00000000;
	div.rn.f32 	%f279, %f14, %f252;
	mul.f32 	%f280, %f279, %f421;
	sub.f32 	%f433, %f277, %f280;
	fma.rn.f32 	%f434, %f279, %f420, %f278;

$L__BB0_15:
	add.s32 	%r12, %r1, 1;
	@%p7 bra 	$L__BB0_17;
	bra.uni 	$L__BB0_16;

$L__BB0_17:
	add.s32 	%r79, %r55, -1;
	min.s32 	%r157, %r12, %r79;
	bra.uni 	$L__BB0_18;

$L__BB0_16:
	rem.s32 	%r77, %r12, %r55;
	add.s32 	%r78, %r77, %r55;
	rem.s32 	%r157, %r78, %r55;

$L__BB0_18:
	add.s32 	%r16, %r157, %r5;
	setp.ge.s32 	%p18, %r12, %r55;
	mov.f32 	%f431, 0f00000000;
	and.pred  	%p20, %p18, %p7;
	mov.f32 	%f430, %f431;
	mov.f32 	%f429, %f431;
	@%p20 bra 	$L__BB0_20;

	mul.wide.s32 	%rd32, %r16, 4;
	add.s64 	%rd33, %rd6, %rd32;
	ld.global.nc.f32 	%f429, [%rd33];
	add.s64 	%rd34, %rd5, %rd32;
	ld.global.nc.f32 	%f430, [%rd34];
	add.s64 	%rd35, %rd4, %rd32;
	ld.global.nc.f32 	%f431, [%rd35];

$L__BB0_20:
	mul.f32 	%f284, %f429, %f429;
	fma.rn.f32 	%f285, %f430, %f430, %f284;
	fma.rn.f32 	%f43, %f431, %f431, %f285;
	setp.eq.f32 	%p21, %f43, 0f00000000;
	mov.u16 	%rs60, %rs1;
	@%p21 bra 	$L__BB0_22;

	cvt.s64.s32 	%rd36, %r16;
	add.s64 	%rd37, %rd3, %rd36;
	ld.global.nc.u8 	%rs60, [%rd37];

$L__BB0_22:
	min.u16 	%rs31, %rs60, %rs1;
	cvt.u32.u16 	%r80, %rs31;
	max.u16 	%rs32, %rs60, %rs1;
	cvt.u32.u16 	%r81, %rs32;
	add.s32 	%r82, %r81, 1;
	mul.lo.s32 	%r83, %r82, %r81;
	shr.u32 	%r84, %r83, 1;
	add.s32 	%r85, %r84, %r80;
	mul.wide.s32 	%rd38, %r85, 4;
	add.s64 	%rd39, %rd2, %rd38;
	add.s64 	%rd40, %rd1, %rd38;
	ld.global.nc.f32 	%f44, [%rd40];
	ld.global.nc.f32 	%f45, [%rd39];
	setp.leu.f32 	%p22, %f45, 0f00000000;
	mov.f32 	%f428, %f44;
	@%p22 bra 	$L__BB0_24;

	add.f32 	%f286, %f45, %f45;
	div.rn.f32 	%f428, %f44, %f286;

$L__BB0_24:
	and.pred  	%p25, %p13, %p21;
	@%p25 bra 	$L__BB0_28;

	setp.neu.f32 	%p26, %f43, 0f00000000;
	@%p26 bra 	$L__BB0_27;

	mul.f32 	%f287, %f428, %f252;
	mul.f32 	%f288, %f6, %f287;
	sub.f32 	%f430, %f5, %f288;
	fma.rn.f32 	%f431, %f5, %f287, %f6;
	mov.f32 	%f429, %f1;

$L__BB0_27:
	mul.f32 	%f289, %f252, %f252;
	add.f32 	%f290, %f45, %f45;
	div.rn.f32 	%f291, %f290, %f289;
	sub.f32 	%f292, %f429, %f1;
	sub.f32 	%f293, %f430, %f5;
	sub.f32 	%f294, %f431, %f6;
	fma.rn.f32 	%f432, %f291, %f292, %f432;
	fma.rn.f32 	%f295, %f291, %f293, %f433;
	fma.rn.f32 	%f296, %f291, %f294, %f434;
	div.rn.f32 	%f297, %f44, %f252;
	fma.rn.f32 	%f433, %f297, %f431, %f295;
	mul.f32 	%f298, %f297, %f430;
	sub.f32 	%f434, %f296, %f298;

$L__BB0_28:
	and.b16  	%rs7, %rs22, 2;
	setp.eq.s16 	%p27, %rs7, 0;
	add.s32 	%r17, %r2, -1;
	@%p27 bra 	$L__BB0_30;
	bra.uni 	$L__BB0_29;

$L__BB0_30:
	max.s32 	%r158, %r17, 0;
	bra.uni 	$L__BB0_31;

$L__BB0_29:
	rem.s32 	%r86, %r17, %r56;
	add.s32 	%r87, %r86, %r56;
	rem.s32 	%r158, %r87, %r56;

$L__BB0_31:
	add.s32 	%r88, %r158, %r4;
	mad.lo.s32 	%r21, %r88, %r55, %r1;
	setp.lt.s32 	%p29, %r2, 1;
	mov.f32 	%f441, 0f00000000;
	and.pred  	%p30, %p29, %p27;
	mov.f32 	%f440, %f441;
	mov.f32 	%f439, %f441;
	@%p30 bra 	$L__BB0_33;

	mul.wide.s32 	%rd41, %r21, 4;
	add.s64 	%rd42, %rd6, %rd41;
	ld.global.nc.f32 	%f439, [%rd42];
	add.s64 	%rd43, %rd5, %rd41;
	ld.global.nc.f32 	%f440, [%rd43];
	add.s64 	%rd44, %rd4, %rd41;
	ld.global.nc.f32 	%f441, [%rd44];

$L__BB0_33:
	mul.f32 	%f302, %f439, %f439;
	fma.rn.f32 	%f303, %f440, %f440, %f302;
	fma.rn.f32 	%f73, %f441, %f441, %f303;
	setp.eq.f32 	%p31, %f73, 0f00000000;
	mov.u16 	%rs61, %rs1;
	@%p31 bra 	$L__BB0_35;

	cvt.s64.s32 	%rd45, %r21;
	add.s64 	%rd46, %rd3, %rd45;
	ld.global.nc.u8 	%rs61, [%rd46];

$L__BB0_35:
	min.u16 	%rs35, %rs61, %rs1;
	cvt.u32.u16 	%r89, %rs35;
	max.u16 	%rs36, %rs61, %rs1;
	cvt.u32.u16 	%r90, %rs36;
	add.s32 	%r91, %r90, 1;
	mul.lo.s32 	%r92, %r91, %r90;
	shr.u32 	%r93, %r92, 1;
	add.s32 	%r94, %r93, %r89;
	mul.wide.s32 	%rd47, %r94, 4;
	add.s64 	%rd48, %rd2, %rd47;
	add.s64 	%rd49, %rd1, %rd47;
	ld.global.nc.f32 	%f74, [%rd49];
	ld.global.nc.f32 	%f75, [%rd48];
	setp.leu.f32 	%p32, %f75, 0f00000000;
	mov.f32 	%f438, %f74;
	@%p32 bra 	$L__BB0_37;

	add.f32 	%f304, %f75, %f75;
	div.rn.f32 	%f438, %f74, %f304;

$L__BB0_37:
	and.pred  	%p35, %p13, %p31;
	@%p35 bra 	$L__BB0_41;

	setp.neu.f32 	%p36, %f73, 0f00000000;
	@%p36 bra 	$L__BB0_40;

	mul.f32 	%f305, %f438, %f253;
	mul.f32 	%f306, %f6, %f305;
	sub.f32 	%f439, %f1, %f306;
	fma.rn.f32 	%f441, %f1, %f305, %f6;
	mov.f32 	%f440, %f5;

$L__BB0_40:
	mul.f32 	%f307, %f253, %f253;
	add.f32 	%f308, %f75, %f75;
	div.rn.f32 	%f309, %f308, %f307;
	sub.f32 	%f310, %f439, %f1;
	sub.f32 	%f311, %f440, %f5;
	sub.f32 	%f312, %f441, %f6;
	fma.rn.f32 	%f313, %f309, %f310, %f432;
	fma.rn.f32 	%f433, %f309, %f311, %f433;
	fma.rn.f32 	%f314, %f309, %f312, %f434;
	div.rn.f32 	%f315, %f74, %f253;
	fma.rn.f32 	%f432, %f315, %f441, %f313;
	mul.f32 	%f316, %f315, %f439;
	sub.f32 	%f434, %f314, %f316;

$L__BB0_41:
	add.s32 	%r22, %r2, 1;
	@%p27 bra 	$L__BB0_43;
	bra.uni 	$L__BB0_42;

$L__BB0_43:
	add.s32 	%r97, %r56, -1;
	min.s32 	%r159, %r22, %r97;
	bra.uni 	$L__BB0_44;

$L__BB0_42:
	rem.s32 	%r95, %r22, %r56;
	add.s32 	%r96, %r95, %r56;
	rem.s32 	%r159, %r96, %r56;

$L__BB0_44:
	add.s32 	%r98, %r159, %r4;
	mad.lo.s32 	%r26, %r98, %r55, %r1;
	setp.ge.s32 	%p38, %r22, %r56;
	mov.f32 	%f451, 0f00000000;
	and.pred  	%p40, %p38, %p27;
	mov.f32 	%f450, %f451;
	mov.f32 	%f449, %f451;
	@%p40 bra 	$L__BB0_46;

	mul.wide.s32 	%rd50, %r26, 4;
	add.s64 	%rd51, %rd6, %rd50;
	ld.global.nc.f32 	%f449, [%rd51];
	add.s64 	%rd52, %rd5, %rd50;
	ld.global.nc.f32 	%f450, [%rd52];
	add.s64 	%rd53, %rd4, %rd50;
	ld.global.nc.f32 	%f451, [%rd53];

$L__BB0_46:
	mul.f32 	%f320, %f449, %f449;
	fma.rn.f32 	%f321, %f450, %f450, %f320;
	fma.rn.f32 	%f103, %f451, %f451, %f321;
	setp.eq.f32 	%p41, %f103, 0f00000000;
	mov.u16 	%rs62, %rs1;
	@%p41 bra 	$L__BB0_48;

	cvt.s64.s32 	%rd54, %r26;
	add.s64 	%rd55, %rd3, %rd54;
	ld.global.nc.u8 	%rs62, [%rd55];

$L__BB0_48:
	min.u16 	%rs39, %rs62, %rs1;
	cvt.u32.u16 	%r99, %rs39;
	max.u16 	%rs40, %rs62, %rs1;
	cvt.u32.u16 	%r100, %rs40;
	add.s32 	%r101, %r100, 1;
	mul.lo.s32 	%r102, %r101, %r100;
	shr.u32 	%r103, %r102, 1;
	add.s32 	%r104, %r103, %r99;
	mul.wide.s32 	%rd56, %r104, 4;
	add.s64 	%rd57, %rd2, %rd56;
	add.s64 	%rd58, %rd1, %rd56;
	ld.global.nc.f32 	%f104, [%rd58];
	ld.global.nc.f32 	%f105, [%rd57];
	setp.leu.f32 	%p42, %f105, 0f00000000;
	mov.f32 	%f448, %f104;
	@%p42 bra 	$L__BB0_50;

	add.f32 	%f322, %f105, %f105;
	div.rn.f32 	%f448, %f104, %f322;

$L__BB0_50:
	and.pred  	%p45, %p13, %p41;
	@%p45 bra 	$L__BB0_54;

	setp.neu.f32 	%p46, %f103, 0f00000000;
	@%p46 bra 	$L__BB0_53;

	mul.f32 	%f323, %f448, %f253;
	fma.rn.f32 	%f449, %f6, %f323, %f1;
	mul.f32 	%f324, %f1, %f323;
	sub.f32 	%f451, %f6, %f324;
	mov.f32 	%f450, %f5;

$L__BB0_53:
	mul.f32 	%f325, %f253, %f253;
	add.f32 	%f326, %f105, %f105;
	div.rn.f32 	%f327, %f326, %f325;
	sub.f32 	%f328, %f449, %f1;
	sub.f32 	%f329, %f450, %f5;
	sub.f32 	%f330, %f451, %f6;
	fma.rn.f32 	%f331, %f327, %f328, %f432;
	fma.rn.f32 	%f433, %f327, %f329, %f433;
	fma.rn.f32 	%f332, %f327, %f330, %f434;
	div.rn.f32 	%f333, %f104, %f253;
	mul.f32 	%f334, %f333, %f451;
	sub.f32 	%f432, %f331, %f334;
	fma.rn.f32 	%f434, %f333, %f449, %f332;

$L__BB0_54:
	setp.eq.s32 	%p47, %r57, 1;
	@%p47 bra 	$L__BB0_115;

	and.b16  	%rs41, %rs24, 3;
	setp.ne.s16 	%p48, %rs41, 0;
	@%p48 bra 	$L__BB0_82;

	and.b16  	%rs12, %rs22, 4;
	setp.eq.s16 	%p49, %rs12, 0;
	add.s32 	%r27, %r3, -1;
	@%p49 bra 	$L__BB0_58;
	bra.uni 	$L__BB0_57;

$L__BB0_58:
	max.s32 	%r160, %r27, 0;
	bra.uni 	$L__BB0_59;

$L__BB0_57:
	rem.s32 	%r105, %r27, %r57;
	add.s32 	%r106, %r105, %r57;
	rem.s32 	%r160, %r106, %r57;

$L__BB0_59:
	mad.lo.s32 	%r107, %r160, %r56, %r2;
	mad.lo.s32 	%r31, %r107, %r55, %r1;
	setp.lt.s32 	%p51, %r3, 1;
	mov.f32 	%f461, 0f00000000;
	and.pred  	%p52, %p51, %p49;
	mov.f32 	%f460, %f461;
	mov.f32 	%f459, %f461;
	@%p52 bra 	$L__BB0_61;

	mul.wide.s32 	%rd59, %r31, 4;
	add.s64 	%rd60, %rd6, %rd59;
	ld.global.nc.f32 	%f459, [%rd60];
	add.s64 	%rd61, %rd5, %rd59;
	ld.global.nc.f32 	%f460, [%rd61];
	add.s64 	%rd62, %rd4, %rd59;
	ld.global.nc.f32 	%f461, [%rd62];

$L__BB0_61:
	mul.f32 	%f338, %f459, %f459;
	fma.rn.f32 	%f339, %f460, %f460, %f338;
	fma.rn.f32 	%f133, %f461, %f461, %f339;
	setp.eq.f32 	%p53, %f133, 0f00000000;
	mov.u16 	%rs63, %rs1;
	@%p53 bra 	$L__BB0_63;

	cvt.s64.s32 	%rd63, %r31;
	add.s64 	%rd64, %rd3, %rd63;
	ld.global.nc.u8 	%rs63, [%rd64];

$L__BB0_63:
	min.u16 	%rs44, %rs63, %rs1;
	cvt.u32.u16 	%r108, %rs44;
	max.u16 	%rs45, %rs63, %rs1;
	cvt.u32.u16 	%r109, %rs45;
	add.s32 	%r110, %r109, 1;
	mul.lo.s32 	%r111, %r110, %r109;
	shr.u32 	%r112, %r111, 1;
	add.s32 	%r113, %r112, %r108;
	mul.wide.s32 	%rd65, %r113, 4;
	add.s64 	%rd66, %rd2, %rd65;
	add.s64 	%rd67, %rd1, %rd65;
	ld.global.nc.f32 	%f134, [%rd67];
	ld.global.nc.f32 	%f135, [%rd66];
	setp.leu.f32 	%p54, %f135, 0f00000000;
	mov.f32 	%f458, %f134;
	@%p54 bra 	$L__BB0_65;

	add.f32 	%f340, %f135, %f135;
	div.rn.f32 	%f458, %f134, %f340;

$L__BB0_65:
	and.pred  	%p57, %p13, %p53;
	@%p57 bra 	$L__BB0_69;

	setp.neu.f32 	%p58, %f133, 0f00000000;
	@%p58 bra 	$L__BB0_68;

	mul.f32 	%f341, %f458, %f254;
	fma.rn.f32 	%f459, %f5, %f341, %f1;
	mul.f32 	%f342, %f1, %f341;
	sub.f32 	%f460, %f5, %f342;
	mov.f32 	%f461, %f6;

$L__BB0_68:
	mul.f32 	%f343, %f254, %f254;
	add.f32 	%f344, %f135, %f135;
	div.rn.f32 	%f345, %f344, %f343;
	sub.f32 	%f346, %f459, %f1;
	sub.f32 	%f347, %f460, %f5;
	sub.f32 	%f348, %f461, %f6;
	fma.rn.f32 	%f349, %f345, %f346, %f432;
	fma.rn.f32 	%f350, %f345, %f347, %f433;
	fma.rn.f32 	%f434, %f345, %f348, %f434;
	div.rn.f32 	%f351, %f134, %f254;
	mul.f32 	%f352, %f351, %f460;
	sub.f32 	%f432, %f349, %f352;
	fma.rn.f32 	%f433, %f351, %f459, %f350;

$L__BB0_69:
	add.s32 	%r32, %r3, 1;
	@%p49 bra 	$L__BB0_71;
	bra.uni 	$L__BB0_70;

$L__BB0_71:
	add.s32 	%r116, %r57, -1;
	min.s32 	%r161, %r32, %r116;
	bra.uni 	$L__BB0_72;

$L__BB0_70:
	rem.s32 	%r114, %r32, %r57;
	add.s32 	%r115, %r114, %r57;
	rem.s32 	%r161, %r115, %r57;

$L__BB0_72:
	mad.lo.s32 	%r117, %r161, %r56, %r2;
	mad.lo.s32 	%r36, %r117, %r55, %r1;
	setp.ge.s32 	%p60, %r32, %r57;
	mov.f32 	%f471, 0f00000000;
	and.pred  	%p62, %p60, %p49;
	mov.f32 	%f470, %f471;
	mov.f32 	%f469, %f471;
	@%p62 bra 	$L__BB0_74;

	mul.wide.s32 	%rd68, %r36, 4;
	add.s64 	%rd69, %rd6, %rd68;
	ld.global.nc.f32 	%f469, [%rd69];
	add.s64 	%rd70, %rd5, %rd68;
	ld.global.nc.f32 	%f470, [%rd70];
	add.s64 	%rd71, %rd4, %rd68;
	ld.global.nc.f32 	%f471, [%rd71];

$L__BB0_74:
	mul.f32 	%f356, %f469, %f469;
	fma.rn.f32 	%f357, %f470, %f470, %f356;
	fma.rn.f32 	%f163, %f471, %f471, %f357;
	setp.eq.f32 	%p63, %f163, 0f00000000;
	mov.u16 	%rs64, %rs1;
	@%p63 bra 	$L__BB0_76;

	cvt.s64.s32 	%rd72, %r36;
	add.s64 	%rd73, %rd3, %rd72;
	ld.global.nc.u8 	%rs64, [%rd73];

$L__BB0_76:
	min.u16 	%rs48, %rs64, %rs1;
	cvt.u32.u16 	%r118, %rs48;
	max.u16 	%rs49, %rs64, %rs1;
	cvt.u32.u16 	%r119, %rs49;
	add.s32 	%r120, %r119, 1;
	mul.lo.s32 	%r121, %r120, %r119;
	shr.u32 	%r122, %r121, 1;
	add.s32 	%r123, %r122, %r118;
	mul.wide.s32 	%rd74, %r123, 4;
	add.s64 	%rd75, %rd2, %rd74;
	add.s64 	%rd76, %rd1, %rd74;
	ld.global.nc.f32 	%f164, [%rd76];
	ld.global.nc.f32 	%f165, [%rd75];
	setp.leu.f32 	%p64, %f165, 0f00000000;
	mov.f32 	%f468, %f164;
	@%p64 bra 	$L__BB0_78;

	add.f32 	%f358, %f165, %f165;
	div.rn.f32 	%f468, %f164, %f358;

$L__BB0_78:
	and.pred  	%p67, %p13, %p63;
	@%p67 bra 	$L__BB0_82;

	setp.neu.f32 	%p68, %f163, 0f00000000;
	@%p68 bra 	$L__BB0_81;

	mul.f32 	%f359, %f468, %f254;
	mul.f32 	%f360, %f5, %f359;
	sub.f32 	%f469, %f1, %f360;
	fma.rn.f32 	%f470, %f1, %f359, %f5;
	mov.f32 	%f471, %f6;

$L__BB0_81:
	mul.f32 	%f361, %f254, %f254;
	add.f32 	%f362, %f165, %f165;
	div.rn.f32 	%f363, %f362, %f361;
	sub.f32 	%f364, %f469, %f1;
	sub.f32 	%f365, %f470, %f5;
	sub.f32 	%f366, %f471, %f6;
	fma.rn.f32 	%f367, %f363, %f364, %f432;
	fma.rn.f32 	%f368, %f363, %f365, %f433;
	fma.rn.f32 	%f434, %f363, %f366, %f434;
	div.rn.f32 	%f369, %f164, %f254;
	fma.rn.f32 	%f432, %f369, %f470, %f367;
	mul.f32 	%f370, %f369, %f469;
	sub.f32 	%f433, %f368, %f370;

$L__BB0_82:
	and.b16  	%rs50, %rs24, 2;
	setp.eq.s16 	%p69, %rs50, 0;
	@%p69 bra 	$L__BB0_115;

	and.b16  	%rs17, %rs22, 4;
	setp.eq.s16 	%p70, %rs17, 0;
	div.s32 	%r37, %r57, %r58;
	rem.s32 	%r38, %r3, %r37;
	@%p70 bra 	$L__BB0_85;
	bra.uni 	$L__BB0_84;

$L__BB0_85:
	add.s32 	%r127, %r38, -1;
	max.s32 	%r162, %r127, 0;
	bra.uni 	$L__BB0_86;

$L__BB0_84:
	add.s32 	%r124, %r38, -1;
	rem.s32 	%r125, %r124, %r37;
	add.s32 	%r126, %r125, %r37;
	rem.s32 	%r162, %r126, %r37;

$L__BB0_86:
	div.s32 	%r128, %r3, %r37;
	mul.lo.s32 	%r129, %r128, %r57;
	div.s32 	%r42, %r129, %r58;
	add.s32 	%r130, %r42, %r162;
	mad.lo.s32 	%r131, %r130, %r56, %r2;
	mad.lo.s32 	%r43, %r131, %r55, %r1;
	@%p70 bra 	$L__BB0_88;
	bra.uni 	$L__BB0_87;

$L__BB0_88:
	add.s32 	%r135, %r38, -1;
	max.s32 	%r163, %r135, 0;
	bra.uni 	$L__BB0_89;

$L__BB0_87:
	add.s32 	%r132, %r38, -1;
	rem.s32 	%r133, %r132, %r37;
	add.s32 	%r134, %r133, %r37;
	rem.s32 	%r163, %r134, %r37;

$L__BB0_89:
	setp.lt.s32 	%p73, %r163, 0;
	mov.f32 	%f481, 0f00000000;
	and.pred  	%p74, %p70, %p73;
	mov.f32 	%f480, %f481;
	mov.f32 	%f479, %f481;
	@%p74 bra 	$L__BB0_91;

	mul.wide.s32 	%rd77, %r43, 4;
	add.s64 	%rd78, %rd6, %rd77;
	ld.global.nc.f32 	%f479, [%rd78];
	add.s64 	%rd79, %rd5, %rd77;
	ld.global.nc.f32 	%f480, [%rd79];
	add.s64 	%rd80, %rd4, %rd77;
	ld.global.nc.f32 	%f481, [%rd80];

$L__BB0_91:
	mul.f32 	%f374, %f479, %f479;
	fma.rn.f32 	%f375, %f480, %f480, %f374;
	fma.rn.f32 	%f193, %f481, %f481, %f375;
	setp.eq.f32 	%p75, %f193, 0f00000000;
	mov.u16 	%rs65, %rs1;
	@%p75 bra 	$L__BB0_93;

	cvt.s64.s32 	%rd81, %r43;
	add.s64 	%rd82, %rd3, %rd81;
	ld.global.nc.u8 	%rs65, [%rd82];

$L__BB0_93:
	min.u16 	%rs53, %rs65, %rs1;
	cvt.u32.u16 	%r136, %rs53;
	max.u16 	%rs54, %rs65, %rs1;
	cvt.u32.u16 	%r137, %rs54;
	add.s32 	%r138, %r137, 1;
	mul.lo.s32 	%r139, %r138, %r137;
	shr.u32 	%r140, %r139, 1;
	add.s32 	%r141, %r140, %r136;
	mul.wide.s32 	%rd83, %r141, 4;
	add.s64 	%rd84, %rd2, %rd83;
	add.s64 	%rd85, %rd1, %rd83;
	ld.global.nc.f32 	%f194, [%rd85];
	ld.global.nc.f32 	%f195, [%rd84];
	setp.leu.f32 	%p76, %f195, 0f00000000;
	mov.f32 	%f478, %f194;
	@%p76 bra 	$L__BB0_95;

	add.f32 	%f376, %f195, %f195;
	div.rn.f32 	%f478, %f194, %f376;

$L__BB0_95:
	and.pred  	%p79, %p13, %p75;
	@%p79 bra 	$L__BB0_99;

	setp.neu.f32 	%p80, %f193, 0f00000000;
	@%p80 bra 	$L__BB0_98;

	mul.f32 	%f377, %f478, %f254;
	fma.rn.f32 	%f479, %f5, %f377, %f1;
	mul.f32 	%f378, %f1, %f377;
	sub.f32 	%f480, %f5, %f378;
	mov.f32 	%f481, %f6;

$L__BB0_98:
	mul.f32 	%f379, %f254, %f254;
	add.f32 	%f380, %f195, %f195;
	div.rn.f32 	%f381, %f380, %f379;
	sub.f32 	%f382, %f479, %f1;
	sub.f32 	%f383, %f480, %f5;
	sub.f32 	%f384, %f481, %f6;
	fma.rn.f32 	%f385, %f381, %f382, %f432;
	fma.rn.f32 	%f386, %f381, %f383, %f433;
	fma.rn.f32 	%f434, %f381, %f384, %f434;
	div.rn.f32 	%f387, %f194, %f254;
	mul.f32 	%f388, %f387, %f480;
	sub.f32 	%f432, %f385, %f388;
	fma.rn.f32 	%f433, %f387, %f479, %f386;

$L__BB0_99:
	add.s32 	%r47, %r38, 1;
	@%p70 bra 	$L__BB0_101;
	bra.uni 	$L__BB0_100;

$L__BB0_101:
	add.s32 	%r144, %r37, -1;
	min.s32 	%r164, %r47, %r144;
	bra.uni 	$L__BB0_102;

$L__BB0_100:
	rem.s32 	%r142, %r47, %r37;
	add.s32 	%r143, %r142, %r37;
	rem.s32 	%r164, %r143, %r37;

$L__BB0_102:
	add.s32 	%r145, %r164, %r42;
	mad.lo.s32 	%r146, %r145, %r56, %r2;
	mad.lo.s32 	%r51, %r146, %r55, %r1;
	@%p70 bra 	$L__BB0_104;
	bra.uni 	$L__BB0_103;

$L__BB0_104:
	add.s32 	%r149, %r37, -1;
	min.s32 	%r165, %r47, %r149;
	bra.uni 	$L__BB0_105;

$L__BB0_103:
	rem.s32 	%r147, %r47, %r37;
	add.s32 	%r148, %r147, %r37;
	rem.s32 	%r165, %r148, %r37;

$L__BB0_105:
	setp.ge.s32 	%p83, %r165, %r37;
	mov.f32 	%f491, 0f00000000;
	and.pred  	%p85, %p70, %p83;
	mov.f32 	%f490, %f491;
	mov.f32 	%f489, %f491;
	@%p85 bra 	$L__BB0_107;

	mul.wide.s32 	%rd86, %r51, 4;
	add.s64 	%rd87, %rd6, %rd86;
	ld.global.nc.f32 	%f489, [%rd87];
	add.s64 	%rd88, %rd5, %rd86;
	ld.global.nc.f32 	%f490, [%rd88];
	add.s64 	%rd89, %rd4, %rd86;
	ld.global.nc.f32 	%f491, [%rd89];

$L__BB0_107:
	mul.f32 	%f392, %f489, %f489;
	fma.rn.f32 	%f393, %f490, %f490, %f392;
	fma.rn.f32 	%f223, %f491, %f491, %f393;
	setp.eq.f32 	%p86, %f223, 0f00000000;
	mov.u16 	%rs66, %rs1;
	@%p86 bra 	$L__BB0_109;

	cvt.s64.s32 	%rd90, %r51;
	add.s64 	%rd91, %rd3, %rd90;
	ld.global.nc.u8 	%rs66, [%rd91];

$L__BB0_109:
	min.u16 	%rs57, %rs66, %rs1;
	cvt.u32.u16 	%r150, %rs57;
	max.u16 	%rs58, %rs66, %rs1;
	cvt.u32.u16 	%r151, %rs58;
	add.s32 	%r152, %r151, 1;
	mul.lo.s32 	%r153, %r152, %r151;
	shr.u32 	%r154, %r153, 1;
	add.s32 	%r155, %r154, %r150;
	mul.wide.s32 	%rd92, %r155, 4;
	add.s64 	%rd93, %rd2, %rd92;
	add.s64 	%rd94, %rd1, %rd92;
	ld.global.nc.f32 	%f224, [%rd94];
	ld.global.nc.f32 	%f225, [%rd93];
	setp.leu.f32 	%p87, %f225, 0f00000000;
	mov.f32 	%f488, %f224;
	@%p87 bra 	$L__BB0_111;

	add.f32 	%f394, %f225, %f225;
	div.rn.f32 	%f488, %f224, %f394;

$L__BB0_111:
	and.pred  	%p90, %p13, %p86;
	@%p90 bra 	$L__BB0_115;

	setp.neu.f32 	%p91, %f223, 0f00000000;
	@%p91 bra 	$L__BB0_114;

	mul.f32 	%f395, %f488, %f254;
	mul.f32 	%f396, %f5, %f395;
	sub.f32 	%f489, %f1, %f396;
	fma.rn.f32 	%f490, %f1, %f395, %f5;
	mov.f32 	%f491, %f6;

$L__BB0_114:
	mul.f32 	%f397, %f254, %f254;
	add.f32 	%f398, %f225, %f225;
	div.rn.f32 	%f399, %f398, %f397;
	sub.f32 	%f400, %f489, %f1;
	sub.f32 	%f401, %f490, %f5;
	sub.f32 	%f402, %f491, %f6;
	fma.rn.f32 	%f403, %f399, %f400, %f432;
	fma.rn.f32 	%f404, %f399, %f401, %f433;
	fma.rn.f32 	%f434, %f399, %f402, %f434;
	div.rn.f32 	%f405, %f224, %f254;
	fma.rn.f32 	%f432, %f405, %f490, %f403;
	mul.f32 	%f406, %f405, %f489;
	sub.f32 	%f433, %f404, %f406;

$L__BB0_115:
	setp.eq.s64 	%p92, %rd10, 0;
	@%p92 bra 	$L__BB0_117;

	cvta.to.global.u64 	%rd95, %rd10;
	add.s64 	%rd97, %rd95, %rd17;
	ld.global.nc.f32 	%f407, [%rd97];
	mul.f32 	%f495, %f407, %f495;

$L__BB0_117:
	setp.eq.f32 	%p93, %f495, 0f00000000;
	mov.f32 	%f496, 0f00000000;
	@%p93 bra 	$L__BB0_119;

	rcp.rn.f32 	%f496, %f495;

$L__BB0_119:
	cvta.to.global.u64 	%rd98, %rd7;
	add.s64 	%rd100, %rd98, %rd17;
	ld.global.f32 	%f409, [%rd100];
	fma.rn.f32 	%f410, %f432, %f496, %f409;
	st.global.f32 	[%rd100], %f410;
	cvta.to.global.u64 	%rd101, %rd8;
	add.s64 	%rd102, %rd101, %rd17;
	ld.global.f32 	%f411, [%rd102];
	fma.rn.f32 	%f412, %f433, %f496, %f411;
	st.global.f32 	[%rd102], %f412;
	cvta.to.global.u64 	%rd103, %rd9;
	add.s64 	%rd104, %rd103, %rd17;
	ld.global.f32 	%f413, [%rd104];
	fma.rn.f32 	%f414, %f434, %f496, %f413;
	st.global.f32 	[%rd104], %f414;

$L__BB0_120:
	ret;

}

`
)
