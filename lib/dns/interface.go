package dns

type DNSResolverIf interface {
	LookupAddr(addr string) (names []string, err error)
}
