package director 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Env struct {

	/*HttpsProxy - Descr: HTTPS proxy that the director, scheduler and workers should use Default: <nil>
*/
	HttpsProxy interface{} `yaml:"https_proxy,omitempty"`

	/*HttpProxy - Descr: HTTP proxy that the director, scheduler and workers should use Default: <nil>
*/
	HttpProxy interface{} `yaml:"http_proxy,omitempty"`

	/*NoProxy - Descr: List of comma-separated hosts that should skip connecting to the proxy in the director, scheduler and workers Default: <nil>
*/
	NoProxy interface{} `yaml:"no_proxy,omitempty"`

}