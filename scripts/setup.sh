
# python3.7 -m pip install mysql-connector
# encryption
# python3.7 -m pip install cffi
# python3.7 -m pip install cryptography
# python3.7 -m pip install python-dotenv

# export .env files
env:
  export $(egrep -v '^#' .env | xargs)