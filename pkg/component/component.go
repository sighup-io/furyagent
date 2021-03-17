// Copyright © 2018 Sighup SRL support@sighup.io
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

package component

import (
	"crypto/x509"
	"net"
	"time"

	"github.com/sighupio/furyagent/pkg/storage"
	certutil "k8s.io/client-go/util/cert"
)

var (
	CertConfig = certutil.Config{
		CommonName:   "SIGHUP s.r.l. Server",
		Organization: []string{"SIGHUP s.r.l."},
		AltNames:     certutil.AltNames{DNSNames: []string{}, IPs: []net.IP{}},
		Usages:       []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
)

type ClusterComponentData struct {
	*ClusterConfig
	*storage.Data
}

// ClusterComponent interface represent the basic concept of the component: etcd, master, node
type ClusterComponent interface {
	Backup() error
	Restore() error
	Configure(bool) error
	Init(string) error
}

// ClusterConfig represents the configuration for the whole cluster
type ClusterConfig struct {
	NodeName string        `mapstructure:"nodeName"`
	Etcd     EtcdConfig    `mapstructure:"etcd"`
	Master   MasterConfig  `mapstructure:"master"`
	Node     NodeConfig    `mapstructure:"node"`
	OpenVPN  OpenVPNConfig `mapstructure:"openvpn"`
	SSH      SSHConfig     `mapstructure:"sshkeys"`
}

// EtcdConfig is used to backup/restore/configure etcd nodes
type EtcdConfig struct {
	DataDir             string `mapstructure:"dataDir"`
	CertDir             string `mapstructure:"certDir"`
	CaCertFilename      string `mapstructure:"caCertFilename"`
	CaKeyFilename       string `mapstructure:"caKeyFilename"`
	ClientCertFilename  string `mapstructure:"clientCertFilename"`
	InitialClusterToken string `mapstructure:"initialClusterToken"`
	SnapshotFile        string `mapstructure:"snapshotFile"`
	ClientKeyFilename   string `mapstructure:"clientKeyFilename"`
	Endpoint            string `mapstructure:"endpoint"`
}

// MasterConfig is used to backup/restore/configure master nodes
type MasterConfig struct {
	CertDir          string `mapstructure:"certDir"`
	CaCertFile       string `mapstructure:"caCertFilename"`
	CaKeyFile        string `mapstructure:"caKeyFilename"`
	SaPubFile        string `mapstructure:"saPubFilename"`
	SaKeyFile        string `mapstructure:"saKeyFilename"`
	ProxyCaCertFile  string `mapstructure:"proxyCaCertFilename"`
	ProxyKeyCertFile string `mapstructure:"proxyKeyCertFilename"`
}

// NodeConfig is used to backup/restore/configure worker nodes (backup and restore have an empty implementation right now)
type NodeConfig struct {
	CloudProvider string        `mapstructure:"caKeyFilename"`
	joinTimeout   time.Duration `mapstructure:""joinTimeout`
}

type OpenVPNConfig struct {
	CertDir string   `mapstructure:"certDir"`
	Servers []string `mapstructure:"servers"`
}

type SSHConfig struct {
	User            string         `mapstructure:"user"`
	TempDir         string         `mapstructure:"tempDir"`
	LocalDirConfigs string         `mapstructure:"localDirConfigs"`
	Adapter         HTTPAdapterSet `mapstructure:"adapter"`
	Privileged      string           `mapstructure:"privileged"`
}
