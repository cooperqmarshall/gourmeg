postgres:
	podman run -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -d postgres:alpine

createdb:
	podman exec -it postgres createdb --username=root --owner=root gourmeg_2

dropdb:
	podman exec -it postgres dropdb gourmeg_2

tailwind:
	tailwindcss -i public/css/styles.css -o public/css/output.css --watch

