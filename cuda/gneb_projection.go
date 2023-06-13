package cuda

import (
	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
)

// dst += prefactor * dot(a, b), as used for energy density
func Projection(k, m *data.Slice) {

	util.Argument(k.Len() == m.Len())

	N := k.Len()
	cfg := make1DConf(N)
	k_projection_async(k.DevPtr(X), k.DevPtr(Y), k.DevPtr(Z),
		m.DevPtr(X), m.DevPtr(Y), m.DevPtr(Z),
		N, cfg)
}
