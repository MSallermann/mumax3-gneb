package engine

import (
	"code.google.com/p/mx3/cuda"
	"code.google.com/p/mx3/data"
	"log"
	"runtime"
)

const VERSION = "mx3.0.11 α "

var UNAME = VERSION + runtime.GOOS + "_" + runtime.GOARCH + " " + runtime.Version() + "(" + runtime.Compiler + ")"

var (
	M      magnetization // reduced magnetization (unit length)
	M_full setterQuant   // non-reduced magnetization in T
	B_eff  setterQuant   // effective field (T) output handle
	Torque setterQuant   // total torque/γ0, in T
	//Table  DataTable
)

var (
	globalmesh data.Mesh
	inited     = make(chan int, 1) // fires when engine ready to serve GUI
)

func init() {
	DeclFunc("setgridsize", setGridSize, `Sets the number of cells for X,Y,Z`)
	DeclFunc("setcellsize", setCellSize, `Sets the X,Y,Z cell size in meters`)
	DeclLValue("m", &M, `Reduced magnetization (unit length)`)

	//DeclROnly("table", &Table, `Provides methods for tabular output`)
	//Table = *newTable("datatable") // output handle for tabular data (average magnetization etc.)
}

func init() {

	M_full.init(3, &globalmesh, "m_full", "T", "Unnormalized magnetization", func(dst *data.Slice) {
		msat, r := Msat.Get()
		if r {
			defer cuda.RecycleBuffer(msat)
		}
		for c := 0; c < 3; c++ {
			cuda.Mul(dst.Comp(c), M.buffer.Comp(c), msat)
		}
	})

}

func initialize() {
	M.init()
	FFTM.init()

	regions.init()

	//Table.Add(&M)

	initDemag()
	initBExt()

	// effective field
	B_eff.init(3, Mesh(), "B_eff", "T", "Effective field", func(dst *data.Slice) {
		B_demag.set(dst)
		B_exch.addTo(dst)
		B_anis.addTo(dst)
		B_ext.addTo(dst)
	})

	torquebuffer := cuda.NewSlice(3, Mesh())
	torqueFn := func() *data.Slice {
		Torque.set(torquebuffer)
		return torquebuffer
	}
	Solver = *cuda.NewHeun(M.buffer, torqueFn, cuda.Normalize, 1e-15, Gamma0, &Time)

	inited <- 1
}

//func sanitycheck() {
//	if Msat() == 0 {
//		log.Fatal("Msat should be nonzero")
//	}
//}

func Mesh() *data.Mesh {
	checkMesh()
	return &globalmesh
}

func WorldSize() [3]float64 {
	w := Mesh().WorldSize()
	return [3]float64{w[2], w[1], w[0]} // swaps XYZ
}

// Set the simulation mesh to Nx x Ny x Nz cells of given size.
// Can be set only once at the beginning of the simulation.
func SetMesh(Nx, Ny, Nz int, cellSizeX, cellSizeY, cellSizeZ float64) {
	if Nx <= 1 {
		log.Fatal("mesh size X should be > 1, have: ", Nx)
	}
	globalmesh = *data.NewMesh(Nz, Ny, Nx, cellSizeZ, cellSizeY, cellSizeX)
	log.Println("set mesh:", Mesh().UserString())
	initialize()
}

// for lazy setmesh: set gridsize and cellsize in separate calls
var (
	gridsize []int
	cellsize []float64
)

func setGridSize(Nx, Ny, Nz int) {
	gridsize = []int{Nx, Ny, Nz}
	if cellsize != nil {
		SetMesh(Nx, Ny, Nz, cellsize[0], cellsize[1], cellsize[2])
	}
}

func setCellSize(cx, cy, cz float64) {
	cellsize = []float64{cx, cy, cz}
	if gridsize != nil {
		SetMesh(gridsize[0], gridsize[1], gridsize[2], cx, cy, cz)
	}
}

// check if mesh is set
func checkMesh() {
	if globalmesh.Size() == [3]int{0, 0, 0} {
		log.Panic("need to set mesh first")
	}
}

// check if m is set
func checkM() {
	checkMesh()
	if M.buffer.DevPtr(0) == nil {
		log.Fatal("need to initialize magnetization first")
	}
	if cuda.MaxVecNorm(M.buffer) == 0 {
		log.Fatal("need to initialize magnetization first")
	}
}

// Cleanly exits the simulation, assuring all output is flushed.
func Close() {
	log.Println("shutting down")
	drainOutput()
	log.Println("TODO: FLUSH TABLE")
	//Table.flush()
}
