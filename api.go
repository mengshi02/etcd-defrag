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

func RunWithoutCA(endpoints []string) error {
	defrag, err := NewDefrag(endpoints, nil)
	if  err != nil {
		return err
	}
	return defrag.Clean()
}
