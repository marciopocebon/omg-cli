package registry

/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
 */
type Registry struct {

	/*Http - Descr: Username clients must use to access Registry via HTTP Basic Auth Default: <nil>
	 */
	Http *Http `yaml:"http,omitempty"`

	/*Db - Descr: Name of the registry database Default: bosh_registry
	 */
	Db *Db `yaml:"db,omitempty"`

	/* Hacked into this file as cannot specify to enaml-gen structs in same keyed map */

	/*Password - Descr: Password to access the Registry Default: <nil>
	 */
	Password interface{} `yaml:"password,omitempty"`

	/*Username - Descr: User to access the Registry Default: <nil>
	 */
	Username interface{} `yaml:"username,omitempty"`

	/*Host - Descr: Address of the Registry to connect to Default: <nil>
	 */
	Host interface{} `yaml:"host,omitempty"`

	/*Address - Descr: Address of the Registry to connect to Default: <nil>
	 */
	Address interface{} `yaml:"address,omitempty"`
}
