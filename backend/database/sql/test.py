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
    host = "121.199.33.44"
    database = "dolphin"
    user = "root"
    password = "610878"

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
        cursor = self.conn.cursor()
        sql="INSERT INTO `user`(`id`,`name`,`password`,`self_group_id`) VALUES((null,'zyh',null,null), (null,'yyj','990731',null),(null,'yyn','123456','1'),(null,'asd','990731','3'))
        
        )
        "
        try:
            cursor.execute(sql)
            result = cursor.fetchall()
            assert(len(row) > 0)
            for row in result:
                id=row(0)
                assert(id == 1)
        except Error as e:
            print(e)


        # sql = "SELECT * FROM `user`;"
        # cursor.execute(sql)
        # result = cursor.fetchall()
        # for row in result:
        #     id = row[0]
        #     print(id)
        # insert to user
        # check user self_group_id

        pass
    def test_group(self):
        pass
    # add more tests

    def tearDown(self):
        self.conn.close()


if __name__ == '__main__':
    unittest.main()
