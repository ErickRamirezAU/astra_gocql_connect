# astra_gocql_connect
Sample gocql code to connect to a [**DataStax Astra**](https://astra.datastax.com/) database.

This code was created in response to a post on the [**DataStax Community**](https://community.datastax.com/) site _"How do I connect to Astra with the gocql driver?"_.

## Build
Satisfy the required [Prerequisites](#prerequisites) in the next section.

Download [`astra_gocql_connect.go`](astra_gocql_connect.go) then build it with:
```
$ go build astra_gocql_connect.go
```

Finally, run the executable:
```
$ ./astra_gocql_connect
```

## Prerequisites
1. To compile and run the code, download and install the Go language from https://golang.org/.

2. An [**Astra** database](https://astra.datastax.com). Launching an Apache Cassandra™ database only takes a few clicks.

   It's free to try with no credit card required. Visit https://www.datastax.com/products/datastax-astra-beta for details.
   
3. [Download the secure connect bundle](https://docs.datastax.com/en/astra/aws/doc/dscloud/astra/dscloudObtainingCredentials.html) for your Astra database.

   Unzip your copy of `secure-connect-your_astra_db.zip` which will contain the following files:
   ```
   ca.crt
   cert
   cert.pfx
   config.json
   cqlshrc
   identity.jks
   key
   trustStore.jks
   ```
   
   You will need these files to configure the SSL/TLS options.
   
   The `cqlshrc` file contains the connection details for `_cqlshrc_host` and `_cqlshrc_port` in the code. Here is an example:
   ```
   [connection]
   hostname = 31fecf38-2491-4d43-b6ce-22562679f1b8-us-east1.db.astra.datastax.com
   port = 34567
   ssl = true
   ```
   
4. To run the sample code as-is, create the table `cities_by_rank`. Note that you can use your own keyspace so modify the code accordingly.
   ```
   CREATE KEYSPACE community WITH replication = {'class': 'NetworkTopologyStrategy', 'caas-dc': '1'};
   ```
   ```
   CREATE TABLE community.cities_by_rank (
       rank int PRIMARY KEY,
       city text,
       country text
   );
   ```
   Load the data in [`cities_by_rank.csv`](cities_by_rank.csv) using the cqlsh `COPY FROM` command:
   ```
   cqlsh> COPY cities_by_rank (rank,city,country) FROM './cities_by_rank.csv' WITH header = true;
   ```

## Sample data
The CSV file contains the top 10 most liveable cities of 2019 rated by The Independent. 
```
 rank | city       | country
------+------------+-----------
    1 |     Vienna |   Austria
    2 |  Melbourne | Australia
    3 |     Sydney | Australia
    4 |      Osaka |     Japan
    5 |    Calgary |    Canada
    6 |  Vancouver |    Canada
    7 |    Toronto |    Canada
    8 |      Tokyo |     Japan
    9 | Copenhagen |   Denmark
   10 |   Adelaide | Australia
```

Source: Helen Coffey, 'This is the world's most liveable city', _The Independent_, 4 September 2019 (accessed 2 May 2020), https://www.independent.co.uk/travel/news-and-advice/vienna-city-best-quality-life-study-ranked-austria-sydney-melbourne-osaka-a9091016.html

## Credit
Huge thanks to [@dougwettlaufer](https://github.com/dougwettlaufer) from DataStax Cloud Engineering for his help with the solution.

## Contact us
If you have any questions, hit me up at [**DataStax Community**](https://community.datastax.com/). Cheers!
