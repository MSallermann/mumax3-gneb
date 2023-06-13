package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/kuchkin/mumax3-gneb/cuda/cu"
	"github.com/kuchkin/mumax3-gneb/timer"
	"sync"
	"unsafe"
)

// CUDA handle for velocity_rodrig kernel
var velocity_rodrig_code cu.Function

// Stores the arguments for velocity_rodrig kernel invocation
type velocity_rodrig_args_t struct {
	arg_vx  unsafe.Pointer
	arg_vy  unsafe.Pointer
	arg_vz  unsafe.Pointer
	arg_kx  unsafe.Pointer
	arg_ky  unsafe.Pointer
	arg_kz  unsafe.Pointer
	arg_mx  unsafe.Pointer
	arg_my  unsafe.Pointer
	arg_mz  unsafe.Pointer
	arg_m0x unsafe.Pointer
	arg_m0y unsafe.Pointer
	arg_m0z unsafe.Pointer
	arg_N   int
	arg_Nz  int
	arg_dt  float32
	argptr  [15]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for velocity_rodrig kernel invocation
var velocity_rodrig_args velocity_rodrig_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	velocity_rodrig_args.argptr[0] = unsafe.Pointer(&velocity_rodrig_args.arg_vx)
	velocity_rodrig_args.argptr[1] = unsafe.Pointer(&velocity_rodrig_args.arg_vy)
	velocity_rodrig_args.argptr[2] = unsafe.Pointer(&velocity_rodrig_args.arg_vz)
	velocity_rodrig_args.argptr[3] = unsafe.Pointer(&velocity_rodrig_args.arg_kx)
	velocity_rodrig_args.argptr[4] = unsafe.Pointer(&velocity_rodrig_args.arg_ky)
	velocity_rodrig_args.argptr[5] = unsafe.Pointer(&velocity_rodrig_args.arg_kz)
	velocity_rodrig_args.argptr[6] = unsafe.Pointer(&velocity_rodrig_args.arg_mx)
	velocity_rodrig_args.argptr[7] = unsafe.Pointer(&velocity_rodrig_args.arg_my)
	velocity_rodrig_args.argptr[8] = unsafe.Pointer(&velocity_rodrig_args.arg_mz)
	velocity_rodrig_args.argptr[9] = unsafe.Pointer(&velocity_rodrig_args.arg_m0x)
	velocity_rodrig_args.argptr[10] = unsafe.Pointer(&velocity_rodrig_args.arg_m0y)
	velocity_rodrig_args.argptr[11] = unsafe.Pointer(&velocity_rodrig_args.arg_m0z)
	velocity_rodrig_args.argptr[12] = unsafe.Pointer(&velocity_rodrig_args.arg_N)
	velocity_rodrig_args.argptr[13] = unsafe.Pointer(&velocity_rodrig_args.arg_Nz)
	velocity_rodrig_args.argptr[14] = unsafe.Pointer(&velocity_rodrig_args.arg_dt)
}

// Wrapper for velocity_rodrig CUDA kernel, asynchronous.
func k_velocity_rodrig_async(vx unsafe.Pointer, vy unsafe.Pointer, vz unsafe.Pointer, kx unsafe.Pointer, ky unsafe.Pointer, kz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, m0x unsafe.Pointer, m0y unsafe.Pointer, m0z unsafe.Pointer, N int, Nz int, dt float32, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("velocity_rodrig")
	}

	velocity_rodrig_args.Lock()
	defer velocity_rodrig_args.Unlock()

	if velocity_rodrig_code == 0 {
		velocity_rodrig_code = fatbinLoad(velocity_rodrig_map, "velocity_rodrig")
	}

	velocity_rodrig_args.arg_vx = vx
	velocity_rodrig_args.arg_vy = vy
	velocity_rodrig_args.arg_vz = vz
	velocity_rodrig_args.arg_kx = kx
	velocity_rodrig_args.arg_ky = ky
	velocity_rodrig_args.arg_kz = kz
	velocity_rodrig_args.arg_mx = mx
	velocity_rodrig_args.arg_my = my
	velocity_rodrig_args.arg_mz = mz
	velocity_rodrig_args.arg_m0x = m0x
	velocity_rodrig_args.arg_m0y = m0y
	velocity_rodrig_args.arg_m0z = m0z
	velocity_rodrig_args.arg_N = N
	velocity_rodrig_args.arg_Nz = Nz
	velocity_rodrig_args.arg_dt = dt

	args := velocity_rodrig_args.argptr[:]
	cu.LaunchKernel(velocity_rodrig_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("velocity_rodrig")
	}
}

// maps compute capability on PTX code for velocity_rodrig kernel.
var velocity_rodrig_map = map[int]string{0: "",
	50: velocity_rodrig_ptx_50}

// velocity_rodrig PTX code for various compute capabilities.
const (
	velocity_rodrig_ptx_50 = `
.version 7.5
.target sm_50
.address_size 64

	// .globl	velocity_rodrig
.global .align 4 .b8 __cudart_i2opi_f[24] = {65, 144, 67, 60, 153, 149, 98, 219, 192, 221, 52, 245, 209, 87, 39, 252, 41, 21, 68, 78, 110, 131, 249, 162};

.visible .entry velocity_rodrig(
	.param .u64 velocity_rodrig_param_0,
	.param .u64 velocity_rodrig_param_1,
	.param .u64 velocity_rodrig_param_2,
	.param .u64 velocity_rodrig_param_3,
	.param .u64 velocity_rodrig_param_4,
	.param .u64 velocity_rodrig_param_5,
	.param .u64 velocity_rodrig_param_6,
	.param .u64 velocity_rodrig_param_7,
	.param .u64 velocity_rodrig_param_8,
	.param .u64 velocity_rodrig_param_9,
	.param .u64 velocity_rodrig_param_10,
	.param .u64 velocity_rodrig_param_11,
	.param .u32 velocity_rodrig_param_12,
	.param .u32 velocity_rodrig_param_13,
	.param .f32 velocity_rodrig_param_14
)
{
	.local .align 4 .b8 	__local_depot0[28];
	.reg .b64 	%SP;
	.reg .b64 	%SPL;
	.reg .pred 	%p<53>;
	.reg .f32 	%f<199>;
	.reg .b32 	%r<264>;
	.reg .f64 	%fd<30>;
	.reg .b64 	%rd<131>;


	mov.u64 	%SPL, __local_depot0;
	ld.param.u64 	%rd34, [velocity_rodrig_param_0];
	ld.param.u64 	%rd35, [velocity_rodrig_param_1];
	ld.param.u64 	%rd36, [velocity_rodrig_param_2];
	ld.param.u64 	%rd37, [velocity_rodrig_param_3];
	ld.param.u64 	%rd38, [velocity_rodrig_param_4];
	ld.param.u64 	%rd39, [velocity_rodrig_param_5];
	ld.param.u64 	%rd40, [velocity_rodrig_param_6];
	ld.param.u64 	%rd41, [velocity_rodrig_param_7];
	ld.param.u64 	%rd42, [velocity_rodrig_param_8];
	ld.param.u32 	%r92, [velocity_rodrig_param_12];
	ld.param.f32 	%f69, [velocity_rodrig_param_14];
	add.u64 	%rd1, %SPL, 0;
	mov.u32 	%r93, %nctaid.x;
	mov.u32 	%r94, %ctaid.y;
	mov.u32 	%r95, %ctaid.x;
	mad.lo.s32 	%r96, %r94, %r93, %r95;
	mov.u32 	%r97, %ntid.x;
	mov.u32 	%r98, %tid.x;
	mad.lo.s32 	%r1, %r96, %r97, %r98;
	setp.ge.s32 	%p1, %r1, %r92;
	@%p1 bra 	$L__BB0_65;

	cvta.to.global.u64 	%rd44, %rd40;
	cvt.s64.s32 	%rd2, %r1;
	mul.wide.s32 	%rd45, %r1, 4;
	add.s64 	%rd46, %rd44, %rd45;
	cvta.to.global.u64 	%rd47, %rd41;
	add.s64 	%rd48, %rd47, %rd45;
	cvta.to.global.u64 	%rd49, %rd42;
	add.s64 	%rd50, %rd49, %rd45;
	cvta.to.global.u64 	%rd51, %rd37;
	add.s64 	%rd52, %rd51, %rd45;
	cvta.to.global.u64 	%rd53, %rd38;
	add.s64 	%rd54, %rd53, %rd45;
	cvta.to.global.u64 	%rd55, %rd39;
	add.s64 	%rd56, %rd55, %rd45;
	ld.global.nc.f32 	%f1, [%rd52];
	mul.f32 	%f70, %f1, %f69;
	ld.global.nc.f32 	%f2, [%rd54];
	mul.f32 	%f71, %f2, %f69;
	ld.global.nc.f32 	%f3, [%rd56];
	mul.f32 	%f72, %f3, %f69;
	mul.f32 	%f73, %f71, %f71;
	fma.rn.f32 	%f74, %f70, %f70, %f73;
	fma.rn.f32 	%f75, %f72, %f72, %f74;
	sqrt.rn.f32 	%f4, %f75;
	ld.global.nc.f32 	%f5, [%rd48];
	mul.f32 	%f76, %f5, %f72;
	ld.global.nc.f32 	%f6, [%rd50];
	mul.f32 	%f77, %f6, %f71;
	sub.f32 	%f7, %f76, %f77;
	mul.f32 	%f78, %f6, %f70;
	ld.global.nc.f32 	%f8, [%rd46];
	mul.f32 	%f79, %f8, %f72;
	sub.f32 	%f9, %f78, %f79;
	mul.f32 	%f80, %f8, %f71;
	mul.f32 	%f81, %f5, %f70;
	sub.f32 	%f10, %f80, %f81;
	mul.f32 	%f82, %f2, %f71;
	fma.rn.f32 	%f83, %f1, %f70, %f82;
	fma.rn.f32 	%f11, %f3, %f72, %f83;
	mul.f32 	%f84, %f2, %f9;
	fma.rn.f32 	%f85, %f1, %f7, %f84;
	fma.rn.f32 	%f12, %f3, %f10, %f85;
	cvt.f64.f32 	%fd4, %f4;
	setp.gt.f64 	%p2, %fd4, 0d3F747AE147AE147B;
	mul.f32 	%f86, %f4, 0f3F22F983;
	cvt.rni.s32.f32 	%r263, %f86;
	cvt.rn.f32.s32 	%f87, %r263;
	mov.f32 	%f88, 0fBFC90FDA;
	fma.rn.f32 	%f89, %f87, %f88, %f4;
	mov.f32 	%f90, 0fB3A22168;
	fma.rn.f32 	%f91, %f87, %f90, %f89;
	mov.f32 	%f92, 0fA7C234C5;
	fma.rn.f32 	%f196, %f87, %f92, %f91;
	abs.f32 	%f14, %f4;
	add.s64 	%rd3, %rd1, 24;
	@%p2 bra 	$L__BB0_27;
	bra.uni 	$L__BB0_2;

$L__BB0_27:
	mul.f32 	%f122, %f4, %f4;
	cvt.f64.f32 	%fd13, %f122;
	div.rn.f64 	%fd14, %fd13, 0dC034000000000000;
	add.f64 	%fd15, %fd14, 0d3FF0000000000000;
	mul.f64 	%fd16, %fd15, %fd13;
	div.rn.f64 	%fd17, %fd16, 0dC018000000000000;
	add.f64 	%fd18, %fd17, 0d3FF0000000000000;
	cvt.rn.f32.f64 	%f189, %fd18;
	div.rn.f64 	%fd19, %fd13, 0dC03E000000000000;
	add.f64 	%fd20, %fd19, 0d3FF0000000000000;
	mul.f64 	%fd21, %fd20, %fd13;
	div.rn.f64 	%fd22, %fd21, 0dC038000000000000;
	add.f64 	%fd29, %fd22, 0d3FE0000000000000;
	bra.uni 	$L__BB0_28;

$L__BB0_2:
	setp.leu.f32 	%p3, %f14, 0f47CE4780;
	mov.u32 	%r247, %r263;
	mov.f32 	%f183, %f196;
	@%p3 bra 	$L__BB0_10;

	setp.eq.f32 	%p4, %f14, 0f7F800000;
	@%p4 bra 	$L__BB0_9;
	bra.uni 	$L__BB0_4;

$L__BB0_9:
	mov.f32 	%f95, 0f00000000;
	mul.rn.f32 	%f183, %f4, %f95;
	mov.u32 	%r247, %r263;
	bra.uni 	$L__BB0_10;

$L__BB0_4:
	mov.b32 	%r3, %f4;
	bfe.u32 	%r100, %r3, 23, 8;
	add.s32 	%r4, %r100, -128;
	shl.b32 	%r101, %r3, 8;
	or.b32  	%r5, %r101, -2147483648;
	shr.u32 	%r6, %r4, 5;
	mov.u64 	%rd118, 0;
	mov.u32 	%r244, 0;
	mov.u64 	%rd116, __cudart_i2opi_f;
	mov.u64 	%rd117, %rd1;

$L__BB0_5:
	.pragma "nounroll";
	ld.global.nc.u32 	%r102, [%rd116];
	mad.wide.u32 	%rd59, %r102, %r5, %rd118;
	shr.u64 	%rd118, %rd59, 32;
	st.local.u32 	[%rd117], %rd59;
	add.s64 	%rd117, %rd117, 4;
	add.s64 	%rd116, %rd116, 4;
	add.s32 	%r244, %r244, 1;
	setp.ne.s32 	%p5, %r244, 6;
	@%p5 bra 	$L__BB0_5;

	st.local.u32 	[%rd3], %rd118;
	mov.u32 	%r103, 4;
	sub.s32 	%r9, %r103, %r6;
	mov.u32 	%r104, 6;
	sub.s32 	%r105, %r104, %r6;
	mul.wide.s32 	%rd60, %r105, 4;
	add.s64 	%rd61, %rd1, %rd60;
	ld.local.u32 	%r245, [%rd61];
	ld.local.u32 	%r246, [%rd61+-4];
	and.b32  	%r12, %r4, 31;
	setp.eq.s32 	%p6, %r12, 0;
	@%p6 bra 	$L__BB0_8;

	mov.u32 	%r106, 32;
	sub.s32 	%r107, %r106, %r12;
	shr.u32 	%r108, %r246, %r107;
	shl.b32 	%r109, %r245, %r12;
	add.s32 	%r245, %r108, %r109;
	mul.wide.s32 	%rd62, %r9, 4;
	add.s64 	%rd63, %rd1, %rd62;
	ld.local.u32 	%r110, [%rd63];
	shr.u32 	%r111, %r110, %r107;
	shl.b32 	%r112, %r246, %r12;
	add.s32 	%r246, %r111, %r112;

$L__BB0_8:
	and.b32  	%r113, %r3, -2147483648;
	shr.u32 	%r114, %r246, 30;
	shl.b32 	%r115, %r245, 2;
	or.b32  	%r116, %r114, %r115;
	shr.u32 	%r117, %r116, 31;
	shr.u32 	%r118, %r245, 30;
	add.s32 	%r119, %r117, %r118;
	neg.s32 	%r120, %r119;
	setp.eq.s32 	%p7, %r113, 0;
	selp.b32 	%r247, %r119, %r120, %p7;
	setp.ne.s32 	%p8, %r117, 0;
	xor.b32  	%r121, %r113, -2147483648;
	selp.b32 	%r122, %r121, %r113, %p8;
	selp.b32 	%r123, -1, 0, %p8;
	xor.b32  	%r124, %r116, %r123;
	shl.b32 	%r125, %r246, 2;
	xor.b32  	%r126, %r125, %r123;
	cvt.u64.u32 	%rd64, %r124;
	cvt.u64.u32 	%rd65, %r126;
	bfi.b64 	%rd66, %rd64, %rd65, 32, 32;
	cvt.rn.f64.s64 	%fd5, %rd66;
	mul.f64 	%fd6, %fd5, 0d3BF921FB54442D19;
	cvt.rn.f32.f64 	%f93, %fd6;
	setp.eq.s32 	%p9, %r122, 0;
	neg.f32 	%f94, %f93;
	selp.f32 	%f183, %f93, %f94, %p9;

$L__BB0_10:
	and.b32  	%r19, %r247, 1;
	setp.eq.s32 	%p10, %r19, 0;
	selp.f32 	%f18, %f183, 0f3F800000, %p10;
	mul.rn.f32 	%f19, %f183, %f183;
	mov.f32 	%f184, 0fB94D4153;
	@%p10 bra 	$L__BB0_12;

	mov.f32 	%f97, 0fBAB607ED;
	mov.f32 	%f98, 0f37CBAC00;
	fma.rn.f32 	%f184, %f98, %f19, %f97;

$L__BB0_12:
	selp.f32 	%f99, 0f3C0885E4, 0f3D2AAABB, %p10;
	fma.rn.f32 	%f100, %f184, %f19, %f99;
	selp.f32 	%f101, 0fBE2AAAA8, 0fBEFFFFFF, %p10;
	fma.rn.f32 	%f102, %f100, %f19, %f101;
	mov.f32 	%f103, 0f00000000;
	fma.rn.f32 	%f104, %f19, %f18, %f103;
	fma.rn.f32 	%f185, %f102, %f104, %f18;
	and.b32  	%r127, %r247, 2;
	setp.eq.s32 	%p12, %r127, 0;
	@%p12 bra 	$L__BB0_14;

	mov.f32 	%f106, 0fBF800000;
	fma.rn.f32 	%f185, %f185, %f106, %f103;

$L__BB0_14:
	mov.u32 	%r251, %r263;
	mov.f32 	%f186, %f196;
	@%p3 bra 	$L__BB0_22;

	setp.eq.f32 	%p14, %f14, 0f7F800000;
	@%p14 bra 	$L__BB0_21;
	bra.uni 	$L__BB0_16;

$L__BB0_21:
	mul.rn.f32 	%f186, %f4, %f103;
	mov.u32 	%r251, %r263;
	bra.uni 	$L__BB0_22;

$L__BB0_16:
	mov.b32 	%r20, %f4;
	bfe.u32 	%r129, %r20, 23, 8;
	add.s32 	%r21, %r129, -128;
	shl.b32 	%r130, %r20, 8;
	or.b32  	%r22, %r130, -2147483648;
	shr.u32 	%r23, %r21, 5;
	mov.u64 	%rd121, 0;
	mov.u32 	%r248, 0;
	mov.u64 	%rd119, __cudart_i2opi_f;
	mov.u64 	%rd120, %rd1;

$L__BB0_17:
	.pragma "nounroll";
	ld.global.nc.u32 	%r131, [%rd119];
	mad.wide.u32 	%rd69, %r131, %r22, %rd121;
	shr.u64 	%rd121, %rd69, 32;
	st.local.u32 	[%rd120], %rd69;
	add.s64 	%rd120, %rd120, 4;
	add.s64 	%rd119, %rd119, 4;
	add.s32 	%r248, %r248, 1;
	setp.ne.s32 	%p15, %r248, 6;
	@%p15 bra 	$L__BB0_17;

	st.local.u32 	[%rd3], %rd121;
	mov.u32 	%r132, 4;
	sub.s32 	%r26, %r132, %r23;
	mov.u32 	%r133, 6;
	sub.s32 	%r134, %r133, %r23;
	mul.wide.s32 	%rd70, %r134, 4;
	add.s64 	%rd71, %rd1, %rd70;
	ld.local.u32 	%r249, [%rd71];
	ld.local.u32 	%r250, [%rd71+-4];
	and.b32  	%r29, %r21, 31;
	setp.eq.s32 	%p16, %r29, 0;
	@%p16 bra 	$L__BB0_20;

	mov.u32 	%r135, 32;
	sub.s32 	%r136, %r135, %r29;
	shr.u32 	%r137, %r250, %r136;
	shl.b32 	%r138, %r249, %r29;
	add.s32 	%r249, %r137, %r138;
	mul.wide.s32 	%rd72, %r26, 4;
	add.s64 	%rd73, %rd1, %rd72;
	ld.local.u32 	%r139, [%rd73];
	shr.u32 	%r140, %r139, %r136;
	shl.b32 	%r141, %r250, %r29;
	add.s32 	%r250, %r140, %r141;

$L__BB0_20:
	and.b32  	%r142, %r20, -2147483648;
	shr.u32 	%r143, %r250, 30;
	shl.b32 	%r144, %r249, 2;
	or.b32  	%r145, %r143, %r144;
	shr.u32 	%r146, %r145, 31;
	shr.u32 	%r147, %r249, 30;
	add.s32 	%r148, %r146, %r147;
	neg.s32 	%r149, %r148;
	setp.eq.s32 	%p17, %r142, 0;
	selp.b32 	%r251, %r148, %r149, %p17;
	setp.ne.s32 	%p18, %r146, 0;
	xor.b32  	%r150, %r142, -2147483648;
	selp.b32 	%r151, %r150, %r142, %p18;
	selp.b32 	%r152, -1, 0, %p18;
	xor.b32  	%r153, %r145, %r152;
	shl.b32 	%r154, %r250, 2;
	xor.b32  	%r155, %r154, %r152;
	cvt.u64.u32 	%rd74, %r153;
	cvt.u64.u32 	%rd75, %r155;
	bfi.b64 	%rd76, %rd74, %rd75, 32, 32;
	cvt.rn.f64.s64 	%fd7, %rd76;
	mul.f64 	%fd8, %fd7, 0d3BF921FB54442D19;
	cvt.rn.f32.f64 	%f107, %fd8;
	setp.eq.s32 	%p19, %r151, 0;
	neg.f32 	%f108, %f107;
	selp.f32 	%f186, %f107, %f108, %p19;

$L__BB0_22:
	add.s32 	%r36, %r251, 1;
	and.b32  	%r37, %r36, 1;
	setp.eq.s32 	%p20, %r37, 0;
	selp.f32 	%f28, %f186, 0f3F800000, %p20;
	mul.rn.f32 	%f29, %f186, %f186;
	mov.f32 	%f187, 0fB94D4153;
	@%p20 bra 	$L__BB0_24;

	mov.f32 	%f111, 0fBAB607ED;
	mov.f32 	%f112, 0f37CBAC00;
	fma.rn.f32 	%f187, %f112, %f29, %f111;

$L__BB0_24:
	selp.f32 	%f113, 0f3C0885E4, 0f3D2AAABB, %p20;
	fma.rn.f32 	%f114, %f187, %f29, %f113;
	selp.f32 	%f115, 0fBE2AAAA8, 0fBEFFFFFF, %p20;
	fma.rn.f32 	%f116, %f114, %f29, %f115;
	fma.rn.f32 	%f118, %f29, %f28, %f103;
	fma.rn.f32 	%f188, %f116, %f118, %f28;
	and.b32  	%r156, %r36, 2;
	setp.eq.s32 	%p22, %r156, 0;
	@%p22 bra 	$L__BB0_26;

	mov.f32 	%f120, 0fBF800000;
	fma.rn.f32 	%f188, %f188, %f120, %f103;

$L__BB0_26:
	cvt.f64.f32 	%fd9, %f188;
	mov.f64 	%fd10, 0d3FF0000000000000;
	sub.f64 	%fd11, %fd10, %fd9;
	mul.f32 	%f121, %f4, %f4;
	cvt.f64.f32 	%fd12, %f121;
	div.rn.f64 	%fd29, %fd11, %fd12;
	div.rn.f32 	%f189, %f185, %f4;

$L__BB0_28:
	cvt.rn.f32.f64 	%f38, %fd29;
	setp.leu.f32 	%p23, %f14, 0f47CE4780;
	mov.u32 	%r255, %r263;
	mov.f32 	%f190, %f196;
	@%p23 bra 	$L__BB0_36;

	setp.eq.f32 	%p24, %f14, 0f7F800000;
	@%p24 bra 	$L__BB0_35;
	bra.uni 	$L__BB0_30;

$L__BB0_35:
	mov.f32 	%f125, 0f00000000;
	mul.rn.f32 	%f190, %f4, %f125;
	mov.u32 	%r255, %r263;
	bra.uni 	$L__BB0_36;

$L__BB0_30:
	mov.b32 	%r38, %f4;
	bfe.u32 	%r158, %r38, 23, 8;
	add.s32 	%r39, %r158, -128;
	shl.b32 	%r159, %r38, 8;
	or.b32  	%r40, %r159, -2147483648;
	shr.u32 	%r41, %r39, 5;
	mov.u64 	%rd124, 0;
	mov.u32 	%r252, 0;
	mov.u64 	%rd122, __cudart_i2opi_f;
	mov.u64 	%rd123, %rd1;

$L__BB0_31:
	.pragma "nounroll";
	ld.global.nc.u32 	%r160, [%rd122];
	mad.wide.u32 	%rd79, %r160, %r40, %rd124;
	shr.u64 	%rd124, %rd79, 32;
	st.local.u32 	[%rd123], %rd79;
	add.s64 	%rd123, %rd123, 4;
	add.s64 	%rd122, %rd122, 4;
	add.s32 	%r252, %r252, 1;
	setp.ne.s32 	%p25, %r252, 6;
	@%p25 bra 	$L__BB0_31;

	st.local.u32 	[%rd3], %rd124;
	mov.u32 	%r161, 4;
	sub.s32 	%r44, %r161, %r41;
	mov.u32 	%r162, 6;
	sub.s32 	%r163, %r162, %r41;
	mul.wide.s32 	%rd80, %r163, 4;
	add.s64 	%rd81, %rd1, %rd80;
	ld.local.u32 	%r253, [%rd81];
	ld.local.u32 	%r254, [%rd81+-4];
	and.b32  	%r47, %r39, 31;
	setp.eq.s32 	%p26, %r47, 0;
	@%p26 bra 	$L__BB0_34;

	mov.u32 	%r164, 32;
	sub.s32 	%r165, %r164, %r47;
	shr.u32 	%r166, %r254, %r165;
	shl.b32 	%r167, %r253, %r47;
	add.s32 	%r253, %r166, %r167;
	mul.wide.s32 	%rd82, %r44, 4;
	add.s64 	%rd83, %rd1, %rd82;
	ld.local.u32 	%r168, [%rd83];
	shr.u32 	%r169, %r168, %r165;
	shl.b32 	%r170, %r254, %r47;
	add.s32 	%r254, %r169, %r170;

$L__BB0_34:
	and.b32  	%r171, %r38, -2147483648;
	shr.u32 	%r172, %r254, 30;
	shl.b32 	%r173, %r253, 2;
	or.b32  	%r174, %r172, %r173;
	shr.u32 	%r175, %r174, 31;
	shr.u32 	%r176, %r253, 30;
	add.s32 	%r177, %r175, %r176;
	neg.s32 	%r178, %r177;
	setp.eq.s32 	%p27, %r171, 0;
	selp.b32 	%r255, %r177, %r178, %p27;
	setp.ne.s32 	%p28, %r175, 0;
	xor.b32  	%r179, %r171, -2147483648;
	selp.b32 	%r180, %r179, %r171, %p28;
	selp.b32 	%r181, -1, 0, %p28;
	xor.b32  	%r182, %r174, %r181;
	shl.b32 	%r183, %r254, 2;
	xor.b32  	%r184, %r183, %r181;
	cvt.u64.u32 	%rd84, %r182;
	cvt.u64.u32 	%rd85, %r184;
	bfi.b64 	%rd86, %rd84, %rd85, 32, 32;
	cvt.rn.f64.s64 	%fd23, %rd86;
	mul.f64 	%fd24, %fd23, 0d3BF921FB54442D19;
	cvt.rn.f32.f64 	%f123, %fd24;
	setp.eq.s32 	%p29, %r180, 0;
	neg.f32 	%f124, %f123;
	selp.f32 	%f190, %f123, %f124, %p29;

$L__BB0_36:
	add.s32 	%r54, %r255, 1;
	and.b32  	%r55, %r54, 1;
	setp.eq.s32 	%p30, %r55, 0;
	selp.f32 	%f42, %f190, 0f3F800000, %p30;
	mul.rn.f32 	%f43, %f190, %f190;
	mov.f32 	%f191, 0fB94D4153;
	@%p30 bra 	$L__BB0_38;

	mov.f32 	%f127, 0fBAB607ED;
	mov.f32 	%f128, 0f37CBAC00;
	fma.rn.f32 	%f191, %f128, %f43, %f127;

$L__BB0_38:
	selp.f32 	%f129, 0f3C0885E4, 0f3D2AAABB, %p30;
	fma.rn.f32 	%f130, %f191, %f43, %f129;
	selp.f32 	%f131, 0fBE2AAAA8, 0fBEFFFFFF, %p30;
	fma.rn.f32 	%f132, %f130, %f43, %f131;
	mov.f32 	%f133, 0f00000000;
	fma.rn.f32 	%f134, %f43, %f42, %f133;
	fma.rn.f32 	%f192, %f132, %f134, %f42;
	and.b32  	%r185, %r54, 2;
	setp.eq.s32 	%p32, %r185, 0;
	@%p32 bra 	$L__BB0_40;

	mov.f32 	%f136, 0fBF800000;
	fma.rn.f32 	%f192, %f192, %f136, %f133;

$L__BB0_40:
	mul.f32 	%f137, %f11, %f8;
	mul.f32 	%f138, %f189, %f137;
	mul.f32 	%f139, %f1, %f192;
	sub.f32 	%f140, %f139, %f138;
	mul.f32 	%f141, %f7, %f12;
	fma.rn.f32 	%f142, %f141, %f38, %f140;
	cvta.to.global.u64 	%rd87, %rd34;
	shl.b64 	%rd88, %rd2, 2;
	add.s64 	%rd89, %rd87, %rd88;
	st.global.f32 	[%rd89], %f142;
	mov.u32 	%r259, %r263;
	mov.f32 	%f193, %f196;
	@%p23 bra 	$L__BB0_48;

	setp.eq.f32 	%p34, %f14, 0f7F800000;
	@%p34 bra 	$L__BB0_47;
	bra.uni 	$L__BB0_42;

$L__BB0_47:
	mul.rn.f32 	%f193, %f4, %f133;
	mov.u32 	%r259, %r263;
	bra.uni 	$L__BB0_48;

$L__BB0_42:
	mov.b32 	%r56, %f4;
	bfe.u32 	%r187, %r56, 23, 8;
	add.s32 	%r57, %r187, -128;
	shl.b32 	%r188, %r56, 8;
	or.b32  	%r58, %r188, -2147483648;
	shr.u32 	%r59, %r57, 5;
	mov.u64 	%rd127, 0;
	mov.u32 	%r256, 0;
	mov.u64 	%rd125, __cudart_i2opi_f;
	mov.u64 	%rd126, %rd1;

$L__BB0_43:
	.pragma "nounroll";
	ld.global.nc.u32 	%r189, [%rd125];
	mad.wide.u32 	%rd92, %r189, %r58, %rd127;
	shr.u64 	%rd127, %rd92, 32;
	st.local.u32 	[%rd126], %rd92;
	add.s64 	%rd126, %rd126, 4;
	add.s64 	%rd125, %rd125, 4;
	add.s32 	%r256, %r256, 1;
	setp.ne.s32 	%p35, %r256, 6;
	@%p35 bra 	$L__BB0_43;

	st.local.u32 	[%rd3], %rd127;
	mov.u32 	%r190, 4;
	sub.s32 	%r62, %r190, %r59;
	mov.u32 	%r191, 6;
	sub.s32 	%r192, %r191, %r59;
	mul.wide.s32 	%rd93, %r192, 4;
	add.s64 	%rd94, %rd1, %rd93;
	ld.local.u32 	%r257, [%rd94];
	ld.local.u32 	%r258, [%rd94+-4];
	and.b32  	%r65, %r57, 31;
	setp.eq.s32 	%p36, %r65, 0;
	@%p36 bra 	$L__BB0_46;

	mov.u32 	%r193, 32;
	sub.s32 	%r194, %r193, %r65;
	shr.u32 	%r195, %r258, %r194;
	shl.b32 	%r196, %r257, %r65;
	add.s32 	%r257, %r195, %r196;
	mul.wide.s32 	%rd95, %r62, 4;
	add.s64 	%rd96, %rd1, %rd95;
	ld.local.u32 	%r197, [%rd96];
	shr.u32 	%r198, %r197, %r194;
	shl.b32 	%r199, %r258, %r65;
	add.s32 	%r258, %r198, %r199;

$L__BB0_46:
	and.b32  	%r200, %r56, -2147483648;
	shr.u32 	%r201, %r258, 30;
	shl.b32 	%r202, %r257, 2;
	or.b32  	%r203, %r201, %r202;
	shr.u32 	%r204, %r203, 31;
	shr.u32 	%r205, %r257, 30;
	add.s32 	%r206, %r204, %r205;
	neg.s32 	%r207, %r206;
	setp.eq.s32 	%p37, %r200, 0;
	selp.b32 	%r259, %r206, %r207, %p37;
	setp.ne.s32 	%p38, %r204, 0;
	xor.b32  	%r208, %r200, -2147483648;
	selp.b32 	%r209, %r208, %r200, %p38;
	selp.b32 	%r210, -1, 0, %p38;
	xor.b32  	%r211, %r203, %r210;
	shl.b32 	%r212, %r258, 2;
	xor.b32  	%r213, %r212, %r210;
	cvt.u64.u32 	%rd97, %r211;
	cvt.u64.u32 	%rd98, %r213;
	bfi.b64 	%rd99, %rd97, %rd98, 32, 32;
	cvt.rn.f64.s64 	%fd25, %rd99;
	mul.f64 	%fd26, %fd25, 0d3BF921FB54442D19;
	cvt.rn.f32.f64 	%f143, %fd26;
	setp.eq.s32 	%p39, %r209, 0;
	neg.f32 	%f144, %f143;
	selp.f32 	%f193, %f143, %f144, %p39;

$L__BB0_48:
	add.s32 	%r72, %r259, 1;
	and.b32  	%r73, %r72, 1;
	setp.eq.s32 	%p40, %r73, 0;
	selp.f32 	%f52, %f193, 0f3F800000, %p40;
	mul.rn.f32 	%f53, %f193, %f193;
	mov.f32 	%f194, 0fB94D4153;
	@%p40 bra 	$L__BB0_50;

	mov.f32 	%f147, 0fBAB607ED;
	mov.f32 	%f148, 0f37CBAC00;
	fma.rn.f32 	%f194, %f148, %f53, %f147;

$L__BB0_50:
	selp.f32 	%f149, 0f3C0885E4, 0f3D2AAABB, %p40;
	fma.rn.f32 	%f150, %f194, %f53, %f149;
	selp.f32 	%f151, 0fBE2AAAA8, 0fBEFFFFFF, %p40;
	fma.rn.f32 	%f152, %f150, %f53, %f151;
	fma.rn.f32 	%f154, %f53, %f52, %f133;
	fma.rn.f32 	%f195, %f152, %f154, %f52;
	and.b32  	%r214, %r72, 2;
	setp.eq.s32 	%p42, %r214, 0;
	@%p42 bra 	$L__BB0_52;

	mov.f32 	%f156, 0fBF800000;
	fma.rn.f32 	%f195, %f195, %f156, %f133;

$L__BB0_52:
	mul.f32 	%f157, %f11, %f5;
	mul.f32 	%f158, %f189, %f157;
	mul.f32 	%f159, %f2, %f195;
	sub.f32 	%f160, %f159, %f158;
	mul.f32 	%f161, %f9, %f12;
	fma.rn.f32 	%f162, %f161, %f38, %f160;
	cvta.to.global.u64 	%rd100, %rd35;
	add.s64 	%rd102, %rd100, %rd88;
	st.global.f32 	[%rd102], %f162;
	@%p23 bra 	$L__BB0_60;

	setp.eq.f32 	%p44, %f14, 0f7F800000;
	@%p44 bra 	$L__BB0_59;
	bra.uni 	$L__BB0_54;

$L__BB0_59:
	mul.rn.f32 	%f196, %f4, %f133;
	bra.uni 	$L__BB0_60;

$L__BB0_54:
	mov.b32 	%r74, %f4;
	bfe.u32 	%r216, %r74, 23, 8;
	add.s32 	%r75, %r216, -128;
	shl.b32 	%r217, %r74, 8;
	or.b32  	%r76, %r217, -2147483648;
	shr.u32 	%r77, %r75, 5;
	mov.u64 	%rd130, 0;
	mov.u32 	%r260, 0;
	mov.u64 	%rd128, __cudart_i2opi_f;
	mov.u64 	%rd129, %rd1;

$L__BB0_55:
	.pragma "nounroll";
	ld.global.nc.u32 	%r218, [%rd128];
	mad.wide.u32 	%rd105, %r218, %r76, %rd130;
	shr.u64 	%rd130, %rd105, 32;
	st.local.u32 	[%rd129], %rd105;
	add.s64 	%rd129, %rd129, 4;
	add.s64 	%rd128, %rd128, 4;
	add.s32 	%r260, %r260, 1;
	setp.ne.s32 	%p45, %r260, 6;
	@%p45 bra 	$L__BB0_55;

	st.local.u32 	[%rd3], %rd130;
	mov.u32 	%r219, 4;
	sub.s32 	%r80, %r219, %r77;
	mov.u32 	%r220, 6;
	sub.s32 	%r221, %r220, %r77;
	mul.wide.s32 	%rd106, %r221, 4;
	add.s64 	%rd107, %rd1, %rd106;
	ld.local.u32 	%r261, [%rd107];
	ld.local.u32 	%r262, [%rd107+-4];
	and.b32  	%r83, %r75, 31;
	setp.eq.s32 	%p46, %r83, 0;
	@%p46 bra 	$L__BB0_58;

	mov.u32 	%r222, 32;
	sub.s32 	%r223, %r222, %r83;
	shr.u32 	%r224, %r262, %r223;
	shl.b32 	%r225, %r261, %r83;
	add.s32 	%r261, %r224, %r225;
	mul.wide.s32 	%rd108, %r80, 4;
	add.s64 	%rd109, %rd1, %rd108;
	ld.local.u32 	%r226, [%rd109];
	shr.u32 	%r227, %r226, %r223;
	shl.b32 	%r228, %r262, %r83;
	add.s32 	%r262, %r227, %r228;

$L__BB0_58:
	and.b32  	%r229, %r74, -2147483648;
	shr.u32 	%r230, %r262, 30;
	shl.b32 	%r231, %r261, 2;
	or.b32  	%r232, %r230, %r231;
	shr.u32 	%r233, %r232, 31;
	shr.u32 	%r234, %r261, 30;
	add.s32 	%r235, %r233, %r234;
	neg.s32 	%r236, %r235;
	setp.eq.s32 	%p47, %r229, 0;
	selp.b32 	%r263, %r235, %r236, %p47;
	setp.ne.s32 	%p48, %r233, 0;
	xor.b32  	%r237, %r229, -2147483648;
	selp.b32 	%r238, %r237, %r229, %p48;
	selp.b32 	%r239, -1, 0, %p48;
	xor.b32  	%r240, %r232, %r239;
	shl.b32 	%r241, %r262, 2;
	xor.b32  	%r242, %r241, %r239;
	cvt.u64.u32 	%rd110, %r240;
	cvt.u64.u32 	%rd111, %r242;
	bfi.b64 	%rd112, %rd110, %rd111, 32, 32;
	cvt.rn.f64.s64 	%fd27, %rd112;
	mul.f64 	%fd28, %fd27, 0d3BF921FB54442D19;
	cvt.rn.f32.f64 	%f163, %fd28;
	setp.eq.s32 	%p49, %r238, 0;
	neg.f32 	%f164, %f163;
	selp.f32 	%f196, %f163, %f164, %p49;

$L__BB0_60:
	add.s32 	%r90, %r263, 1;
	and.b32  	%r91, %r90, 1;
	setp.eq.s32 	%p50, %r91, 0;
	selp.f32 	%f62, %f196, 0f3F800000, %p50;
	mul.rn.f32 	%f63, %f196, %f196;
	mov.f32 	%f197, 0fB94D4153;
	@%p50 bra 	$L__BB0_62;

	mov.f32 	%f167, 0fBAB607ED;
	mov.f32 	%f168, 0f37CBAC00;
	fma.rn.f32 	%f197, %f168, %f63, %f167;

$L__BB0_62:
	selp.f32 	%f169, 0f3C0885E4, 0f3D2AAABB, %p50;
	fma.rn.f32 	%f170, %f197, %f63, %f169;
	selp.f32 	%f171, 0fBE2AAAA8, 0fBEFFFFFF, %p50;
	fma.rn.f32 	%f172, %f170, %f63, %f171;
	fma.rn.f32 	%f174, %f63, %f62, %f133;
	fma.rn.f32 	%f198, %f172, %f174, %f62;
	and.b32  	%r243, %r90, 2;
	setp.eq.s32 	%p52, %r243, 0;
	@%p52 bra 	$L__BB0_64;

	mov.f32 	%f176, 0fBF800000;
	fma.rn.f32 	%f198, %f198, %f176, %f133;

$L__BB0_64:
	mul.f32 	%f177, %f11, %f6;
	mul.f32 	%f178, %f189, %f177;
	mul.f32 	%f179, %f3, %f198;
	sub.f32 	%f180, %f179, %f178;
	mul.f32 	%f181, %f10, %f12;
	fma.rn.f32 	%f182, %f181, %f38, %f180;
	cvta.to.global.u64 	%rd113, %rd36;
	add.s64 	%rd115, %rd113, %rd88;
	st.global.f32 	[%rd115], %f182;

$L__BB0_65:
	ret;

}

`
)
