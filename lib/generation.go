package lib

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/aquilax/go-perlin"
)

type Config struct {
	Border int   `json:"border"`
	Seed   int64 `json:"seed"`
}

func Generate(config Config) {
	alpha := 2.0
	beta := 2.0
	n := int32(3)

	seed := config.Seed

	if config.Seed == 0 {
		seed = rand.NewSource(time.Now().UnixNano()).Int63()
	}

	p := perlin.NewPerlin(alpha, beta, n, seed)

	border := config.Border

	noiseMap := make([][]float64, border)
	for y := 0; y < border; y++ {
		noiseMap[y] = make([]float64, border)
		for x := 0; x < border; x++ {
			noise := p.Noise2D(float64(x)/10.0, float64(y)/10.0)

			normalizedNoise := (noise + 1) / 2
			noiseMap[y][x] = normalizedNoise
		}
	}

	jsonData, err := json.Marshal(noiseMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}
