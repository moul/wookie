package wookie

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("Testing Basic", t, func() {
		input := `abcd efgh cdef ghij`
		expected := `abcdefghij`
		wookie := NewWookie(input)
		So(wookie.Compute(), ShouldBeNil)
		So(wookie.Genome.String(), ShouldEqual, expected)
	})

	Convey("Testing Wookie", t, func() {
		input := `AAAAACAAAAGAAAATAAAC
AAGCAAAGGAAAGTAAA
AAGTGAAGTTAATACAATAGAAT
ACACGACACTACAGCACAGGACAGTACAT
ACCTGACCTTACGAGACGATACGCCACGCG
ACGCGACGCTACGGCACGGGACGGTAC
AGACTATACTCCACTCGACTCT
AGGCTAGGGCAGGGGAGGGTAGGTCAGGT
AGTAGATCAGATGAGATTAGCATAGCCC
ATTAGCATAGCCCAGCCGAGCCTAGCG
ATTCAATTGAATTTACACCACACG
CAACAGAACATAACCCAACCGA
CAGTCGAGTCTAGTGCAGTGGAGTGTAGTTCAG
CCATACCCCACCCGACCCTACC
CCATGCGATGCTATGGCATGGGATGGTATGT
CCTACCGCACCGGACCGTACCTCACCT
CCTAGCGCAGCGGAGCGTAGCTCAGCTGAGCTTA
CCTCCCGGCCCGTCCCTGCCCTTCCGCGCCGCTCC
CGAACCTAACGCAACGGAAC
CTTAAGACAAGAGAAGATAAGCC
CTTTGCTTTTGGGGGTGGGTTGGTGTGGT
GAACGTAACTCAACTGAACTTA
GATATTATCCCATCCGATCCTATCGCATCGGAT
GCTCCGGGCCGGTCCGTGCCGTTCCTCGCCT
GCTGAGCTTAGGATAGGCCAGGCGAGGCT
GGTACGTCACGTGACGTTACTAGACT
GTAAATCAAATGAAATTAACACAACAG
GTACATCACATGACATTACCAGACCATACC
GTATGTCATGTGATGTTATTCCATT
GTGTGGTTTGTGTTGTTTT
TAAACCAAACGAAACTAAAGCA
TAAGCCAAGCGAAGCTAAGG
TAAGGCAAGGGAAGGTAAGTCAAGTGA
TAGAATATAATCCAATCGAATCTAATGC
TAGTTCAGTTGAGTTTATATCATATGATATTA
TCAGGTGAGGTTAGTATAGTCCAG
TCGCCTCTCCTGGCCTGTCCTTGCCTTTCGCG
TCGGATCGTATCTCATCTGATCTTATGCCAT
TCGTCTCGTGGCGTGTCGTTGCG
TCTAATGCAATGGAATGTAATTCAA
TCTACTGCACTGGACTGTACTTCACTTG
TCTGGGCTGGTCTGTGCTGTTCTTGGCTTGTCTTTGCT
TGTCGTTGCGTTTCTCTGCTCTTCTGGGC
TTCCATTCGATTCTATTGCATTGGATTGTAT
TTCGCGGCGCGTCGCTGCGCTTCGG
TTCGGCTCGGGGCGGGTCGGTGCGGTTCGTCTCG
TTGACTTTAGAGCAGAGGAGAGTAG
TTGTATTTCATTTGATTTTCCCCCGCCCCTCCC`

		expected := `AAAAACAAAAGAAAATAAACCAAACGAAACTAAAGCAAAGGAAAGTAAATCAAATGAAATTAACACAACAGAACATAACCCAACCGAACCTAACGCAACGGAACGTAACTCAACTGAACTTAAGACAAGAGAAGATAAGCCAAGCGAAGCTAAGGCAAGGGAAGGTAAGTCAAGTGAAGTTAATACAATAGAATATAATCCAATCGAATCTAATGCAATGGAATGTAATTCAATTGAATTTACACCACACGACACTACAGCACAGGACAGTACATCACATGACATTACCAGACCATACCCCACCCGACCCTACCGCACCGGACCGTACCTCACCTGACCTTACGAGACGATACGCCACGCGACGCTACGGCACGGGACGGTACGTCACGTGACGTTACTAGACTATACTCCACTCGACTCTACTGCACTGGACTGTACTTCACTTGACTTTAGAGCAGAGGAGAGTAGATCAGATGAGATTAGCATAGCCCAGCCGAGCCTAGCGCAGCGGAGCGTAGCTCAGCTGAGCTTAGGATAGGCCAGGCGAGGCTAGGGCAGGGGAGGGTAGGTCAGGTGAGGTTAGTATAGTCCAGTCGAGTCTAGTGCAGTGGAGTGTAGTTCAGTTGAGTTTATATCATATGATATTATCCCATCCGATCCTATCGCATCGGATCGTATCTCATCTGATCTTATGCCATGCGATGCTATGGCATGGGATGGTATGTCATGTGATGTTATTCCATTCGATTCTATTGCATTGGATTGTATTTCATTTGATTTTCCCCCGCCCCTCCCGGCCCGTCCCTGCCCTTCCGCGCCGCTCCGGGCCGGTCCGTGCCGTTCCTCGCCTCTCCTGGCCTGTCCTTGCCTTTCGCGGCGCGTCGCTGCGCTTCGGCTCGGGGCGGGTCGGTGCGGTTCGTCTCGTGGCGTGTCGTTGCGTTTCTCTGCTCTTCTGGGCTGGTCTGTGCTGTTCTTGGCTTGTCTTTGCTTTTGGGGGTGGGTTGGTGTGGTTTGTGTTGTTTT`

		wookie := NewWookie(input)
		So(wookie.Compute(), ShouldBeNil)
		So(wookie.Genome.String(), ShouldEqual, expected)
	})

	Convey("Testing Impossible", t, func() {
		input := `abcd efgh`
		wookie := NewWookie(input)
		So(wookie.Compute(), ShouldNotBeNil)
	})
}

func TestWookie_Graphviz(t *testing.T) {
	Convey("Testing Wookie.Graphviz()", t, func() {
		input := `abcd efgh cdef ghij bcde`
		wookie := NewWookie(input)
		So(wookie.Compute(), ShouldBeNil)
		expected := `digraph G {
	start->abcd;
	ghij->end;
	abcd->bcde[ color=red, label=idx1_len3 ];
	bcde->cdef[ color=red, label=idx2_len3 ];
	cdef->efgh[ color=red, label=idx3_len2 ];
	efgh->ghij[ color=red, label=idx4_len2 ];
	abcd->cdef[ label=len2 ];
	abcd [ color=blue ];
	bcde [ color=blue ];
	cdef [ color=blue ];
	efgh [ color=blue ];
	end [ shape=Msquare ];
	ghij [ color=blue ];
	start [ shape=Mdiamond ];

}
`
		So(wookie.Graphviz(), ShouldEqual, expected)
	})
}
