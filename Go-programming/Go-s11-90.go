obj, dist := t.scene.Surface.Intersect(ray, infinity)

if obj == nil {
	return t.scene.Env.At(ray.Dir)
}

func (m *Material) absorb(wi, wo Vector3) (bool, Vector3, Vector3)

type Direction Vector3

// Unit normalizes a Vector3 into a Direction.
func (a Vector3) Unit() Direction {
	d := a.Len()
	return Direction{a.X / d, a.Y / d, a.Z / d}
}

type Energy geom.Vector3

func (e Energy) GammaCorrected() Energy {
	return Energy{gamma(e.X), gamma(e.Y), gamma(e.Z)}
}

func (m *Material) absorb(wi, wo Direction) (bool, Vector3, Energy)

type Stringer interface {
	String() string
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

package io

type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}
