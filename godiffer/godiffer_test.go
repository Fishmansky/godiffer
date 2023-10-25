package godiffer_test

import (
	"testing"

	"github.com/Fishmansky/godiffer/godiffer"
	"github.com/stretchr/testify/suite"
)

type GodifferTestSuite struct {
	suite.Suite
}

type Tests struct {
	str1 string
	str2 string
	want string
}

type LineTests struct {
	arr1 []string
	arr2 []string
	want string
}

type CotainsTests struct {
	str1 []byte
	str2 byte
	want bool
}

type CotainsLineTests struct {
	arr  []string
	str  string
	want bool
}

func (suite *GodifferTestSuite) TestContains() {
	tests := []CotainsTests{
		{str1: []byte{'a', 's', 'd', 'f'}, str2: 'a', want: true},
		{str1: []byte{'a', 's', 'd', 'f'}, str2: 'f', want: true},
		{str1: []byte{'a', 's', 'd', 'f'}, str2: 's', want: true},
		{str1: []byte{'a', 's', 'd', 'f'}, str2: 'w', want: false},
		{str1: []byte{}, str2: 'f', want: false},
		{str1: []byte{}, str2: ' ', want: false},
	}
	for _, tc := range tests {
		got := godiffer.Contains(tc.str1, tc.str2)
		suite.Suite.Equal(got, tc.want)
	}
}

func (suite *GodifferTestSuite) TestContainsLine() {
	tests := []CotainsLineTests{
		{arr: []string{"test", "test1", "test2"}, str: "test2", want: true},
		{arr: []string{"test", "test1", "test2"}, str: "test1", want: true},
		{arr: []string{"test", "test1", "test2"}, str: "tes", want: false},
		{arr: []string{"test", "test1", "test2"}, str: "", want: false},
		{arr: []string{}, str: "test2", want: false},
		{arr: []string{}, str: "", want: false},
	}
	for _, tc := range tests {
		got := godiffer.ContainsLine(tc.arr, tc.str)
		suite.Suite.Equal(got, tc.want)
	}
}

func (suite *GodifferTestSuite) TestUnion() {
	tests := []Tests{
		{str1: "asdf", str2: "sdfg", want: "asdfg"},
		{str1: "asdf", str2: "dfgh", want: "asdfgh"},
		{str1: "asdf", str2: "ghjk", want: "asdfghjk"},
	}
	for _, tc := range tests {
		got := godiffer.Union(tc.str1, tc.str2)
		suite.Suite.Equal(got, tc.want)
	}
}

func (suite *GodifferTestSuite) TestUnionLines() {
	tests := []LineTests{
		{arr1: []string{"test", "test1", "test2"}, arr2: []string{"test3", "test4", "test5"}, want: "test\ntest1\ntest2\ntest3\ntest4\ntest5"},
		{arr1: []string{"test", "test1", "test2"}, arr2: []string{"test1", "test2", "test3"}, want: "test\ntest1\ntest2\ntest3"},
		{arr1: []string{"test2", "test1", "test4"}, arr2: []string{"test", "test2", "test5"}, want: "test2\ntest1\ntest4\ntest\ntest5"},
		{arr1: []string{}, arr2: []string{"test", "test2", "test5"}, want: "test\ntest2\ntest5"},
		{arr1: []string{"test2", "test1", "test4"}, arr2: []string{}, want: "test2\ntest1\ntest4"},
		{arr1: []string{}, arr2: []string{}, want: ""},
	}
	for _, tc := range tests {
		got := godiffer.UnionLines(tc.arr1, tc.arr2)
		suite.Suite.Equal(got, tc.want)
	}
}

func (suite *GodifferTestSuite) TestIntersection() {
	tests := []Tests{
		{str1: "asdf", str2: "sdfg", want: "sdf"},
		{str1: "asdf", str2: "dfgh", want: "df"},
		{str1: "asdf", str2: "ghjk", want: ""},
	}
	for _, tc := range tests {
		got := godiffer.Intersection(tc.str1, tc.str2)
		suite.Suite.Equal(got, tc.want)
	}
}

func (suite *GodifferTestSuite) TestIntersectionLines() {
	tests := []LineTests{
		{arr1: []string{"test", "test1", "test2"}, arr2: []string{"test3", "test4", "test5"}, want: ""},
		{arr1: []string{"test", "test1", "test2"}, arr2: []string{"test1", "test2", "test3"}, want: "test1\ntest2"},
		{arr1: []string{"test2", "test1", "test4"}, arr2: []string{"test", "test2", "test5"}, want: "test2"},
		{arr1: []string{}, arr2: []string{"test", "test2", "test5"}, want: ""},
		{arr1: []string{"test2", "test1", "test4"}, arr2: []string{}, want: ""},
		{arr1: []string{}, arr2: []string{}, want: ""},
	}
	for _, tc := range tests {
		got := godiffer.IntersectionLines(tc.arr1, tc.arr2)
		suite.Suite.Equal(got, tc.want)
	}
}

func (suite *GodifferTestSuite) TestDifference() {
	tests := []Tests{
		{str1: "asdf", str2: "sdfg", want: "a"},
		{str1: "asdf", str2: "dfgh", want: "as"},
		{str1: "asdf", str2: "fghj", want: "asd"},
	}
	for _, tc := range tests {
		got := godiffer.Difference(tc.str1, tc.str2)
		suite.Suite.Equal(got, tc.want)
	}
}

func (suite *GodifferTestSuite) TestDifferenceLines() {
	tests := []LineTests{
		{arr1: []string{"test", "test1", "test2"}, arr2: []string{"test3", "test4", "test5"}, want: "test\ntest1\ntest2"},
		{arr1: []string{"test", "test1", "test2"}, arr2: []string{"test1", "test2", "test3"}, want: "test"},
		{arr1: []string{"test2", "test1", "test4"}, arr2: []string{"test", "test2", "test5"}, want: "test1\ntest4"},
		{arr1: []string{}, arr2: []string{"test", "test2", "test5"}, want: ""},
		{arr1: []string{"test2", "test1", "test4"}, arr2: []string{}, want: "test2\ntest1\ntest4"},
		{arr1: []string{}, arr2: []string{}, want: ""},
	}
	for _, tc := range tests {
		got := godiffer.DifferenceLines(tc.arr1, tc.arr2)
		suite.Suite.Equal(got, tc.want)
	}
}

func (suite *GodifferTestSuite) TestSymetricDifference() {
	tests := []struct {
		str1 string
		str2 string
		want string
	}{
		{str1: "asdf", str2: "sdfg", want: "ag"},
		{str1: "asdf", str2: "dfgh", want: "asgh"},
		{str1: "asdf", str2: "ghjk", want: "asdfghjk"},
	}
	for _, tc := range tests {
		got := godiffer.SymetricDifference(tc.str1, tc.str2)
		suite.Suite.Equal(got, tc.want)
	}
}

func (suite *GodifferTestSuite) TestSymetricDifferenceLines() {
	tests := []LineTests{
		{arr1: []string{"test", "test1", "test2"}, arr2: []string{"test3", "test4", "test5"}, want: "test\ntest1\ntest2\ntest3\ntest4\ntest5"},
		{arr1: []string{"test", "test1", "test2"}, arr2: []string{"test1", "test2", "test3"}, want: "test\ntest3"},
		{arr1: []string{"test2", "test1", "test4"}, arr2: []string{"test", "test2", "test5"}, want: "test1\ntest4\ntest\ntest5"},
		{arr1: []string{}, arr2: []string{"test", "test2", "test5"}, want: "test\ntest2\ntest5"},
		{arr1: []string{"test2", "test1", "test4"}, arr2: []string{}, want: "test2\ntest1\ntest4"},
		{arr1: []string{}, arr2: []string{}, want: ""},
	}
	for _, tc := range tests {
		got := godiffer.SymetricDifferenceLines(tc.arr1, tc.arr2)
		suite.Suite.Equal(got, tc.want)
	}
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(GodifferTestSuite))
}
