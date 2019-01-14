package main

func handleConfigCluster(urllist []string, cib Cib) (bool, interface{}) {

	nv := make(map[string]string)
	for _, item := range cib.Configuration.CrmConfig.ClusterPropertySet {
		for _, nv_item := range item.Nvpair {
			nv[nv_item.Name] = nv_item.Value
		}
	}
	return true, nv
}

func handleConfigRscDefaults(urllist []string, cib Cib) (bool, interface{}) {

	nv := make(map[string]string)
	for _, item := range cib.Configuration.RscDefaults.MetaAttributes {
		for _, nv_item := range item.Nvpair {
			nv[nv_item.Name] = nv_item.Value
		}
	}
	return true, nv
}

func handleConfigOpDefaults(urllist []string, cib Cib) (bool, interface{}) {

	nv := make(map[string]string)
	for _, item := range cib.Configuration.OpDefaults.MetaAttributes {
		for _, nv_item := range item.Nvpair {
			nv[nv_item.Name] = nv_item.Value
		}
	}
	return true, nv
}

func handleStateSummary(urllist []string, cib Cib) (bool, interface{}) {
	return true, nil
}
