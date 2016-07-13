package dedicated_node 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Ssl struct {

	/*Key - Descr: SSL private key for broker (PEM encoded) Default: <nil>
*/
	Key interface{} `yaml:"key,omitempty"`

	/*Pem - Descr: SSL Certificate for broker (PEM encoded) Default: <nil>
*/
	Pem interface{} `yaml:"pem,omitempty"`

}