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
	"log"

	"k8s.io/client-go/util/keyutil"
	"k8s.io/kubernetes/cmd/kubeadm/app/util/pkiutil"
)

const (
	MasterSaKey     = "sa.key"
	MasterSaPub     = "sa.pub"
	MasterFProxyCrt = "front-proxy-ca.crt"
	MasterFProxyKey = "front-proxy-ca.key"
	MasterCaKey     = "ca.key"
	MasterCaCrt     = "ca.crt"
	masterPath      = "pki/master"
)

// Master implements the ClusterComponent interface
type Master struct {
	ClusterComponentData
}

// Backup implements
func (m Master) Backup() error {
	return nil
}

// Restore implements
func (m Master) Restore() error {
	return nil
}

func (m Master) getFileMappings() [][]string {
	return [][]string{
		[]string{m.Master.CaCertFile, MasterCaCrt},
		[]string{m.Master.CaKeyFile, MasterCaKey},
		[]string{m.Master.SaKeyFile, MasterSaKey},
		[]string{m.Master.SaPubFile, MasterSaPub},
		[]string{m.Master.ProxyCaCertFile, MasterFProxyCrt},
		[]string{m.Master.ProxyKeyCertFile, MasterFProxyKey},
	}
}

// Configure implements
func (m Master) Configure(overwrite bool) error {
	// remove, create and download new certs
	files := m.getFileMappings()
	bucketDir := "pki/master"
	return m.DownloadFilesToDirectory(files, m.Master.CertDir, bucketDir, overwrite)
}

func (m Master) Init(dir string) error {
	// remove, create and download new certs
	caCert, caKey, err := pkiutil.NewCertificateAuthority(&CertConfig)
	if err != nil {
		log.Fatal(err)
	}
	saCert, saKey, err := pkiutil.NewCertificateAuthority(&CertConfig)
	if err != nil {
		log.Fatal(err)
	}
	fpCert, fpKey, err := pkiutil.NewCertificateAuthority(&CertConfig)
	if err != nil {
		log.Fatal(err)
	}
	masterCAKeyPEM, err := keyutil.MarshalPrivateKeyToPEM(caKey)
	if err != nil {
		log.Fatal(err)
	}
	masterSAKeyPEM, err := keyutil.MarshalPrivateKeyToPEM(saKey)
	if err != nil {
		log.Fatal(err)
	}
	masterFProxyKeyPEM, err := keyutil.MarshalPrivateKeyToPEM(fpKey)
	if err != nil {
		log.Fatal(err)
	}
	certs := map[string][]byte{
		MasterCaCrt:     pkiutil.EncodeCertPEM(caCert),
		MasterCaKey:     masterCAKeyPEM,
		MasterSaPub:     pkiutil.EncodeCertPEM(saCert),
		MasterSaKey:     masterSAKeyPEM,
		MasterFProxyCrt: pkiutil.EncodeCertPEM(fpCert),
		MasterFProxyKey: masterFProxyKeyPEM,
	}
	log.Printf("Writing files to %s ", masterPath)
	return m.UploadFilesFromMemory(certs, masterPath)
}
