package firewall

import (
	"fmt"
	"github.com/coreos/go-iptables/iptables"
	"github.com/crazy-canux/go-devops/utils"
	"log"
	"strconv"
)

var ipts = map[string][]string{
	"filter": {
		"INPUT",
		"OUTPUT",
		"FORWARD",
	},
	"nat": {
		"INPUT",
		"OUTPUT",
		"PREROUTING",
		"POSTROUTING",
	},
	"mangle": {
		"INPUT",
		"OUTPUT",
		"FORWARD",
		"PREROUTING",
		"POSTROUTING",
	},
}

var ipt *iptables.IPTables

func init() {
	var err error
	ipt, err = iptables.New()
	if err != nil {
		log.Fatalf("get iptables failed: %s", err.Error())
	}
}

func CleanUDC() {
	for table := range ipts {
		chains := ListChain(table)
		for _, chain := range chains {
			if !utils.In(chain, ipts[table]) {
				ClearChain(table, chain)
				DeleteChain(table, chain)
			}
		}
	}
}

func CleanRules(table, chain string) {
	if table != "" {
		if chain != "" {
			ClearChain(table, chain)
		} else {
			for _, chain := range ipts[table] {
				ClearChain(table, chain)
			}
		}
	} else {
		for table := range ipts {
			for _, chain := range ipts[table] {
				ClearChain(table, chain)
			}
		}
	}
}

func AppendRules(table, chain string, rule []string) {
	err := ipt.AppendUnique(table, chain, rule...)
	if err != nil {
		log.Fatalf("Appendunique failed: %s", err.Error())
	}
}

func InsertRules(table, chain string, position int, rule []string) {
	err := ipt.Insert(table, chain, position, rule...)
	if err != nil {
		log.Fatalf("Insert failed: %s", err.Error())
	}
}

func DeleteRules(table, chain string, rule []string) {
	// delete rules in table-chain, iptables -t table -D chain rulenum
	err := ipt.Delete(table, chain, rule...)
	if err != nil {
		log.Fatalf("Delete failed: %s", err.Error())
	}
}

func Version() string {
	// get version
	major, minor, patch := ipt.GetIptablesVersion()
	return strconv.Itoa(major) + "." + strconv.Itoa(minor) + "." + strconv.Itoa(patch)
}

func List(table, chain string) []string {
	// iptables -t <table> -S <chain>
	rules, err := ipt.List(table, chain)
	if err != nil {
		log.Fatalf("list failed: %s", err.Error())
	}
	for _, rule := range rules {
		fmt.Println(rule)
	}
	return rules
}

func ListWithCounters(table, chain string) []string {
	// iptables -t <table> -L <chain> --line-number
	rules, err := ipt.ListWithCounters(table, chain)
	if err != nil {
		log.Fatalf("listwithcounters failed: %s", err.Error())
	}
	for _, rule := range rules {
		fmt.Println(rule)
	}
	return rules
}

func ListChain(table string) []string {
	chains, err := ipt.ListChains(table)
	if err != nil {
		log.Fatalf("listchains in %s failed: %s", table, err.Error())
	}
	for _, chain := range chains {
		fmt.Println(chain)
	}
	return chains
}

func ClearChain(table, chain string) {
	// clear rules in table-chain, iptables -t table -F chain
	err := ipt.ClearChain(table, chain)
	if err != nil {
		log.Fatalf("Clear rules in %s-%s failed: %s", table, chain, err.Error())
	}
}

func DeleteChain(table, chain string) {
	// delete chain in table, iptables -t table -X [chain]
	err := ipt.DeleteChain(table, chain)
	if err != nil {
		log.Fatalf("deletechain %s-%s failed: %s", table, chain, err.Error())
	}
}
