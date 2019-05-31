package main

import (
	"io/ioutil"
	"fmt"
	"github.com/chikamim/go-resvg"
)

const svg = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-52 -53 100 100" stroke-width="2">
 <g fill="none">
  <ellipse stroke="#66899a" rx="6" ry="44"/>
  <ellipse stroke="#e1d85d" rx="6" ry="44" transform="rotate(-66)"/>
  <ellipse stroke="#80a3cf" rx="6" ry="44" transform="rotate(66)"/>
  <circle  stroke="#4b541f" r="44"/>
 </g>
 <g fill="#66899a" stroke="white">
  <circle fill="#80a3cf" r="13"/>
  <circle cy="-44" r="9"/>
  <circle cx="-40" cy="18" r="9"/>
  <circle cx="40" cy="18" r="9"/>
 </g>
</svg>`
func main() {
	ioutil.WriteFile("test.svg", []byte(svg), 0666)
	err := resvg.RenderPNGFromFile("test.svg", "test.png", &resvg.Options{BackgroundColor: "#eeddcc"})
	if err != nil {
		fmt.Println(err)
	}
	bb, _ := ioutil.ReadFile("test.png")
	fmt.Println(string(bb))
}
