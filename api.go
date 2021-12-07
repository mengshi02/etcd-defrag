package etcd_defrag

func Run(endpoints []string, cert, key, trusted string) error {
	ca := NewCA(cert, key, trusted)
	cfg, err := ca.GetConfig()
	if err != nil {
		return err
	}
	defrag, err := NewDefrag(endpoints, cfg)
	if  err != nil {
		return err
	}
	return defrag.Clean()
}
