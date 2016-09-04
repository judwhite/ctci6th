package cubesums

import "testing"

func TestCubeTable(t *testing.T) {
	n := 100
	bruteResult := bruteForce(n)
	result := cubeTable(n)

	if len(bruteResult) != len(result) {
		t.Errorf("want: %d got: %d", len(bruteResult), len(result))
		return
	}

	for _, val := range bruteResult {
		foundAt := -1
		for i := 0; i < len(result); i++ {
			r := result[i]
			if val.a == r.a && val.b == r.b && val.c == r.c && val.d == r.d {
				foundAt = i
				break
			}
		}
		if foundAt == -1 {
			t.Errorf("couldn't find %v", val)
			return
		}
		tmp := result[:foundAt:foundAt]
		tmp = append(tmp, result[foundAt+1:]...)
		result = tmp
	}
}

func BenchmarkBruteForce300(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bruteForce(300)
	}
}

func BenchmarkCubeTable300(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cubeTable(300)
	}
}

func BenchmarkCubeTable1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cubeTable(1000)
	}
}
