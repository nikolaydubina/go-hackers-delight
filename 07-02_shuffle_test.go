package hd_test

import (
	"fmt"
	"testing"

	hd "github.com/nikolaydubina/go-hackers-delight"
)

func ExampleShuffleOuter() {
	// each symbol is a bit
	// in: abcd_efgh_ijkl_mnop_ABCD_EFGH_IJKL_MNPO
	// out: aAbB_cCdD_eEfF_gGhH_iI_jJ_kKlL_mMnN_oOpP
	fmt.Printf("%032b", hd.ShuffleOuter(0b1111_1111_0000_0000_0000_0000_1111_1111))
	// Output: 10101010101010100101010101010101
}

func ExampleShuffleInner() {
	// each symbol is a bit
	// in: abcd_efgh_ijkl_mnop_ABCD_EFGH_IJKL_MNPO
	// out: AaBb_CcDd_EeFf_GgHh_IiJj_KkLl_MmNn_OoPp
	fmt.Printf("%032b", hd.ShuffleInner(0b1111_1111_1111_1111_0000_0000_0000_0000))
	// Output: 01010101010101010101010101010101
}

func FuzzShuffleOuter(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		if got := hd.UnShuffleOuter(hd.ShuffleOuter(x)); x != got {
			t.Errorf("x=%0X got=%0X", x, got)
		}
	})
}

func FuzzShuffleInner(f *testing.F) {
	for _, x := range fuzzUint32 {
		f.Add(x)
	}
	f.Fuzz(func(t *testing.T, x uint32) {
		if got := hd.UnShuffleInner(hd.ShuffleInner(x)); x != got {
			t.Errorf("x=%0X got=%0X", x, got)
		}
	})
}
