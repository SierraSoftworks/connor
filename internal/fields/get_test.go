package fields_test

import (
	"encoding/json"
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SierraSoftworks/connor/internal/fields"
)

var _ = Describe("Get", func() {
	type TestCase struct {
		field string
		value string
		found bool
	}

	cases := map[string][]TestCase{
		`{"x": 1}`: {
			{
				"x",
				"1",
				true,
			},
			{
				"y",
				"null",
				false,
			},
			{
				"x.y",
				`null`,
				false,
			},
		},
		`{"x": null}`: {
			{
				"x",
				"null",
				true,
			},
		},
		`{"x": { "y": 1 }}`: {
			{
				"x.y",
				"1",
				true,
			},
			{
				"x",
				`{ "y": 1 }`,
				true,
			},
		},
		`{"x": [ { "y": 1 }, { "y" : 2 } ]}`: {
			{
				"x.0.y",
				`1`,
				true,
			},
			{
				"x.1.y",
				`2`,
				true,
			},
		},
		`{"x": [ { "y": [ 5,6] }]}`: {
			{
				"x.0.y.0",
				`5`,
				true,
			},
			{
				"x.0.y.1",
				`6`,
				true,
			},
			{
				"x.0.y.2",
				`null`,
				false,
			},
			{
				"x.-1.y.2",
				`null`,
				false,
			},
		},
		`{"x": [ { "y": 1 } ]}`: {
			{
				"x.0.z",
				`null`,
				false,
			},
			{
				"x.3.z",
				`null`,
				false,
			},
		},
	}

	for dataStr, cs := range cases {
		cs := cs
		Describe(fmt.Sprintf("with %s as data", dataStr), func() {
			for _, c := range cs {
				c := c
				Context(fmt.Sprintf("getting the field %s", c.field), func() {
					var value interface{}
					var expected interface{}
					var found bool

					BeforeEach(func() {
						var data map[string]interface{}
						Expect(json.NewDecoder(strings.NewReader(dataStr)).Decode(&data)).To(Succeed())
						Expect(json.NewDecoder(strings.NewReader(c.value)).Decode(&expected)).To(Succeed())
						value, found = fields.Get(data, c.field)
					})

					Context("Get()", func() {
						if c.found {
							It("should find the field", func() {
								Expect(found).To(BeTrue())
							})
						} else {
							It("should not find the field", func() {
								Expect(found).To(BeFalse())
							})
						}

						It(fmt.Sprintf("should return %s", c.value), func() {
							Expect(value).To(BeEquivalentTo(expected))
						})
					})

					Context("TryGet()", func() {
						It(fmt.Sprintf("should return %s", c.value), func() {
							Expect(value).To(BeEquivalentTo(expected))
						})
					})
				})
			}
		})
	}
})
