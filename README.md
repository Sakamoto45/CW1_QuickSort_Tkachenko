# CW1. QuickSort. Ткаченко Егор Андреевич

[Параллельная реализация](parallel/parallel.go)

[Последовательная реализация](sequential/sequentioal.go)

[Бенчмарк для сравнения (с проверкой корректности)](qsort_test.go)

Результат запуска с командой `go test -bench=. -benchtime=1x -benchmem -cpu=4`:
```
goos: linux
goarch: amd64
pkg: quicksort
cpu: AMD Ryzen 7 5800X3D 8-Core Processor           
BenchmarkParQuickSorter/Par-4                  1        11605819187 ns/op
BenchmarkSeqQuickSorter/Seq-4                  1        35913375525 ns/op
PASS
ok      quicksort       79.406s
```

Разница 35913375525ns/11605819187ns = 3.094 раз