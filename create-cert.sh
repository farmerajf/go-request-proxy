#Create a private key (.key) and a certificate signing request (.csr)
openssl req -new -newkey rsa:2048 -nodes -keyout cert.key -out cert.csr

#Use the private key (.key) to sign the certificate signing request (.csr) to create a certificate (.crt)
openssl x509 -req -days 365 -in cert.csr -signkey cert.key -out cert.crt
