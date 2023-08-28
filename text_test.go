package cql2_test

import (
	"github.com/go-geospatial/cql2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Parsing OGC text examples",
	func(exampleNum string, text string) {
		_, err := cql2.ParseText(text)
		Expect(err).To(BeNil())
	},
	Entry("Find the Landsat scene with identifier 'LC82030282019133LGN00'", "example01", "landsat:scene_id = 'LC82030282019133LGN00'"),
	Entry("Find any item where the instrument name starts with 'OLI'", "example02", "eo:instrument LIKE 'OLI%'"),
)
