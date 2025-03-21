# LearningProject1



# PSQL things
[PSQL DB](https://neon.tech/postgresql/postgresql-getting-started/postgresql-sample-database)

```sql
                     List of relations
 Schema |            Name            |   Type   |  Owner
--------+----------------------------+----------+----------
 public | actor                      | table    | postgres
 public | actor_actor_id_seq         | sequence | postgres
 public | actor_info                 | view     | postgres
 public | address                    | table    | postgres
 public | address_address_id_seq     | sequence | postgres
 public | category                   | table    | postgres
 public | category_category_id_seq   | sequence | postgres
 public | city                       | table    | postgres
 public | city_city_id_seq           | sequence | postgres
 public | colors                     | table    | postgres
 public | colors_id_seq              | sequence | postgres
 public | country                    | table    | postgres
 public | country_country_id_seq     | sequence | postgres
 public | customer                   | table    | postgres
 public | customer_customer_id_seq   | sequence | postgres
 public | customer_list              | view     | postgres
 public | film                       | table    | postgres
 public | film_actor                 | table    | postgres
 public | film_category              | table    | postgres
 public | film_film_id_seq           | sequence | postgres
 public | film_list                  | view     | postgres
 public | inventory                  | table    | postgres
 public | inventory_inventory_id_seq | sequence | postgres
 public | language                   | table    | postgres
 public | language_language_id_seq   | sequence | postgres
 public | nicer_but_slower_film_list | view     | postgres
 public | payment                    | table    | postgres
 public | payment_payment_id_seq     | sequence | postgres
 public | rental                     | table    | postgres
 public | rental_rental_id_seq       | sequence | postgres
 public | sales_by_film_category     | view     | postgres
 public | sales_by_store             | view     | postgres
 public | staff                      | table    | postgres
 public | staff_list                 | view     | postgres
 public | staff_staff_id_seq         | sequence | postgres
 public | store                      | table    | postgres
 public | store_store_id_seq         | sequence | postgres
```
## Actor table
```sql
Column      |            Type             | Collation | Nullable |                 Default
-------------+-----------------------------+-----------+----------+-----------------------------------------
actor_id    | integer                     |           | not null | nextval('actor_actor_id_seq'::regclass)
first_name  | character varying(45)       |           | not null |
last_name   | character varying(45)       |           | not null |
last_update | timestamp without time zone |           | not null | now()
```
## Address
```sql
Column      |            Type             | Collation | Nullable |                   Default
-------------+-----------------------------+-----------+----------+---------------------------------------------
address_id  | integer                     |           | not null | nextval('address_address_id_seq'::regclass)
address     | character varying(50)       |           | not null |
address2    | character varying(50)       |           |          |
district    | character varying(20)       |           | not null |
city_id     | smallint                    |           | not null |
postal_code | character varying(10)       |           |          |
phone       | character varying(20)       |           | not null |
last_update | timestamp without time zone |           | not null | now()
```
## category
```sql
Column      |            Type             | Collation | Nullable |                    Default
------------+-----------------------------+-----------+----------+-----------------------------------------------
category_id | integer                     |           | not null | nextval('category_category_id_seq'::regclass)
name        | character varying(25)       |           | not null |
last_update | timestamp without time zone |           | not null | now()
```
## city
```sql
   Column    |            Type             | Collation | Nullable |                Default
-------------+-----------------------------+-----------+----------+---------------------------------------
 city_id     | integer                     |           | not null | nextval('city_city_id_seq'::regclass)
 city        | character varying(50)       |           | not null |
 country_id  | smallint                    |           | not null |
 last_update | timestamp without time zone |           | not null | now()
```
## country
```sql
   Column    |            Type             | Collation | Nullable |                   Default
-------------+-----------------------------+-----------+----------+---------------------------------------------
 country_id  | integer                     |           | not null | nextval('country_country_id_seq'::regclass)
 country     | character varying(50)       |           | not null |
 last_update | timestamp without time zone |           | not null | now()
```
## customer
```sql
   Column    |            Type             | Collation | Nullable |                    Default
-------------+-----------------------------+-----------+----------+-----------------------------------------------
 customer_id | integer                     |           | not null | nextval('customer_customer_id_seq'::regclass)
 store_id    | smallint                    |           | not null |
 first_name  | character varying(45)       |           | not null |
 last_name   | character varying(45)       |           | not null |
 email       | character varying(50)       |           |          |
 address_id  | smallint                    |           | not null |
 activebool  | boolean                     |           | not null | true
 create_date | date                        |           | not null | 'now'::text::date
 last_update | timestamp without time zone |           |          | now()
 active      | integer                     |           |          |
```
## film
```sql
      Column      |            Type             | Collation | Nullable |                Default
------------------+-----------------------------+-----------+----------+---------------------------------------
 film_id          | integer                     |           | not null | nextval('film_film_id_seq'::regclass)
 title            | character varying(255)      |           | not null |
 description      | text                        |           |          |
 release_year     | year                        |           |          |
 language_id      | smallint                    |           | not null |
 rental_duration  | smallint                    |           | not null | 3
 rental_rate      | numeric(4,2)                |           | not null | 4.99
 length           | smallint                    |           |          |
 replacement_cost | numeric(5,2)                |           | not null | 19.99
 rating           | mpaa_rating                 |           |          | 'G'::mpaa_rating
 last_update      | timestamp without time zone |           | not null | now()
 special_features | text[]                      |           |          |
 fulltext         | tsvector                    |           | not null |
```
## film_actor
```sql
   Column    |            Type             | Collation | Nullable | Default
-------------+-----------------------------+-----------+----------+---------
 actor_id    | smallint                    |           | not null |
 film_id     | smallint                    |           | not null |
 last_update | timestamp without time zone |           | not null | now()
```
## film_category
```sql
   Column    |            Type             | Collation | Nullable | Default
-------------+-----------------------------+-----------+----------+---------
 film_id     | smallint                    |           | not null |
 category_id | smallint                    |           | not null |
 last_update | timestamp without time zone |           | not null | now()
```
## film_list
```sql
                        View "public.film_list"
   Column    |          Type          | Collation | Nullable | Default
-------------+------------------------+-----------+----------+---------
 fid         | integer                |           |          |
 title       | character varying(255) |           |          |
 description | text                   |           |          |
 category    | character varying(25)  |           |          |
 price       | numeric(4,2)           |           |          |
 length      | smallint               |           |          |
 rating      | mpaa_rating            |           |          |
 actors      | text                   |           |          |

```
## language
```sql
   Column    |            Type             | Collation | Nullable |                    Default
-------------+-----------------------------+-----------+----------+-----------------------------------------------
 language_id | integer                     |           | not null | nextval('language_language_id_seq'::regclass)
 name        | character(20)               |           | not null |
 last_update | timestamp without time zone |           | not null | now()
```
## nicer_but_slower_film_list
```sql
               View "public.nicer_but_slower_film_list"
   Column    |          Type          | Collation | Nullable | Default
-------------+------------------------+-----------+----------+---------
 fid         | integer                |           |          |
 title       | character varying(255) |           |          |
 description | text                   |           |          |
 category    | character varying(25)  |           |          |
 price       | numeric(4,2)           |           |          |
 length      | smallint               |           |          |
 rating      | mpaa_rating            |           |          |
 actors      | text                   |           |          |
```
## payment
```sql
    Column    |            Type             | Collation | Nullable |                   Default
--------------+-----------------------------+-----------+----------+---------------------------------------------
 payment_id   | integer                     |           | not null | nextval('payment_payment_id_seq'::regclass)
 customer_id  | smallint                    |           | not null |
 staff_id     | smallint                    |           | not null |
 rental_id    | integer                     |           | not null |
 amount       | numeric(5,2)                |           | not null |
 payment_date | timestamp without time zone |           | not null |
```
## sales_by_film_category
```sql
   Column    |         Type          | Collation | Nullable | Default
-------------+-----------------------+-----------+----------+---------
 category    | character varying(25) |           |          |
 total_sales | numeric               |           |          |

```