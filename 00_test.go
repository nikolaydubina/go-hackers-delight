package hd_test

import "testing"

func BenchmarkNoop(b *testing.B) { b.Run("--------------------------------", func(b *testing.B) {}) }
