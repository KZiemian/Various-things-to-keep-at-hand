func MapKeys(m map[string]int) []string {
	var s []string

	for k := range m {
		s = append(s, k)
	}

	return s
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	var s []K

	for k := range m {
		s = append(s, k)
	}

	return s
}

type Tree[T any] struct {
	cmp  func(T, T) int
	root *leaf[T]
}

type leaf[T any] struct {
	val         T
	left, right *leaf[T]
}

func (bt *Tree[T]) find(val T) **leaf[T] {
	pl := &bt.root

	for *pl != nil {
		switch cmp := bt.cmp(val, (*pl).val); {
		case cmp < 0:
			pl = &(*pl).left

		case cmp > 0:
			pl = &(*pl).right

		default:
			return pl
		}
	}

	return pl
}

type SliceFn[T any] struct {
	s   []T
	cmp func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.s)
}

func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}

func (s SliceFn[T]) Less(i, j int) bool {
	return s.cmp(s.s[i], s.s[j])
}

func trace(r Ray, bounces int) (brightness float64) {
	if bounces <= 0 {
		return 0
	}

	hit, ok := r.Intersect(scene)
	if !ok {
		return 0
	}

	if hit.Light > 0 {
		return hit.Light
	}

	bounce, pdf := hit.RandomDirection()
	cos := bounce.Dot(hit.Normal)
	incoming := trace(bounce, bounces - 1)

	return (incoming * hit.Reflectance * cos) / pdf
}

// brdf(a, b) == brdf(b, a)
// R, G, B are each <= 1

func (m *Material) brdf(a, b Direction) RGB {
	// ...
	return m.Color
}

import "github.com/hunterloftis/pbr/obj"

traingles, err := obj.ReadFile("fixtures/model/lambo/lambo.obj")

func Glass(roughness float64) *Uniform {
	return &Uniform{
		Color:        rgb.Energy{1, 1, 1},
		Roughness:    roughness,
		Specularity:  0.042,
		Transmission: 0.91339,
	}
}
