postgres:
	docker run -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=postgres -d postgres:alpine

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres gourmegdb

dropdb:
	docker exec -it postgres dropdb -U postgres gourmegdb

tailwind:
	tailwindcss -i public/css/styles.css -o public/css/output.css --watch

