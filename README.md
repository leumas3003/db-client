# db-client
DB Client for postgresql

## Installation and SetUp
 1. Download and install Golang. 
 2. Donwload and install PostgreSQL. 
 3. On PostgreSQL up and running use this script to create and add a new table into a database

```sql
CREATE TABLE IF NOT EXISTS public.movies
(
    id integer NOT NULL DEFAULT nextval('movies_id_seq'::regclass),
    movie_name text COLLATE pg_catalog."default",
    genre text COLLATE pg_catalog."default",
    duration text COLLATE pg_catalog."default",
    CONSTRAINT movies_pkey PRIMARY KEY (id)
)
```

 4. Replace in main.go ln 27 with your PostgreSQL data
```go
dbClient, err = sql.Open("postgres", fmt.Sprintf(DbConn, "localhost", "5432", "USERNAME", "", "DATABASENAME"))
```

## Package Used

| Function| Package|
|-------|--|
| Database | https://github.com/lib/pq |


## License

MIT License

Copyright (c) [2021] [Samuel Martel]

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
