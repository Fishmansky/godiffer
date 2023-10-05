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

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(GodifferTestSuite))
}
