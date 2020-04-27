from cryptography.fernet import Fernet
import pickle


def direct_store(name: str, passwd: str):
    dic = {}
    import os
    if os.path.exists("passwd.pkl"):
        with open('passwd.pkl', 'rb') as f:
            dic = pickle.load(f)
    dic.update({name: key})
    with open('passwd.pkl', 'wb+') as f:
        pickle.dump(dic, f, pickle.HIGHEST_PROTOCOL)
    return ciphered_text


def encrypt(name: str, passwd: str):
    key = Fernet.generate_key()
    cipher_suite = Fernet(key)
    ciphered_text = cipher_suite.encrypt(str.encode(passwd))
    dic = {}
    import os
    if os.path.exists("passwd.pkl"):
        with open('passwd.pkl', 'rb') as f:
            dic = pickle.load(f)
    dic.update({name: key})
    with open('passwd.pkl', 'wb+') as f:
        pickle.dump(dic, f, pickle.HIGHEST_PROTOCOL)
    return ciphered_text


def decrypt(name: str, ciphered_text: bytes):
    with open('passwd.pkl', 'rb') as f:
        dic = pickle.load(f)
        key = dic.get(name, None)
        assert key is not None
        cipher_suite = Fernet(key)
        passwd = cipher_suite.decrypt(ciphered_text)
        return passwd


if __name__ == "__main__":
    pwd = "test"
    name = "test"
    ciphered_text = encrypt(name, pwd)
    pwd = bytes.decode(decrypt(name, ciphered_text))
    assert pwd == "test"
