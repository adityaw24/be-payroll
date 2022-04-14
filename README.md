# PSM_Be-payroll

before run docker-compose, create network first

### step by step running docker compose

- create network with subnet 172.10.0.0 :

  - docker network create psm-payroll-network --subnet=172.10.0.0/16

menjalankan docker compose
docker-compose up
