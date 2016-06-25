package etcd_metrics_server 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Etcd struct {

	/*Port - Descr: port of ETCD server to instrument Default: 4001
*/
	Port interface{} `yaml:"port,omitempty"`

	/*Machine - Descr: address of ETCD server to instrument Default: 127.0.0.1
*/
	Machine interface{} `yaml:"machine,omitempty"`

}