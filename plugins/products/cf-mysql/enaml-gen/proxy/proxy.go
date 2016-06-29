package proxy 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Proxy struct {

	/*Proxy - Descr: List of IP addresses for all proxy jobs Default: <nil>
*/
	Proxy *Proxy `yaml:"proxy,omitempty"`

	/*ClusterIps - Descr: List of nodes.  Must have the same number of ips as there are nodes in the cluster Default: <nil>
*/
	ClusterIps interface{} `yaml:"cluster_ips,omitempty"`

	/*ApiForceHttps - Descr: Redirect all HTTP requests to the API to HTTPS Default: true
*/
	ApiForceHttps interface{} `yaml:"api_force_https,omitempty"`

	/*SyslogAggregator - Descr: IP address for syslog aggregator Default: <nil>
*/
	SyslogAggregator *SyslogAggregator `yaml:"syslog_aggregator,omitempty"`

	/*Standalone - Descr: Standalone Mode: Are you deploying MySQL without a CloudFoundry deployment? Default: false
*/
	Standalone interface{} `yaml:"standalone,omitempty"`

	/*NetworkName - Descr: The name of the network (needed for the syslog aggregator) Default: <nil>
*/
	NetworkName interface{} `yaml:"network_name,omitempty"`

	/*Port - Descr: Port for the proxy to listen on Default: 3306
*/
	Port interface{} `yaml:"port,omitempty"`

	/*ExternalHost - Descr: Domain of the route registered for the UI via NATS (with the router in cf-release) Default: <nil>
*/
	ExternalHost interface{} `yaml:"external_host,omitempty"`

	/*ApiPort - Descr: Port for the proxy API to listen on Default: 80
*/
	ApiPort interface{} `yaml:"api_port,omitempty"`

	/*HealthPort - Descr: Port for checking the health of the proxy process Default: 1936
*/
	HealthPort interface{} `yaml:"health_port,omitempty"`

	/*ApiUsername - Descr: Username for Basic Auth used to secure API Default: <nil>
*/
	ApiUsername interface{} `yaml:"api_username,omitempty"`

	/*HealthcheckTimeoutMillis - Descr: Timeout (milliseconds) before assuming a backend is unhealthy Default: 5000
*/
	HealthcheckTimeoutMillis interface{} `yaml:"healthcheck_timeout_millis,omitempty"`

	/*Nats - Descr: IP of each NATS cluster member. Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*ProxyIps - Descr: List of IP addresses for all proxy jobs Default: <nil>
*/
	ProxyIps interface{} `yaml:"proxy_ips,omitempty"`

	/*ApiPassword - Descr: Password for Basic Auth used to secure API Default: <nil>
*/
	ApiPassword interface{} `yaml:"api_password,omitempty"`

}