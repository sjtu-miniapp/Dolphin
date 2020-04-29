import sys
import mysql.connector
from mysql.connector import Error
import unittest


# from dotenv import load_dotenv
# from ...utils.encryption.encrypt import decrypt


def connect():
    import os
    # load_dotenv('../../.env')
    # host = os.getenv("SQL_HOST")
    # database = os.getenv("SQL_DATABASE")
    # user = os.getenv("SQL_USER")
    # password = os.getenv("SQL_PASSWORD")

    """ Connect to MySQL database """
    conn = None
    try:
        conn = mysql.connector.connect(host=host,
                                       database=database,
                                       user=user,
                                       password=password)
    except Error as e:
        print(e)
    finally:
        return conn


class TestCase(unittest.TestCase):
    conn: mysql.connector.MySQLConnection

    def setUp(self):
        self.conn = connect()
        if not self.conn.is_connected():
            print("Connection to sql server failed!")
            sys.exit(-1)

    # TODO: test the database
    def test_user(self):
        pass

    def test_group(self):
        pass

    # add more tests

    def tearDown(self):
        self.conn.close()


if __name__ == '__main__':
    unittest.main()
