import ssl
from cryptography.hazmat import backends
from cryptography.hazmat.primitives.asymmetric import dsa
from cryptography.hazmat.primitives.ciphers.modes import ECB

from pyOpenSSL import SSL


ssl.wrap_socket(ssl_version=ssl.PROTOCOL_SSLv3)


dsa.generate_private_key(key_size=2048,
                         backend=backends.default_backend())


dsa.generate_private_key(key_size=4096,
                         backend=backends.default_backend())

herp_derp(ssl_version=ssl.PROTOCOL_SSLv3)

ssl.wrap_socket()
