<h1>Inventory Managment API</h1>

###

<p>This is a small project to show my Golang skills, using PostreSQL(though there is no SQL code)</p>

###

<p>Here you can log in as admin(to allow editing/adding stock), read, edit, delete and add stock</p>

###

<h2>Tech stack</h2>

###

<div align="left">
  <img width="12"  />
  <img src="https://cdn.simpleicons.org/go/00ADD8" height="40" alt="go logo"  />
</div>

###

<h2>Setup</h2>

###

<ul>
  <li>Firstly, you have to have <a href="https://www.postgresql.org/download/">PostgreSQL</a> installed with psql shell</li>
  <li>Now create a dotenv file</li>
  <li>Paste this in:</li>
  <pre>
    <code>DB_HOST=localhost </code>
    <code>DB_PORT=5432 </code>
    <code>DB_USER=postgres</code>
    <code>DB_PASSWORD=</code>
    <code>DB_NAME=inventory_db</code>
    <code>JWT_SECRET= </code>
  </pre>
  <li>In the DB_PASSWORD enter the password that you made when downloading PostgreSQL</li>
  <li>Uncomment the <code>utils.Token()</code> and <code>"github.com/VladVozhzhov/inventory-managment-api/utils"</code></li>
  <li>Paste the token fron the console into <code>JWT_SECRET=</code></li>
  <li>Open the postgres shell (psql), log in, and paste in this: <code>CREATE DATABASE inventory_db;</code></li>
</ul>

###

<h2>Start</h2>

###

<pre>
  <code>go mod tidy</code>
  <code>go run cmd/main.go</code>
</pre>  

###

<h2>Sample cURL's to test:</h2>

###

<h3>Register</h3>

###

<pre>
  <code>
    curl --location 'http://localhost:3500/register' \
    --header 'Content-Type: application/json' \
    --data '{ "username": "user", "password": "password" }'
  </code>
</pre>

###

<h3>Log in</h3>

###

<pre>
  <code>
    curl --location 'http://localhost:3500/login' \
    --header 'Content-Type: application/json' \
    --data '{ "username": "user", "password": "password" }'
  </code>
</pre>

###

<h3>Log out</h3>

<pre>
  <code>
    curl --location 'http://localhost:3500/logout' 
  </code>
</pre>

###

<h3>POST</h3>

###

<pre>
  <code>
    curl --location 'http://localhost:3500/admin/products' \
    --header 'Content-Type: application/json' \
    --data '{
        "name": "Phones",
        "sku": "SKU12345",
        "category": "Electronics",
        "quantity": 100,
        "supplier": "Acme Corp",
        "description": "A sample product for testing."
    }'
  </code>
</pre>

###

<h3>GET</h3>

###

<pre>
  <code>
    curl --location 'http://localhost:3500/products' 
  </code>
</pre>

###

<p>NOTE: there is no need to log in for this</p>

###

<h3>PUT</h3>

###

<pre>
  <code>
    curl --location --request PUT 'http://localhost:3500/admin/products' \
    --header 'Content-Type: application/json' \
    --data '{
        "name": "Phones",
        "sku": "SKU12345",
        "category": "Electronics",
        "quantity": 200,
        "supplier": "Acme Corp",
        "description": "A sample product for testing."
    }'
  </code>
</pre>

###

<p>NOTE: the id in the param will change in your product, so change it accordingly</p>

###

<h3>DELETE</h3>

###

<pre>
  <code>
    curl --location --request DELETE 'http://localhost:3500/admin/products?id=uivzib95uvj6jciakl85' \
    --header 'Content-Type: application/json' \
    --data '{
        "name": "Phones",
        "sku": "SKU12345",
        "category": "Electronics",
        "quantity": 200,
        "supplier": "Acme Corp",
        "description": "A sample product for testing."
    }'
  </code>
</pre>

###

<p>NOTE: the id in the param will change in your product, so change it accordingly</p>

###
