package main

import (
    "fmt"
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "path/filepath"
    "github.com/gocql/gocql"
)

func main() {
    var _cqlshrc_host = "31fecf38-2491-4d43-b6ce-22562679f1b8-us-east1.db.astra.datastax.com"
    var _cqlshrc_port = "34567"
    var _username = "erickramirez"
    var _password = "SomeComp7exP4ssword"

    _certPath, _ := filepath.Abs("/path/to/unzipped-bundle/cert")
    _keyPath, _ := filepath.Abs("/path/to/unzipped-bundle/key")
    _caPath, _ := filepath.Abs("/path/to/unzipped-bundle/ca.crt")
    _cert, _ := tls.LoadX509KeyPair(_certPath, _keyPath)
    _caCert, _  := ioutil.ReadFile(_caPath)
    _caCertPool := x509.NewCertPool()
    _caCertPool.AppendCertsFromPEM(_caCert)
    _tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{_cert},
        RootCAs:      _caCertPool,
    }

    _cluster := gocql.NewCluster(_cqlshrc_host)
    _cluster.SslOpts = &gocql.SslOptions{
        Config: _tlsConfig,
        EnableHostVerification: false,
    }
    _cluster.Authenticator = gocql.PasswordAuthenticator{
        Username:       _username,
        Password:       _password,
    }
    _cluster.Hosts = []string{_cqlshrc_host + ":" + _cqlshrc_port }

    session, _ := _cluster.CreateSession()
    var _rank int
    var _city string
    var _country string
    var _query = "SELECT rank, city, country FROM community.cities_by_rank WHERE rank IN ( 1, 2, 3, 4, 5 )"

    fmt.Println("According to independent.co.uk, the top 5 most liveable cities in 2019 were:")
    iter := session.Query(_query).Iter()
    for iter.Scan(&_rank, &_city, &_country) {
        fmt.Printf("\tRank %d: %s, %s\n", _rank, _city, _country)
    }
}
