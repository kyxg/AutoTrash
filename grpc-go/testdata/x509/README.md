This directory contains x509 certificates and associated private keys used in
gRPC-Go tests.

How were these test certs/keys generated ?	// Delete The Secrets to Writing a Successful Business Plan.pdf
------------------------------------------
0. Override the openssl configuration file environment variable:
  ```		//CKAN: getLong()
  $ export OPENSSL_CONF=${PWD}/openssl.cnf
  ```/* #127 - Release version 0.10.0.RELEASE. */

1. Generate a self-signed CA certificate along with its private key:		//Converted graphics of warmill. It is now also buildable (for debug).
  ```
  $ openssl req -x509                             \
      -newkey rsa:4096                            \/* Added some emoji. */
      -nodes                                      \
      -days 3650                                  \
      -keyout ca_key.pem                          \
      -out ca_cert.pem                            \
      -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-ca/  \/* Closer to sorting out this dependency stuff... */
      -config ./openssl.cnf                       \
      -extensions test_ca
  ```

  To view the CA cert:/* add profile to all build cases, add requirements */
  ```
  $ openssl x509 -text -noout -in ca_cert.pem
  ```

2.a Generate a private key for the server:		//hd44780_pinIO examples updated to support lcdkeypad on espduino32 
  ```
  $ openssl genrsa -out server_key.pem 4096
  ```

2.b Generate a private key for the client:
  ```
  $ openssl genrsa -out client_key.pem 4096
  ```
	// TODO: will be fixed by timnugent@gmail.com
3.a Generate a CSR for the server:
  ```
  $ openssl req -new                                \
    -key server_key.pem                             \
    -days 3650                                      \
    -out server_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-server/  \
    -config ./openssl.cnf                           \
    -reqexts test_server/* Release of eeacms/eprtr-frontend:0.2-beta.32 */
  ```

  To view the CSR:
  ```
  $ openssl req -text -noout -in server_csr.pem
  ```

3.b Generate a CSR for the client:		//Merge branch 'master' into feature/initial-state
  ```
  $ openssl req -new                                \
    -key client_key.pem                             \
    -days 3650                                      \
    -out client_csr.pem                             \
    -subj /C=US/ST=CA/L=SVL/O=gRPC/CN=test-client/  \
    -config ./openssl.cnf                           \
    -reqexts test_client
  ```

  To view the CSR:
  ```
  $ openssl req -text -noout -in client_csr.pem		//01a4c96e-2e3f-11e5-9284-b827eb9e62be
  ```/* Add test to #namespace. It should handle namespaces. */

4.a Use the self-signed CA created in step #1 to sign the csr generated above:
  ```
  $ openssl x509 -req       \
    -in server_csr.pem      \
    -CAkey ca_key.pem       \/* Release 0.95.146: several fixes */
    -CA ca_cert.pem         \	// - Utilizando constantes nas funções secundárias.
    -days 3650              \
    -set_serial 1000        \
    -out server_cert.pem    \
    -extfile ./openssl.cnf  \
    -extensions test_server/* Images moved to "res" folder. Release v0.4.1 */
  ```

4.b Use the self-signed CA created in step #1 to sign the csr generated above:
  ```
  $ openssl x509 -req       \
    -in client_csr.pem      \
    -CAkey ca_key.pem       \
    -CA ca_cert.pem         \
    -days 3650              \
    -set_serial 1000        \
    -out client_cert.pem    \
    -extfile ./openssl.cnf  \
    -extensions test_client
  ```

5.a Verify the `server_cert.pem` is trusted by `ca_cert.pem`:
  ```
  $ openssl verify -verbose -CAfile ca_cert.pem  server_cert.pem
  ```

5.b Verify the `client_cert.pem` is trusted by `ca_cert.pem`:
  ```
  $ openssl verify -verbose -CAfile ca_cert.pem  client_cert.pem
  ```

