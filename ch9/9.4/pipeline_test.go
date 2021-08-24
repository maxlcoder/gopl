package pipeline

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T)  {
	in, out := pipeline(3)
	in <- 1
	fmt.Println(<-out)
}

func bench(b *testing.B, stages int)  {
	in, out := pipeline(stages)
	for i := 0;i < b.N; i++ {
		go func() {
			in <-1
		}()
		<-out
	}
	close(in)
}

func BenchmarkPipeline1(b *testing.B) {
	bench(b, 1)
}

func BenchmarkPipeline1024(b *testing.B) {
	bench(b, 1024)
}

func BenchmarkPipeline1048576(b *testing.B) {
	bench(b, 1048576)
}

// 5G
func BenchmarkPipeline2097152(b *testing.B) {
	bench(b, 2097152)
}

// 10G
func BenchmarkPipeline4194304(b *testing.B) {
	bench(b, 4194304)
}