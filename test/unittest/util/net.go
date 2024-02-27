package util

import (
	"net"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/kubeovn/kube-ovn/pkg/util"
)

var _ = ginkgo.Describe("[Net]", func() {
	ginkgo.It("AddressCount", func() {
		args := []*net.IPNet{
			{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(32, 32)},
			{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(31, 32)},
			{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(30, 32)},
			{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(24, 32)},
		}
		wants := []float64{
			0,
			0,
			2,
			254,
		}
		gomega.Expect(args).To(gomega.HaveLen(len(wants)))

		for i := range args {
			gomega.Expect(args[i].IP).NotTo(gomega.BeNil())
			gomega.Expect(util.AddressCount(args[i])).To(gomega.Equal(wants[i]))
		}
	})

	ginkgo.It("CountIPNums", func() {
		args := [][]string{
			{"10.0.0.101"},
			{"10.0.0.101..10.0.0.105"},
			{"10.0.0.101..10.0.0.105", "10.0.0.111..10.0.0.120"},
		}
		wants := []float64{
			1,
			5,
			15,
		}
		gomega.Expect(args).To(gomega.HaveLen(len(wants)))

		for i := range args {
			gomega.Expect(util.CountIPNums(args[i])).To(gomega.Equal(wants[i]))
		}
	})

	ginkgo.It("ExpandExcludeIPs", func() {
		type arg struct {
			cidr       string
			excludeIps []string
		}

		args := []arg{
			{"10.0.1.0/24", []string{"10.0.0.255"}},
			{"10.0.1.0/24", []string{"10.0.1.0"}},
			{"10.0.1.0/24", []string{"10.0.1.1"}},
			{"10.0.1.0/24", []string{"10.0.1.254"}},
			{"10.0.1.0/24", []string{"10.0.1.255"}},
			{"10.0.1.0/24", []string{"10.0.2.0"}},
			{"10.0.1.0/24", []string{"10.0.1.101..10.0.1.105"}},
			{"10.0.1.0/24", []string{"10.0.0.101..10.0.1.105"}},
			{"10.0.1.0/24", []string{"10.0.1.101..10.0.2.105"}},
			{"10.0.1.0/24", []string{"10.0.1.101..10.0.1.101"}},
			{"10.0.1.0/24", []string{"10.0.1.105..10.0.1.101"}},
			{"10.0.1.0/24", []string{"10.0.1.1", "10.0.1.101"}},
			{"10.0.1.0/24", []string{"10.0.1.1", "10.0.1.101..10.0.1.105"}},
			{"10.0.1.0/24", []string{"10.0.1.1", "10.0.1.101..10.0.1.105", "10.0.1.111..10.0.1.120"}},
			{"10.0.1.0/30", []string{"10.0.1.1", "179.17.0.0..179.17.0.10"}},
			{"10.0.1.0/31", []string{"10.0.1.1", "179.17.0.0..179.17.0.10"}},
			{"10.0.1.0/32", []string{"10.0.1.1", "179.17.0.0..179.17.0.10"}},

			{"fe00::100/120", []string{"fe00::ff"}},
			{"fe00::100/120", []string{"fe00::100"}},
			{"fe00::100/120", []string{"fe00::101"}},
			{"fe00::100/120", []string{"fe00::1fe"}},
			{"fe00::100/120", []string{"fe00::1ff"}},
			{"fe00::100/120", []string{"fe00::200"}},
			{"fe00::100/120", []string{"fe00::1a1..fe00::1a5"}},
			{"fe00::100/120", []string{"fe00::a1..fe00::1a5"}},
			{"fe00::100/120", []string{"fe00::1a1..fe00::2a5"}},
			{"fe00::100/120", []string{"fe00::1a1..fe00::1a1"}},
			{"fe00::100/120", []string{"fe00::1a5..fe00::1a1"}},
			{"fe00::100/120", []string{"fe00::101", "fe00::1a1"}},
			{"fe00::100/120", []string{"fe00::101", "fe00::1a1..fe00::1a5"}},
			{"fe00::100/120", []string{"fe00::101", "fe00::1a1..fe00::1a5", "fe00::1b1..fe00::1c0"}},
			{"fe00::100/126", []string{"fe00::101", "feff::..feff::a"}},
			{"fe00::100/127", []string{"fe00::101", "feff::..feff::a"}},
			{"fe00::100/128", []string{"fe00::101", "feff::..feff::a"}},

			{"10.0.1.0/24,fe00::100/120", []string{"10.0.1.1", "10.0.1.101..10.0.1.105"}},
			{"10.0.1.0/24,fe00::100/120", []string{"fe00::101", "fe00::1a1..fe00::1a5"}},
			{"10.0.1.0/24,fe00::100/120", []string{"10.0.1.1", "10.0.1.101..10.0.1.105", "fe00::101", "fe00::1a1..fe00::1a5"}},
		}
		wants := [][]string{
			{},
			{},
			{"10.0.1.1"},
			{"10.0.1.254"},
			{},
			{},
			{"10.0.1.101..10.0.1.105"},
			{"10.0.1.1..10.0.1.105"},
			{"10.0.1.101..10.0.1.254"},
			{"10.0.1.101"},
			{},
			{"10.0.1.1", "10.0.1.101"},
			{"10.0.1.1", "10.0.1.101..10.0.1.105"},
			{"10.0.1.1", "10.0.1.101..10.0.1.105", "10.0.1.111..10.0.1.120"},
			{"10.0.1.1"},
			{},
			{},

			{},
			{},
			{"fe00::101"},
			{"fe00::1fe"},
			{},
			{},
			{"fe00::1a1..fe00::1a5"},
			{"fe00::101..fe00::1a5"},
			{"fe00::1a1..fe00::1fe"},
			{"fe00::1a1"},
			{},
			{"fe00::101", "fe00::1a1"},
			{"fe00::101", "fe00::1a1..fe00::1a5"},
			{"fe00::101", "fe00::1a1..fe00::1a5", "fe00::1b1..fe00::1c0"},
			{"fe00::101"},
			{},
			{},

			{"10.0.1.1", "10.0.1.101..10.0.1.105"},
			{"fe00::101", "fe00::1a1..fe00::1a5"},
			{"10.0.1.1", "10.0.1.101..10.0.1.105", "fe00::101", "fe00::1a1..fe00::1a5"},
		}
		gomega.Expect(args).To(gomega.HaveLen(len(wants)))

		for i := range args {
			gomega.Expect(util.ExpandExcludeIPs(args[i].excludeIps, args[i].cidr)).To(gomega.Equal(wants[i]))
		}
	})
})
