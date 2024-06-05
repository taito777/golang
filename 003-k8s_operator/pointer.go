package main

import (
	"fmt"
	"math"
)

// interfaceの定義
type Abser interface {
	Abs() float64
}

// MyFloatはfloat64の構造体
type MyFloat float64

// MyFloatの値"f"をレシーバーとするAbs関数
// 使用する際は、f.Abs() ※fの部分はMyFloat型の値
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertexという構造体を定義
type Vertex struct {
	X, Y float64
}

// Vertexのポインタの"v"をレシーバーとするScale関数
// 使用する際は、v.Scale(f) ※vの部分はVertexのポインタ、fはfloat64型の引数
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Vertexのポインタの"v"をレシーバーとするAbs関数
// 使用する際は、v.Abs() ※vの部分はVertexのポインタ
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// main関数
func main() {
	var a Abser // aはAbserインターフェース
	var b Abser // bはAbserインターフェース
	// var c Abser // cはAbserインターフェース

	f := MyFloat(-math.Sqrt2) // fはfloat64型の値
	v := Vertex{3, 4}         // vはVertex構造体。{X, Y} = {3, 4}
	// k := float64(-math.Sqrt2)

	a = &v // aにvのメモリアドレスを代入。　※aはVertexのポインタ
	b = f  // bにfの値を代入。
	// c := &v // cにvの値を代入。
	//c.Scale(k)

	// Tips
	// c = v とするとエラーが発生。
	// Abs関数はVertexのポインタをレシーバーとするように定義されているため、
	// Vertexの値を直接レシーバーにすることはできない。

	fmt.Println(a.Abs())
	fmt.Println(b.Abs())
	fmt.Println(v)
	fmt.Println(a)
}
