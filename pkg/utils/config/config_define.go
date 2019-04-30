// Copyright 2019 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import "time"

type Default struct{}

type OsdsApiServer struct {
	ApiEndpoint        string        `conf:"api_endpoint,localhost:50040"`
	AuthStrategy       string        `conf:"auth_strategy,noauth"`
	Daemon             bool          `conf:"daemon,false"`
	PolicyPath         string        `conf:"policy_path,/etc/opensds/policy.json"`
	LogFlushFrequency  time.Duration `conf:"log_flush_frequency,5s"` // Default value is 5s
	HTTPSEnabled       bool          `conf:"https_enabled,false"`
	BeegoHTTPSCertFile string        `conf:"beego_https_cert_file,/opt/opensds-security/opensds/opensds-cert.pem"`
	BeegoHTTPSKeyFile  string        `conf:"beego_https_key_file,/opt/opensds-security/opensds/opensds-key.pem"`
}

type OsdsLet struct {
	ApiEndpoint       string        `conf:"api_endpoint,localhost:50049"`
	Daemon            bool          `conf:"daemon,false"`
	LogFlushFrequency time.Duration `conf:"log_flush_frequency,5s"` // Default value is 5s
}

type OsdsfileLet struct {
	ApiEndpoint       string        `conf:"api_endpoint,localhost:50061"`
	Daemon            bool          `conf:"daemon,false"`
	LogFlushFrequency time.Duration `conf:"log_flush_frequency,5s"` // Default value is 5s
}

type OsdsDock struct {
	ApiEndpoint                string        `conf:"api_endpoint,localhost:50050"`
	DockType                   string        `conf:"dock_type,provisioner"`
	EnabledBackends            []string      `conf:"enabled_backends,lvm"`
	Daemon                     bool          `conf:"daemon,false"`
	BindIp                     string        `conf:"bind_ip"` // Just used for attacher dock
	HostBasedReplicationDriver string        `conf:"host_based_replication_driver,drbd"`
	LogFlushFrequency          time.Duration `conf:"log_flush_frequency,5s"` // Default value is 5s
	Backends
}

type FileShareBackends struct {
	NFSNative BackendProperties `conf:"nfsnative"`
}

type OsdsfileDock struct {
	ApiEndpoint                string        `conf:"api_endpoint,localhost:50051"`
	DockType                   string        `conf:"dock_type,provisioner"`
	FileShareEnabledBackends   []string      `conf:"enabled_backends,lvm"`
	Daemon                     bool          `conf:"daemon,false"`
	HostBasedReplicationDriver string        `conf:"host_based_replication_driver,drbd"`
	LogFlushFrequency          time.Duration `conf:"log_flush_frequency,5s"` // Default value is 5s
	FileShareBackends
}

type Database struct {
	Credential string `conf:"credential,username:password@tcp(ip:port)/dbname"`
	Driver     string `conf:"driver,etcd"`
	Endpoint   string `conf:"endpoint,localhost:2379,localhost:2380"`
}

type BackendProperties struct {
	Name               string `conf:"name"`
	Description        string `conf:"description"`
	DriverName         string `conf:"driver_name"`
	ConfigPath         string `conf:"config_path"`
	SupportReplication bool   `conf:"support_replication,false"`
}

type Backends struct {
	Ceph                BackendProperties `conf:"ceph"`
	Cinder              BackendProperties `conf:"cinder"`
	Sample              BackendProperties `conf:"sample"`
	LVM                 BackendProperties `conf:"lvm"`
	HuaweiDorado        BackendProperties `conf:"huawei_dorado"`
	HuaweiFusionStorage BackendProperties `conf:"huawei_fusionstorage"`
}

type KeystoneAuthToken struct {
	MemcachedServers  string `conf:"memcached_servers"`
	SigningDir        string `conf:"signing_dir"`
	Cafile            string `conf:"cafile"`
	AuthUri           string `conf:"auth_uri"`
	ProjectDomainName string `conf:"project_domain_name"`
	ProjectName       string `conf:"project_name"`
	UserDomainName    string `conf:"user_domain_name"`
	Password          string `conf:"password"`
	// Encryption and decryption tool. Default value is aes. The decryption tool can only decrypt the corresponding ciphertext.
	PwdEncrypter string `conf:"pwd_encrypter,aes"`
	// Whether to encrypt the password. If enabled, the value of the password must be ciphertext.
	EnableEncrypted bool   `conf:"enable_encrypted,false"`
	Username        string `conf:"username"`
	AuthUrl         string `conf:"auth_url"`
	AuthType        string `conf:"auth_type"`
}

type Config struct {
	Default           `conf:"default"`
	OsdsApiServer     `conf:"osdsapiserver"`
	OsdsLet           `conf:"osdslet"`
	OsdsfileLet       `conf:"osdsflet"`
	OsdsDock          `conf:"osdsdock"`
	OsdsfileDock      `conf:"osdsfdock"`
	Database          `conf:"database"`
	KeystoneAuthToken `conf:"keystone_authtoken"`
}
