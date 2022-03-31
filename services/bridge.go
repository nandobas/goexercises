package services

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct {

}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

type RasterRenderer struct {
	dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

type Circle struct {
	renderer Renderer
	radius float32
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}